import React from "react"

import ITrack from "types/track"
import TrackListBase from "./TrackListBase"


interface ITrackList {
    tracks: ITrack[]
    onLikedChanged?: (liked: boolean, trackId: number) => void
}

export default function TrackList({ tracks, onLikedChanged }: ITrackList) {
    return <TrackListBase
        tracks={tracks}
        columns={[
            { display: "#", key: "trackNumber", width: 50 },
            { display: "TITLE", key: "title", width: "grow" },
            { display: "", key: "liked", width: 50 },
            { display: "🕒", key: "length", width: 60, centered: true }
        ]}
        initalSortColumn="trackNumber"
        onLikedChanged={onLikedChanged}
    />
}

