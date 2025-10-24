package dds

import (
	"fmt"
)

type DDSBody struct {
	// The number of bytes in the message
	SizeBytes string `json:"size_bytes"`

	// Base64 encoded data
	Token string `json:"token"`

	// Complete file information
	FullDocument string `json:"full_document"`

	// Column Name
	Ns DDSNSBody `json:"ns"`
}

func (b *DDSBody) String() string {
	return fmt.Sprintf(`DDSBody{
                                 size_bytes=%v,
                                 token=%v,
                                 full_document=%v,
                                 ns=%+v
                               }`, b.SizeBytes, b.Token, b.FullDocument, b.Ns)
}
