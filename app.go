package main

func main() {
	conf := NewConfig(
		All,
		Encoder{
			T: "@{2006/01/02 15:04:05} [T] @N - @M.",
			D: "[DEBUG] @N - @M.",
		},
	)

	conf.FileName = "2006-01-02 15:04"
	conf.FileSize = 1
	conf.Fpath = "./logs"

	l, err := NewLog(&conf)
	if err != nil {
		panic(err)
	}

	conf.Dissemble()

	l.W("WHAT", "qweowqjeopw")
}
