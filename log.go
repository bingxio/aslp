package main

import (
	"fmt"
)

const (
	T = iota
	D
	I
	W
	E
)

var c *Config = nil

type Log struct{}

func NewLog(conf *Config) (*Log, error) {
	c = conf

	if c.Mode == File || c.Mode == All {
		Check(*c)
	}
	return &Log{}, nil
}

func (l Log) T(comp, msg string) {
	Process(T, comp, msg)
}

func (l Log) D(comp, msg string) {
	Process(D, comp, msg)
}

func (l Log) I(comp, msg string) {
	Process(I, comp, msg)
}

func (l Log) W(comp, msg string) {
	Process(W, comp, msg)
}

func (l Log) E(comp, msg string) {
	Process(E, comp, msg)
}

func Process(lm int, comp, msg string) {
	fmt.Println(c)
}
