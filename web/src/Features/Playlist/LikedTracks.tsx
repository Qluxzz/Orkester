import CenteredDotLoader from "CenteredDotLoader";
import { useEffect, useState } from "react";
import TrackList from "TrackList";
import ITrack from "types/track";

async function getLikedTracks(): Promise<ITrack[]> {
    const response = await fetch("/api/v1/playlist/liked")

    if (!response.ok)
        throw new Error(`Http request failed with status, ${response.status}`)

    return await response.json()
}

export default function LikedTracks() {
    const [tracks, setTracks] = useState<ITrack[]>()

    useEffect(() => {
        let isCanceled = false

        getLikedTracks()
            .then(tracks => {
                if (isCanceled)
                    return

                setTracks(tracks)
            })
            .catch(error => {
                console.error("Failed to get liked tracks", error)
            })
    }, [])

    if (!tracks)
        return <CenteredDotLoader />

    return <div>
        <h1>Liked tracks</h1>
        {tracks.length === 0
            ? <p>You have no liked tracks</p>
            : <TrackList
                tracks={tracks}
                onLikeStatusChanged={(status, trackId) => {
                    if (status === "notliked")
                        setTracks(tracks.filter(x => x.id !== trackId))
                }}
                showAlbumCover
            />
        }
    </div>
}