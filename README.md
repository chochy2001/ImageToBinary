# English Version
# Video Animation to Binary Converter

This project is a video animation to binary converter developed in Go. It allows you to convert a series of images representing a video animation into a binary representation suitable for use in VHDL with a VGA controller using ROMs.

## Project Structure

The project has the following file structure:

```
.
├── BenchGoroutines/
│   ├── Frames/
│   ├── Output/
│   ├── go.mod
│   ├── go.sum
│   └── main_test.go
├── Frames/
├── Output/
├── animacion.txt.zip
├── concatenacion*
├── concatenacion.go
├── go.mod
├── go.sum
├── image_binary.go
└── README.md
```

## Functionality

The project consists of several components that work together to convert a video animation into a binary representation:

1. `Frames/` directory: Contains a series of images representing the frames of the video animation.
2. `image_binary.go`: Performs the conversion of each image to a binary format suitable for use in VHDL.
3. `BenchGoroutines/main_test.go`: Performs a benchmark to determine the optimal number of goroutines for processing the images.
4. `concatenacion.go`: Concatenates all the generated binary files into a single final file.

## Usage

To use the converter, follow these steps:

1. Place the images of the video animation in the `Frames/` directory.
2. Run the `image_binary.go` file to convert the images to a binary format. This will generate text files in the `Output/` directory.
3. (Optional) If you want to optimize performance, run the benchmark in `BenchGoroutines/main_test.go` to determine the appropriate number of goroutines to use.
4. Run the `concatenacion.go` file to concatenate all the generated binary files into a single final file.
5. The final binary file can be used in VHDL with a VGA controller using ROMs to display the video animation.

## Customization

You can customize the behavior of the converter by modifying the relevant parameters in the `image_binary.go` file:

- `threshold`: Threshold used to determine whether a pixel is considered "0" or "1". Adjust this value according to your needs.
- `inputDir`: Path to the directory containing the input image frames.
- `outputDir`: Path to the directory where the generated binary files will be saved.
- `imgWidth`, `imgHeight`: Desired dimensions of the resized images.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.



# Spanish Version

# Conversor de Animación de Video a Binario

Este proyecto es un conversor de animación de video a binario desarrollado en Go. Permite convertir una serie de imágenes que representan una animación de video en una representación binaria adecuada para su uso en VHDL con un controlador VGA utilizando ROMs.

## Estructura del Proyecto

El proyecto tiene la siguiente estructura de archivos:

```
.
├── BenchGoroutines/
│   ├── Frames/
│   ├── Output/
│   ├── go.mod
│   ├── go.sum
│   └── main_test.go
├── Frames/
├── Output/
├── animacion.txt.zip
├── concatenacion*
├── concatenacion.go
├── go.mod
├── go.sum
├── image_binary.go
└── README.md
```

## Funcionalidad

El proyecto consta de varios componentes que trabajan juntos para convertir una animación de video en una representación binaria:

1. Directorio `Frames/`: Contiene una serie de imágenes que representan los frames de la animación de video.
2. `image_binary.go`: Realiza la conversión de cada imagen a un formato binario adecuado para su uso en VHDL.
3. `BenchGoroutines/main_test.go`: Realiza un benchmark para determinar el número óptimo de goroutines para procesar las imágenes.
4. `concatenacion.go`: Concatena todos los archivos binarios generados en un solo archivo final.

## Uso

Para utilizar el conversor, sigue estos pasos:

1. Coloca las imágenes de la animación de video en el directorio `Frames/`.
2. Ejecuta el archivo `image_binary.go` para convertir las imágenes a un formato binario. Esto generará archivos de texto en el directorio `Output/`.
3. (Opcional) Si deseas optimizar el rendimiento, ejecuta el benchmark en `BenchGoroutines/main_test.go` para determinar el número apropiado de goroutines a utilizar.
4. Ejecuta el archivo `concatenacion.go` para concatenar todos los archivos binarios generados en un solo archivo final.
5. El archivo binario final se puede utilizar en VHDL con un controlador VGA utilizando ROMs para mostrar la animación de video.

## Personalización

Puedes personalizar el comportamiento del conversor modificando los parámetros relevantes en el archivo `image_binary.go`:

- `threshold`: Umbral utilizado para determinar si un píxel se considera "0" o "1". Ajusta este valor según tus necesidades.
- `inputDir`: Ruta del directorio que contiene los frames de las imágenes de entrada.
- `outputDir`: Ruta del directorio donde se guardarán los archivos binarios generados.
- `imgWidth`, `imgHeight`: Dimensiones deseadas de las imágenes redimensionadas.

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.
