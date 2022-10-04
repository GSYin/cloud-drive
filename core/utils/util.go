package utils

import (
	"bytes"
	"cloud-drive/core/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func GenerateToken(id int, identity, username string, second int) (string, error) {
	userClaim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	signedString, err := token.SignedString([]byte(define.GetJetKey()))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func AnalyzeToken(tokenString string) (define.UserClaim, error) {
	var userClaim define.UserClaim
	token, err := jwt.ParseWithClaims(tokenString, &userClaim, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.GetJetKey()), nil
	})
	if err != nil {
		return userClaim, err
	}
	if token.Valid {
		return userClaim, err
	} else {
		return userClaim, errors.New("token is invalid")
	}
}

func MailCodeSend(mail, code string) error {
	e := email.NewEmail()
	e.From = "Cloud Drive <wencungsy@126.com>"
	e.To = []string{mail}
	e.Subject = "Cloud drive verification code"
	e.HTML = []byte("<p>【Cloud Drive】 验证码:</p>  <h2>" + code + "</h2> 您正在进行邮箱验证码登录，切勿将验证码泄露于他人，验证码10分钟内有效。")
	err := e.SendWithTLS("smtp.126.com:465", smtp.PlainAuth("", "wencungsy@126.com", define.GetSentEmailInfo(), "smtp.126.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.126.com"})
	if err != nil {
		return err
	}
	return nil

}

func GenerateRandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func CosUploadFile(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosAddr)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 环境变量 SECRETID 表示用户的 SecretId
			SecretID:  define.GetCosSecretId(),
			SecretKey: define.GetCosSecretKey(),
		},
	})
	file, fileHeader, err := r.FormFile("file")
	key := "cloud-drive/" + GenerateUUID() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}
	return define.CosAddr + "/" + key, nil
}

// CosInitPartUpload  init cos part upload
func CosInitPartUpload(fileext string) (string, string, error) {
	u, _ := url.Parse(define.CosAddr)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.GetCosSecretId(),
			SecretKey: define.GetCosSecretKey(),
		},
	})
	key := "cloud-drive/" + GenerateUUID() + fileext
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}
	return key, v.UploadID, nil
}

// CosPartUpload : Uploading multipart files
func CosPartUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosAddr)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 环境变量 SECRETID 表示用户的 SecretId
			SecretID:  define.GetCosSecretId(),
			SecretKey: define.GetCosSecretKey(),
		},
	})
	key := r.PostForm.Get("key")
	uploadId := r.PostForm.Get("upload_id")
	chunkNumber, err := strconv.Atoi(r.PostForm.Get("chunk_number"))
	if err != nil {
		return "", err
	}
	f, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	if err != nil {
		return "", err
	}

	resp, err := client.Object.UploadPart(
		context.Background(), key, uploadId, chunkNumber, bytes.NewReader(buf.Bytes()), nil,
	)
	if err != nil {
		return "", err
	}
	return strings.Trim(resp.Header.Get("ETag"), "\""), nil
}

// CosPartUploadComplete : Complete file fragment upload
func CosPartUploadComplete(key, uploadId string, cosObject []cos.Object) error {
	u, _ := url.Parse(define.CosAddr)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.GetCosSecretId(),
			SecretKey: define.GetCosSecretKey(),
		},
	})
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cosObject...)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, opt,
	)
	return err
}
