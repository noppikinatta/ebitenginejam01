package result

type state int

const (
	stateFadein state = iota
	stateResultAnimation
	stateWaitClick
	stateFadeout
)
