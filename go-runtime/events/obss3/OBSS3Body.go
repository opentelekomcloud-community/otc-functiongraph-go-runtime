package obs

import (
	"fmt"
)

type OBSS3Body struct {
	// Schema Version: 1.0
	S3SchemaVersion string `json:"s3SchemaVersion"`

	// Configuration Id. This ID can be found in the bucket notification configuration.
	ConfigurationId string `json:"configurationId"`

	// Bucket property
	Bucket OBSS3Bucket `json:"bucket"`

	// Object property.
	Object OBSS3Object `json:"object"`
}

func (b *OBSS3Body) String() string {
	return fmt.Sprintf(`OBSS3Body{
                                 s3SchemaVersion=%v,
																 configurationId=%v,
																 bucket=%+v,
                                 object=%+v,
                               }`,
		b.S3SchemaVersion,
		b.ConfigurationId,
		b.Bucket,
		b.Object)
}
