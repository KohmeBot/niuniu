package util

import "encoding/gob"

func RegGob[T any]() {
	gob.Register(new(T))
}
