package ports

type Consumer interface {
	Consume(topic string) string
}
