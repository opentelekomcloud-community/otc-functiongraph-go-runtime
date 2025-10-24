package cts

import (
	"fmt"
)

type CTS struct {
	// time in epoch
	Time int64 `json:"time"`

	// Information about the user who initiated this request
	User User `json:"user"`

	// Event request content
	Request map[string]string `json:"request"`

	// Incident response content
	Response map[string]string `json:"response"`

	// 	Event response code, such as 200, 400
	Code int `json:"code"`

	// Abbreviation of the sender, such as vpc, ecs, etc.
	ServiceType string `json:"service_type"`

	// The sender resource type, such as vm, vpn, etc.
	ResourceType string `json:"resource_type"`

	// Resource name, such as the name of a virtual machine in the ecs service
	ResourceName string `json:"resource_name"`

	// Resource Id
	ResourceId string `json:"resource_id"`

	// Event name, such as: startServer, shutDown, etc.
	TraceName string `json:"trace_name"`

	// The event source type, such as ApiCall
	TraceType string `json:"trace_type"`

	// The time when the cts service receives this trace (Epoch timestamp in milliseconds)
	RecordTime int64 `json:"record_time"`

	// Unique identifier for the event
	TraceId string `json:"trace_id"`

	// Status of the event
	TraceStatus string `json:"trace_status"`
}

func (cts *CTS) String() string {
	return fmt.Sprintf(`CTS{
                                  time='%v',
                                  user=%+v,
                                  request=%v,
                                  response=%v,
                                  code=%v,
                                  service_type='%v',
                                  resource_type='%v',
                                  resource_name='%v',
                                  resource_id='%v',
                                  trace_name='%v',
                                  trace_type='%v',
                                  record_time='%v',
                                  trace_id='%v',
                                  trace_status='%v'
                               }`, cts.Time, cts.User, cts.Request, cts.Response, cts.Code, cts.ServiceType, cts.ResourceType,
		cts.ResourceName, cts.ResourceId, cts.TraceName, cts.TraceType, cts.RecordTime, cts.TraceId,
		cts.TraceStatus)
}
