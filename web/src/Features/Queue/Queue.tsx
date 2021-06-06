import React from "react"
import { usePlayerContext } from "Context"
import TrackList from "TrackList"

export default function QueueView() {
    const { queue } = usePlayerContext()

    return <div>
        <h1>Queue</h1>
        {queue.length === 0
            ? <p>No tracks in queue</p>
            : <TrackList tracks={queue} />
        }
    </div>
}