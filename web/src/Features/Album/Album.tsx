import { useEffect, useState } from "react";
import styled from "styled-components";

import ITrack from "types/track";
import { secondsToTimeFormat } from "Utilities/secondsToTimeFormat";
import { usePlayerContext } from "Context";
import CenteredDotLoader from "CenteredDotLoader";

interface IAlbum {
    name: string
    urlName: string
    tracks: ITrack[]
}

export function GetAlbumWithId({ id }: { id: number }) {
    const [album, setAlbum] = useState<IAlbum>()
    const history = useHistory()

    useEffect(() => {
        let isCanceled = false

        fetchAlbumInfo(id)
            .then(album => {
                if (isCanceled)
                    return

                setAlbum(album)

                history.replace(`/album/${id}/${album.urlName}`)

            })
            .catch(error => {
                console.error("Failed to get album info!", error)
            })

        return () => { isCanceled = true }
    }, [id, history])

    if (!album)
        return <CenteredDotLoader />

    return <AlbumView {...album} />
}


const LinkButton = styled.button`
    background: none;
    color: white;
    border: none;
    text-decoration: underline;

    :hover {
        cursor: pointer;
    }
`

const TableStyle = styled.table`
    border: 0px;
    * {
        font-size: 16px;
    }
`

const TableData = styled.td<{ align?: "left" | "center" | "right" }>`
    padding: 5px;
    border: none;
    text-align: ${props => props.align}
`

const TableRow = styled.tr<{ striped?: boolean }>`
    margin: 0 5px;
    border: none;
    background : ${props => props.striped ? "#333" : "#444"}
`

const AlbumViewContainer = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
`


function AlbumView({ name, tracks }: IAlbum) {
    const { play } = usePlayerContext()

    return <AlbumViewContainer>
        {name}
        <TableStyle>
            <thead>
                <TableRow>
                    <th>#</th>
                    <th>Name</th>
                    <th>Length</th>
                </TableRow>
            </thead>
            <tbody>
                {tracks.map((track, i) =>
                    <TableRow key={i} striped={i % 2 === 0}>
                        <TableData>{track.trackNumber}</TableData>
                        <TableData>
                            <LinkButton
                                type="button"
                                onClick={() => play(track.id)}
                            >
                                {track.title}
                            </LinkButton>
                        </TableData>
                        <TableData align="center">{secondsToTimeFormat(track.length)}</TableData>
                    </TableRow>
                )}
            </tbody>
        </TableStyle>
    </AlbumViewContainer>
}

async function fetchAlbumInfo(id: number): Promise<IAlbum> {
    const response = await fetch(`/api/v1/album/${id}`)

    if (!response.ok)
        throw new Error(`Http request failed with status ${response.status}`)

    return await response.json()
}