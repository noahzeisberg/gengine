package requests

import (
	"io"
	"net/http"
	"os"

	"github.com/NoahOnFyre/gengine/logging"
)

func Get(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		logging.Error("Failed to make GET request:", err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		logging.Error("Failed to read body contents:", err.Error())
		os.Exit(1)
	}

	return content
}
