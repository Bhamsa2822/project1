package main

import "errors"

var (
	errInvalidId     = errors.New("invalid id")
	errInvalidRating = errors.New("invalid imdb rating")
)

func validateId(id int) error {
	if id < 0 {
		return errInvalidId
	}
	return nil
}

func validateMovie(movie Movie) error {
	if movie.ID < 1 {
		return errInvalidId
	}

	if movie.IMDb < 1 || movie.IMDb > 10 {
		return errInvalidRating
	}
	return nil
}

type movieService interface {
	CreateMovie(newmovie Movie) error
	GetAllMovie() ([]Movie, error)
	GetMovieById(id int) (Movie, error)
	UpdateMovie(id int, updatedmovie Movie) (Movie, error)
	DeleteMovie(id int) (Movie, error)
}

type service struct {
	repo Repo
}

func Newservice(r Repo) *service {
	return &service{repo: r}
}

func (s *service) CreateMovie(newmovie Movie) error {
	if err := validateMovie(newmovie); err != nil {
		return err
	}

	if err := s.repo.createMovie(newmovie); err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllMovie() ([]Movie, error) {
	movies, err := s.repo.getAllMovie()
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (s *service) GetMovieById(id int) (Movie, error) {
	if err := validateId(id); err != nil {
		return Movie{}, err
	}

	movie, err := s.repo.getMovieById(id)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (s *service) UpdateMovie(id int, updatedmovie Movie) (Movie, error) {
	if err := validateMovie(updatedmovie); err != nil {
		return Movie{}, err
	}

	movie, err := s.repo.updateMovie(id, updatedmovie)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (s *service) DeleteMovie(id int) (Movie, error) {
	if err := validateId(id); err != nil {
		return Movie{}, err
	}

	movie, err := s.repo.deleteMovie(id)
	if err != nil {
		return movie, err
	}

	return movie, nil
}
