// routes.go

package api

import "net/http"

func SetupRoutes() {
	http.HandleFunc("/register", serveRegisterForm)
	http.HandleFunc("/users/create", createUser)
	http.HandleFunc("/login", login)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/profile/update", updateProfile)
	http.HandleFunc("/profile/delete", deleteUser)
	http.HandleFunc("/home_connected", homeConnected)
	http.HandleFunc("/post/create", createPost)
	http.HandleFunc("/post/update", updatePost)
	http.HandleFunc("/post/delete", deletePost)
	http.HandleFunc("/posts", getPosts)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/home.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
