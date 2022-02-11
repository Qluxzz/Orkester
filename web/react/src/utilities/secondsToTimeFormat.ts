function padWithTwoLeadingZeroes(value: number): string {
    return value.toString(10).padStart(2, "0")
}

/**
 * Convert seconds to (hours):(minutes):(seconds)
 * @param {Number} value 
 * @returns {string} formatted value
 * @example // returns 22:17
 * secondsToTimeFormat(1337)
 */
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

/**
 * Convert seconds to (hours) hr (minutes?) min (seconds?) sec
 * @param {Number} value
 * @returns {string} formatted value
 * @example // returns 32 min 38 sec
 * secondsToTimeFormat(1958)
 */
export function secondsToHumanReadableFormat(value: number): string {
    const hours = Math.floor(value / 3600)
    value -= hours * 3600
    const minutes = Math.floor(value / 60)
    value -= minutes * 60
    const seconds = Math.round(value)

    if (hours > 24) {
        return "over 24 hr"
    }

    if (hours > 0) {
        return `${hours} hr ${minutes} min`
    }

    if (minutes > 0) {
        return `${minutes} min ${seconds} sec`
    }

    return `${seconds} sec`
}