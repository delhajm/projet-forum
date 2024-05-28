# Utiliser une image de base contenant Go
FROM golang:1.22-alpine

RUN apk update && apk add --no-cache gcc musl-dev
# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers nécessaires dans le conteneur
COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Compiler l'application Go
RUN go build -o /front

# Exposer le port sur lequel l'application écoute
EXPOSE 8080

# Commande pour démarrer l'application
CMD ["/front"]

