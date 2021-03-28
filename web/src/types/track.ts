type ITrack = {
    Id: number
    Title: string
    TrackNumber: string
    Date: string
    Album: IAlbum
    Artist: IArtist
    Genre: string
}

type IArtist = {
    Id: number
    Name: string
    UrlName: string
}

type IAlbum = {
    Id: number
    Name: string
    UrlName: string
}

export default ITrack