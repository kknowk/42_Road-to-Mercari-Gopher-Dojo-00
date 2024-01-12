// godoc -http=:6060 
// this package converts png to jpg and jpg to png
package pngjpeg

import (
	"fmt"
	"os"
	"strings"
	"image/jpeg"
	"image/png"
)

// ImageConverter is a struct
type ImageConverter struct {
	Quality int // jpeg quality
}

// convert image to jpg
func (ic *ImageConverter) ToJpg(path string) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error: cannot open file")
		return
	}
	defer file.Close()

	// decode image
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("error: cannot decode image")
		return
	}

	// create new file
	newFile, err := os.Create(strings.Replace(path, ".png", ".jpg", -1))
	if err != nil {
		fmt.Println("error: cannot create file")
		return
	}
	defer newFile.Close()

	// encode image
	err = jpeg.Encode(newFile, img, &jpeg.Options{Quality: ic.Quality})
	if err != nil {
		fmt.Println("error: cannot encode image")
		return
	}
}

// convert image to png
func (ic *ImageConverter) ToPng(path string) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error: cannot open file")
		return
	}
	defer file.Close()

	// decode image
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("error: cannot decode image")
		return
	}

	// create new file
	newFile, err := os.Create(strings.Replace(path, ".jpg", ".png", -1))
	if err != nil {
		fmt.Println("error: cannot create file")
		return
	}
	defer newFile.Close()

	// encode image
	err = png.Encode(newFile, img)
	if err != nil {
		fmt.Println("error: cannot encode image")
		return
	}
}
