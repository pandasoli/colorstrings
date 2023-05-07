package errors

import "fmt"


type Error struct {
  Kind ErrorKind
  Err error
}

func NewError(kind ErrorKind, msg string) Error {
  return Error { kind, fmt.Errorf(msg) }
}
