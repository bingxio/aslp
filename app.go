// Copyright 2021 bingxio. All rights reserved.
//
// Gnu Public License v3
// license that can be found in the LICENSE file.
//

// ASLP
// A Go language based log library, simple, convenient and concise.

package main

func main() {
	// 1. create a configuration structure.
	conf := NewConfig(
		Both,
		Encoder{
			T: "@{2006/01/02 15:04:05} [T]: @N - @M",
			D: "LOG: @{2006-01-02} [DEBUG] @N - @M",
			I: "<INFO> --> @N @M",
			W: "",
			E: "@{15:04:05} [ERROR](@N): @M",
		},
	)

	conf.FileName = "2006-01-02 15:04"
	conf.FileSize = 3
	conf.Fpath = "./logs"

	// 2. create a log structure.
	l, err := NewLog(&conf)
	if err != nil {
		panic(err)
	}

	conf.Dissemble()

	// 3. five modes of using logs.
	l.T("TEST", "my log")
	l.D("TITLE", "my log message")
	l.I("WHAT", "it's good")
	l.E("LOGIN", "database connection failed")
}
