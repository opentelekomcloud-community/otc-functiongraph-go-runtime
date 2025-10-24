/*
For more information about how to use Kafka trigger,
see: https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_kafka_trigger.html
*/
package kafka

import (
	"fmt"
)

type KAFKATriggerEvent struct {
	// Kafka instance ID
	InstanceId string `json:"instance_id"`

	Records []KAFKARecord `json:"records"`

	// Event type
	TriggerType string `json:"trigger_type"`

	// Region where a Kafka instance resides
	Region string `json:"region"`

	// Time when an event occurs
	EventTime int64 `json:"event_time"`

	// Event version
	EventVersion string `json:"event_version"`
}

func (e *KAFKATriggerEvent) String() string {
	return fmt.Sprintf(`KAFKATriggerEvent{
                                 instance_id=%v,
                                 records=%+v,
                                 trigger_type=%v,
                                 region=%v,
                                 event_time=%v,
                                 event_version=%v
                               }`, e.InstanceId, e.Records, e.TriggerType, e.Region, e.EventTime, e.EventVersion)
}
