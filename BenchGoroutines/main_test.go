package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"testing"

	"github.com/nfnt/resize"
)

const inputDir = "./Frames"
const outputDir = "./Output"
const combinedFile = "output.txt"
const threshold = 50
const maxGoroutines = 100 // Número máximo de goroutines simultáneas
const imgWidth, imgHeight = 640, 480

func processImage(index int, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println("Procesando frame: ", index)

	filename := fmt.Sprintf("badApple%05d.png", index)
	// fmt.Println("Procesando frame: ", filename)

	// Abre la imagen
	file, err := os.Open(filepath.Join(inputDir, filename))
	// fmt.Println("Se abrio la imagen: ", filename)
	if err != nil {
		fmt.Printf("Error al abrir el archivo %s: %v\n", filename, err)
		return
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
		fmt.Printf("Error al decodificar el archivo %s: %v\n", filename, err)
		return
	}
	// fmt.Println("Se decodifico la imagen: ", filename)

	// Redimensiona la imagen
	img = resize.Resize(imgWidth, imgHeight, img, resize.Lanczos3)
	// fmt.Println("Se redimensiono la imagen: ", filename)

	// Escribe el encabezado para cada frame en un archivo independiente
	txtFilename := fmt.Sprintf("badApple%05d.txt", index)
	outputFile, err := os.Create(filepath.Join(outputDir, txtFilename))
	if err != nil {
		fmt.Printf("Error al crear el archivo %s: %v\n", txtFilename, err)
		return
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			panic(err)
		}
	}(outputFile)

	_, err1 := fmt.Fprintf(outputFile, "CONSTANT ROM%d: MEMORY%d := (", index, index)
	if err1 != nil {
		return
	}
	for y := 0; y < img.Bounds().Dy(); y++ {
		_, err := outputFile.WriteString("\"")
		if err != nil {
			return
		}
		for x := 0; x < img.Bounds().Dx(); x++ {
			grayColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			if grayColor.Y > uint8(threshold) {
				_, err := outputFile.WriteString("0")
				if err != nil {
					return
				}
			} else {
				_, err := outputFile.WriteString("1")
				if err != nil {
					return
				}
			}
		}
		_, err2 := outputFile.WriteString("\"")
		if err2 != nil {
			return
		}
		if y != img.Bounds().Dy()-1 {
			_, err := outputFile.WriteString(",\n")
			if err != nil {
				return
			}
		}
	}
	_, err3 := outputFile.WriteString(");\n")
	if err3 != nil {
		return
	}
	// fmt.Printf("Procesado: %s\n", filename)
	// fmt.Println("Se escribio el archivo: ", txtFilename)
}

func combineFiles() {
	output, err := os.Create(combinedFile)
	if err != nil {
		panic(err)
	}
	defer func(output *os.File) {
		err := output.Close()
		if err != nil {
			panic(err)
		}
	}(output)

	files, err := filepath.Glob(filepath.Join(outputDir, "badApple*.txt"))
	if err != nil {
		panic(err)
	}
	sort.Strings(files)

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error al leer el archivo %s: %v\n", file, err)
			continue
		}
		_, err = output.Write(data)
		if err != nil {
			fmt.Printf("Error al escribir en el archivo combinado: %v\n", err)
			continue
		}
	}
}

func BenchmarkGoroutines(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("Goroutines-%d", i), func(b *testing.B) {
			var wg sync.WaitGroup
			semaphore := make(chan struct{}, i)

			err := os.MkdirAll(outputDir, os.ModePerm)
			if err != nil {
				b.Fatal(err)
			}

			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				for k := 1; k <= 100; k++ {
					wg.Add(1)
					semaphore <- struct{}{}
					go func(index int) {
						processImage(index, &wg)
						<-semaphore
					}(k)
				}
				wg.Wait()
			}

			b.StopTimer()
			err = os.RemoveAll(outputDir)
			if err != nil {
				b.Fatal(err)
			}
		})
	}
}
