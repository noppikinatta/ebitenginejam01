package prologue

type state int

const (
	stateMsgFadeIn state = iota
	stateMsgWaitClick
	stateMsgFadeOut
	stateLaunchFadeIn
	stateLaunchWaitClick
	stateLaunchFadeOut
)

func (s state) FadingIn() bool {
	switch s {
	case stateMsgFadeIn, stateLaunchFadeIn:
		return true
	}
	return false
}

func (s state) FadingOut() bool {
	switch s {
	case stateMsgFadeOut, stateLaunchFadeOut:
		return true
	}
	return false
}

func (s state) Msg() bool {
	switch s {
	case stateMsgFadeIn, stateMsgWaitClick, stateMsgFadeOut:
		return true
	}
	return false
}

func (s state) Launching() bool {
	return !s.Msg()
}
