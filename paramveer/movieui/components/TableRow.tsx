import { Box, Tr, Td, Button } from '@chakra-ui/react'
import UpdateMovie from './UpdateMovie'
import { Movie } from '../movies'

export interface TableRowProps {
    movie: Movie
    DeleteMovie: (id: string) => Promise<void>
    loadMovies(): void
}

export default function TableRow({ movie, loadMovies, DeleteMovie }: TableRowProps): JSX.Element {
    const id = movie.id.toString()
    return (
        <Tr>
            <Td>{movie.id}</Td>
            <Td>{movie.title}</Td>
            <Td>{movie.director}</Td>
            <Td>{movie.imdb}</Td>
            <Td>{movie.hollywood}</Td>
            <Td>{movie.bollywood}</Td>
            <Td display="flex">
                <Button margin="10px" onClick={(() => DeleteMovie(id))} colorScheme="red">Delete</Button>
                <UpdateMovie loadMovies={loadMovies} movie={movie} />
            </Td>
        </Tr>
    )
}