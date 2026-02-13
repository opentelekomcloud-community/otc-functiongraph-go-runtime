/*
You can write FunctionGraph functions to process logs subscribed to Cloud Log Service.
When Cloud Log Service collects subscribed logs, you can call FunctionGraph functions
by passing the collected logs as parameters ( LTS sample events ).
FunctionGraph function code can be customized, analyzed, or loaded into other systems.

For more information about how to use LTS trigger,
see https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_an_lts_trigger.html
*/
package lts

import (
	"fmt"
)

type LTSTriggerEvent struct {
	Lts LTSBody `json:"lts"`
}

func (e *LTSTriggerEvent) String() string {
	return fmt.Sprintf(`LTSTriggerEvent{
                                 lts=%+v
                               }`, e.Lts)
}
