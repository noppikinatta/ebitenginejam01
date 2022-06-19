package result

type state int

const (
	stateFadeIn state = iota
	stateResultAnimation
	stateWaitClick
	stateFadeOut
)
