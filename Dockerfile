# Etapa 1: compilación
# Imagen con Go para compilar la aplicación
FROM golang:alpine AS build
WORKDIR /app
# Primero el go.mod y las dependencias, para aprovechar la caché
COPY go.mod ./
RUN go mod download
# Después el resto del proyecto y la compilación
COPY . .
RUN CGO_ENABLED=0 go build -o servidor .

# Etapa 2: imagen final
# Imagen ligera, sin Go
FROM alpine:latest
WORKDIR /app
# Solo el ejecutable, la plantilla y la imagen
COPY --from=build /app/servidor .
COPY templates ./templates
COPY static ./static
# Puerto en el que escucha la aplicación
EXPOSE 8080
# Comando que arranca el servidor
CMD ["./servidor"]