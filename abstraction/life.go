package abstraction

type ILife interface {
	Toggle()
	GetStatus() bool
	NewFrom(other ILife) ILife
}
