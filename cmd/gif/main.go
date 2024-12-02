package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	endpoint   = "br-se1.magaluobjects.com"
	accessKey  = os.Getenv("S3_ACCESS_KEY_ID")
	secretKey  = os.Getenv("S3_SECRET_KEY")
	useSSL     = true
	bucketName = "aulao-cloud"
)

func main() {
	// err := VideoToGIF("./input.mp4", "output.gif", 5.0, 10.0, 15)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	mux := http.NewServeMux()
	mux.HandleFunc("POST /convert", ConvertToGIF)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}

}

type GIFRequest struct {
	StartTime float32 `json:"start_time"`
	EndTime   float32 `json:"end_time"`
	FPS       int     `json:"fps"`
}

func ConvertToGIF(w http.ResponseWriter, r *http.Request) {
	// receive query parameters
	startTime := r.URL.Query().Get("start_time")
	endTime := r.URL.Query().Get("end_time")
	fps := r.URL.Query().Get("fps")

	outputPathVideo, err := getFileFromReq(r)
	if err != nil {
		fmt.Printf("error uploading file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// convert all values to int and save on GIFRequest struct
	sTime, err := strconv.ParseFloat(startTime, 32)
	if err != nil {
		fmt.Println(err)
	}

	eTime, err := strconv.ParseFloat(endTime, 32)
	if err != nil {
		fmt.Println(err)
	}
	fpsi, err := strconv.Atoi(fps)
	if err != nil {
		fmt.Println(err)
	}

	gifReq := GIFRequest{
		StartTime: float32(sTime),
		EndTime:   float32(eTime),
		FPS:       fpsi,
	}

	// take the last parameters split by / and change .mp4 to .gif
	outputPath := strings.Split(outputPathVideo, "/")
	outputPath[len(outputPath)-1] = strings.Replace(outputPath[len(outputPath)-1], ".mp4", ".gif", 1)
	path := "./tmp/" + outputPath[len(outputPath)-1]
	fmt.Println("output path", path)

	VideoToGIF(outputPathVideo, path, float64(gifReq.StartTime), float64(gifReq.EndTime), gifReq.FPS)

	path = path[2:]
	// fazer upload do arquivo que esta em 'path' pra obj store da magalu
	err = uploadFile(endpoint, accessKey, secretKey, bucketName, path, path, useSSL)
	if err != nil {
		fmt.Println(err)
	}

	// retornar a url assinada
	signedUrl, err := generatePresignedURL(endpoint, accessKey, secretKey, bucketName, path, 30*time.Minute, useSSL)
	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(signedUrl + "\n"))
	w.WriteHeader(http.StatusOK)
}

func getFileFromReq(r *http.Request) (string, error) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tmpFile, err := os.CreateTemp("./tmp", "video-*.mp4")
	if err != nil {
		return "", err
	}

	defer tmpFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	fmt.Println(tmpFile.Name())
	_, err = tmpFile.Write(fileBytes)
	if err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}
