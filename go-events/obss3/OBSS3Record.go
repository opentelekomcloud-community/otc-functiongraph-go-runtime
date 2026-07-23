package obs

import (
	"fmt"
)

type OBSS3Record struct {
	// event version: 2.0
	EventVersion string `json:"eventVersion"`
	EventSource  string `json:"eventSource"`

	// AwsRegion.
	AWSRegion string `json:"awsRegion"`

	// The time, in ISO-8601 format, for example, 1970-01-01T00:00:00.000Z, when OTC finished processing the request
	EventTime string `json:"eventTime"`

	/*
	  Event name:
	  - ObjectCreated:Put, ObjectCreated:Post, ObjectCreated:Copy
	      Operations such as PUT, POST, and COPY can create an object.
	      With these event types, you can enable notifications when an object is created using a specific API operation.

	  - Created:CompleteMultipartUpload
	      ObjectCreated:CompleteMultipartUpload includes objects that are created using UploadPartCopy for Copy operations.

	  - ObjectRemoved:Delete, ObjectRemoved:DeleteMarkerCreated
	      By using the ObjectRemoved event types, you can enable notification
	      when an object or a batch of objects is removed from a bucket.
	      You can request notification when an object is deleted or a versioned object is permanently
	      deleted by using the s3:ObjectRemoved:Delete event type.
	      Alternatively, you can request notification when a delete marker is created for a versioned
	      object using s3:ObjectRemoved:DeleteMarkerCreated
	      These event notifications don't alert you for automatic deletes from lifecycle
	      configurations or from failed operations.</description>
	*/
	EventName string `json:"eventName"`

	UserIdentity OBSS3UserIdentity `json:"userIdentity"`

	// Gets and sets the RequestParameters property
	RequestParameters map[string]string `json:"requestParameters"`

	// Gets and sets the ResponseElements property
	ResponseElements map[string]string `json:"responseElements"`

	// Gets and sets the S3 property.
	S3 OBSS3Body `json:"s3"`
}

func (r *OBSS3Record) String() string {
	return fmt.Sprintf(`OBSS3Record{
                                 eventVersion=%v,
                                 eventSource=%v,
                                 awsRegion=%v,
                                 eventTime=%v,
                                 eventName=%v,
                                 userIdentity=%+v,
                                 requestParameters=%v,
                                 responseElements=%v,
                                 s3=%+v
                               }`,
		r.EventVersion,
		r.EventSource,
		r.AWSRegion,
		r.EventTime,
		r.EventName,
		r.UserIdentity,
		r.RequestParameters,
		r.ResponseElements,
		r.S3)
}
