package domain

type MessagingTopics string

const (
	SendEmailTopic MessagingTopics = "email.send"
)

var TopicList = []string{
	string(SendEmailTopic),
}

type MessageEvent struct {
	EventTopic MessagingTopics `json:"eventTopic"`
	// Topic      AwsTopics
	Data map[string]interface{} `json:"data"`
}

/* type AwsTopics string

const (
	HandleEventsTopic AwsTopics = "handle-events"
)
*/
