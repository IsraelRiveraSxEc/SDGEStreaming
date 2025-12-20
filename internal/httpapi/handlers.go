package httpapi

import (
	"SDGEStreaming/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

// Referencias a los servicios
var (
	userService         *services.UserService
	contentService      *services.ContentService
	subscriptionService *services.SubscriptionService
	playbackService     *services.PlaybackService
)

func RegisterHandlers(
	mux *http.ServeMux,
	uSvc *services.UserService,
	cSvc *services.ContentService,
	sSvc *services.SubscriptionService,
	pSvc *services.PlaybackService,
) {
	userService = uSvc
	contentService = cSvc
	subscriptionService = sSvc
	playbackService = pSvc

	// Auth
	mux.HandleFunc("/api/register", registerHandler)
	mux.HandleFunc("/api/login", loginHandler)

	// Planes
	mux.HandleFunc("/api/plans", getPlansHandler)
	mux.HandleFunc("/api/subscriptions/change-plan", changePlanHandler)

	// Contenido
	mux.HandleFunc("/api/content/audiovisual", getAudiovisualContentHandler)
	mux.HandleFunc("/api/content/audio", getAudioContentHandler)

	// Subida de contenido
	mux.HandleFunc("/api/admin/content/audiovisual/upload", createAudiovisualHandler)
	mux.HandleFunc("/api/admin/content/audio/upload", createAudioHandler)

	// Ratings (SIN id en la URL, igual que en pruebas PS)
	mux.HandleFunc("/api/content/audiovisual/rate", rateAudiovisualHandler)
	mux.HandleFunc("/api/content/audio/rate", rateAudioHandler)
}

// Helpers JSON
func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":"Error interno del servidor"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, map[string]string{"error": message})
}

// ------------------- Auth -------------------

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	var req struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Todos los campos son obligatorios")
		return
	}

	user, err := userService.Register(req.Name, req.Age, req.Email, req.Password, false)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error en el registro: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Usuario registrado exitosamente",
		"user_id": user.ID,
		"email":   user.Email,
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	user, err := userService.Login(req.Email, req.Password)
	if err != nil || user == nil {
		respondWithError(w, http.StatusUnauthorized, "Email o contraseña incorrectos")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message":   "Inicio de sesión exitoso",
		"user_id":   user.ID,
		"email":     user.Email,
		"plan_id":   user.PlanID,
		"plan_name": getPlanName(user.PlanID),
		"is_admin":  user.IsAdmin,
	})
}

// ------------------- Planes -------------------

func getPlansHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	plans, err := subscriptionService.GetAvailablePlans()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error al obtener planes: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, plans)
}

func changePlanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	var req struct {
		UserID      int    `json:"user_id"`
		PlanID      int    `json:"plan_id"`
		CardHolder  string `json:"card_holder"`
		CardNumber  string `json:"card_number"`
		ExpiryMonth int    `json:"expiry_month"`
		ExpiryYear  int    `json:"expiry_year"`
		CVV         int    `json:"cvv"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if req.UserID <= 0 || req.PlanID <= 0 {
		respondWithError(w, http.StatusBadRequest, "user_id y plan_id deben ser mayores a 0")
		return
	}

	if err := subscriptionService.ProcessPayment(
		req.UserID,
		req.PlanID,
		req.CardHolder,
		req.CardNumber,
		req.ExpiryMonth,
		req.ExpiryYear,
		req.CVV,
	); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error en el pago: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Plan actualizado exitosamente",
	})
}

// ------------------- Contenido -------------------

func getAudiovisualContentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	contents, err := contentService.GetAllAudiovisual()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error al obtener contenido audiovisual: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, contents)
}

func getAudioContentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	contents, err := contentService.GetAllAudio()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error al obtener contenido de audio: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, contents)
}

// ------------------- Upload Content -----------
func createAudiovisualHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	var req struct {
		Title       string `json:"title"`
		Type        string `json:"type"`
		Genre       string `json:"genre"`
		Duration    int    `json:"duration"`
		AgeRating   string `json:"age_rating"`
		Synopsis    string `json:"synopsis"`
		ReleaseYear int    `json:"release_year"`
		Director    string `json:"director"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if req.Title == "" || req.Type == "" {
		respondWithError(w, http.StatusBadRequest, "Campos obligatorios faltantes")
		return
	}

	err := contentService.CreateAudiovisual(
		req.Title,
		req.Type,
		req.Genre,
		req.Duration,
		req.AgeRating,
		req.Synopsis,
		req.ReleaseYear,
		req.Director,
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{
		"message": "Contenido audiovisual creado",
	})
}

// ---------- Upload Audio Content ---------------
func createAudioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	var req struct {
		Title       string `json:"title"`
		Type        string `json:"type"`
		Genre       string `json:"genre"`
		Duration    int    `json:"duration"`
		AgeRating   string `json:"age_rating"`
		Artist      string `json:"artist"`
		Album       string `json:"album"`
		TrackNumber int    `json:"track_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if req.Title == "" || req.Type == "" {
		respondWithError(w, http.StatusBadRequest, "Campos obligatorios faltantes")
		return
	}

	err := contentService.CreateAudio(
		req.Title,
		req.Type,
		req.Genre,
		req.Duration,
		req.AgeRating,
		req.Artist,
		req.Album,
		req.TrackNumber,
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{
		"message": "Contenido de audio creado",
	})
}
// ------------------- Ratings -------------------

// POST /api/content/audiovisual/rate
// Body:
// { "user_id": 3, "content_id": 1, "rating": 8.5 }
func rateAudiovisualHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	var req struct {
		UserID    int     `json:"user_id"`
		ContentID int     `json:"content_id"`
		Rating    float64 `json:"rating"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if req.UserID <= 0 || req.ContentID <= 0 {
		respondWithError(w, http.StatusBadRequest, "user_id y content_id deben ser mayores a 0")
		return
	}
	if req.Rating < 1.0 || req.Rating > 10.0 {
		respondWithError(w, http.StatusBadRequest, "rating debe estar entre 1.0 y 10.0")
		return
	}

	if err := contentService.RateContent(req.UserID, req.ContentID, "audiovisual", req.Rating); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error al valorar contenido audiovisual: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Valoración registrada correctamente",
	})
}

// POST /api/content/audio/rate
// Body:
// { "user_id": 3, "content_id": 1, "rating": 9.0 }
func rateAudioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	var req struct {
		UserID    int     `json:"user_id"`
		ContentID int     `json:"content_id"`
		Rating    float64 `json:"rating"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if req.UserID <= 0 || req.ContentID <= 0 {
		respondWithError(w, http.StatusBadRequest, "user_id y content_id deben ser mayores a 0")
		return
	}
	if req.Rating < 1.0 || req.Rating > 10.0 {
		respondWithError(w, http.StatusBadRequest, "rating debe estar entre 1.0 y 10.0")
		return
	}

	if err := contentService.RateContent(req.UserID, req.ContentID, "audio", req.Rating); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error al valorar contenido de audio: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Valoración registrada correctamente",
	})
}

// ------------------- Helper -------------------

func getPlanName(planID int) string {
	switch planID {
	case 1:
		return "Free"
	case 2:
		return "Estándar"
	case 3:
		return "Premium 4K"
	default:
		return "Desconocido"
	}
}
