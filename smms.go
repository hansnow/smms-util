package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// SmmsData is upload result
type SmmsData struct {
	Filename string // 上传的文件名
	Size     int    // 文件大小Byte
	Width    int    // 宽度
	Height   int    // 高度
	Delete   string // 删除地址
	URL      string // 外链地址
}

// SmmsResponse is the response of sm.ms upload api
type SmmsResponse struct {
	Code string   // success | error
	Data SmmsData // upload result
}

func parseJSON(s string) *SmmsResponse {
	var r SmmsResponse
	// r := &SmmsResponse{}
	err := json.Unmarshal([]byte(s), &r)
	if err != nil {
		panic(err)
	}
	return &r
}

func newUploadFileRequest(path string) (*http.Request, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	part, err := w.CreateFormFile("smfile", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, f)
	err = w.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "https://sm.ms/api/upload", body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	return req, err
}

func uploadFile(path string) (result string, err error) {
	req, err := newUploadFileRequest(path)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	t := parseJSON(body.String())
	s := fmt.Sprintf("Filename: %s\nSize    : %d\nW × H   : %d × %d\n"+
		"Delete  : %s\nURL     : %s\n===========",
		t.Data.Filename, t.Data.Size, t.Data.Width, t.Data.Height,
		t.Data.Delete, t.Data.URL)
	return s, nil
}

func hasSuffixs(s string, sfxs []string) bool {
	for _, sfx := range sfxs {
		if strings.HasSuffix(s, sfx) {
			return true
		}
	}
	return false
}

func main() {
	var path = os.Args[1]
	flag.Parse()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	// allowed suffixs
	sfxs := []string{".jpg", ".png"}
	for _, file := range files {
		if n := file.Name(); hasSuffixs(n, sfxs) {
			var fullPath string
			if strings.HasSuffix(path, "/") {
				fullPath = path + file.Name()
			} else {
				fullPath = path + "/" + file.Name()
			}
			result, _ := uploadFile(fullPath)
			fmt.Println(result)
		}
	}
}
