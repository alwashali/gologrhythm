package gologrhythm

type LogRhtyhm struct {
	Name  string
	Token string
	Host  string
}

func Create_App(name, token, host string) *LogRhtyhm {

	app := LogRhtyhm{
		Name:  name,
		Token: token,
		Host:  host,
	}

	return &app
}
