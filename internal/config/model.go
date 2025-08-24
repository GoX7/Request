package config

type (
	Config struct {
		NET      *Net
		Test     *Test
		Protocol string
		Export   string
		result   string
	}
	Net struct {
		Method  string
		Address string
		Header  string
		Body    string
		Timeout int64
	}
	Test struct {
		Request bool
		Export  bool
	}
)
