package httpapi

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/content/audiovisual", getAudiovisualContentHandler)
	mux.HandleFunc("/api/content/audio", getAudioContentHandler)

	return mux
}
