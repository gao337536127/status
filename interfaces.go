package status

type Register interface {
	SetStatus(name string, statusFunc func() (string, error))
	GetStatus(name string) (string, error)
}
