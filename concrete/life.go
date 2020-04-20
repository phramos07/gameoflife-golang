package concrete

import (
	abstraction "gol/abstraction"
)

type Life struct {
	abstraction.ILife

	status bool
}

func NewLife() abstraction.ILife {
	return &Life{
		status: false,
	}
}

func (l *Life) NewFrom(other abstraction.ILife) abstraction.ILife {
	return &Life{
		status: other.GetStatus(),
	}
}

func (l *Life) GetStatus() bool {
	return l.status
}

func (l *Life) Toggle() {
	l.status = !l.status
}
