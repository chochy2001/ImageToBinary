package main

import (
	"image"
	"image/color"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	// Abre el archivo de imagen
	file, err := os.Open("./images/imagen6.png")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Decodifica la imagen
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// Redimensiona la imagen a 640x480
	img = resize.Resize(640, 480, img, resize.Lanczos3)

	// Crear archivo de salida
	outFile, err := os.Create("imagen6.txt")
	if err != nil {
		panic(err)
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			panic(err)
		}
	}(outFile)

	// Convierte la imagen a binario con umbral
	threshold := 50
	for y := 0; y < img.Bounds().Dy(); y++ {
		_, err := outFile.WriteString("\"")
		if err != nil {
			panic(err)
			return
		}
		for x := 0; x < img.Bounds().Dx(); x++ {
			// Obtiene el valor del pixel en escala de grises
			grayColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			if grayColor.Y > uint8(threshold) {
				_, err := outFile.WriteString("0")
				if err != nil {
					panic(err)
					return
				}
			} else {
				_, err := outFile.WriteString("1")
				if err != nil {
					panic(err)
					return
				}
			}
		}
		_, err1 := outFile.WriteString("\",\n")
		if err1 != nil {
			panic(err1)
			return
		}
	}
}
