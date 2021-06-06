import ITrack from "types/track";

export default function useQueue() {
    let queue: ITrack[] = []

    function queueTracks(tracks: ITrack[]) {
        queue.push(...tracks)
    }

    function queueTrack(track: ITrack) {
        queueTracks([track])
    }

    function clearQueue() {
        queue = []
    }

    function removeIndexFromQueue(index: number) {
        queue = queue.filter((_, i) => i !== index)
    }

    function getNextTrackInQueue() {
        const nextTrack = queue.shift()

        return nextTrack
    }

    function getQueueTracks() {
        return [...queue]
    }

    return {
        queue: getQueueTracks(),
        queueTracks,
        queueTrack,
        clearQueue,
        removeTrackFromQueue: removeIndexFromQueue,
        getNextTrackInQueue
    }
}