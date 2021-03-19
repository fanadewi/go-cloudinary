package examples

import (
	"log"

	goCld "github.com/fanadewi/go-cloudinary"
)

func UploadByte() {
	byteData := []byte("it's a byte")
	e := goCld.CloudinaryRequest{
		File:     byteData,
		FileName: "example.jpg",
		Name:     "cloudinaryName",
		Key:      "CloudinaryKey",
		Secret:   "CloudinarySecret",
	}
	response, err := e.Upload()
	if err != nil {
		panic(err)
	}

	log.Println(response)
}
