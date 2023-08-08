package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ginuerzh/weedo"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
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
	client := weedo.NewClient("127.0.0.1:9333", "")
	for i := 0; i < 1024; i++ {
		f, err := os.Open(filedir + "small" + fmt.Sprint(i))
		if err != nil {
			panic(err)
		}
		client.AssignUpload(randStr(4)+"small"+fmt.Sprint(i), "text/plain", f)
		f.Close()
	}
	fmt.Print(1024 / time.Since(start).Seconds())
	fmt.Println(" MiB/s")
	fmt.Println("(Large File Upload Speed)")
}

func main() {
	master := flag.String("master", "", "Specify master address")
	filedir := flag.String("f", "", "Specify data directory")
	op := flag.String("op", "", "Specify operation (s: small upload; l: large upload)")
	flag.Parse()

	if *op == "s" {
		smallFileUploadTest(*filedir, *master)
	}
}
