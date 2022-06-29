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

package magnet

type Location struct {
	X float64
	Y float64
}

func (l Location) Move(v Velocity) Location {
	return Location{
		X: l.X + v.X,
		Y: l.Y + v.Y,
	}
}

type Velocity struct {
	X float64
	Y float64
}

func (v Velocity) Avarage(other Velocity) Velocity {
	return Velocity{
		X: (v.X + other.X) / 2,
		Y: (v.Y + other.Y) / 2,
	}
}

func (v Velocity) Accelerate(p Power) Velocity {
	return Velocity{
		X: v.X + p.X,
		Y: v.Y + p.Y,
	}
}

type Power struct {
	X float64
	Y float64
}
