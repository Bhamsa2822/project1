package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_movieHandler_createMovie(t *testing.T) {
	tests := []struct {
		name             string
		existingMovies   []Movie
		requestBody      string
		wantResponseBody string
		wantStatusCode   int
	}{
		{
			name:           "new movie",
			existingMovies: []Movie{},
			requestBody: `
			{
				"id":        1,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      8,
				"hollywood": "no",
				"bollywood": "yes"
				}`,
			wantResponseBody: `
			{
				"id": 1,
				"title": "bhamsa",
				"director": "paramveer",
				"imdb": 8,
				"hollywood": "no",
				"bollywood": "yes"
			}`,
			wantStatusCode: http.StatusOK,
		},
		{
			name: "conflict",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			requestBody: `
			{
				"id":        1,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      8,
				"hollywood": "no",
				"bollywood": "yes"
				}`,
			wantResponseBody: `"movie already exist"`,
			wantStatusCode:   http.StatusConflict,
		},
		{
			name:           "invalid id",
			existingMovies: []Movie{},
			requestBody: `
			{
				"id":        0,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      8,
				"hollywood": "no",
				"bollywood": "yes"
				}`,
			wantResponseBody: `"invalid id"`,
			wantStatusCode:   http.StatusBadRequest,
		},
		{
			name:           "invalid rating",
			existingMovies: []Movie{},
			requestBody: `
			{
				"id":        1,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      11,
				"hollywood": "no",
				"bollywood": "yes"
				}`,
			wantResponseBody: `"invalid rating"`,
			wantStatusCode:   http.StatusBadRequest,
		},
		{
			name:           "invalid body",
			existingMovies: []Movie{},
			requestBody: `
			{
				"id":        1,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      11,
				"hollywood": "no",
				"bollywood": "yes" invalid body
				}`,
			wantResponseBody: `"invalid body"`,
			wantStatusCode:   http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			repo.movies = tt.existingMovies

			serv := Newservice(repo)
			transport := NewMovieHandler(serv)

			reqBody := strings.NewReader(tt.requestBody)
			req := httptest.NewRequest("POST", "/api/movies", reqBody)

			res := httptest.NewRecorder()
			transport.createMovie(res, req)

			assert.JSONEqf(t, tt.wantResponseBody, res.Body.String(), "want  %s but got %s", tt.wantResponseBody, res.Body.String())

			if res.Code != tt.wantStatusCode {
				t.Errorf("want statuscode %d but got %d", tt.wantStatusCode, res.Code)
			}
		})
	}
}

