package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func DownloadFile(url string, output string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	fileSize := resp.ContentLength

	buffer := make([]byte, 1024)
	var downloadedBytes int64

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			progress := float64(downloadedBytes) / float64(fileSize) * 100
			fmt.Printf("\rDownloading... %.2f%% complete", progress)

			if downloadedBytes >= fileSize {
				break
			}
		}
		fmt.Println()
	}()

	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			_, err := out.Write(buffer[:n])
			if err != nil {
				return err
			}

			downloadedBytes += int64(n)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	wg.Wait()

	return nil
}
