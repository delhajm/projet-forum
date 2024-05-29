# Utiliser une image de base contenant Go
FROM golang:1.22-alpine

RUN apk update && apk add --no-cache gcc musl-dev
# Définir le répertoire de travail dans le conteneur
WORKDIR /api

# Copier les fichiers nécessaires dans le conteneur
COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Compiler l'application Go
RUN go build -o /go-api

# Exposer le port sur lequel l'application écoute
EXPOSE 8181

# Commande pour démarrer l'application
CMD ["/go-api"]

