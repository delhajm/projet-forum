package main

import (
    "log"
    "net/http"
    "path/filepath"
)

func main() {
    // Définir les répertoires des fichiers statiques et HTML
    staticDir := "../static"
    webDir := "../web"

    // Servir les fichiers statiques (CSS, JS, etc.)
    fs := http.FileServer(http.Dir(staticDir))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Servir les fichiers HTML
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var filePath string

        switch r.URL.Path {
        case "/":
            filePath = filepath.Join(webDir, "home.html") // Servir home.html par défaut
        case "/login":
            filePath = filepath.Join(webDir, "login.html")
        case "/register":
            filePath = filepath.Join(webDir, "register.html")
        case "/post":
            filePath = filepath.Join(webDir, "post.html")
        case "/home":
            filePath = filepath.Join(webDir, "home.html")
        default:
            // Si la route n'est pas définie, renvoyer une erreur 404
            http.NotFound(w, r)
            return
        }

        http.ServeFile(w, r, filePath)
    })

    log.Println("Serveur démarré sur le port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
