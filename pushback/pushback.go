package pushback

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type HandlerConfig struct {
	Path string
}

// Creates a new Pushback handler
func NewHandler(c *HandlerConfig) func(http.ResponseWriter, *http.Request) {

	// pushback receives the push request and writes it into a file
	// according to a mapping provided by a json configuration
	return func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Content-Type")

		if t != "binary/octet-stream" {
			log.Printf("Wrong Content-Type %s", t)
			w.Write([]byte(fmt.Sprintf("%s is not a supported Content-Type", t)))
			return
		}

		// Open test file
		f, err := os.Create(fmt.Sprintf("%s/%s.pushback", c.Path, "test"))

		if err != nil {
			log.Printf("Could not open file %e", err)
		}

		n, err := io.Copy(f, r.Body)
		if err != nil {
			log.Printf("Could only receive %d", n)
		}

		w.Write([]byte("OK"))
	}
}

// Send a file to the pushback server
func PushFile(url, filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	res, err := http.Post(url, "binary/octet-stream", f)
	if err != nil {
		return "", err
	}

	r, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(r), nil
}
