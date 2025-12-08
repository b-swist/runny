package modes

type Entry interface {
	DefaultName() string
	Description() string
	Launch() error
}
