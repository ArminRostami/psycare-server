package http

import (
	"net/http"
	"psycare/domain"
)

func (h *Handler) makeAppointment(w http.ResponseWriter, r *http.Request) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	appt := &domain.Appointment{}
	httpErr = h.decodeAndValidate(r, appt)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}
	appt.UserID = id
	err := h.CreateAppointment(appt)
	if err != nil {
		renderError(w, r, &httpError{"failed to add appointment", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, "appointment added")
}

func (h *Handler) appointmentsHandler(w http.ResponseWriter, r *http.Request, forUser bool) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	appts, err := h.GetAppointments(id, forUser)
	if err != nil {
		renderError(w, r, &httpError{"failed to get appointments", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, appts)
}

func (h *Handler) getUserAppointments(w http.ResponseWriter, r *http.Request) {
	h.appointmentsHandler(w, r, true)
}

func (h *Handler) getAdvisorAppointments(w http.ResponseWriter, r *http.Request) {
	h.appointmentsHandler(w, r, false)
}

func (h *Handler) rateAppointment(w http.ResponseWriter, r *http.Request) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	rating := &domain.Rating{UserID: id}
	httpErr = h.decodeAndValidate(r, rating)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	err := h.AddRating(rating)
	if err != nil {
		renderError(w, r, &httpError{"failed to add rating", http.StatusInternalServerError, err})
		return
	}

	renderData(w, r, rating)

}
