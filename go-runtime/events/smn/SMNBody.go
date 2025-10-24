package smn

import (
	"fmt"
)

type SMNBody struct {
	// ID of an SMN event
	TopicUrn string `json:"topic_urn"`

	// time when event occured
	TimeStamp string `json:"timestamp"`

	// message attributes
	MessageAttributes map[string]string `json:"message_attributes"`

	// message body
	Message string `json:"message"`

	// Message format
	Type string `json:"type"`

	// Message ID. The ID of each message is unique.
	MessageId string `json:"message_id"`

	// Subject of message
	Subject string `json:"subject"`
}

func (b *SMNBody) String() string {
	return fmt.Sprintf(`SMNBody{
                                 topic_urn=%v,
                                 timestamp=%v,
                                 message_attributes=%v,
                                 message=%v,
                                 type=%v,
                                 message_id=%v,
                                 subject=%v
                               }`, b.TopicUrn, b.TimeStamp, b.MessageAttributes, b.Message, b.Type, b.MessageId, b.Subject)
}
