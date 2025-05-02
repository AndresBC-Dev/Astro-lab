# Usa una imagen oficial de Go
FROM golang

# Define el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del proyecto
COPY . .

# Descarga dependencias
RUN go mod download

# Compila el código
RUN go build -o astro-lab

# Expone el puerto donde correrá la API
EXPOSE 8080

# Comando que inicia la API
CMD ["./astro-lab"]
