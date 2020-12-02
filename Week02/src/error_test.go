package test

import (
	"testing"
)

func TestError(t *testing.T) {
	err := Try()
	t.Errorf("%v", err)
	t.Errorf("%+v", err)
}
