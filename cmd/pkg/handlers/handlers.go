package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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

	render.DisplayTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.DisplayTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation is the reservation page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{})
}

// Generals is the general page handler
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors is the about major page handler
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Contact is the about contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// SearchAvailability renders the search availability page
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	_, err := w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
	if err != nil {
		return
	}
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      false,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		return
	}
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.DisplayTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}
