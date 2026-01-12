package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"strings"

	"encoding/json"
	"regexp"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	obsevent "github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/obss3"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
	"golang.org/x/image/draw"
)

var validImageExtensions = regexp.MustCompile(`(?i)\.(jpg|jpeg|png)$`)

// handleRequest is your function handler method
func handleRequest(event []byte, ctx context.RuntimeContext) (string, error) {
	logger := ctx.GetLogger()

	var endpoint_url = ctx.GetUserData("OBS_ENDPOINT")
	if endpoint_url == "" {
		endpoint_url = "https://obs.otc.t-systems.com"
	}

	logger.Logf("OBS Endpoint from env: %s", endpoint_url)

	var obsEvent obsevent.OBSS3TriggerEvent
	err := json.Unmarshal(event, &obsEvent)
	if err != nil {
		logger.Logf("Unmarshal failed: %v", err)
		return "invalid data", err
	}

	logger.Logf("Function invoked with event: %v", obsEvent)

	var record = obsEvent.Records[0]
	var bucketName = record.S3.Bucket.Name
	var srcKey = record.S3.Object.Key

	logger.Logf("Bucket: %s, Object Key: %s", bucketName, srcKey)

	if !validImageExtensions.MatchString(srcKey) {
		logger.Logf("The object %s is not a valid image file. Skipping thumbnail creation.", srcKey)
		return "not a valid image file", errors.New("not a valid image file")
	}

	var dstBucket = ctx.GetUserData("OUTPUT_BUCKET")
	var dstKey = "resized-" + srcKey

	logger.Logf("Thumbnail will be saved to Bucket: %s, Object Key: %s", dstBucket, dstKey)

	var sAccessKeyId = ctx.GetSecurityAccessKey()
	var sAccessKeySecret = ctx.GetSecuritySecretKey()
	var sSessionToken = ctx.GetSecurityToken()

	logger.Logf("Access Key: %s", sAccessKeyId)
	logger.Logf("Secret Key: %s", sAccessKeySecret)
	logger.Logf("Session Token: %s", sSessionToken)

	obsClient, err := obs.New(sAccessKeyId, sAccessKeySecret, endpoint_url, obs.WithSecurityToken(sSessionToken))
	if err != nil {
		logger.Logf("client failed: %v", err)
		return "client failed", err
	}

	// Download the object from OBS
	input := &obs.GetObjectInput{}
	input.Bucket = bucketName
	input.Key = srcKey

	output, err := obsClient.GetObject(input)
	if err != nil {
		logger.Logf("getObject failed: %v", err)
		return "getObject failed", err
	}
	defer func() {
		errMsg := output.Body.Close()
		if errMsg != nil {
			logger.Logf("close failed: %v", errMsg)
			panic(errMsg)
		}
	}()

	body, err := io.ReadAll(output.Body)
	if err != nil {
		logger.Logf("readall failed: %v", err)
		return "readall failed", err
	}

	// Resize the image to create a thumbnail (200x200 pixels, maintaining aspect ratio)
	resizedData, format, err := resizeImage(body, 200, 200, true)
	if err != nil {
		logger.Logf("resize image failed: %v", err)
		return "resize image failed", err
	}
	logger.Logf("Image resized successfully. Original size: %d bytes, Resized size: %d bytes, Format: %s", len(body), len(resizedData), format)

	//  Save file to OBS
	save := &obs.PutObjectInput{}
	save.Bucket = dstBucket
	save.Key = dstKey

	// Set content type based on image format
	if strings.ToLower(format) == "png" {
		save.ContentType = "image/png"
	} else {
		save.ContentType = "image/jpeg"
	}

	save.Body = bytes.NewReader(resizedData)

	// Upload the object to OBS
	_, err = obsClient.PutObject(save)
	if err != nil {
		logger.Logf("putObject failed: %v", err)
		return "putObject failed", err
	}

	return fmt.Sprintf("Hello, %v!", string(event)), nil
}

// resizeImage resizes an image to the specified width and height
// It maintains aspect ratio if keepAspectRatio is true
func resizeImage(data []byte, width, height int, keepAspectRatio bool) ([]byte, string, error) {
	// Decode the image
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode image: %w", err)
	}

	// Get original dimensions
	bounds := img.Bounds()
	origWidth := bounds.Dx()
	origHeight := bounds.Dy()

	// Calculate new dimensions if keeping aspect ratio
	if keepAspectRatio {
		aspectRatio := float64(origWidth) / float64(origHeight)
		if width > 0 && height > 0 {
			// Both dimensions specified, adjust to fit
			targetAspectRatio := float64(width) / float64(height)
			if aspectRatio > targetAspectRatio {
				// Image is wider, fit to width
				height = int(float64(width) / aspectRatio)
			} else {
				// Image is taller, fit to height
				width = int(float64(height) * aspectRatio)
			}
		} else if width > 0 {
			// Only width specified
			height = int(float64(width) / aspectRatio)
		} else if height > 0 {
			// Only height specified
			width = int(float64(height) * aspectRatio)
		}
	}

	// Create a new image with the target dimensions
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// Use high-quality scaling
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, bounds, draw.Over, nil)

	// Encode the resized image
	var buf bytes.Buffer
	switch strings.ToLower(format) {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, dst, &jpeg.Options{Quality: 90})
	case "png":
		err = png.Encode(&buf, dst)
	default:
		// Default to JPEG for unknown formats
		err = jpeg.Encode(&buf, dst, &jpeg.Options{Quality: 90})
		format = "jpeg"
	}

	if err != nil {
		return nil, "", fmt.Errorf("failed to encode image: %w", err)
	}

	return buf.Bytes(), format, nil
}

// main function starts the runtime with your handler
func main() {
	runtime.Register(handleRequest)
}
