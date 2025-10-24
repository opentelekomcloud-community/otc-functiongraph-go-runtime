/*
You can schedule a timer to invoke your code based on a fixed rate of minutes, hours, or days or a cron expression.

For more information about how to use Timer trigger,
see https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_timer_trigger.html
*/
package timer

import (
	"fmt"
)

type TimerTriggerEvent struct {
	// Event version
	Version string `json:"version"`

	// Time when an event occurs. (e.g. 2018-06-01T08:30:00+08:00)
	Time string `json:"time"`

	// Trigger name
	TriggerName string `json:"trigger_name"`

	// Trigger type
	TriggerType string `json:"trigger_type"`

	// Additional information of the trigger
	UserEvent string `json:"user_event"`
}

func (e *TimerTriggerEvent) String() string {
	return fmt.Sprintf(`TimerTriggerEvent{
                                 version=%v,
                                 time=%v,
                                 trigger_name=%v,
                                 trigger_type=%v,
                                 user_event=%v
                               }`, e.Version, e.Time, e.TriggerName, e.TriggerType, e.UserEvent)
}
