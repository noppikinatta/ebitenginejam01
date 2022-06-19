package title

type state int

const (
	stateFadeIn state = iota
	stateWaitClick
	stateFadeOut
)
