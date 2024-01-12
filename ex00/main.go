// godoc -http=:6060
package main

import (
	"convert/pngjpeg"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// this is main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: invaild argument")
		return
	}

	// default format
	inputFormat := flag.String("i", "jpg", "input format")
	outputFormat := flag.String("o", "png", "input format")
	qualityFormat := flag.Int("q", 100, "jpeg quality")

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("error: invaild argument")
		fmt.Println("usage: ./convert -i=inputFormat -o=outputFormat <directory>")
		return
	}

	dir := flag.Arg(0)
	ic := pngjpeg.ImageConverter{Quality: *qualityFormat} // set quality

	err := filepath.Walk(dir, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fbyte, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		ftype := http.DetectContentType(fbyte)

		switch {
		case ftype == "image/png" && *inputFormat == "png" && *outputFormat == "jpg":
			ic.ToJpg(path)
		case ftype == "image/jpeg" && *inputFormat == "jpg" && *outputFormat == "png":
			ic.ToPng(path)
		case ftype != "image/png" && ftype != "image/jpeg":
			fmt.Printf("error: %s is not a valid file\n", path)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("error: %s: no such file or directory", dir)
	}
}
