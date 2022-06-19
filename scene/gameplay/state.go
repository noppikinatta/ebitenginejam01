package gameplay

type state int

const (
	stateFadein state = iota
	stateCombine
	stateCombineComplete
	stateFadeout
)
