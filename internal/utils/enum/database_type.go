package enum

type DatabaseType int

const (
	Main DatabaseType = iota
)

func (e DatabaseType) String() string {
	switch e {
	case Main:
		return "main"
	default:
		return ""
	}
}
