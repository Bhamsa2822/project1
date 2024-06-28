package main

import "errors"

var errConflict = errors.New("movie already exist")
var errNotFound = errors.New("movie doesn't found")

type Repo interface {
	createMovie(newmovie Movie) error
	getAllMovie() ([]Movie, error)
	getMovieById(id int) (Movie, error)
	updateMovie(id int, newmovie Movie) (Movie, error)
	deleteMovie(id int) (Movie, error)
}

type InMemoryRepo struct {
	movies []Movie
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		movies: []Movie{},
	}
}

func (m *InMemoryRepo) createMovie(newmovie Movie) error {
	for _, existingmovie := range m.movies {
		if existingmovie.ID == newmovie.ID {
			return errConflict
		}
	}
	m.movies = append(m.movies, newmovie)
	return nil
}

func (m *InMemoryRepo) getAllMovie() ([]Movie, error) {
	return m.movies, nil
}

func (m *InMemoryRepo) getMovieById(id int) (Movie, error) {
	for _, existingmovie := range m.movies {
		if id == existingmovie.ID {
			return existingmovie, nil
		}
	}
	return Movie{}, errNotFound
}

func (m *InMemoryRepo) updateMovie(id int, newmovie Movie) (Movie, error) {
	for i, movietoupdate := range m.movies {
		if id == movietoupdate.ID {
			m.movies[i] = newmovie
			return newmovie, nil
		}
	}
	return Movie{}, errNotFound
}

func (m *InMemoryRepo) deleteMovie(id int) (Movie, error) {
	for i, deletedmovie := range m.movies {
		if id == deletedmovie.ID {
			m.movies = append(m.movies[:i], m.movies[i+1:]...)
			return deletedmovie, nil
		}
	}
	return Movie{}, errNotFound
}
