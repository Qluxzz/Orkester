import React, { useState } from "react"
import styled from "styled-components"
import { ILikeStatus } from "types/track"

const Button = styled.button`
    background: none;
    border: none;
    color: white;
`

async function likeTrack(trackId: number): Promise<ILikeStatus> {
    const response = await fetch(`/api/v1/track/${trackId}/like`)

    if (!response.ok)
        throw new Error(`Http request was not successful, status code: ${response.status}`)

    return "liked"
}

async function unlikeTrack(trackId: number): Promise<ILikeStatus> {
    const response = await fetch(`/api/v1/track/${trackId}/unlike`)

    if (!response.ok)
        throw new Error(`Http request was not successful, status code: ${response.status}`)

    return "notliked"
}

interface ILikeButtonProps {
    trackId: number
    likeStatus: ILikeStatus
    onLikeStatusChanged?: (status: ILikeStatus) => void
}

export default function LikeButton({ trackId, likeStatus: originalLikeStatus, onLikeStatusChanged }: ILikeButtonProps) {
    const [likeStatus, setLikeStatus] = useState<ILikeStatus>(originalLikeStatus)

    const onClickFn = (() => {
        switch (likeStatus) {
            case "liked":
                return unlikeTrack
            case "notliked":
                return likeTrack
            default:
                throw new Error(`Unknown likestatus: ${likeStatus}`)
        }
    })()

    const buttonText = (() => {
        switch (likeStatus) {
            case "liked":
                return "unlike"
            case "notliked":
                return "like"
            default:
                throw new Error(`Unknown likestatus: ${likeStatus}`)
        }
    })()

    return <Button
        onClick={e => {
            e.preventDefault()
            onClickFn(trackId)
                .then(likeStatus => {
                    setLikeStatus(likeStatus)

                    if (typeof (onLikeStatusChanged) === "function")
                        onLikeStatusChanged(likeStatus)
                })
                .catch(error => {
                    console.error(error)
                })
        }}
    >
        {buttonText}
    </Button>
}