package gameplay

type state int

const (
	stateFadeIn state = iota
	stateCombine
	stateCombineComplete
	stateFadeOut
)
