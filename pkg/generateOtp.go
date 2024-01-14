package pkg

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/pquerna/otp/totp"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GenerateOtp() (string, string) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@example.com",
	})
	if err != nil {
		fmt.Println("Error after totp.Generate")
	}
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	png.Encode(&buf, img)
	os.WriteFile("qr-code.png", buf.Bytes(), 0644)
	bytes, err := ioutil.ReadFile("qr-code.png")
	if err != nil {
		log.Fatal(err)
	}
	mimeType := http.DetectContentType(bytes)
	var base64Encoding string
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}
	base64Encoding += toBase64(bytes)
	log.Println(key.Secret(), "\n"+base64Encoding)
	return key.Secret(), base64Encoding
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
