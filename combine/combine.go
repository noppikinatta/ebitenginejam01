package combine

type CombinedType int

const (
	CombinedTypeNone CombinedType = iota
	CombinedTypeCorrectArm
	CombinedTypeInverseArm
	CombinedTypeCorrectLeg
	CombinedTypeInverseLeg
	CombinedTypeTNT
	CombinedTypeEbitenN
	CombinedTypeEbitenS
)

type RandingType int

const (
	RandingTypeFailure RandingType = iota
	RandingTypeSuccess
)

type AttackType int

const (
	AttackTypeFall AttackType = iota
	AttackTypeTNT
	AttackTypeSuccess
)

type CombinedResult struct {
	LeftArm  CombinedType
	RightArm CombinedType
	LeftLeg  CombinedType
	RightLeg CombinedType
}

func (r *CombinedResult) Complete() bool {
	if r.LeftArm == CombinedTypeNone {
		return false
	}
	if r.RightArm == CombinedTypeNone {
		return false
	}
	if r.LeftLeg == CombinedTypeNone {
		return false
	}
	if r.RightLeg == CombinedTypeNone {
		return false
	}
	return true
}

func (r *CombinedResult) Randing() RandingType {
	if r.LeftLeg != CombinedTypeCorrectLeg {
		return RandingTypeFailure
	}
	if r.RightLeg != CombinedTypeCorrectLeg {
		return RandingTypeFailure
	}
	return RandingTypeSuccess
}

func (r *CombinedResult) Attack() AttackType {
	if r.LeftArm == CombinedTypeCorrectArm {
		return AttackTypeSuccess
	}
	if r.LeftArm == CombinedTypeTNT {
		return AttackTypeTNT
	}
	if r.RightArm == CombinedTypeTNT {
		return AttackTypeTNT
	}

	return AttackTypeFall
}

func (r *CombinedResult) Reset() {
	r.LeftArm = CombinedTypeNone
	r.RightArm = CombinedTypeNone
	r.LeftLeg = CombinedTypeNone
	r.RightLeg = CombinedTypeNone
}
