package api

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type Post struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func serveRegisterForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "web/register.html")
		return
	}
	http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	jsonResponse, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	/* trunk-ignore(golangci-lint/errcheck) */
	w.Write(jsonResponse)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "web/login.html")
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "L'email et le mot de passe sont requis", http.StatusBadRequest)
		return
	}

	var user User
	err := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/home_connected?id="+strconv.Itoa(user.ID), http.StatusSeeOther)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if name == "" || email == "" || password == "" {
		http.Error(w, "Le nom, l'email et le mot de passe sont requis", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("INSERT INTO users(name, email, password) VALUES(?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email, string(hashedPassword))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func profile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID utilisateur requis", http.StatusBadRequest)
		return
	}

	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	tmpl, err := template.ParseFiles("web/profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if id == "" || name == "" || email == "" {
		http.Error(w, "L'ID, le nom et l'email sont requis", http.StatusBadRequest)
		return
	}

	var stmt *sql.Stmt
	var err error
	if password != "" {
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if hashErr != nil {
			http.Error(w, hashErr.Error(), http.StatusInternalServerError)
			return
		}
		stmt, err = db.Prepare("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = stmt.Exec(name, email, string(hashedPassword), id)
	} else {
		stmt, err = db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = stmt.Exec(name, email, id)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home_connected?id="+id, http.StatusSeeOther)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "ID utilisateur requis", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func homeConnected(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	tmpl, err := template.ParseFiles("web/home_connected.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	userID := r.FormValue("user_id")
	content := r.FormValue("content")

	if userID == "" || content == "" {
		http.Error(w, "L'ID utilisateur et le contenu sont requis", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO posts(user_id, content) VALUES(?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home_connected?id="+userID, http.StatusSeeOther)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	postID := r.FormValue("post_id")
	userID := r.FormValue("user_id")
	content := r.FormValue("content")

	if postID == "" || userID == "" || content == "" {
		http.Error(w, "L'ID du post, l'ID utilisateur et le contenu sont requis", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("UPDATE posts SET content = ? WHERE id = ? AND user_id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(content, postID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home_connected?id="+userID, http.StatusSeeOther)
}
func deletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	postID := r.FormValue("post_id")
	userID := r.FormValue("user_id")

	if postID == "" || userID == "" {
		http.Error(w, "L'ID du post et l'ID utilisateur sont requis", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ? AND user_id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(postID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home_connected?id="+userID, http.StatusSeeOther)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
			SELECT posts.id, posts.user_id, users.name, posts.content, posts.created_at
			FROM posts
			JOIN users ON posts.user_id = users.id
			ORDER BY posts.created_at DESC
		`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.UserName, &post.Content, &post.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	jsonResponse, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
