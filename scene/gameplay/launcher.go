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
