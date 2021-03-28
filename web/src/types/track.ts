type ITrack = {
    id: number
    title: string
    trackNumber: string
    date: string
    length: number
    album: IAlbum
    artist: IArtist
    genre: string
}

type IArtist = {
    id: number
    name: string
    urlName: string
}

type IAlbum = {
    id: number
    name: string
    urlName: string
}

export default ITrack