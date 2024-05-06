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
	defer file.Close()

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
	defer outFile.Close()

	// Convierte la imagen a binario con umbral
	threshold := 55
	for y := 0; y < img.Bounds().Dy(); y++ {
		outFile.WriteString("\"")
		for x := 0; x < img.Bounds().Dx(); x++ {
			// Obtiene el valor del pixel en escala de grises
			grayColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			if grayColor.Y > uint8(threshold) {
				outFile.WriteString("0")
			} else {
				outFile.WriteString("1")
			}
		}
		outFile.WriteString("\",\n")
	}
}
