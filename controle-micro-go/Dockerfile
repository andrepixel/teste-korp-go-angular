# Usa a imagem oficial do Golang
FROM golang:1.23 as builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos do projeto para o container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila o binário do projeto
RUN go build -o cmd/main .

# Usa uma imagem mais leve para rodar o app
FROM alpine:latest

WORKDIR /root/

# Copia o binário da fase anterior
COPY --from=builder /app/main .

# Define a porta que o container vai expor
EXPOSE 3000

# Comando para rodar o binário
CMD ["./main"]
