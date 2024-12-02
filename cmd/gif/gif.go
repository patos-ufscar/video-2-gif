package main

import (
	"fmt"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func VideoToGIF(inputPath, outputPath string, startTime, endTime float64, fps int) error {
	duration := endTime - startTime

	return ffmpeg.Input(inputPath, ffmpeg.KwArgs{
		"ss": fmt.Sprintf("%f", startTime),
	}).
		Output(outputPath, ffmpeg.KwArgs{
			"vf": fmt.Sprintf("fps=%d,scale=320:-1:flags=lanczos", fps),
			"t":  fmt.Sprintf("%f", duration),
		}).
		OverWriteOutput().
		Run()
}
