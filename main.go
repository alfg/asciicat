package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"reflect"
)

var (
	file  = flag.String("i", "gopher.png", "Use -i <filename>")
	width = flag.Int("w", 80, "Use -w <width>")
)

func main() {
	const ASCIISTR = " `'.,:;i+o*%&$#@"

	flag.Parse()

	// Open file into buffer.
	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	_ = f.Close()

	// Scale Image.
	sz := img.Bounds()
	height := (sz.Max.Y * *width * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(*width), uint(height), img, resize.Lanczos3)

	// Convert to ASCII
	table := []byte(ASCIISTR)
	buf := new(bytes.Buffer)

	for i := 0; i < height; i++ {
		for j := 0; j < *width; j++ {

			// Get grayscale intensity.
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()

			// Write to ascii map.
			pos := int(y * uint64(len(ASCIISTR)) / 256)
			_ = buf.WriteByte(table[pos])
		}
		_ = buf.WriteByte('\n')
	}

	// Write to stdout.
	fmt.Println(string(buf.Bytes()))

}
