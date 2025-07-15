package handlers

import (
	"myweb/cmd/pkg/config"
	"myweb/cmd/pkg/models"
	"myweb/cmd/pkg/render"
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

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.DisplayTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.DisplayTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, "reservation.page.tmpl", &models.TemplateData{})
}

// Generals is the general page handler
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

// Majors is the about major page handler
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

// Contact is the about major page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// SearchAvailability is the about major page handler
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}
