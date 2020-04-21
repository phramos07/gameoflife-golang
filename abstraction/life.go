package abstraction

type ILife interface {
	Toggle()
	GetStatus() bool
	NewFrom(other interface{}) ILife
}
