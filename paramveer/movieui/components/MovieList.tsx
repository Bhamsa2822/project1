import { Alert, AlertIcon, AlertTitle, Box, Center } from '@chakra-ui/react'
import { Movie } from "../movies";
import getMovies, { createMovie, deleteMovie } from "@/api";
import { useEffect, useState } from "react";
import MovieTable from "./MovieTable";
import PopoverForm from "./PopoverForm";

const initialMovie: Movie = {
    id: 0,
    title: "",
    director: "",
    imdb: 0,
    hollywood: "",
    bollywood: "",
}

export default function MovieList() {
    const [movies, setMovies] = useState<Movie[]>([])
    const [error, setError] = useState("")

    function loadMovies(): void {
        getMovies()
            .then((movies) => setMovies(movies))
            .catch((error) => {
                setError(error)
            })
    }

    function onDelete(id: string): Promise<void> {
        return deleteMovie(id)
            .then(() => loadMovies())
            .catch((error) => { setError(error); console.log(error) });
    }

    function addMovie(movie: Movie): Promise<void> {
        return createMovie(movie).then(() => loadMovies())
    }

    useEffect(() => {
        loadMovies()
    }, [])

    return (
        <Box>
            <Center>
                <PopoverForm saveMovie={addMovie} buttonText="ADD" movie={initialMovie} buttonColour="teal" />
            </Center>

            {error && (<Alert status='error'>
                <AlertIcon />
                <AlertTitle>{error}</AlertTitle>
            </Alert>
            )}

            <Center>
                <MovieTable movies={movies} loadMovies={loadMovies} DeleteMovie={onDelete} />
            </Center>
        </Box>
    )
}
