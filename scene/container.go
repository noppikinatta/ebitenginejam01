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

package scene

import "errors"

type Container struct {
	scenes      []Scene
	current     int
	transitions map[int]int
}

func NewContainer(scenes []Scene) *Container {
	c := Container{
		scenes:      scenes,
		current:     0,
		transitions: make(map[int]int),
	}
	return &c
}

func (c *Container) AddTransition(from, to Scene) error {
	fi, ok := c.indexOf(from)
	if !ok {
		return errors.New("from scene for transition is not in container")
	}

	ti, ok := c.indexOf(to)
	if !ok {
		return errors.New("to scene for transition is not in container")
	}

	c.transitions[fi] = ti
	return nil
}

func (c *Container) indexOf(s Scene) (int, bool) {
	for i := range c.scenes {
		if s == c.scenes[i] {
			return i, true
		}
	}

	return 0, false
}

func (c *Container) Current() Scene {
	s := c.scenes[c.current]
	if s.End() {
		c.current = c.next()
		s = c.scenes[c.current]
		s.Reset()
	}
	return s
}

func (c *Container) next() int {
	idx, ok := c.transitions[c.current]
	if ok {
		return idx
	}

	idx = c.current + 1
	idx = idx % len(c.scenes)
	return idx
}
