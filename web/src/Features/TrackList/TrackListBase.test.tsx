import React from "react"
import { render, screen } from "@testing-library/react"

import TrackListBase from "Features/TrackList/TrackListBase"

test("returns a list with one column", () => {
    render(<TrackListBase
        columns={
            [{
                display: "Test",
                key: "album"
            }]
        }
        tracks={[]}
    />)

    expect(screen.getByText("Test")).toBeTruthy()
})