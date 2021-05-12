function padWithTwoLeadingZeroes(value: number): string {
    return value.toString(10).padStart(2, "0")
}

// Convert seconds to (hours):(minutes):(seconds)
// Example: secondsToTimeFormat(1337) => 22:17
export function secondsToTimeFormat(value: number): string {
    const hours = Math.floor(value / 3600)
    value -= hours * 3600
    const minutes = Math.floor(value / 60)
    value -= minutes * 60
    const seconds = Math.round(value)

    const parts = []

    if (hours > 0)
        parts.push(padWithTwoLeadingZeroes(hours))

    parts.push(padWithTwoLeadingZeroes(minutes))
    parts.push(padWithTwoLeadingZeroes(seconds))

    return parts.join(":")
}