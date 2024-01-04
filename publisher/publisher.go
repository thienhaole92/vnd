package publisher

type Publisher interface {
	PublishMessage(message ...string) error
	// Close all underlying connections and resources
	Close() error
	// Topic that it is published
	Topic() string
}
