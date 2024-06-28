import { Table, Thead, Tbody, Tr, Th, TableCaption, TableContainer } from '@chakra-ui/react'
import TableRow from './TableRow'
import { Movie } from '../movies'

export interface MovieTableProps {
    movies: Movie[]
    DeleteMovie: (id: string) => Promise<void>
    loadMovies(): void
}

export default function MovieTable({ movies, loadMovies, DeleteMovie }: MovieTableProps,): JSX.Element {
    return (
        <TableContainer height={"100%"} width={"80%"}>
            <Table colorScheme="telegram" height={"100%"}>
                <TableCaption placement='top'> Movie List</TableCaption>
                <Thead>
                    <Tr>
                        <Th>ID</Th>
                        <Th>Title</Th>
                        <Th>Director</Th>
                        <Th>IMDb</Th>
                        <Th>Hollywood</Th>
                        <Th>Bollywood</Th>
                    </Tr>
                </Thead>
                <Tbody>
                    {movies.map((movie, index) => (
                        <TableRow key={index} movie={movie} loadMovies={loadMovies} DeleteMovie={DeleteMovie} />
                    ))}
                </Tbody>
            </Table>
        </TableContainer>
    )
}