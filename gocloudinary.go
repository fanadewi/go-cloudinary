package gocloudinary

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type CloudinaryRequest struct {
	File        string
	CloudName   string
	CloudKey    string
	CloudSecret string
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

func (cloudReq *CloudinaryRequest) Upload() *CloudinaryResponse {
	baseUrl := fmt.Sprintf("https://api.cloudinary.com/v1_1/%s/auto/upload", cloudReq.CloudName)

	var param = url.Values{}
	timestamp := time.Now().Unix()
	stringify := []byte(fmt.Sprintf("timestamp=%d%s", timestamp, cloudReq.CloudSecret))
	signature := sha256.Sum256(stringify)

	param.Set("file", cloudReq.File)
	param.Set("api_key", cloudReq.CloudKey)
	param.Set("timestamp", fmt.Sprintf("%d", timestamp))
	param.Set("signature", fmt.Sprintf("%x", signature[:]))

	resp, err := http.PostForm(baseUrl, param)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	cld := &CloudinaryResponse{}
	err = json.NewDecoder(resp.Body).Decode(cld)
	if err != nil {
		panic(err)
	}

	return cld
}
