package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_service_createMovie(t *testing.T) {
	type args struct {
		newmovie Movie
	}
	tests := []struct {
		name           string
		existingMovies []Movie
		args           args
		wantMovies     []Movie
		wantErr        error
	}{
		{
			name:           "invalid id",
			existingMovies: []Movie{},
			args: args{
				newmovie: Movie{
					ID:        0,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantMovies: []Movie{},
			wantErr:    errInvalidId,
		},
		{
			name:           "invalid rating",
			existingMovies: []Movie{},
			args: args{
				newmovie: Movie{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      11,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantMovies: []Movie{},
			wantErr:    errInvalidRating,
		},
		{
			name: "id conflict",
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
					Title:     "Hardik",
					Director:  "Sharma",
					IMDb:      9,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			args: args{
				newmovie: Movie{
					ID:        1,
					Title:     "confliced movie",
					Director:  "conflicted director",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantMovies: []Movie{
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
					Title:     "Hardik",
					Director:  "Sharma",
					IMDb:      9,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantErr: errConflict,
		},
		{
			name: "test for new movie",
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
			args: args{
				newmovie: Movie{
					ID:        2,
					Title:     "Singh",
					Director:  "Paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantMovies: []Movie{
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
					Title:     "Singh",
					Director:  "Paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			serv := Newservice(repo)

			repo.movies = tt.existingMovies
			getErr := serv.CreateMovie(tt.args.newmovie)

			if !errors.Is(getErr, tt.wantErr) {
				t.Errorf("want error %q but got %q", tt.wantErr, getErr)
			}

			if !reflect.DeepEqual(repo.movies, tt.wantMovies) {
				t.Errorf("got %+v \n but want %+v", repo.movies, tt.wantMovies)
			}
		})
	}
}

func Test_service_updateMovie(t *testing.T) {
	type args struct {
		id           int
		updatedMovie Movie
	}
	tests := []struct {
		name           string
		existingMovies []Movie
		args           args
		wantMovies     []Movie
		want           Movie
		wantErr        error
	}{
		{
			name: "invalid id",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: 1,
				updatedMovie: Movie{
					ID:        0,
					Title:     "Bhamsa",
					Director:  "Paramveer",
					IMDb:      9,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			want:    Movie{},
			wantErr: errInvalidId,
		},
		{
			name: "invalid rating",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: 1,
				updatedMovie: Movie{
					ID:        1,
					Title:     "Bhamsa",
					Director:  "Hardik",
					IMDb:      11,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			wantMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			want:    Movie{},
			wantErr: errInvalidRating,
		},
		{
			name: "not found",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: 2,
				updatedMovie: Movie{
					ID:        1,
					Title:     "Bhamsa",
					Director:  "Hardik",
					IMDb:      10,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			wantMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			want:    Movie{},
			wantErr: errNotFound,
		},
		{
			name: "updated movie",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
				{
					ID:        2,
					Title:     "Singh",
					Director:  "Sharma",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: 2,
				updatedMovie: Movie{
					ID:        2,
					Title:     "updated movie",
					Director:  "updated director",
					IMDb:      10,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			wantMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
				{
					ID:        2,
					Title:     "updated movie",
					Director:  "updated director",
					IMDb:      10,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			want: Movie{
				ID:        2,
				Title:     "updated movie",
				Director:  "updated director",
				IMDb:      10,
				Hollywood: "yes",
				Bollywood: "no",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			serv := Newservice(repo)

			repo.movies = tt.existingMovies

			getMovie, getErr := serv.UpdateMovie(tt.args.id, tt.args.updatedMovie)

			if !errors.Is(getErr, tt.wantErr) {
				t.Errorf("want error %q but got %q", tt.wantErr, getErr)
			}

			if getMovie != tt.want {
				t.Errorf("want %+v but got %+v", tt.want, getMovie)
			}

			if !reflect.DeepEqual(repo.movies, tt.wantMovies) {
				t.Errorf("got %+v \n but want %+v", repo.movies, tt.wantMovies)
			}
		})
	}
}

func Test_service_GetMovieById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name           string
		existingMovies []Movie
		args           args
		wantMovie      Movie
		wantErr        error
	}{
		{
			name: "invalid id",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "Paramveer",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: -1,
			},
			wantMovie: Movie{},
			wantErr:   errInvalidId,
		},
		{
			name: "movie not found",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "Paramveer",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: 2,
			},
			wantMovie: Movie{},
			wantErr:   errNotFound,
		},
		{
			name: "Valid id",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "Paramveer",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
				{
					ID:        2,
					Title:     "Paramveer singh ",
					Director:  "hardik sharma",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: 1,
			},
			wantMovie: Movie{
				ID:        1,
				Title:     "Paramveer",
				Director:  "hardik",
				IMDb:      8,
				Hollywood: "yes",
				Bollywood: "no",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			serv := Newservice(repo)

			repo.movies = tt.existingMovies

			getMovie, getErr := serv.GetMovieById(tt.args.id)

			if !errors.Is(getErr, tt.wantErr) {
				t.Errorf("want error %q but got %q", tt.wantErr, getErr)
			}

			if getMovie != tt.wantMovie {
				t.Errorf("got movies %+v but want movies %+v", getMovie, tt.wantMovie)
			}
		})
	}
}

func Test_service_GetAllMovie(t *testing.T) {
	tests := []struct {
		name           string
		existingMovies []Movie
		wantMovies     []Movie
		wantErr        error
	}{
		{
			name: "movies exist",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "Bhamsa",
					Director:  "Paramveer",
					IMDb:      7.7,
					Hollywood: "yes",
					Bollywood: "no",
				},
				{
					ID:        2,
					Title:     "Paramveer",
					Director:  "Bhamsa",
					IMDb:      8.9,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantMovies: []Movie{
				{
					ID:        1,
					Title:     "Bhamsa",
					Director:  "Paramveer",
					IMDb:      7.7,
					Hollywood: "yes",
					Bollywood: "no",
				},
				{
					ID:        2,
					Title:     "Paramveer",
					Director:  "Bhamsa",
					IMDb:      8.9,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantErr: nil,
		},
		{
			name:           "movies not present",
			existingMovies: []Movie{},
			wantMovies:     []Movie{},
			wantErr:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			serv := Newservice(repo)

			repo.movies = tt.existingMovies

			getMovies, getErr := serv.GetAllMovie()

			if !errors.Is(getErr, tt.wantErr) {
				t.Errorf("want error %q but got %q", tt.wantErr, getErr)
			}

			if !reflect.DeepEqual(getMovies, tt.wantMovies) {
				t.Errorf("want movies %+v but got movies %+v", tt.wantMovies, getMovies)
			}
		})
	}
}

func Test_service_DeleteMovie(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name           string
		existingMovies []Movie
		args           args
		want           Movie
		wantMovies     []Movie
		wantErr        error
	}{
		{
			name: "valid id",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      9,
					Hollywood: "yes",
					Bollywood: "no",
				},
				{
					ID:        2,
					Title:     "Hardik Sharma",
					Director:  "Paramveer singh sarangdevot",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			args: args{
				id: 1,
			},
			want: Movie{
				ID:        1,
				Title:     "bhamsa",
				Director:  "paramveer",
				IMDb:      9,
				Hollywood: "yes",
				Bollywood: "no",
			},
			wantMovies: []Movie{
				{
					ID:        2,
					Title:     "Hardik Sharma",
					Director:  "Paramveer singh sarangdevot",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantErr: nil,
		},
		{
			name: "invalid id",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      9,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: -1,
			},
			want: Movie{},
			wantMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      9,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			wantErr: errInvalidId,
		},
		{
			name: "movie not found",
			existingMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      9,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: 2,
			},
			want: Movie{},
			wantMovies: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      9,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			wantErr: errNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			serv := Newservice(repo)

			repo.movies = tt.existingMovies

			getMovie, getErr := serv.DeleteMovie(tt.args.id)

			if !errors.Is(getErr, tt.wantErr) {
				t.Errorf("want error %q but got %q", tt.wantErr, getErr)
			}

			if getMovie != tt.want {
				t.Errorf("want %+v but got %+v", tt.want, getMovie)
			}

			if !reflect.DeepEqual(repo.movies, tt.wantMovies) {
				t.Errorf("got movies %+v\n but want movies %+v", repo.movies, tt.wantMovies)
			}
		})
	}
}
