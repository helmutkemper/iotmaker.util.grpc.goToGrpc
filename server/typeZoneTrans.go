package server

type ZoneTrans struct {
	When  int64
	Index uint8
	IsStd bool
	IsUtc bool
}
