package obs

import (
	"fmt"
)

type OBSS3BucketOwnerIdentity struct {
	// customer-ID-of-the-bucket-owner
	PrincipalId string `json:"PrincipalId"`
}

func (b *OBSS3BucketOwnerIdentity) String() string {
	return fmt.Sprintf(`OBSS3BucketOwnerIdentity{
                                 PrincipalId=%v                                 
                               }`,
		b.PrincipalId)
}
