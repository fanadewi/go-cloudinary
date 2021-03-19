package gocloudinary

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

type CloudinaryRequest struct {
	File      interface{}
	FileName  string
	Name      string
	Key       string
	Secret    string
	Timestamp int64
	Signature [32]byte
}

type CloudinaryResponse struct {
	AssetID          string        `json:"asset_id"`
	Bytes            int           `json:"bytes"`
	CreatedAt        time.Time     `json:"created_at"`
	Etag             string        `json:"etag"`
	Format           string        `json:"format"`
	Height           int           `json:"height"`
	OriginalFilename string        `json:"original_filename"`
	Pages            int           `json:"pages"`
	Placeholder      bool          `json:"placeholder"`
	PublicID         string        `json:"public_id"`
	ResourceType     string        `json:"resource_type"`
	SecureURL        string        `json:"secure_url"`
	Signature        string        `json:"signature"`
	Tags             []interface{} `json:"tags"`
	Type             string        `json:"type"`
	URL              string        `json:"url"`
	Version          int           `json:"version"`
	VersionID        string        `json:"version_id"`
	Width            int           `json:"width"`
}

func (cloud *CloudinaryRequest) Upload() (*CloudinaryResponse, error) {

	baseUrl := fmt.Sprintf("https://api.cloudinary.com/v1_1/%s/auto/upload", cloud.Name)

	cloud.Timestamp = time.Now().Unix()
	cloud.Signature = sha256.Sum256([]byte(fmt.Sprintf("timestamp=%d%s", cloud.Timestamp, cloud.Secret)))

	byteFile, isFileByte := isByte(cloud.File)
	stringFile, isFileString := isString(cloud.File)

	if isFileByte {
		return multiUpload(cloud, byteFile, baseUrl)
	} else if isFileString {
		return urlEncodedUpload(cloud, stringFile, baseUrl)
	} else {
		return nil, fmt.Errorf("File unknown")
	}
}

func isByte(file interface{}) ([]byte, bool) {
	theFile, ok := file.([]byte)
	if !ok {
		return nil, false
	}
	return theFile, true
}

func isSupportedFile(fileName string) bool {
	mediaType := mime.TypeByExtension(filepath.Ext(fileName))
	if !(strings.Contains(mediaType, "image") || strings.Contains(mediaType, "pdf")) {
		return false
	}
	return true
}

func isString(file interface{}) (string, bool) {
	theFile, ok := file.(string)
	if !ok {
		return "", false
	}
	return theFile, true
}

func urlEncodedUpload(cloud *CloudinaryRequest, stringFile string, baseUrl string) (*CloudinaryResponse, error) {
	var param = url.Values{}
	param.Set("file", stringFile)
	param.Set("api_key", cloud.Key)
	param.Set("timestamp", fmt.Sprintf("%d", cloud.Timestamp))
	param.Set("signature", fmt.Sprintf("%x", cloud.Signature[:]))
	response, err := http.PostForm(baseUrl, param)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	cloudinaryResponse := &CloudinaryResponse{}
	err = json.NewDecoder(response.Body).Decode(cloudinaryResponse)
	if err != nil {
		return nil, err
	}

	return cloudinaryResponse, nil
}

func multiUpload(cloud *CloudinaryRequest, byteFile []byte, baseUrl string) (*CloudinaryResponse, error) {
	if !isSupportedFile(cloud.FileName) {
		return nil, fmt.Errorf("Only PDF and images supported")
	}

	extraParams := map[string]string{
		"api_key":   cloud.Key,
		"timestamp": fmt.Sprintf("%d", cloud.Timestamp),
		"signature": fmt.Sprintf("%x", cloud.Signature[:]),
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", cloud.FileName)
	if err != nil {
		return nil, err
	}
	part.Write(byteFile)
	for key, val := range extraParams {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", baseUrl, body)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	cloudinaryResponse := &CloudinaryResponse{}
	err = json.NewDecoder(response.Body).Decode(cloudinaryResponse)
	if err != nil {
		return nil, err
	}

	return cloudinaryResponse, nil
}
