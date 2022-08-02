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

package asset

import (
	"bytes"
	_ "embed"
	"errors"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

//go:embed sound/bgm1.wav
var bgm1 []byte

//go:embed sound/bgm2.wav
var bgm2 []byte

//go:embed sound/bgm3.wav
var bgm3 []byte

//go:embed sound/atari_boom.wav
var seFire []byte

//go:embed sound/atari_boom4.wav
var seFly []byte

//go:embed sound/explosion1.mp3
var seExplosion []byte

//go:embed sound/qubodupPunch05.ogg
var seCombined []byte

const sampleRate int = 48000

var context *audio.Context

func init() {
	context = audio.NewContext(sampleRate)
	soundCache = map[Sound]*audio.Player{}
}

type Sound int

const (
	BGM1 Sound = iota
	BGM2
	BGM3
	SEFire
	SEFly
	SEExplosion
	SECombined
)

func LoadSounds() error {
	ss := []struct {
		Resource []byte
		Sound    Sound
		FileType fileType
		Volume   float64
	}{
		{seFire, SEFire, fileTypeWav, 0.8},
		{seFly, SEFly, fileTypeWav, 0.8},
		{seExplosion, SEExplosion, fileTypeMp3, 0.8},
		{seCombined, SECombined, fileTypeOgg, 0.8},
		{bgm1, BGM1, fileTypeWav, 0.08},
		{bgm2, BGM2, fileTypeWav, 0.4},
		{bgm3, BGM3, fileTypeWav, 0.08},
	}

	for _, s := range ss {
		err := load(s.Resource, s.Sound, s.FileType, s.Volume)
		if err != nil {
			return err
		}
	}

	return nil
}

type fileType int

const (
	fileTypeWav fileType = iota
	fileTypeMp3
	fileTypeOgg
)

func load(resource []byte, sound Sound, ftype fileType, vol float64) error {
	var s io.ReadSeeker
	var err error

	switch ftype {
	case fileTypeWav:
		s, err = wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(resource))
		if err != nil {
			return err
		}
	case fileTypeMp3:
		s, err = mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(resource))
		if err != nil {
			return err
		}
	case fileTypeOgg:
		s, err = vorbis.DecodeWithSampleRate(sampleRate, bytes.NewReader(resource))
		if err != nil {
			return err
		}
	default:
		return errors.New("not supported filetype")
	}

	// BGM2 loops
	if sound == BGM2 {
		s = audio.NewInfiniteLoop(s, int64(len(resource)))
	}

	p, err := context.NewPlayer(s)
	if err != nil {
		return err
	}
	p.SetVolume(vol)
	soundCache[sound] = p

	return nil
}

var soundCache map[Sound]*audio.Player

func PlaySound(s Sound) {
	p := soundCache[s]
	err := p.Rewind()
	if err != nil {
		log.Println(err)
	}
	p.Play()
}

func StopSound(s Sound) {
	p := soundCache[s]
	p.Pause()
}
