export type ILikeStatus = "liked" | "unliked"

type ITrack = {
    readonly id: number
    readonly title: string
    readonly trackNumber: string
    readonly date: string
    readonly length: number
    readonly album: IAlbum
    readonly artists: IArtist[]
    readonly genre: string
    readonly likeStatus: ILikeStatus
}

export type IArtist = {
    readonly id: number
    readonly name: string
    readonly urlName: string
}

export type IAlbum = {
    readonly id: number
    readonly name: string
    readonly urlName: string
}

export default ITrack