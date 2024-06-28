package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestInMemory_createMovie(t *testing.T) {
	type args struct {
		newmovie Movie
	}

	tests := []struct {
		name          string
		existingmovie []Movie
		args          args
		wantErr       error
	}{
		{
			name: "when movie already exist",
			existingmovie: []Movie{
				{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			args: args{
				newmovie: Movie{
					ID:        1,
					Title:     "bhamsa",
					Director:  "paramveer",
					IMDb:      8,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			wantErr: errConflict,
		},
		{
			name:          "movie does't exist",
			existingmovie: []Movie{},
			args: args{
				newmovie: Movie{
					ID:        1,
					Title:     "paramveer",
					Director:  "bhamsa",
					IMDb:      9,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			repo.movies = tt.existingmovie

			gotErr := repo.createMovie(tt.args.newmovie)

			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("want error %q but got error %q", tt.wantErr, gotErr)
			}
		})
	}
}

func TestInMemory_getMovieById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name          string
		existingmovie []Movie
		args          args
		wantRes       Movie
		wantErr       error
	}{
		{
			name: "movie exist",
			existingmovie: []Movie{
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
			wantRes: Movie{
				ID:        1,
				Title:     "bhamsa",
				Director:  "paramveer",
				IMDb:      9,
				Hollywood: "yes",
				Bollywood: "no",
			},
			wantErr: nil,
		},
		{
			name: "when movie doesn't found",
			existingmovie: []Movie{
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
			wantRes: Movie{},
			wantErr: errNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			repo.movies = tt.existingmovie

			getRes, getErr := repo.getMovieById(tt.args.id)

			if !errors.Is(getErr, tt.wantErr) {
				t.Errorf("want error %q but got %q", tt.wantErr, getErr)
			}

			if getRes != tt.wantRes {
				t.Errorf("got %+v but want %+v", getRes, tt.wantRes)
			}
		})
	}
}

func TestInMemory_getAllMovies(t *testing.T) {
	tests := []struct {
		name          string
		existingmovie []Movie
		wantRes       []Movie
		wantErr       error
	}{
		{
			name: "movies exists",
			existingmovie: []Movie{
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
			wantRes: []Movie{
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
			name:          "movies not present",
			existingmovie: []Movie{},
			wantRes:       []Movie{},
			wantErr:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			repo.movies = tt.existingmovie

			gotRes, gotErr := repo.getAllMovie()

			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("want error %q but got %q", tt.wantErr, gotErr)
			}

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("got %+v but want %+v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestInMemory_deleteMovie(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name          string
		existingmovie []Movie
		wantMovies    []Movie
		args          args
		wantRes       Movie
		wantErr       error
	}{
		{
			name: "when movie exist",
			existingmovie: []Movie{
				{
					ID:        1,
					Title:     "Paramveer",
					Director:  "Bhamsa",
					IMDb:      9,
					Hollywood: "yes",
					Bollywood: "no",
				},
				{
					ID:        2,
					Title:     "Paramveer singh",
					Director:  "Bhamsa sarangdevot",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},

			wantMovies: []Movie{
				{
					ID:        2,
					Title:     "Paramveer singh",
					Director:  "Bhamsa sarangdevot",
					IMDb:      10,
					Hollywood: "no",
					Bollywood: "yes",
				},
			},
			args: args{
				id: 1,
			},
			wantRes: Movie{
				ID:        1,
				Title:     "Paramveer",
				Director:  "Bhamsa",
				IMDb:      9,
				Hollywood: "yes",
				Bollywood: "no",
			},
			wantErr: nil,
		},
		{
			name: "when movie doesn't found",
			existingmovie: []Movie{
				{
					ID:        1,
					Title:     "Paramveer",
					Director:  "Bhamsa",
					IMDb:      9,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},

			wantMovies: []Movie{
				{
					ID:        1,
					Title:     "Paramveer",
					Director:  "Bhamsa",
					IMDb:      9,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: 2,
			},
			wantRes: Movie{},
			wantErr: errNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			repo.movies = tt.existingmovie

			getRes, getErr := repo.deleteMovie(tt.args.id)

			if !errors.Is(getErr, tt.wantErr) {
				t.Errorf("got error %q but want %q", getErr, tt.wantErr)
			}

			if getRes != tt.wantRes {
				t.Errorf("got %+v but want %+v", getRes, tt.wantRes)
			}

			if !reflect.DeepEqual(repo.movies, tt.wantMovies) {
				t.Errorf("got %+v\n but want %+v", repo.movies, tt.wantMovies)
			}
		})
	}
}

func TestInMemoryRepo_updateMovie(t *testing.T) {
	type args struct {
		id       int
		newmovie Movie
	}
	tests := []struct {
		name          string
		existingmovie []Movie
		wantMovies    []Movie
		args          args
		wantRes       Movie
		wantErr       error
	}{
		{
			name: "movie exist",
			existingmovie: []Movie{
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
					Title:     "Paramveer singh ",
					Director:  "hardik sharma",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			wantMovies: []Movie{
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
				newmovie: Movie{
					ID:        1,
					Title:     "Paramveer",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			wantRes: Movie{
				ID:        1,
				Title:     "Paramveer",
				Director:  "hardik",
				IMDb:      8,
				Hollywood: "yes",
				Bollywood: "no",
			},
			wantErr: nil,
		},
		{
			name: "movie doesn't exist",
			existingmovie: []Movie{
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
					Title:     "Paramveer singh ",
					Director:  "hardik sharma",
					IMDb:      8,
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
					Title:     "Paramveer singh ",
					Director:  "hardik sharma",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			args: args{
				id: 3,
				newmovie: Movie{
					ID:        1,
					Title:     "Paramveer",
					Director:  "hardik",
					IMDb:      8,
					Hollywood: "yes",
					Bollywood: "no",
				},
			},
			wantRes: Movie{},
			wantErr: errNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepo()
			repo.movies = tt.existingmovie

			getRes, getErr := repo.updateMovie(tt.args.id, tt.args.newmovie)

			if !errors.Is(getErr, tt.wantErr) {
				t.Errorf("want error %q but got %q", tt.wantErr, getErr)
			}

			if getRes != tt.wantRes {
				t.Errorf("got %+v want %+v", getRes, tt.wantRes)
			}

			if !reflect.DeepEqual(repo.movies, tt.wantMovies) {
				t.Errorf("got %+v \n but want %+v", repo.movies, tt.wantMovies)
			}

		})
	}
}
