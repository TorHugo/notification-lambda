package ports

type Consumer interface {
	Consume() string
}
