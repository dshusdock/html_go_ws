package handlers

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/apis"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/render"
	"fmt"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

/**
 * 	HandleClickEvents
 */
func (m *Repository) HandleClickEvents(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("This is a test")
	data := r.PostForm
	data.Add("event", constants.EVENT_CLICK)
	v_id := data.Get("view_id")

	if v_id == "" {
		_ = fmt.Errorf("no handler for route")
		return
	}

	// route request to appropriate handler
	ptr := m.App.ViewCache[v_id]
	ptr.ProcessRequest(w, data)
}

/**
 * 	Home is the handler for the home page
 */
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, m.App)
}

func (m *Repository) Test(w http.ResponseWriter, r *http.Request) {

	fmt.Println("This is a tester")

	dbh, err := apis.OpenDB("127.0.0.1", "testdb")
	if err != nil {
		log.Fatal(err)
	}

	uu, _ := apis.GetUserInfo(dbh)

	apis.DisplayUsers(uu)
}
