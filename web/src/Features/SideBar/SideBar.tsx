import React from "react"
import { useTrackContext } from "Context/TrackContext"
import styled from "styled-components"

import { AlbumLink } from "utilities/Links"
import { Link } from "react-router-dom"
import AlbumImageWithToggle from "Features/AlbumImageWithToggle"

const Container = styled.div`
    display: flex;
    flex-direction: column;
    position: relative;
    background: #333;
    margin: 0;
    overflow: hidden;

    width: 200px;
`

const List = styled.ul`
    margin: 0;
    padding: 0;
    list-style: none;

    flex-grow: 1;

    li {
        padding: 10px 5px;
        border: 1px solid grey;
    }
`

export interface IAlbumCoverSettings {
    showAlbumCover: boolean
    hideAlbumCover: () => void
}


export default function SideBar(props: IAlbumCoverSettings) {
    const albumArtSize = 220
    const track = useTrackContext()

    return <Container>
        <List>
            <Link to="/collection/tracks"><li>Liked tracks</li></Link>
        </List>
        {track &&
            <div
                style={{
                    transition: "all 250ms",
                    transform: `translateY(${props.showAlbumCover ? 0 : albumArtSize}px)`,
                }}
            >
                <AlbumLink {...track.album}>
                    <AlbumImageWithToggle
                        album={track.album}
                        onClick={props.hideAlbumCover}
                        size={albumArtSize}
                        icon="⬇️"
                    />
                </AlbumLink>
            </div>
        }
    </Container>
}

