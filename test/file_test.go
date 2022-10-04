package test

import (
	"crypto/md5"
	"io"
	"math"
	"os"
	"strconv"
	"testing"
)

// 100MB
//const chunkSize = 100 * 1024 * 1024

// 10MB
const chunkSize = 10 * 1024 * 1024

// File slicing
func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("test.jpg")
	if err != nil {
		t.Fatal(err)
	}

	// print file size and chunk size
	t.Log(fileInfo.Size()/1024/1024, "MB")
	chunkNumber := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	t.Log(chunkNumber)

	// open file
	myFile, err := os.OpenFile("test.jpg", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNumber); i++ {
		// Read the start position of the file
		_, err := myFile.Seek(int64(i*chunkSize), 0)
		if err != nil {
			t.Fatal(err)
		}
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		_, err = myFile.Read(b)
		if err != nil {
			t.Fatal(err)
		}
		chunkFile, err := os.OpenFile("./static/"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		_, err = chunkFile.Write(b)
		if err != nil {
			t.Fatal(err)
		}
		err = chunkFile.Close()
		if err != nil {
			t.Fatal(err)
		}
	}
	err = myFile.Close()
	if err != nil {
		t.Fatal(err)
	}
}

// Fragment file merging
func TestMergeChunkFiles(t *testing.T) {
	myFile, err := os.OpenFile("output.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fileInfo, err := os.Stat("test.mp4")
	if err != nil {
		t.Fatal(err)
	}

	// print file size and chunk size
	t.Log(fileInfo.Size()/1024/1024, "MB")
	chunkNumber := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	for i := 0; i < int(chunkNumber); i++ {
		chunkFile, err := os.OpenFile("./static/"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		// 'ioutil.ReadAll' is deprecated
		// ioutil.ReadAll(...) --> io.ReadAll(...)
		all, err := io.ReadAll(chunkFile)
		if err != nil {
			t.Fatal(err)
		}
		_, err = myFile.Write(all)
		if err != nil {
			t.Fatal(err)
		}
		err = chunkFile.Close()
		if err != nil {
			t.Fatal(err)
		}
	}
	err = myFile.Close()
	if err != nil {
		t.Fatal(err)
	}

}

// File consistency check
func TestCheck(t *testing.T) {
	// first file
	file1, err := os.OpenFile("test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	f1, err := io.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}

	// second file
	file2, err := os.OpenFile("output.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	f2, err := io.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%x", md5.Sum(f1))
	t.Logf("%x", md5.Sum(f2))
}
