// Copyright 2022 noppikinatta
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	// AttackTypeTNT
	AttackTypeSuccess
)

type CombinedResult struct {
	Drawer *Drawer
	types  map[part.PartType]CombinedType
}

func NewCombinedResult() *CombinedResult {
	r := CombinedResult{
		types: make(map[part.PartType]CombinedType),
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
	return r.typeOrDefault(part.PartTypeRightLeg)
}

func (r *CombinedResult) typeOrDefault(t part.PartType) CombinedType {
	ct, ok := r.types[t]
	if !ok {
		return CombinedTypeNone
	}
	return ct
}

func (r *CombinedResult) Combined(t part.PartType) bool {
	return r.typeOrDefault(t) != CombinedTypeNone
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
	if r.RightArm() == CombinedTypeCorrectArm {
		return AttackTypeSuccess
	}
	// if r.LeftArm() == CombinedTypeTNT {
	// 	return AttackTypeTNT
	// }
	// if r.RightArm() == CombinedTypeTNT {
	// 	return AttackTypeTNT
	// }

	return AttackTypeFall
}

func (r *CombinedResult) Reset() {
	r.types = make(map[part.PartType]CombinedType)
	r.Drawer.Reset()
}
