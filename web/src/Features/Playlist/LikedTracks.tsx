import CenteredDotLoader from "CenteredDotLoader";
import { useEffect, useState } from "react";
import TrackListBase from "Features/TrackList/TrackListBase";
import ITrack from "types/track";

async function getLikedTracks(): Promise<ITrack[]> {
    const response = await fetch("/api/v1/playlist/liked")

    if (!response.ok)
        throw new Error(`Http request failed with status, ${response.status}`)

    return await response.json()
}

export default function LikedTracks() {
    const [tracks, setTracks] = useState<ITrack[] | "loading">("loading")

    function removeLikedTrack(trackId: number) {
        if (tracks === "loading")
            return

        setTracks(tracks.filter(x => x.id !== trackId))
    }

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

    if (tracks === "loading")
        return <CenteredDotLoader />

    return <div>
        <h1>Liked tracks</h1>
        {tracks.length === 0
            ? <p>You have no liked tracks</p>
            : <TrackListBase
                tracks={tracks}
                onLikedChanged={(liked, trackId) => {
                    if (!liked)
                        removeLikedTrack(trackId)
                }}
                columns={[
                    { display: "#", key: "index", width: 50 },
                    { display: "", key: "albumCover" },
                    { display: "TITLE", key: "title", width: 400 },
                    { display: "ALBUM", key: "album", width: "grow" },
                    { display: "ADDED", key: "date", width: 100, centered: true },
                    { display: "", key: "liked", width: 100, centered: true },
                    { display: "🕒", key: "length", width: 60, centered: true }
                ]}
            />
        }
    </div>
}