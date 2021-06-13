import React, { useState } from "react"
import styled from "styled-components"

const Button = styled.button`
    background: none;
    border: none;
    color: white;
`

async function likeTrack(trackId: number): Promise<boolean> {
    const response = await fetch(`/api/v1/track/${trackId}/like`)

    if (!response.ok)
        throw new Error(`Http request was not successful, status code: ${response.status}`)

    return true
}

async function unlikeTrack(trackId: number): Promise<boolean> {
    const response = await fetch(`/api/v1/track/${trackId}/unlike`)

    if (!response.ok)
        throw new Error(`Http request was not successful, status code: ${response.status}`)

    return false
}

interface ILikeButtonProps {
    trackId: number
    liked: boolean
    onLikeChanged?: (status: boolean, trackId: number) => void
}

export default function LikeButton({ trackId, liked: originalLiked, onLikeChanged }: ILikeButtonProps) {
    const [liked, setLiked] = useState<boolean>(originalLiked)

    const onClickFn = liked ? unlikeTrack : likeTrack

    const buttonText = liked ? "unlike" : "like"

    return <Button
        onClick={e => {
            e.preventDefault()
            e.stopPropagation()

            onClickFn(trackId)
                .then(liked => {
                    setLiked(liked)

                    if (typeof (onLikeChanged) === "function")
                        onLikeChanged(liked, trackId)
                })
                .catch(error => {
                    console.error("Failed to toggle like status", error)
                })
        }}
    >
        {buttonText}
    </Button>
}