package test

import (
	"golang.org/x/xerrors"
)

// ErrCustomMessage var.
var ErrCustomMessage = xerrors.New("err: Error custom messagae")

// Foo func.
func Foo() error {
	return ErrCustomMessage
}

// Bar func.
func Bar() error {
	return xerrors.Errorf("bar: %w", Foo())
}

// Try func.
func Try() error {
	return xerrors.Errorf("try : %w", Bar())
}
