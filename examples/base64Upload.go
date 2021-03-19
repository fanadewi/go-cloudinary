package examples

import (
	"fmt"

	goCld "github.com/fanadewi/go-cloudinary"
)

func UploadBase64() {
	e := goCld.CloudinaryRequest{
		File:   "data:image/jpg;base64,frgcsgcrclfhfrdgvh",
		Name:   "cloudinaryName",
		Key:    "CloudinaryKey",
		Secret: "CloudinarySecret",
	}
	response, err := e.Upload()
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
