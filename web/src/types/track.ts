type ITrack = {
    id: number
    title: string
    trackNumber: string
    date: string
    length: number
    album: IAlbum
    artists: IArtist[]
    genre: string
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