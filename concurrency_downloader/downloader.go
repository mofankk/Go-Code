package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

type Downloader struct {
	concurrency int
}

func NewDownloader(concurrency int) *Downloader {
	return &Downloader{concurrency: concurrency}
}

func (d *Downloader) Download(strURL, filename string) error {
	if filename == "" {
		filename = path.Base(strURL)
	}

	resp, err := http.Head(strURL)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK && resp.Header.Get("Accept-Ranges") == "bytes" {
		return d.mutiDownload(strURL, filename, int(resp.ContentLength))
	}

	return d.singleDownload(strURL, filename)
}

func (d *Downloader) mutiDownload(strURL, filename string, contentLen int) error {
	partSize := contentLen / d.concurrency

	// 创建部分文件存放目录
	partDir := d.getPartDir(filename)
	os.Mkdir(partDir, 0777)
	defer os.RemoveAll(partDir)

	var wg sync.WaitGroup
	wg.Add(d.concurrency)

	rangeStart := 0

	for i := 0; i < d.concurrency; i++ {

		// 并发请求
		go func(i, rangeStart int) {
			defer wg.Done()

			rangeEnd := rangeStart + partSize

			// 最后一部分，总长度不能超过ContentLength
			if i == d.concurrency-1 {
				rangeEnd = contentLen
			}

			d.downloadPartial(strURL, filename, rangeStart, rangeEnd, i)
		}(i, rangeStart)

		rangeStart += partSize + 1
	}
	wg.Wait()

	d.merge(filename)

	return nil
}

func (d *Downloader) merge(filename string) error {
	destFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer destFile.Close()

	for i := 0; i < d.concurrency; i++ {
		partFileName := d.getPartFilename(filename, i)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}

		io.Copy(destFile, partFile)
		partFile.Close()
		os.Remove(partFileName)
	}

	return nil
}

func (d *Downloader) singleDownload(strURL, filename string) error {
	return nil
}

func (d *Downloader) downloadPartial(strURL, filename string, rangeStart, rangeEnd, i int) {
	if rangeEnd > rangeStart {
		return
	}

	req, err := http.NewRequest("GET", strURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}

func (d *Downloader) getPartDir(filename string) string {
	return strings.SplitN(filename, ".", 2)[0]
}

func (d *Downloader) getPartFilename(filename string, partNum int) string {
	partDir := d.getPartDir(filename)
	return fmt.Sprintf("%s/%s-%d", partDir, filename, partNum)
}
