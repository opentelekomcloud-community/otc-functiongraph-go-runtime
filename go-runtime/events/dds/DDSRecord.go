package dds

import (
	"fmt"
)

type DDSRecord struct {
	// Event name
	EventName string `json:"event_name"`

	// Event protocol version
	EventVersion string `json:"event_version"`

	// Source of the event
	EventSource string `json:"event_source"`

	// The region where the DDS instance is located
	Region string `json:"region"`

	// DDS body
	Dds DDSBody `json:"dds"`

	// Event source unique identifier
	EventSourceId string `json:"event_source_id"`
}

func (r *DDSRecord) String() string {
	return fmt.Sprintf(`DDSRecord{
                                 event_name=%v,
                                 event_version=%v,
                                 event_source=%v,
                                 region=%v,
                                 dds=%+v,
                                 event_source_id=%v
                               }`, r.EventName, r.EventVersion, r.EventSource, r.Region, r.Dds, r.EventSourceId)
}
