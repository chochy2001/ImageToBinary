package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	outputFolder := "Output"
	outputFile := "animacion.txt"

	// Obtener la lista de archivos en la carpeta Output
	files, err := ioutil.ReadDir(outputFolder)
	if err != nil {
		fmt.Println("Error al leer la carpeta:", err)
		return
	}

	// Filtrar solo los archivos de texto
	var txtFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			txtFiles = append(txtFiles, file.Name())
		}
	}

	// Ordenar los archivos alfabéticamente
	sort.Strings(txtFiles)

	// Crear el archivo de salida
	animacionFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error al crear el archivo de salida:", err)
		return
	}
	defer animacionFile.Close()

	// Crear un escritor para el archivo de salida
	writer := bufio.NewWriter(animacionFile)

	// Recorrer los archivos en orden
	for _, fileName := range txtFiles {
		filePath := filepath.Join(outputFolder, fileName)

		// Leer el contenido del archivo
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error al leer el archivo %s: %v\n", fileName, err)
			continue
		}

		// Escribir el contenido en el archivo de salida
		_, err = writer.Write(content)
		if err != nil {
			fmt.Printf("Error al escribir en el archivo de salida: %v\n", err)
			return
		}

		// Agregar una nueva línea después de cada archivo
		_, err = writer.WriteString("\n")
		if err != nil {
			fmt.Printf("Error al escribir en el archivo de salida: %v\n", err)
			return
		}
	}

	// Asegurar que todo el contenido se escriba en el archivo de salida
	writer.Flush()

	fmt.Println("Concatenación de archivos completada.")
}
