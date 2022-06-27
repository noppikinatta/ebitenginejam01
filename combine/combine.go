package combine

import "github.com/noppikinatta/ebitenginejam01/part"

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
	Drawer *Drawer
	types  map[part.PartType]CombinedType
}

func NewCombinedResult(drawer *Drawer) *CombinedResult {
	r := CombinedResult{
		Drawer: drawer,
		types:  make(map[part.PartType]CombinedType),
	}
	return &r
}

func (r *CombinedResult) Set(pt part.PartType, ct CombinedType) {
	r.types[pt] = ct
}

func (r *CombinedResult) LeftArm() CombinedType {
	return r.typeOrDefault(part.PartTypeLeftArm)
}

func (r *CombinedResult) RightArm() CombinedType {
	return r.typeOrDefault(part.PartTypeRightArm)
}

func (r *CombinedResult) LeftLeg() CombinedType {
	return r.typeOrDefault(part.PartTypeLeftLeg)
}

func (r *CombinedResult) RightLeg() CombinedType {
	return r.typeOrDefault(part.PartTypeRightArm)
}

func (r *CombinedResult) typeOrDefault(t part.PartType) CombinedType {
	ct, ok := r.types[t]
	if !ok {
		return CombinedTypeNone
	}
	return ct
}

func (r *CombinedResult) Complete() bool {
	if r.LeftArm() == CombinedTypeNone {
		return false
	}
	if r.RightArm() == CombinedTypeNone {
		return false
	}
	if r.LeftLeg() == CombinedTypeNone {
		return false
	}
	if r.RightLeg() == CombinedTypeNone {
		return false
	}
	return true
}

func (r *CombinedResult) Randing() RandingType {
	if r.LeftLeg() != CombinedTypeCorrectLeg {
		return RandingTypeFailure
	}
	if r.RightLeg() != CombinedTypeCorrectLeg {
		return RandingTypeFailure
	}
	return RandingTypeSuccess
}

func (r *CombinedResult) Attack() AttackType {
	if r.LeftArm() == CombinedTypeCorrectArm {
		return AttackTypeSuccess
	}
	if r.LeftArm() == CombinedTypeTNT {
		return AttackTypeTNT
	}
	if r.RightArm() == CombinedTypeTNT {
		return AttackTypeTNT
	}

	return AttackTypeFall
}

func (r *CombinedResult) Reset() {
	r.types = make(map[part.PartType]CombinedType)
}