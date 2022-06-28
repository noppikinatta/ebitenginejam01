package result

type state int

const (
	stateFadeIn state = iota
	stateRanding
	stateRobotAttack
	stateEnemyAttack
	stateRobotExplode
	stateEnemyExplode
	stateShowResult
	stateFadeOut
)

func (s state) DrawAttack() bool {
	switch s {
	case stateRobotAttack, stateEnemyAttack, stateEnemyExplode:
		return true
	}
	return false
}

func (s state) DrawEnemy() bool {
	switch s {
	case stateEnemyExplode:
		return false
	}
	return true
}
