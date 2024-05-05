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
├── .DS_Store
├── flutter.png
├── go.mod
├── go.sum
├── golang.png
├── image_binary
├── image_binary.go
├── imagen1.txt
├── imagen2.txt
└── imagen3.txt
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
