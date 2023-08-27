package utils

import "github.com/NoahOnFyre/gengine/logging"

func Catch(thing any, err error) any {
	if err != nil {
		logging.Error("An error occured during the process:", err.Error())
	}
	return thing
}
