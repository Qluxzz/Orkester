import React from "react";
import { IArtist } from "types/track";
import { ArtistLink } from "./Links";

export default function ArtistList({ artists }: { artists: IArtist[] }) {
    return <>
        {artists.map((artist, i, arr) =>
            <React.Fragment key={artist.id}>
                <ArtistLink {...artist} >
                    {artist.name}
                </ArtistLink>
                {i !== arr.length - 1 && ", "}
            </React.Fragment>)
        }
    </>
}