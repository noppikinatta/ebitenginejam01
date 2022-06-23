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

//go:build ignore

package main

import (
	"bytes"
	_ "embed"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed gentemplate.txt
var tmpl string

func main() {
	err := exec()
	if err != nil {
		log.Fatal(err)
	}
}

func exec() error {
	fnames, err := filenames()
	if err != nil {
		return err
	}

	t := template.New("img")
	t, err = t.Parse(tmpl)
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	err = t.Execute(&buf, fnames)
	if err != nil {
		return err
	}

	err = out("imgnames.go", buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func filenames() ([]Filename, error) {
	ret := make([]Filename, 0)

	fn := func(path string, info fs.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".png") {
			return nil
		}
		ret = append(ret, Filename(info.Name()))
		return nil
	}

	err := filepath.Walk(".", fn)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

type Filename string

func (fn Filename) Filename() string {
	return string(fn)
}

func (fn Filename) Scene() string {
	if strings.HasPrefix(string(fn), "t_") {
		return "Title"
	}
	if strings.HasPrefix(string(fn), "p_") {
		return "Prologue"
	}
	if strings.HasPrefix(string(fn), "g_") {
		return "Gameplay"
	}
	if strings.HasPrefix(string(fn), "r_") {
		return "Result"
	}
	return "Other"
}

func (fn Filename) Name() string {
	name := strings.ReplaceAll(string(fn[2:]), ".png", "")
	return strings.ToUpper(name[0:1]) + name[1:]
}

func out(path string, content []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return nil
}
