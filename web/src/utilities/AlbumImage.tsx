interface IAlbumImage {
    album: { id: number, name: string }
    size: string | number | "auto"
}

export default function AlbumImage({ album: { id, name }, size }: IAlbumImage) {
    return <img
        width={size}
        height={size}
        src={`/api/v1/album/${id}/image`}
        alt={`Album cover for album titled ${name}`}
        style={{ aspectRatio: "1 / 1", display: "block" }}
    />
}