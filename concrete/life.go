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

func (l *Life) NewFrom(other interface{}) abstraction.ILife {
	switch life := other.(type) {
	case abstraction.ILife:
		return &Life{
			status: life.GetStatus(),
		}
	default:
		panic("Couldn't create new life from provided object.")
	}
}

func (l *Life) GetStatus() bool {
	return l.status
}

func (l *Life) Toggle() {
	l.status = !l.status
}
