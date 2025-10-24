/*
Using DDS triggers, each time a table in the database is updated, a FunctionGraph can be triggered to perform additional work.

For more information about how to use DDS trigger,
see https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_dds_trigger.html.
*/
package dds

import (
	"fmt"
)

type DDSTriggerEvent struct {
	Records []DDSRecord `json:"records"`
}

func (e *DDSTriggerEvent) String() string {
	return fmt.Sprintf(`DDSTriggerEvent{
                                  records=%+v
                               }`, e.Records)
}
