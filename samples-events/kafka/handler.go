package main

import (
	"encoding/json"
	"fmt"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/kafka"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// Example for Kafka (Opensourcekafka) handler
func KafkaTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var kafkaEvent kafka.KAFKATriggerEvent
	err := json.Unmarshal(payload, &kafkaEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", kafkaEvent.String())
	return "ok", nil
}

func main() {
	runtime.Register(KafkaTest)
}
