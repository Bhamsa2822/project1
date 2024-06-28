import axios from "axios";
import { Movie } from "../movies";

export function createMovie(movie: Movie): Promise<void> {
  return axios.post("/api/movies", movie)
}

export default function getMovies():Promise<Movie[]> {
  return axios.get("/api/movies").then((res)=>{return res.data})
}

export function deleteMovie(id: string): Promise<void> {
  return axios.delete(`/api/movies/${id}`)
}

export function updateMovie(movie: Movie,): Promise<Movie> {
  return axios.put(`/api/movies/${movie.id}`, movie)
}