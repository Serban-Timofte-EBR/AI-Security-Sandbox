package models

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
	"gocv.io/x/gocv"
)

func LoadAndPreprocessImage(path string) ([]float32, error)  {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}

	resized := resized.Resize(224, 224, img, resize.Bilinear)

	bounds := resized.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	data := make([]float32, 0, width*height*3)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := resized.At(x, y).RGBA()
			index := y*width + x
			data[index] = float32(r>>8) / 255.0
			data[width*height+index] = float32(g>>8) / 255.0
			data[2*width*height+index] = float32(b>>8) / 255.0
		}
	}

	return data, nil
}
