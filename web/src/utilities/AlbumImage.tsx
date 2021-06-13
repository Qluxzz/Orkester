interface IAlbumImage {
    album: { id: number, name: string }
    size?: string | number
}

export default function AlbumImage({ album: { id, name }, size }: IAlbumImage) {
    return <img width={size} height={size} src={`/api/v1/album/${id}/image`} alt={name} />
}