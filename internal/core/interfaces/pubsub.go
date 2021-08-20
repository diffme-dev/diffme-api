package interfaces

type PubSubProvider interface {
	publish()
	subscribe()
}
