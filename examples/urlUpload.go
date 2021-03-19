package examples

import (
	"fmt"

	goCld "github.com/fanadewi/go-cloudinary"
)

func UploadUrlFile() {
	e := goCld.CloudinaryRequest{
		File:   "https://upload.wikimedia.org/wikipedia/en/a/a9/Example.jpg",
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
