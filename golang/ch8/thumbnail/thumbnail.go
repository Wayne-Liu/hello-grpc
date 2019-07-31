package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ImageFile(infile string) (string, error)  {
	ext := filepath.Ext(infile)
	outfile := strings.TrimPrefix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)

}

func ImageFile2(outfile, infile string ) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return out.Close()
}

func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

func Image(src image.Image) image.Image {
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128,128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) // portrait
	} else {
		height = int(128 / aspect) // landscape
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// a very crude scaling algorithm
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst

}

func main2() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}

func main()  {
	 filenames := []string{"/Users/liuweipeng/Desktop/tmp/1.jpg","/Users/liuweipeng/Desktop/tmp/2.jpg","/Users/liuweipeng/Desktop/tmp/3.jpg","/Users/liuweipeng/Desktop/tmp/4.jpg","/Users/liuweipeng/Desktop/tmp/5.jpg",}
	 fmt.Println(time.Now().String())
	 makeThumbnails3(filenames)
	 fmt.Println(time.Now().String())
}

func makeThumbnail2(filenames []string)  {
	fmt.Println(len(filenames))
	for _, f := range filenames {
		ImageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}
	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}



