package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg" // Añade esta línea para registrar el decodificador PNJGE
	_ "image/png"  // Añade esta línea para registrar el decodificador JPGN
	"os"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	// Número total de imágenes a procesar
	totalImages := 10 // Ajusta este número según la cantidad de imágenes

	// Genera el nombre del archivo de salida único
	outputFileName := "imagenes.vhd"
	outFile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = outFile.WriteString(
		"LIBRARY IEEE;\n" +
			"USE IEEE.STD_LOGIC_1164.ALL;\n" +
			"USE IEEE.STD_LOGIC_UNSIGNED.ALL;\n" +
			"\n" +
			"ENTITY ima IS\n" +
			"PORT(CLK, RST2, VIDEO: IN STD_LOGIC; -- RELOJ DE 25 MHz (PLL)\n" +
			"	HPOS: IN INTEGER RANGE 0 TO 799;\n" +
			"	VPOS: IN INTEGER RANGE 0 TO 524;-- SALIDAS V y H\n" +
			"	R,G,B: OUT STD_LOGIC_VECTOR(3 DOWNTO 0)); -- SALIDA RGB\n" +
			"END ima;\n" +
			"\n" +
			"ARCHITECTURE BEAS OF ima IS\n")
	if err != nil {
		panic(err)
	}

	//todo: añadir la memoria del laberinto

	for i := 1; i <= totalImages; i++ {
		_, err = outFile.WriteString(fmt.Sprintf("TYPE MEMORY%d IS ARRAY(0 TO 479) OF STD_LOGIC_VECTOR(639 DOWNTO 0);\n", i))
		if err != nil {
			panic(err)
		}
	}

	for i := 1; i <= totalImages; i++ {
		// Escribe la declaración inicial
		memoryName := fmt.Sprintf("CONSTANT ROM%d: MEMORY%d := (\n", i, i)
		_, err = outFile.WriteString(memoryName)
		if err != nil {
			panic(err)
		}

		// Genera el nombre del archivo de entrada

		//todo cambiar el nombre de la imagen
		//inputFileName := fmt.Sprintf("laberinto_imagen.jpg")
		inputFileName := fmt.Sprintf("./imagenes_video/videoanimacionfinal_000/videoanimacionfinal_%03d.jpg", i)

		file, err := os.Open(inputFileName)
		if err != nil {
			panic(err)
		}

		// Decodifica la imagen
		img, _, err := image.Decode(file)
		if err != nil {
			panic(err)
		}
		file.Close()

		// Redimensiona la imagen a 640x480
		img = resize.Resize(640, 480, img, resize.Lanczos3)

		// Convierte la imagen a binario con umbral
		threshold := 190
		var outputLines []string
		for y := 0; y < img.Bounds().Dy(); y++ {
			line := "\""
			for x := 0; x < img.Bounds().Dx(); x++ {
				// Obtiene el valor del pixel en escala de grises
				grayColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
				if grayColor.Y > uint8(threshold) {
					line += "0"
				} else {
					line += "1"
				}
			}
			line += "\","
			outputLines = append(outputLines, line)
		}

		// Elimina la última coma y escribe el contenido al archivo
		if len(outputLines) > 0 {
			outputLines[len(outputLines)-1] = strings.TrimSuffix(outputLines[len(outputLines)-1], ",")
		}
		for _, line := range outputLines {
			_, err := outFile.WriteString(line + "\n")
			if err != nil {
				panic(err)
			}
		}

		// Escribe el cierre del archivo
		_, err = outFile.WriteString(");\n")
		if err != nil {
			panic(err)
		}
		println("Procesando imagen", i)
	}
	_, err = outFile.WriteString(
		fmt.Sprintf(
			"SIGNAL SAL: STD_LOGIC_VECTOR(639 DOWNTO 0);\n"+
				"SIGNAL PIX: STD_LOGIC;\n"+
				"SIGNAL COLOR: INTEGER RANGE 0 TO 2;\n"+
				"SIGNAL CLK2: STD_LOGIC;\n"+
				"SIGNAL CONTA: INTEGER RANGE 0 TO %d;\n"+ // Cambiar a el valor de totalImages
				"SIGNAL AUX: INTEGER RANGE 0 TO 1249999;\n"+
				"BEGIN\n"+
				"\n"+
				"    PROCESS(CLK, RST2, HPOS, VPOS, VIDEO) -- PROCESO PARA COLOREAR LOS PIXELES\n"+
				"    BEGIN\n"+
				"        IF RST2 = '1' THEN\n"+
				"            COLOR <= 2;\n"+
				"        ELSIF FALLING_EDGE(CLK) THEN\n"+
				"            IF VIDEO = '1' THEN\n"+
				"                IF (HPOS >= 0 AND HPOS <= 639) AND (VPOS >= 0 AND VPOS <= 479) THEN\n", totalImages))
	if err != nil {
		panic(err)
	}

	_, err = outFile.WriteString("IF CONTA = 0 THEN\n\tSAL <= ROM1(VPOS);\n\tPIX <= SAL(HPOS);\n")
	if err != nil {
		panic(err)
	}
	for i := 1; i <= totalImages-1; i++ {
		// Escribe la declaración inicial
		_, err = outFile.WriteString(fmt.Sprintf("ELSIF CONTA = %d THEN\nSAL <= ROM%d(VPOS);\nPIX <= SAL(HPOS);\n", i, i+1))
		if err != nil {
			panic(err)
		}
	}

	_, err = outFile.WriteString("END IF;")
	if err != nil {
		panic(err)
	}

	_, err = outFile.WriteString(
		fmt.Sprintf(
			"			IF PIX = '0' THEN\n"+
				"					COLOR <= 0;\n"+
				"				ELSE\n"+
				"					COLOR <= 1;\n"+
				"				END IF;\n"+
				"                ELSE\n"+
				"                    COLOR <= 2;\n"+
				"                END IF;\n"+
				"            END IF;\n"+
				"        END IF;\n"+
				"    END PROCESS;\n"+
				"\n"+
				"    PROCESS(CLK)\n"+
				"    BEGIN\n"+
				"        IF RISING_EDGE(CLK) THEN\n"+
				"            IF AUX = 0 THEN\n"+
				"                CLK2 <= NOT CLK2;\n"+
				"                AUX <= 1249999;\n"+
				"            ELSE\n"+
				"                AUX <= AUX - 1;\n"+
				"            END IF;\n"+
				"        END IF;\n"+
				"    END PROCESS;\n"+
				"\n"+
				"    PROCESS(CLK2)\n"+
				"    BEGIN\n"+
				"        IF FALLING_EDGE(CLK2) THEN\n"+
				"            IF CONTA = %d THEN\n"+ // Cambiar a el valor de totalImages
				"                CONTA <= 0;\n"+
				"            ELSE\n"+
				"                CONTA <= CONTA + 1;\n"+
				"            END IF;\n"+
				"        END IF;\n"+
				"    END PROCESS;\n"+
				"\n"+
				"    WITH COLOR SELECT\n"+
				"        R <= \"0001\" WHEN 0,\n"+
				"              \"0011\" WHEN 1,\n"+
				"              \"0000\" WHEN OTHERS;\n"+
				"    WITH COLOR SELECT\n"+
				"        B <= \"0001\" WHEN 0,\n"+
				"              \"0011\" WHEN 1,\n"+
				"              \"0000\" WHEN OTHERS;\n"+
				"    WITH COLOR SELECT\n"+
				"        G <= \"0001\" WHEN 0,\n"+
				"              \"0011\" WHEN 1,\n"+
				"              \"0000\" WHEN OTHERS;\n"+
				"\n"+
				"END BEAS;\n", totalImages-1))
	if err != nil {
		panic(err)
	}

	println("Proceso completado. Archivo de salida: " + outputFileName)
}
