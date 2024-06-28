package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func resolveError(w http.ResponseWriter, statusCode int, str string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(str); err != nil {
		log.Printf("failed to send response for :%q", err)
	}
}

type movieHandler struct {
	serv movieService
}

func NewMovieHandler(s movieService) *movieHandler {
	return &movieHandler{serv: s}
}

func (h *movieHandler) createMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie Movie
	if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
		resolveError(w, http.StatusBadRequest, "invalid body", err)
		return
	}

	if err := h.serv.CreateMovie(newMovie); err != nil {
		if errors.Is(err, errConflict) {
			resolveError(w, http.StatusConflict, "movie already exist", err)
			return
		}

		if errors.Is(err, errInvalidId) {
			resolveError(w, http.StatusBadRequest, "invalid id", err)
			return
		}

		if errors.Is(err, errInvalidRating) {
			resolveError(w, http.StatusBadRequest, "invalid rating", err)
			return
		}

		resolveError(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newMovie); err != nil {
		log.Println("failed to send response:", err)
	}
}

func (h *movieHandler) getMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.serv.GetAllMovie()
	if err != nil {
		resolveError(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		log.Println("failed to send response:", err)
		return
	}
}

func (h *movieHandler) getMovie(w http.ResponseWriter, r *http.Request) {
	var idStr string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resolveError(w, http.StatusBadRequest, "cannot access id", err)
		return
	}

	movie, err := h.serv.GetMovieById(id)
	if err != nil {
		if errors.Is(err, errInvalidId) {
			resolveError(w, http.StatusBadRequest, "invalid id", err)
			return
		}

		if errors.Is(err, errNotFound) {
			resolveError(w, http.StatusNotFound, "movie not found", err)
			return
		}

		resolveError(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		log.Println("failed to send response:", err)
		return
	}
}

func (h *movieHandler) updateMovie(w http.ResponseWriter, r *http.Request) {
	var idStr string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resolveError(w, http.StatusBadRequest, "cannot access id", err)
		return
	}

	var updatedMovie Movie
	if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
		resolveError(w, http.StatusBadRequest, "invalid body", err)
		return
	}

	movie, err := h.serv.UpdateMovie(id, updatedMovie)
	if err != nil {
		if errors.Is(err, errNotFound) {
			resolveError(w, http.StatusNotFound, "movie not found", err)
			return
		}

		if errors.Is(err, errInvalidId) {
			resolveError(w, http.StatusBadRequest, "invalid id", err)
			return
		}

		if errors.Is(err, errInvalidRating) {
			resolveError(w, http.StatusBadRequest, "invalid rating", err)
			return
		}

		resolveError(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		log.Println("failed to send response for :", err)
		return
	}
}

func (h *movieHandler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	var idStr string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resolveError(w, http.StatusBadRequest, "cannot access id", err)
		return
	}

	deleteMovie, err := h.serv.DeleteMovie(id)
	if err != nil {
		if errors.Is(err, errInvalidId) {
			resolveError(w, http.StatusBadRequest, "invalid id", err)
			return
		}

		if errors.Is(err, errNotFound) {
			resolveError(w, http.StatusNotFound, "movie not found", err)
			return
		}
		resolveError(w, http.StatusInternalServerError, "internal server", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(deleteMovie); err != nil {
		log.Println("failed to send response:", err)
		return
	}
}

func registerRoutes(h *movieHandler) *mux.Router {
	router := mux.NewRouter()
	router.Path("/api/movies").Methods("POST").HandlerFunc(h.createMovie)
	router.Path("/api/movies/{id}").Methods("PUT").HandlerFunc(h.updateMovie)
	router.Path("/api/movies/{id}").Methods("GET").HandlerFunc(h.getMovie)
	router.Path("/api/movies").Methods("GET").HandlerFunc(h.getMovies)
	router.Path("/api/movies/{id}").Methods("DELETE").HandlerFunc(h.deleteMovie)

	return router
}

func main() {
	repo := NewInMemoryRepo()
	serv := Newservice(repo)
	transport := NewMovieHandler(serv)
	router := registerRoutes(transport)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println("http server exited:", err)
	}
}
