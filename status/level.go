package status

import (
	"strconv"
)

// LevelCode -
type LevelCode int64

func (l LevelCode) String() string {
	return strconv.FormatInt(int64(l), 10)
}

//
const (
	LevelNONE LevelCode = iota
	LevelWARNING
	LevelERROR
	LevelFATAL
)
