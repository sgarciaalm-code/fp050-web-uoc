# Utilizamos una imagen de Linux con Go instalado
FROM golang:alpine

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos todos los archivos del proyecto
COPY . .

# Compilamos nuestra aplicación Go
RUN go build -o servidor .

# Indicamos el puerto que utiliza la aplicación
EXPOSE 8080

# Ejecutamos nuestro servidor web
CMD ["./servidor"]