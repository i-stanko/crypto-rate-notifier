package storage

type SubscriberStore interface {
	Add(email string) error
	Exists(email string) (bool, error)
	List() ([]string, error)
}
