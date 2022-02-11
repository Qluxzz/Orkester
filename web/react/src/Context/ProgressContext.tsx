import React, { useContext } from "react"

interface IProgress {
    duration: number
    currentTime: number
}

const ProgressContext = React.createContext<IProgress>({} as IProgress)

export function useProgressContext() {
    const value = useContext(ProgressContext)

    if (!value)
        throw new Error("Tried to access ProgressContext outside of provider")

    return value
}

export const ProgressContextProvider = ProgressContext.Provider