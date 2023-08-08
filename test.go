package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/linxGnu/goseaweedfs"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func largeFileUploadTest(filedir string, master string) {
	f, err := os.Create(filedir + "large")
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(randStr(1024 * 1024 * 1024)))
	if err != nil {
		panic(err)
	}

	f.Close()

	start := time.Now()
	sw, _ := goseaweedfs.NewSeaweed(master, []string{}, 8096, &http.Client{Timeout: 5 * time.Minute})
	sw.UploadFile(filedir+"large", "test", "test")
	fmt.Print(1024 / time.Since(start).Seconds())
	fmt.Println(" MiB/s")
	fmt.Println("(Large File Upload Speed)")
}

func smallFileUploadTest(filedir string, master string) {
	// generate small files
	for i := 0; i < 1024; i++ {
		f, err := os.Create(filedir + "small" + fmt.Sprint(i))
		if err != nil {
			panic(err)
		}
		_, err = f.Write([]byte(randStr(1024 * 1024)))
		if err != nil {
			panic(err)
		}
		f.Close()
	}

	// upload the files
	start := time.Now()
	sw, _ := goseaweedfs.NewSeaweed(master, []string{}, 8096, &http.Client{Timeout: 5 * time.Minute})
	for i := 0; i < 1024; i++ {
		sw.UploadFile(filedir+"small"+fmt.Sprint(i), "", "")
	}
	fmt.Print(1024 / time.Since(start).Seconds())
	fmt.Println(" MiB/s")
	fmt.Println("(Large File Upload Speed)")
}

func main() {
	master := flag.String("master", "", "Specify master address")
	filedir := flag.String("f", "", "Specify data directory")
	flag.Parse()

	largeFileUploadTest(*filedir, *master)
	// smallFileUploadTest(*filedir, *master)
}
