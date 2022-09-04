package drone

import (
	"log"
)

type GenericErr struct {
	Msg string
}

func (g *GenericErr) Error() string {
	return g.Msg
}

func WithLog(err error) *GenericErr {
	ge := New(err.Error())
	log.Println(ge.Msg)
	return &ge
}

func New(s string) GenericErr {
	return GenericErr{
		Msg: s,
	}
}
