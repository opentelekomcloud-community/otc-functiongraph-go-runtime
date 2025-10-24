/*
Triggers on OBS can only be used for FunctionGraphs in the main project, not in sub projects!

For each OBS bucket, only one FunctionGraph can be triggered (no multiple FunctionGraphs listening on same bucket)

For more information about how to use LTS triggers,
see: https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_an_obs_trigger.html
*/
package obs

import (
	"fmt"
)

type OBSS3TriggerEvent struct {
	Records []OBSS3Record `json:"Records"`
}

func (e *OBSS3TriggerEvent) String() string {
	return fmt.Sprintf(`OBSS3TriggerEvent{
                                  records=%+v
                               }`,
		e.Records)
}
