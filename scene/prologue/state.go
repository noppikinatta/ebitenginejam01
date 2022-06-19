package prologue

type state int

const (
	stateMsg1FadeIn state = iota
	stateMsg1WaitClick
	stateMsg1FadeOut
	stateMsg2FadeIn
	stateMsg2WaitClick
	stateMsg2FadeOut
)
