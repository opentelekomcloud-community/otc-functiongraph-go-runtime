package obs

import (
	"fmt"
)

type OBSS3Bucket struct {
	// Bucket name
	Name string `json:"name"`

	// Customer-ID-of-the-bucket-owner
	OwnerIdentity OBSS3BucketOwnerIdentity `json:"ownerIdentity"`

	// Bucket arn
	Arn string `json:"arn"`
}

func (b *OBSS3Bucket) String() string {
	return fmt.Sprintf(`OBSS3Bucket{
                                 name=%v,
                                 ownerIdentity=%+v,
                                 arn=%v
                               }`,
		b.Name,
		b.OwnerIdentity,
		b.Arn)
}
