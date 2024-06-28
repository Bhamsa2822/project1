import { Button, Popover, PopoverArrow, PopoverBody, PopoverCloseButton, PopoverContent, PopoverHeader, PopoverTrigger } from "@chakra-ui/react";
import { useState } from "react";
import { Movie } from "../movies";
import MovieForm from "./MovieForm";

export interface PopoverFormProps {
    saveMovie(movie: Movie): Promise<void>
    buttonText: string
    movie: Movie
    buttonColour: string
}

export default function PopoverForm({ saveMovie, buttonText, movie, buttonColour }: PopoverFormProps) {
    const [open, setOpen] = useState(false)

    function toggle() {
        setOpen(!open)
    }

    function close() {
        setOpen(false)
    }

    async function onSave(movie: Movie) {
        await saveMovie(movie)
        close()
    }

    return (
        <Popover size={"100%"} isOpen={open} onClose={close}>
            <PopoverTrigger>
                <Button onClick={toggle} colorScheme={buttonColour}>{buttonText}</Button>
            </PopoverTrigger>
            <PopoverContent>
                <PopoverArrow />
                <PopoverCloseButton />
                <PopoverHeader></PopoverHeader>
                <PopoverBody>
                    {open && <MovieForm saveMovie={onSave} movie={movie} />}
                </PopoverBody>
            </PopoverContent>
        </Popover>
    )
}