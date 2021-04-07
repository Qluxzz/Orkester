import fc from "fast-check"

import { secondsToTimeFormat } from "./Album"

function padToTwoZeroes(input: number): string {
    return input.toString().padStart(2, "0")
}

describe("padToTwoZeroes", () => {
    it("should pad with zeroes if less than 10", () => fc.assert(
        fc.property(
            fc.integer(0, 10),
            number => padToTwoZeroes(number).length === 2
        )
    ))
})

describe("secondsToTimeFormat", () => {
    it("Returns 00:(seconds) if below one minute", () => {
        expect(
            secondsToTimeFormat(59)
        ).toBe("00:59")
    })

    it("Returns 00:01 if one seconds", () => expect(secondsToTimeFormat(1)).toBe("00:01"))

    it("Returns (minute):(seconds) if above one minute and below an hour", () =>
        fc.assert(
            fc.property(
                fc.integer(0, 59),
                fc.integer(0, 59),
                (minutes, seconds) =>
                    `${padToTwoZeroes(minutes)}:${padToTwoZeroes(seconds)}` === secondsToTimeFormat((minutes * 60) + seconds)
            )
        )
    )

    it("Returns (hours):(minutes):(seconds)", () => {
        fc.assert(
            fc.property(
                fc.integer(1, 1_000_000),
                fc.integer(0, 59),
                fc.integer(0, 59),
                (hours, minutes, seconds) =>
                    `${padToTwoZeroes(hours)}:${padToTwoZeroes(minutes)}:${padToTwoZeroes(seconds)}` === secondsToTimeFormat((hours * 3600) + (minutes * 60) + seconds)
            )
        )
    })
})