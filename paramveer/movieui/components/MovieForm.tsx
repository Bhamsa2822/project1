import { Alert, AlertIcon, AlertTitle, Box, Button, Center, FormLabel, Heading, Input } from '@chakra-ui/react'
import { ChangeEvent, FormEvent, useState } from 'react'
import { Movie } from '../movies'

interface MovieInput {
    id: string,
    title: string,
    director: string
    imdb: string,
    hollywood: string
    bollywood: string
}

export interface MovieFormProps {
    saveMovie(movie: Movie): Promise<void>
    movie: Movie
}

export default function MovieForm({ saveMovie, movie }: MovieFormProps): JSX.Element {
    const initialMovie: MovieInput = {
        id: movie.id.toString(),
        title: movie.title,
        director: movie.director,
        imdb: movie.imdb.toString(),
        hollywood: movie.hollywood,
        bollywood: movie.bollywood,
    }

    const [movieInput, setMovieInput] = useState<MovieInput>(initialMovie)
    const [errorMsg, setErrorMsg] = useState("")

    function handleChange(e: ChangeEvent<HTMLInputElement>) {
        setMovieInput({ ...movieInput, [e.target.name]: e.target.value })
    }

    function handleSubmit(e: FormEvent<HTMLFormElement>) {
        e.preventDefault()

        const movie: Movie = {
            id: parseInt(movieInput.id),
            title: movieInput.title,
            director: movieInput.director,
            imdb: parseInt(movieInput.imdb),
            hollywood: movieInput.hollywood,
            bollywood: movieInput.bollywood
        }

        saveMovie(movie)
            .then(() => {
                setErrorMsg("")
                setMovieInput(initialMovie)
            })
            .catch((error) => {
                setErrorMsg(error?.response?.data)
                console.log(error)
            })
    }

    return (
        <Box>
            <Center>
                {errorMsg && (
                    <Alert status='error'>
                        <AlertIcon />
                        <AlertTitle>{errorMsg}</AlertTitle>
                    </Alert>
                )}
            </Center>

            <form onSubmit={handleSubmit}>
                <Heading>Movie Details</Heading>

                <FormLabel>ID:</FormLabel>
                <Input type="number" placeholder='Enter Id' name="id" value={movieInput.id} onChange={handleChange} />

                <FormLabel>Title:</FormLabel>
                <Input type="text" placeholder='Enter Title' name="title" value={movieInput.title} onChange={handleChange} />

                <FormLabel>Director:</FormLabel>
                <Input type="text" placeholder="Enter Director" name="director" value={movieInput.director} onChange={handleChange} />

                <FormLabel>IMDb:</FormLabel>
                <Input type="number" placeholder='Enter Rating' name="imdb" value={movieInput.imdb} onChange={handleChange} />

                <FormLabel>Hollywood:</FormLabel>
                <Input type="text" placeholder="yes/no" name="hollywood" value={movieInput.hollywood} onChange={handleChange} />

                <FormLabel>Bollywood:</FormLabel>
                <Input type="text" placeholder="yes/no" name="bollywood" value={movieInput.bollywood} onChange={handleChange} />

                <div>
                    <Button type='submit' colorScheme="blue">Submit</Button>
                </div>
            </form>
        </Box>
    )
}