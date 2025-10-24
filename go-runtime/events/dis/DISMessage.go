package dis

import (
	"fmt"
)

type DISMessage struct {
	// Next partition cursor
	NextPartitionCursor string `json:"next_partition_cursor"`

	// Data records stored in a DIS stream
	Records []DISRecord `json:"records"`

	// Reserved Fields
	MillisBehindLatest string `json:"millis_behind_latest"`
}

func (d *DISMessage) String() string {
	return fmt.Sprintf(`DISMessage{
                                 next_partition_cursor=%v,
                                 records=%+v,
                                 millisBehindLatest=%v
                               }`, d.NextPartitionCursor, d.Records, d.MillisBehindLatest)
}
