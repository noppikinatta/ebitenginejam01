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

package gameplay

import (
	"math/rand"
	"time"
)

const (
	maxParts       int = 100
	launchInterval int = 150
)

type launcher struct {
	Parts         []robotPart
	counter       int
	notRndCounter int
	news          []func() robotPart
	rnd           *rand.Rand
}

func newLauncher() *launcher {
	l := launcher{
		news: []func() robotPart{
			newLeftArm,
			newRightArm,
			newEbitenS,
			newLeftLeg,
			newEbitenN,
			newRightLeg,
			//newTNT,
		},
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	return &l
}

func (l *launcher) Update() {
	if l.counter%launchInterval == 0 {
		l.launch()
	}
	l.counter++
}

func (l *launcher) launch() {
	if len(l.Parts) >= maxParts {
		return
	}

	p := l.randomPart()
	l.Parts = append(l.Parts, p)
}

func (l *launcher) randomPart() robotPart {
	if l.notRndCounter < len(l.news) {
		// first not random
		p := l.news[l.notRndCounter]()
		l.notRndCounter++
		return p
	}
	i := l.rnd.Intn(len(l.news))
	return l.news[i]()
}

func (l *launcher) Remove(p robotPart) {
	for i := range l.Parts {
		if l.Parts[i] != p {
			continue
		}

		l.Parts = append(l.Parts[:i], l.Parts[i+1:]...)
		break
	}
}

func (l *launcher) Reset() {
	l.notRndCounter = 0
	l.Parts = make([]robotPart, 0)
}
