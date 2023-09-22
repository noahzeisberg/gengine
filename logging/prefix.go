package logging

import "github.com/NoahOnFyre/gengine/color"

func Prefix(level int) string {
	switch level {
	case 0:
		return color.Black + mainColor + " INFO " + color.Reset + color.Reset + " "
	case 1:
		return color.Black + color.YellowBg + " WARN " + color.Reset + color.Yellow + " "
	case 2:
		return color.Black + color.RedBg + " ERROR " + color.Reset + color.Red + " "
	default:
		return ""
	}
}
