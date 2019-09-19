package hackutils

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
)

// Base64ToBytes converts []byte with encoded base64 data to a raw []byte.
func Base64ToBytes(b64 []byte) []byte {
	b64buf := bytes.NewBuffer(b64)
	dec := base64.NewDecoder(base64.StdEncoding, b64buf)
	bs, err := ioutil.ReadAll(dec)
	if err != nil {
		log.Fatal(err)
	}
	return bs
}
