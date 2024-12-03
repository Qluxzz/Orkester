module Utilities.DurationDisplay exposing (durationDisplay)


type alias Seconds =
    Int


{-| durationDisplay
Returns seconds formatted as hour:min:sec
-}
durationDisplay : Seconds -> String
durationDisplay length =
    let
        oneHour : Seconds
        oneHour =
            3600

        oneMinute : Seconds
        oneMinute =
            60

        hours =
            length // oneHour

        minutes =
            (length - (hours * oneHour)) // oneMinute

        seconds =
            length - (hours * oneHour) - (minutes * oneMinute)

        padTime : Int -> String
        padTime time =
            String.padLeft 2 '0' (String.fromInt time)
    in
    (if hours > 0 then
        padTime hours ++ ":"

     else
        ""
    )
        ++ padTime minutes
        ++ ":"
        ++ padTime seconds
