# Utiliser une image de base contenant Go
FROM golang:1.22-alpine AS builder

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers nécessaires dans le conteneur
COPY . .

# Compiler l'application Go
RUN go build -o main .

# Utiliser une image de base légère pour exécuter l'application compilée
FROM alpine:latest

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier l'exécutable de l'application depuis le premier conteneur
COPY --from=builder /app/main .

# Exposer le port sur lequel l'application écoute
EXPOSE 8080

# Commande pour démarrer l'application
CMD ["./main"]
