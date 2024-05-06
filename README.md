# English Version
# Image to Binary Converter

This project is an image to binary converter developed in Go. It allows you to convert an image in PNG format to a binary representation based on a specified threshold.

## Project Structure

The project has the following file structure:

```
.
├── .git/
├── .idea/
├── images/
│   ├── imagen1.png
│   ├── imagen2.png
│   └── imagen3.png
│   ├── imagen4.png
│   └── imagen5.png
│   └── imagen6.png
├── go.mod
├── go.sum
├── image_binary
├── image_binary.go
├── imagen1.txt
├── imagen2.txt
└── imagen3.txt
└── imagen4.txt
└── imagen5.txt
└── imagen6.txt
```

## Dependencies

The project uses the following dependencies:

- `image`: Standard Go package for image manipulation.
- `image/color`: Standard Go package for color handling.
- `os`: Standard Go package for file operations.
- `github.com/nfnt/resize`: External library for resizing images.

## Functionality

The program performs the following steps:

1. Opens the specified image file (`./images/imagen3.png`).
2. Decodes the image using the `image` package.
3. Resizes the image to a size of 640x480 pixels using the `resize` library.
4. Creates an output file (`imagen3.txt`) to store the binary representation.
5. Iterates over each pixel of the image and converts its value to grayscale.
6. Compares the pixel value with a predefined threshold (50 in this case).
   - If the value is greater than the threshold, it writes a "0" to the output file.
   - If the value is less than or equal to the threshold, it writes a "1" to the output file.
7. Separates each row of the binary representation with a newline.

## Usage

To use the converter, follow these steps:

1. Clone the repository or download the project files.
2. Make sure you have Go installed on your system.
3. Place the image you want to convert in the `images/` folder with the name `imagen3.png`.
4. Open a terminal and navigate to the project folder.
5. Run the following command to compile the program:
   ```
   go build image_binary.go
   ```
6. Run the generated program:
   ```
   ./image_binary
   ```
7. The program will generate an `imagen3.txt` file with the binary representation of the image.

## Customization

You can customize the behavior of the converter by modifying the following parameters in the `image_binary.go` file:

- `threshold`: Threshold used to determine whether a pixel is considered "0" or "1". Adjust this value according to your needs.
- `"./images/imagen3.png"`: Path and name of the input image file. Make sure the image is in the `images/` folder and has the correct name.
- `"imagen3.txt"`: Name of the output file where the binary representation will be saved. You can change it if desired.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.



# Spanish Version 

# Conversor de Imagen a Binario

Este proyecto es un conversor de imagen a binario desarrollado en Go. Permite convertir una imagen en formato PNG a una representación binaria basada en un umbral determinado.

## Estructura del Proyecto

El proyecto tiene la siguiente estructura de archivos:

```
.
├── .git/
├── .idea/
├── images/
│   ├── imagen1.png
│   ├── imagen2.png
│   └── imagen3.png
│   ├── imagen4.png
│   └── imagen5.png
│   └── imagen6.png
├── go.mod
├── go.sum
├── image_binary
├── image_binary.go
├── imagen1.txt
├── imagen2.txt
└── imagen3.txt
└── imagen4.txt
└── imagen5.txt
└── imagen6.txt
```

## Dependencias

El proyecto utiliza las siguientes dependencias:

- `image`: Paquete estándar de Go para manipulación de imágenes.
- `image/color`: Paquete estándar de Go para manejo de colores.
- `os`: Paquete estándar de Go para operaciones de archivo.
- `github.com/nfnt/resize`: Biblioteca externa para redimensionar imágenes.

## Funcionamiento

El programa realiza los siguientes pasos:

1. Abre el archivo de imagen especificado (`./images/imagen3.png`).
2. Decodifica la imagen utilizando el paquete `image`.
3. Redimensiona la imagen a un tamaño de 640x480 píxeles utilizando la biblioteca `resize`.
4. Crea un archivo de salida (`imagen3.txt`) para almacenar la representación binaria.
5. Recorre cada píxel de la imagen y convierte su valor a escala de grises.
6. Compara el valor del píxel con un umbral predefinido (50 en este caso).
   - Si el valor es mayor que el umbral, se escribe un "0" en el archivo de salida.
   - Si el valor es menor o igual que el umbral, se escribe un "1" en el archivo de salida.
7. Separa cada fila de la representación binaria con un salto de línea.

## Uso

Para utilizar el conversor, sigue estos pasos:

1. Clona el repositorio o descarga los archivos del proyecto.
2. Asegúrate de tener Go instalado en tu sistema.
3. Coloca la imagen que deseas convertir en la carpeta `images/` con el nombre `imagen3.png`.
4. Abre una terminal y navega hasta la carpeta del proyecto.
5. Ejecuta el siguiente comando para compilar el programa:
   ```
   go build image_binary.go
   ```
6. Ejecuta el programa generado:
   ```
   ./image_binary
   ```
7. El programa generará un archivo `imagen3.txt` con la representación binaria de la imagen.

## Personalización

Puedes personalizar el comportamiento del conversor modificando los siguientes parámetros en el archivo `image_binary.go`:

- `threshold`: Umbral utilizado para determinar si un píxel se considera "0" o "1". Ajusta este valor según tus necesidades.
- `"./images/imagen3.png"`: Ruta y nombre del archivo de imagen de entrada. Asegúrate de que la imagen esté en la carpeta `images/` y tenga el nombre correcto.
- `"imagen3.txt"`: Nombre del archivo de salida donde se guardará la representación binaria. Puedes cambiarlo si lo deseas.

## Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo `LICENSE` para más detalles.
