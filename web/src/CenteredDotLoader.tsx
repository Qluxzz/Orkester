import React from "react";
import { DotLoader } from "react-spinners";

export default function CenteredDotLoader() {
    return <div style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        flex: "1 1 0"
    }}>
        <DotLoader color="white" />
    </div>
}