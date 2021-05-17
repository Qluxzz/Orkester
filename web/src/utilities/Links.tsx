import { Link } from "react-router-dom";

interface ILink {
    id: number
    urlName: string
    children: React.ReactNode
}

type ILinkType = "artist" | "album"

export function ArtistLink(props: ILink) {
    return createLink("artist", props)
}

export function AlbumLink(props: ILink) {
    return createLink("album", props)
}

function createLink(type: ILinkType, props: ILink) {
    return <Link to={`/${type}/${props.id}/${props.urlName}`}>{props.children}</Link>
}