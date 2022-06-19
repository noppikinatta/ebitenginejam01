package prologue

type state int

const (
	stateMsg1Fadein state = iota
	stateMsg1WaitClick
	stateMsg1Fadeout
	stateMsg2Fadein
	stateMsg2WaitClick
	stateMsg2Fadeout
)
