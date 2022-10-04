package test

import (
	"bytes"
	"cloud-drive/core/define"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestUploadFileByPath(t *testing.T) {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。
	// 替换为用户的 region
	u, _ := url.Parse("https://ryan-gee.cos.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 环境变量 SECRETID 表示用户的 SecretId
			SecretID:  "SecretID",
			SecretKey: "SecretKey",
		},
	})

	key := "exampleobject.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "test.jpg", nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestUploadByReader(t *testing.T) {
	u, _ := url.Parse("https://ryan-gee.cos.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 环境变量 SECRETID 表示用户的 SecretId
			SecretID:  "SecretID",
			SecretKey: "SecretKey",
		},
	})

	key := "exampleobject.jpg"
	f, err := os.ReadFile("test.jpg")
	if err != nil {
		return
	}

	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(f), nil,
	)
	if err != nil {
		panic(err)
	}
}

// Initialization of file fragment upload
func TestInitPartUpload(t *testing.T) {
	u, _ := url.Parse(define.CosAddr)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 环境变量 SECRETID 表示用户的 SecretId
			SecretID:  "SecretID",
			SecretKey: "SecretKey",
		},
	})
	key := "star.jpg"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID //1664847705a8799bb1f6afe8b246c3c5b01cf4a7a7871d4ce60874b01e7865e4b5ab9e325d
	fmt.Println(UploadID)
}

// Uploading multipart files
func TestPartUpload(t *testing.T) {
	u, _ := url.Parse(define.CosAddr)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 环境变量 SECRETID 表示用户的 SecretId
			SecretID:  "SecretID",
			SecretKey: "SecretKey",
		},
	})
	// 注意，上传分块的块数最多10000块
	key := "cloud-drive/star.jpg"
	UploadID := "1664847705a8799bb1f6afe8b246c3c5b01cf4a7a7871d4ce60874b01e7865e4b5ab9e325d"
	ff, err := os.ReadFile("0.chunk") // md5 : 33685b32b9cdbe9e75fe78adbf4ccf71
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, bytes.NewReader(ff), nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag")
	t.Log(PartETag)
}

// Complete file fragment upload
func TestPartUploadComplete(t *testing.T) {
	u, _ := url.Parse(define.CosAddr)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "SecretID",
			SecretKey: "SecretKey",
		},
	})
	key := "star.jpg"
	UploadID := "1664847705a8799bb1f6afe8b246c3c5b01cf4a7a7871d4ce60874b01e7865e4b5ab9e325d"
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: "33685b32b9cdbe9e75fe78adbf4ccf71"},
	)
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 2, ETag: "2837c6b49998ad2f25608f914107b2dc"},
	)
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 3, ETag: "5f7dc5f974ad18052d593d31c050be3f"},
	)
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 4, ETag: "701ca17f1acaca0d2dbb11adf0551098"},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}
