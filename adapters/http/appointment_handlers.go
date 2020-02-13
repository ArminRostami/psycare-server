package http

import (
	"net/http"
	"psycare/domain"
)

func (h *Handler) bookAppointment(w http.ResponseWriter, r *http.Request) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	appt := &domain.Appointment{}
	appt.UserID = id
	httpErr = h.decodeAndValidate(r, appt)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	err := h.CreateAppointment(appt)
	if err != nil {
		renderError(w, r, &httpError{"failed to add appointment", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, appt)
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

func (h *Handler) cancelAppointment(w http.ResponseWriter, r *http.Request) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	req := &struct {
		AppointmentID int64 `json:"appointment_id" validate:"required"`
	}{}

	httpErr = h.decodeAndValidate(r, req)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	err := h.CancelAppointment(id, req.AppointmentID)
	if err != nil {
		renderError(w, r, &httpError{"cannot cancel appointment", http.StatusInternalServerError, err})
		return
	}

	renderData(w, r, "appointment cancelled")
}
