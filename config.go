package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	Stdout = iota
	File
	All
)

var (
	ModeStringer = []string{"STDOUT", "FILE", "ALL"}
)

type Config struct {
	Mode int

	FileSize int
	FileName string
	Fpath    string

	Encoder Encoder
}

type Encoder struct {
	T, D, I, W, E string
}

func NewConfig(mode int, enc Encoder) Config {
	return Config{
		Mode:    mode,
		Encoder: enc,
	}
}

func Check(c Config) error {
	if len(c.FileName) == 0 {
		return errors.New("log file name needs to be specified")
	}
	if len(c.Fpath) == 0 {
		return errors.New("where the log file is saved?")
	}

	err := os.MkdirAll(c.Fpath, os.ModePerm)
	if err != nil {
		return err
	}

	if c.FileSize <= 0 {
		return errors.New("maximum file size is not specified")
	}
	return nil
}

func Exist(p string) string {
	if len(p) == 0 {
		return "none"
	}
	return p
}

func (c Config) Dissemble() {
	encs := fmt.Sprintf(`
    T: %s
    D: %s
    I: %s
    W: %s
    E: %s
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

	if c.Mode == File || c.Mode == All {
		dis = fmt.Sprintf(`ASLP CONFIG -> {
  mode: %s
  encoder: [%s]
  fpath: %s
  fname: %s
  fsize: %d
}`, ModeStringer[c.Mode],
			encs,
			c.Fpath,
			c.FileName,
			c.FileSize,
		)
	}

	fmt.Println(dis)
}
