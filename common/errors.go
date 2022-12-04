package common

import "errors"

var (
	ErrBrokenCommitment = errors.New("proposed and committed payloads differ")
)
