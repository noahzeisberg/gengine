package gengine

var (
	app_name    string
	commandline string
	commands    []Command
	main_color  string
	bg_color    string
)

type Gengine struct {
	Name        string
	CommandLine string
	Commands    []Command
	MainColor   string
	BgColor     string
}

func Initialize(attributes Gengine) {

}
