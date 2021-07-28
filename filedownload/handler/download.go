package handler

import (
	"io"
	"log"
	"net/http"
	"os"
)

type DownloadFile struct {}

func (*DownloadFile) Download(w http.ResponseWriter, r *http.Request) {
	filePath := "/Users/cyp/github/goLearnCode/filedownload/file/"
	fileName := "go1.16.4.linux-amd64.tar.gz"

	path := filePath + fileName

	file, err := os.Open(path)
	if err != nil {
		log.Println("%v", err)
		w.Write([]byte("file open error"))
		return
	}
	defer file.Close()

	log.Println("file start download")
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	_, err = io.Copy(w, file)
	log.Println("file end download")
	if err != nil {
		log.Println("%v", err)
		w.Write([]byte("file copy error"))
		return
	}
	log.Println("function end")
}