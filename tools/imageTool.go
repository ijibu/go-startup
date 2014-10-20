package main

import (
	"code.google.com/p/graphics-go/graphics"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	reszieImage("348.jpg", 100, 150)
}

//缩放图片，proportion为比例
func reszieImage(imageName string, minWidth float32, minHeight float32) {
	src, err := LoadImage(imageName)
	if err != nil {
		log.Fatal(err)
	}

	info := src.Bounds()
	width := info.Dx()  //图片宽度
	height := info.Dy() //图片高度
	var (
		newWidth, newHeight float32
	)

	if float32(width)/float32(height) > minWidth/minHeight {
		newHeight = minHeight
		newWidth = float32(width) * minHeight / float32(height)
	} else {
		newWidth = minWidth
		newHeight = (float32(height) * minWidth) / float32(width)
	}

	// 缩略图的大小
	dst := image.NewRGBA(image.Rect(0, 0, int(newWidth), int(newHeight)))

	// 产生缩略图,等比例缩放
	err = graphics.Scale(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	// 需要保存的文件
	imgcounter := 734
	saveImage(fmt.Sprintf("%03d.jpg", imgcounter), dst)
}

// LoadImage decodes an image from a file.
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

// 保存Png图片
func saveImage(path string, img image.Image) (err error) {
	// 需要保存的文件
	imgfile, err := os.Create(path)
	defer imgfile.Close()

	// 以jpeg格式保存文件
	err = jpeg.Encode(imgfile, img, nil)
	if err != nil {
		log.Fatal(err)
	}
	return
}

//获取图片信息，主要是高度和宽度
func getImageInfo(path string) (img image.Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.DecodeConfig(file)

	return
}
