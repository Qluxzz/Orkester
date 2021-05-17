export type ILikeStatus = "liked" | "notliked"

type ITrack = {
    id: number
    title: string
    trackNumber: string
    date: string
    length: number
    album: IAlbum
    artists: IArtist[]
    genre: string
    likeStatus: ILikeStatus
}

export type IArtist = {
    id: number
    name: string
    urlName: string
}

export type IAlbum = {
    id: number
    name: string
    urlName: string
}

export default ITrack