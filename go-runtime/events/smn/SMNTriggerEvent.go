/*
Simple Message Notification (SMN) sends messages to email addresses, mobile phones, or HTTP/HTTPS URLs.
If you create a function with an SMN trigger, messages published to a specified topic will be passed as a parameter to invoke the function.
Then, the function processes the event, for example, publishing messages to other SMN topics or sending them to other cloud services.

For more information about how to use SMN trigger,
https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_an_smn_trigger.html
*/
package smn

import (
	"fmt"
)

type SMNTriggerEvent struct {
	Record []SMNRecord `json:"record"`
}

func (e *SMNTriggerEvent) String() string {
	return fmt.Sprintf(`SMNTriggerEvent{
                                 record=%+v
                               }`, e.Record)
}
