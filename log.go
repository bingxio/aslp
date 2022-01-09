// Copyright 2021 bingxio. All rights reserved.
//
// Gnu Public License v3
// license that can be found in the LICENSE file.
//

// The main module is used to write files and parse log information.

package main

import (
	"fmt"
	"os"
	"time"
)

const (
	// Type of log.
	T = iota
	D
	I
	W
	E

	megaBytes = 1048576 // Number of megabytes, 1m
)

var c *Config = nil // Global configuration.

// The user needs to give the module name and log content.
type Log struct {
	Comp string
	Msg  string
}

// Generate the log main structure according to the config.
func NewLog(conf *Config) (*Log, error) {
	c = conf

	if c.Mode == File || c.Mode == Both {
		// Check the correctness of configuration.
		CheckFiled(*c)

		f, err := NewLogFile(conf)
		if err != nil {
			return nil, err
		}
		// Currently written file of the configuration.
		c.F = f
	}
	// Empty structure for calling member methods.
	return &Log{}, nil
}

// Create log file according to the settings in the configuration.
// Automatically parse the time format in the configuration.
func NewLogFile(conf *Config) (*os.File, error) {
	// path
	p := c.Fpath + "/" + time.Now().Format(c.FileName)

	// Add suffix as .log file and create it.
	f, err := os.Create(p + ".log")
	if err != nil {
		return nil, err
	}
	return f, nil
}

// LOG: Trace mode.
func (l *Log) T(comp, msg string) {
	l.Comp = comp
	l.Msg = msg
	Process(c.Encoder.T, l)
}

// LOG: Debug mode.
func (l *Log) D(comp, msg string) {
	l.Comp = comp
	l.Msg = msg
	Process(c.Encoder.D, l)
}

// LOG: Info mode.
func (l *Log) I(comp, msg string) {
	l.Comp = comp
	l.Msg = msg
	Process(c.Encoder.I, l)
}

// LOG: Warn mode.
func (l *Log) W(comp, msg string) {
	l.Comp = comp
	l.Msg = msg
	Process(c.Encoder.W, l)
}

// LOG: Error mode.
func (l *Log) E(comp, msg string) {
	l.Comp = comp
	l.Msg = msg
	Process(c.Encoder.E, l)
}

// Process configuration and write file.
func Process(enc string, l *Log) {
	// Parse the overall content of the log according to the parser
	msg := Parse(enc, l)

	if c.Mode == File || c.Mode == Both {
		info, _ := c.F.Stat()
		size := c.FileSize * megaBytes // ? * 1m

		if info.Size() >= int64(size) {
			// Beyond the specified file size, create a new log file.
			f, err := NewLogFile(c)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			c.F = f // Reassign the file pointer.
		}

		// Append write to log file.
		_, err := c.F.WriteString(msg + "\n")
		if err != nil {
			fmt.Println(err.Error())
		}
		if c.Mode != File {
			fmt.Println(msg)
		}
	}
	if c.Mode == Stdout {
		fmt.Println(msg)
	}
}

// Resolve the log structure and splicing module name and content.
func Parse(enc string, g *Log) string {
	p := 0
	l := len(enc)

	if l == 0 {
		return "log type is not defined"
	}

	msg := ""

	for p < l {
		// Specified log variable.
		if enc[p] == '@' {
			p++
			pa := Syntax(&p, l, enc) // Splice the specified content.

			// Formatter for @{} symbol.
			if len(pa) > 1 {
				msg += time.Now().Format(pa)
			}
			if pa == "N" {
				msg += g.Comp
			}
			if pa == "M" {
				msg += g.Msg
			}
		} else {
			// Others.
			msg += string(enc[p])
		}
		p++
	}
	return msg
}

// Resolve variable name.
func Syntax(p *int, l int, enc string) string {
	lit := ""
	now := enc[*p]

	*p++

	// @{2006-01-02 15:04:05}
	for now == '{' && enc[*p] != '}' {
		lit += string(enc[*p])
		*p++
	}
	if lit != "" {
		return lit
	}

	// @N | @M
	*p--
	return string(now)
}
