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
	Parts   []updateCombiner
	counter int
	firsts  []func() updateCombiner
	rnd     *rand.Rand
}

func newLauncher() *launcher {
	l := launcher{
		firsts: []func() updateCombiner{
			func() updateCombiner { return newLeftArm() },
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
	// TODO: set first place
	l.Parts = append(l.Parts, p)
}

func (l *launcher) randomPart() updateCombiner {
	if len(l.Parts) < len(l.firsts) {
		// first not random
		return l.firsts[len(l.Parts)]()
	}
	i := l.rnd.Intn(len(l.firsts))
	return l.firsts[i]()
}

func (l *launcher) Remove(p updateCombiner) {
	for i := range l.Parts {
		if l.Parts[i] != p {
			continue
		}

		l.Parts = append(l.Parts[:i], l.Parts[i+1:]...)
		break
	}
}
