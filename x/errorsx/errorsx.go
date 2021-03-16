package errorsx

import (
	"errors"
)

var (
	Is = errors.Is
	As = errors.As
	New = errors.New
)

type Sentinel string
func (e Sentinel) Error() string { return string(e) }
