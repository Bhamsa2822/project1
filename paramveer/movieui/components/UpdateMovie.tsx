import { updateMovie } from "@/api";
import { Center } from "@chakra-ui/react";
import { Movie } from "../movies";
import PopoverForm from "./PopoverForm";

export interface updateMovieProps {
    loadMovies(): void
    movie: Movie
}

export default function UpdateMovie({ loadMovies, movie }: updateMovieProps) {
    function update(movie: Movie) {
        return updateMovie(movie).then(() => loadMovies())
    }

    return (
        <Center>
            <PopoverForm saveMovie={async (movie) => update(movie)} buttonText="Update" movie={movie} buttonColour="yellow" />
        </Center>
    )
}
