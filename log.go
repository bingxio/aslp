package main

import (
	"fmt"
	"time"
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

	if c.Mode == File || c.Mode == Both {
		Check(*c)
	}
	return &Log{}, nil
}

func (l Log) T(comp, msg string) {
	Process(c.Encoder.T, comp, msg)
}

func (l Log) D(comp, msg string) {
	Process(c.Encoder.D, comp, msg)
}

func (l Log) I(comp, msg string) {
	Process(c.Encoder.I, comp, msg)
}

func (l Log) W(comp, msg string) {
	Process(c.Encoder.W, comp, msg)
}

func (l Log) E(comp, msg string) {
	Process(c.Encoder.E, comp, msg)
}

func Process(lm, comp, msg string) {
	p := 0
	l := len(lm)

	if l == 0 {
		panic("log type is not defined")
	}

	content := ""

	for p < l {
		n := lm[p]

		if n == '@' {
			x := p + 1

			if x < l && lm[x] == '{' {
				p = x + 1
				n = lm[p]

				var lit string

				for n != '}' {
					lit += string(n)
					p++
					n = lm[p]
				}

				p++
				n = lm[p]

				content += time.Now().Format(lit)
			} else {
				lit := lm[p+1]

				if lit == 'N' {
					content += comp
				}
				if lit == 'M' {
					content += msg
				}

				p += 2
				n = lm[p]
			}
		}
		content += string(n)
		p++
	}

	fmt.Println(content)
}
