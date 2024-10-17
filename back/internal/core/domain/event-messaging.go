package domain

type MessagingTopics string

const (
	SendEmailTopic              MessagingTopics = "email.send"
	PropagateUserUnsubscription MessagingTopics = "user.newsletter.unsubscribe"
)

var TopicList = []string{
	string(SendEmailTopic),
	string(PropagateUserUnsubscription),
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
