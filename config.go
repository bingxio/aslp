// Copyright 2021 bingxio. All rights reserved.
//
// Gnu Public License v3
// license that can be found in the LICENSE file.
//

// Configuration structure and tool functions.

package main

import (
	"errors"
	"fmt"
	"os"
)

// Three modes, standard output and file mode, common mode.
const (
	Stdout = iota
	File
	Both
)

var ModeStringer = []string{"STDOUT", "FILE", "BOTH"} // Stringer of modes.

// Configuration.

type Config struct {
	Mode int

	FileSize int // *m
	FileName string
	Fpath    string
	F        *os.File // Log file pointer.

	Encoder Encoder
}

// Customize the contents of the five log modes.
type Encoder struct {
	T, D, I, W, E string
}

// Mode and custom log format.
func NewConfig(mode int, enc Encoder) Config {
	return Config{
		Mode:    mode,
		Encoder: enc,
	}
}

// Check whether the configured member format is correct.
func CheckFiled(c Config) error {
	if len(c.FileName) == 0 {
		return errors.New("log file name needs to be specified")
	}
	if len(c.Fpath) == 0 {
		return errors.New("where the log file is saved?")
	}

	// Create a folder based on the configured directory.
	err := os.MkdirAll(c.Fpath, os.ModePerm)
	if err != nil {
		return err
	}

	if c.FileSize <= 0 {
		return errors.New("maximum file size is not specified")
	}
	return nil
}

// Format output custom log format.
func Exist(p string) string {
	if len(p) == 0 {
		return "none"
	}
	return p
}

// Decorate the structure of the output log.
func (c Config) Dissemble() {
	encs := fmt.Sprintf(`
    T: "%s"
    D: "%s"
    I: "%s"
    W: "%s"
    E: "%s"
  `, Exist(c.Encoder.T),
		Exist(c.Encoder.D),
		Exist(c.Encoder.I),
		Exist(c.Encoder.W),
		Exist(c.Encoder.E),
	)

	var dis string

	if c.Mode == Stdout {
		dis = fmt.Sprintf(`ASLP CONFIG -> {
  mode: %s
  encoder: [%s]
}`, ModeStringer[c.Mode], encs)
	}

	if c.Mode == File || c.Mode == Both {
		dis = fmt.Sprintf(`ASLP CONFIG -> {
  mode: %s
  encoder: [%s]
  fpath: "%s"
  fname: "%s.log"
  fsize: %dm
}`, ModeStringer[c.Mode],
			encs,
			c.Fpath,
			c.FileName,
			c.FileSize,
		)
	}

	fmt.Println(dis)
}
