package web

import (
	"fmt"
	"log"
	"net/http"
	"templtest/internal/services"
	"templtest/internal/storage/sqlite"
	"templtest/views/pages"
	"time"
)

type Handler struct {
	s *services.Service
}

func NewHandlers() *Handler {
	repo, err := sqlite.New()
	if err != nil {
		log.Fatal(err.Error())
	}
	s := services.New(repo)
	return &Handler{s}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	token, err := getAuthCookie(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	claims, err := VerifyJWT(token)
	user, err := h.s.User(claims.UserEmail)
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	todos, err := h.s.Todos(user.ID)
	if err != nil {
		fmt.Println(err.Error())
	}
	pages.Index(user, todos).Render(r.Context(), w)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("pass")
	name := r.FormValue("name")
	fmt.Println(email, password)
	if email != "" && password != "" {
		user, err := h.s.UserLogin(email, name, password)
		if err == nil {
			if user.Password != password {
				pages.LoginPage("Wrong password").Render(r.Context(), w)
				return
			}
			token, _ := CreateJWT(user.ID, user.Email)
			cookie := createCookie(time.Now().Add(24*60*time.Hour), "Authorization", token)
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
	}
	pages.LoginPage("").Render(r.Context(), w)
}

func (h *Handler) LoginForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("some"))
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	token, err := getAuthCookie(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	title := r.FormValue("title")
	claims, err := VerifyJWT(token)
	todo, err := h.s.CreateTodos(title, claims.UserID)
	if err != nil {
		fmt.Println(err.Error())
	}
	pages.TodoComponent(todo).Render(r.Context(), w)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	token, err := getAuthCookie(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	claims, err := VerifyJWT(token)
	id := r.PathValue("id")
	err = h.s.DeleteTodo(id, claims.UserID)

	w.WriteHeader(http.StatusAccepted)
	todos, err := h.s.Todos(claims.UserID)
	if err != nil {
		fmt.Println(err.Error())
	}
	pages.TodoList(todos).Render(r.Context(), w)
}