func Test_movieHandler_updateMovie(t *testing.T) {
	tests := []struct {
		name             string
		existingMovies   []Movie
		requestBody      string
		wantResponseBody string
		wantStatusCode   int
	}{
		{
			name: "invalid id",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			requestBody: `
			{
				"id":        0,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      10,
				"hollywood": "no",
				"bollywood": "yes"
				}`,
			wantResponseBody: `"invalid id"`,
			wantStatusCode:   http.StatusBadRequest,
		},
		{
			name: "invalid rating",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			requestBody: `
			{
				"id":        1,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      101,
				"hollywood": "no",
				"bollywood": "yes"
				}`,
			wantResponseBody: `"invalid rating"`,
			wantStatusCode:   http.StatusBadRequest,
		},
		{
			name: "movie not found",
			existingMovies: []Movie{
				{
					ID:        2,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			requestBody: `
			{
				"id":        1,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      10,
				"hollywood": "no",
				"bollywood": "yes"
				}`,
			wantResponseBody: `"movie not found"`,
			wantStatusCode:   http.StatusNotFound,
		},
		{
			name: "invalid body",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			requestBody: `
			{
				"id":        1,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      10,
				"hollywood": "no",
				"bollywood": "yes"dvbdhvdv
				}`,
			wantResponseBody: `"invalid body"`,
			wantStatusCode:   http.StatusBadRequest,
		},
		{
			name: "movie exist",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
				{
					ID:        2,
					Title:     "Bhamsa",
					Director:  "Paramveer",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			requestBody: `
			{
				"id":        11,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      10,
				"hollywood": "no",
				"bollywood": "yes"
				}`,
			wantResponseBody: `{
				"id":        11,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      10,
				"hollywood": "no",
				"bollywood": "yes"
				}`,
			wantStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			repo.movies = tt.existingMovies
			serv := Newservice(repo)
			transport := NewMovieHandler(serv)
			router := registerRoutes(transport)

			reqBody := strings.NewReader(tt.requestBody)
			req := httptest.NewRequest("PUT", "/api/movies/1", reqBody)

			res := httptest.NewRecorder()

			router.ServeHTTP(res, req)

			assert.JSONEqf(t, tt.wantResponseBody, res.Body.String(), "want  %s but got %s", tt.wantResponseBody, res.Body.String())

			if res.Code != tt.wantStatusCode {
				t.Errorf("want statuscode %d but got %d", tt.wantStatusCode, res.Code)
			}
		})
	}
}

func Test_movieHandler_getMovie(t *testing.T) {
	tests := []struct {
		name             string
		existingMovies   []Movie
		path             string
		wantResponseBody string
		wantStatusCode   int
	}{
		{
			name: "invalid id",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			path:             "/api/movies/-1",
			wantResponseBody: `"invalid id"`,
			wantStatusCode:   http.StatusBadRequest,
		},
		{
			name: "movie not found",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			path:             "/api/movies/2",
			wantResponseBody: `"movie not found"`,
			wantStatusCode:   http.StatusNotFound,
		},
		{
			name: "movie found",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			path: "/api/movies/1",
			wantResponseBody: `
			{
				"id": 1,
				"title": "bhamsa",
				"director": "paramveer",
				"imdb": 8,
				"hollywood": "no",
				"bollywood": "yes"
			}`,
			wantStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			repo.movies = tt.existingMovies
			serv := Newservice(repo)
			transport := NewMovieHandler(serv)
			router := registerRoutes(transport)

			req := httptest.NewRequest("GET", tt.path, nil)

			res := httptest.NewRecorder()

			router.ServeHTTP(res, req)

			assert.JSONEqf(t, tt.wantResponseBody, res.Body.String(), "want  %s but got %s", tt.wantResponseBody, res.Body.String())

			if res.Code != tt.wantStatusCode {
				t.Errorf("want statuscode %d but got %d", tt.wantStatusCode, res.Code)
			}
		})
	}
}

func Test_movieHandler_getMovies(t *testing.T) {
	tests := []struct {
		name             string
		existingMovies   []Movie
		wantResponseBody string
		wantStatusCode   int
	}{
		{
			name:             "no movie exist  ",
			existingMovies:   []Movie{},
			wantResponseBody: `[]`,
			wantStatusCode:   http.StatusOK,
		},
		{
			name: " movies exist  ",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantResponseBody: `
			[
			  {
				"id": 1,
				"title": "bhamsa",
				"director": "paramveer",
				"imdb": 8,
				"hollywood": "no",
				"bollywood": "yes"
			  }
			]`,
			wantStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			repo.movies = tt.existingMovies

			serv := Newservice(repo)
			transport := NewMovieHandler(serv)
			router := registerRoutes(transport)

			req := httptest.NewRequest("GET", "/api/movies", nil)
			res := httptest.NewRecorder()

			router.ServeHTTP(res, req)

			assert.JSONEqf(t, tt.wantResponseBody, res.Body.String(), "want  %s but got %s", tt.wantResponseBody, res.Body.String())

			if res.Code != tt.wantStatusCode {
				t.Errorf("want statuscode %d but got %d", tt.wantStatusCode, res.Code)
			}
		})
	}
}

func Test_movieHandler_deleteMovie(t *testing.T) {
	tests := []struct {
		name             string
		existingMovies   []Movie
		path             string
		wantResponseBody string
		wantStatusCode   int
	}{
		{
			name: "invalid id",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			path:             "/api/movies/-1",
			wantResponseBody: `"invalid id"`,
			wantStatusCode:   http.StatusBadRequest,
		},
		{
			name:             "movie not found",
			existingMovies:   []Movie{},
			path:             "/api/movies/1",
			wantResponseBody: `"movie not found"`,
			wantStatusCode:   http.StatusNotFound,
		},
		{
			name: "delete ok",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			path: "/api/movies/1",
			wantResponseBody: `
				{
				"id":        1,
				"title":     "bhamsa",
				"director":  "paramveer",
				"imdb":      8,
				"hollywood": "no",
				"bollywood": "yes"
				}`,
			wantStatusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			repo.movies = tt.existingMovies

			serv := Newservice(repo)
			transport := NewMovieHandler(serv)
			router := registerRoutes(transport)

			req := httptest.NewRequest("DELETE", tt.path, nil)
			res := httptest.NewRecorder()

			router.ServeHTTP(res, req)

			assert.JSONEqf(t, tt.wantResponseBody, res.Body.String(), "want  %s but got %s", tt.wantResponseBody, res.Body.String())

			if res.Code != tt.wantStatusCode {
				t.Errorf("want statuscode %d but got %d", tt.wantStatusCode, res.Code)
			}
		})
	}
}
