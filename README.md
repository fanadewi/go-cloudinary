# [GoCloudinary](https://pkg.go.dev/github.com/fanadewi/go-cloudinary)

<!-- [![Build Status](https://travis-ci.org/fanadewi/go-cloudinary.svg)](https://travis-ci.org/fanadewi/go-cloudinary) -->
<!-- [![codecov](https://codecov.io/gh/fanadewi/go-cloudinary/branch/master/graph/badge.svg)](https://codecov.io/gh/fanadewi/go-cloudinary) -->
[![Go Report Card](https://goreportcard.com/badge/github.com/fanadewi/go-cloudinary)](https://goreportcard.com/report/github.com/fanadewi/go-cloudinary)
[![GoDoc](https://pkg.go.dev/badge/github.com/fanadewi/go-cloudinary?status.svg)](https://pkg.go.dev/github.com/fanadewi/go-cloudinary?tab=doc)
[![Open Source Helpers](https://www.codetriage.com/fanadewi/go-cloudinary/badges/users.svg)](https://www.codetriage.com/fanadewi/go-cloudinary)
[![Release](https://img.shields.io/github/release/fanadewi/go-cloudinary.svg?style=flat-square)](https://github.com/fanadewi/go-cloudinary/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/fanadewi/go-cloudinary)](https://www.tickgit.com/browse?repo=github.com/fanadewi/go-cloudinary)

Upload files to Cloudinary

### Installing

* install go-cloudinary
```bash
$ go get -u github.com/fanadewi/go-cloudinary
```

* Import it in your code:
```go
import goCld "github.com/fanadewi/go-cloudinary"
```

### Quick start
```go
uploader := goCld.CloudinaryRequest{
		File:   	file,
		FileName:   example.jpg,
		Name:   	"cloudinaryName",
		Key:    	"cloudinaryKey",
		Secret: 	"cloudinarySecret",
	}

 response, err := uploader.Upload())
 if err!=nil{
	 panic(err)
 }
 fmt.Println(response)
```

## Authors

* **Fitri Ana Dewi** - [github](https://github.com/fanadewi) - [medium](https://fanadewi.medium.com) - [f.ana.d@outlook.co.id](mailto:f.ana.d@outlook.co.id)
