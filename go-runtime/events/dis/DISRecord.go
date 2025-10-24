package dis

import (
	"fmt"
)

type DISRecord struct {
	// Partition key
	PartitionKey string `json:"partition_key"`

	// Data blocks, which are added by the data producer to the stream
	Data string `json:"data"`

	// Sequence number (unique identifier for each record)
	SequenceNumber string `json:"sequence_number"`
}

func (r *DISRecord) String() string {
	return fmt.Sprintf(`DISRecord{
                                 partition_key=%v,
                                 data=%v,
                                 sequence_number=%v
                               }`, r.PartitionKey, r.Data, r.SequenceNumber)
}
