package pkg

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"unicode"
)

var (
	jsonPrefix = []byte("{")
	jsonPrefixK = []byte("[")
)

// GetByteStream is read file/url/stdin to []byte stream
func GetByteStream(s string) (bt []byte, err error) {
	if s == "-" {
		// read stdin
		return ReadByteFrom(os.Stdin)
	} else if _, u := isUrl(s); u {
		// read url
		return httpGetBodyByte(s)
	} else if FileExist(s) {
		// read file
		return ioutil.ReadFile(s)
	} else {
		// other case
		return
	}
}

func ReadByteFrom(in io.Reader) (bt []byte, err error) {
	var b bytes.Buffer
	_, err = b.ReadFrom(in)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func HasJSONPrefix(buf []byte) bool {
	return hasPrefix(buf, jsonPrefix) ||hasPrefix(buf, jsonPrefixK)
}
func hasPrefix(buf []byte, prefix []byte) bool {
	trim := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	return bytes.HasPrefix(trim, prefix)
}


func isUrl(u string) (url.URL, bool) {
	if uu, err := url.Parse(u); err == nil && uu != nil && uu.Host != "" {
		return *uu, true
	}
	return url.URL{}, false
}

func httpGetBodyByte(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return ReadByteFrom(resp.Body)
}