import React, { useEffect, useState } from "react"

import { useControlsContext } from "Context/ControlsContext"
import { useProgressContext } from "Context/ProgressContext"

import { secondsToTimeFormat } from "utilities/secondsToTimeFormat"



function Slider() {
    const { duration, currentTime } = useProgressContext()
    const { seek } = useControlsContext()

    const [sliderValue, setSliderValue] = useState(currentTime)
    const [interacting, setInteracting] = useState(false)

    useEffect(() => {
        if (interacting)
            return

        setSliderValue(currentTime)
    }, [currentTime, interacting])

    return <div style={{ display: "flex", alignItems: "center", flexGrow: 1 }}>
        <div style={{ padding: "0 10px" }}>{secondsToTimeFormat(sliderValue)}</div>
        <input
            style={{ width: "100%" }}
            type="range"
            min={0}
            max={duration}
            value={sliderValue}
            onMouseUp={e => {
                seek(e.currentTarget.valueAsNumber)
                setInteracting(false)
            }}
            onChange={e => {
                setSliderValue(e.currentTarget.valueAsNumber)
                setInteracting(true)
            }}
        />
        <DurationOrRemainingTime
            duration={duration}
            currentTime={sliderValue}
        />
    </div>
}

type IState = "duration" | "timeLeft"

function DurationOrRemainingTime({ duration, currentTime }: { duration: number, currentTime: number }) {
    const [state, setState] = useState<IState>("duration")

    function toggle() {
        const newState = (() => {
            switch (state) {
                case "duration":
                    return "timeLeft"
                case "timeLeft":
                    return "duration"
            }
        })()

        setState(newState)
    }

    const time = (() => {
        switch (state) {
            case "duration":
                return secondsToTimeFormat(duration)
            case "timeLeft":
                return `-${secondsToTimeFormat(duration - currentTime)}`
        }
    })()

    return <div
        style={{ padding: "0 0 0 10px", width: "6ch" }}
        onClick={toggle}
    >
        {time}
    </div>
}

export default Slider