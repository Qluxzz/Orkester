import React, { useEffect, useState } from "react";
import { DotLoader } from "react-spinners";

export default function CenteredDotLoader() {
    const [show, setShow] = useState(false)
    const showLoaderTimeout = 500

    useEffect(() => {
        let isCanceled = false

        const timeoutId = setTimeout(() => {
            if (isCanceled) {
                clearTimeout(timeoutId)
                return
            }

            setShow(true)
        }, showLoaderTimeout)

        return () => { isCanceled = true }
    }, [])

    if (!show)
        return null

    return <div style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        flex: "1 1 0"
    }}>
        <DotLoader color="white" />
    </div>
}