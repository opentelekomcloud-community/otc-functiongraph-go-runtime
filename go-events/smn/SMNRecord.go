package smn

import (
	"fmt"
)

type SMNRecord struct {
	// Event version. (Currently, the version is 1.0.)
	EventVersion string `json:"event_version"`

	// Subscription Uniform Resource Name (URN).
	EventSubscriptionUrn string `json:"event_subscription_urn"`

	// Event source
	EventSource string `json:"event_source"`

	// Message body
	Smn SMNBody `json:"smn"`
}

func (r *SMNRecord) String() string {
	return fmt.Sprintf(`SMNRecord{
                                 event_version=%v,
                                 event_subscription_urn=%v,
                                 event_source=%v,
                                 smn=%+v
                               }`, r.EventVersion, r.EventSubscriptionUrn, r.EventSource, r.Smn)
}
