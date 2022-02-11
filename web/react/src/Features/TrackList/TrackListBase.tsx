import React from "react"
import styled from "styled-components"
import ITrack from "types/track"
import getColumnStyle from "./getColumnStyle"
import { TrackRow } from "./TrackRow"



type IColumnKey = ITrackKeys | "albumCover" | "index"

export interface IColumn {
    display: string
    key: IColumnKey
    width?: number | "grow"
    centered?: boolean
}

interface ITrackListBase {
    tracks: ITrack[]
    columns: IColumn[]
    onLikedChanged?: (liked: boolean, trackId: number) => void
}

type ITrackKeys = keyof ITrack


export default function TrackListBase({
    tracks,
    columns,
    onLikedChanged
}: ITrackListBase) {
    return <div>
        <header style={{ display: "flex", padding: "10px" }}>
            {columns.map(column => {
                return <div
                    style={getColumnStyle(column)}
                    key={column.key}
                >
                    {column.display}
                </div>
            })}
        </header>
        <StyledList>
            {tracks.map((track, index) =>
                <TrackRow
                    key={track.id}
                    track={track}
                    columns={columns}
                    onLikedChanged={onLikedChanged}
                    index={index}
                />
            )}
        </StyledList>
    </div>
}

const StyledList = styled.ul`
    padding: 0;
    margin: 0;

    li:nth-child(odd) {
        background-color: #333;
    }
`