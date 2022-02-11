import React from "react"
import { useControlsContext } from "Context/ControlsContext"

import Slider from "./Slider"

function Controls() {
    const { play, pause, playbackState } = useControlsContext()

    return <div>
        <div>
            <button
                onClick={playbackState === "paused"
                    ? play
                    : pause
                }
                style={{ width: 72 }}
            >
                {playbackState === "paused" ? "play" : "pause"}
            </button>
        </div>
        <div>
            <Slider />
        </div>
    </div>
}

export default Controls