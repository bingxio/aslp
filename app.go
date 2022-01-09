package main

func main() {
	conf := NewConfig(
		Both,
		Encoder{
			T: "@{2006/01/02 15:04:05} [T]: @N - @M.",
			D: "@{2006-01-02} [DEBUG] @N - @M.",
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

	l.T("WHAT", "lfmgl;fdgopw")
	l.D("WHAT", "*&$*(@#$wqjeopw")
	l.T("WHAT", "13123213qweowqjeopw")
	l.D("WHAT", "58768eopw")
	l.T("WHAT", "vngheteeopw")
	l.D("WHAT", "sdfweowqjeopw")
	l.T("WHAT", "sgfdw")
	l.T("WHAT", "657635435")
}
