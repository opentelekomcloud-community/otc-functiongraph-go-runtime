package obs

import (
	"fmt"
)

type OBSS3UserIdentity struct {
	// customer-ID-of-the-user-who-caused-the-event
	PrincipalId string `json:"principalId"`
}

func (b *OBSS3UserIdentity) String() string {
	return fmt.Sprintf(`OBSS3UserIdentity{
                                 principalId=%v                                 
                               }`,
		b.PrincipalId)
}
