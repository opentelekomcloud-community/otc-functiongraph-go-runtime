/*
According to the CTS cloud audit service type and the event notification required for

the operation subscription, when the CTS cloud audit service obtains the subscribed
operation record, the collected operation record is passed as a parameter ( CTS sample event )
through the CTS trigger to call the FunctionGraph function.
Through the function, the key information in the log is analyzed and processed, and the system,
network and other business modules are automatically repaired, or alarms are generated through SMS,
email, etc. to notify business personnel to handle.

For more information about how to use CTS trigger,
see: https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_cts_trigger.html
*/
package cts

import (
	"fmt"
)

type CTSTriggerEvent struct {
	Cts CTS `json:"cts"`
}

func (e *CTSTriggerEvent) String() string {
	return fmt.Sprintf(`CTSTriggerEvent{
                                  cts=%+v
                               }`, e.Cts)
}
