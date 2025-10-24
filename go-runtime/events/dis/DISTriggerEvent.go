/*
Data Ingestion Service (DIS) can ingest large amounts of data in real time.
You can create a function to automatically poll a DIS stream and process all new data records,
such as website click streams, financial transactions, social media streams, IT logs, and location-tracking events.
FunctionGraph periodically polls the stream for new data records.

For more information about how to use DIS trigger,
see: https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_dis_trigger.html#functiongraph-01-0206
*/
package dis

import (
	"fmt"
)

type DISTriggerEvent struct {
	// Partition ID
	ShardID string

	// DIS message body ( DISMessage structure )
	Message DISMessage

	// Function version
	Tag string

	// Channel Name
	StreamName string
}

func (e *DISTriggerEvent) String() string {
	return fmt.Sprintf(`DISTriggerEvent{
                                  ShardID=%v,
                                  Message=%+v,
                                  Tag=%v,
                                  StreamName=%v
                               }`, e.ShardID, e.Message, e.Tag, e.StreamName)
}
