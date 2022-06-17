package title

type state int

const (
	stateFadein state = iota
	stateWaitClick
	stateFadeout
)
