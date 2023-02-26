module DurationDisplayTests exposing (suite)

import Expect
import Test exposing (..)
import Utilities.DurationDisplay exposing (durationDisplay)


type alias Seconds =
    Int


oneHour : Seconds
oneHour =
    3600


oneMinute : Seconds
oneMinute =
    60


cases : List ( Seconds, String )
cases =
    [ ( 0, "00:00" )
    , ( 30, "00:30" )
    , ( 60, "01:00" )
    , ( 90, "01:30" )
    , ( oneHour + (23 * oneMinute) + 45, "01:23:45" )
    , ( (oneHour * 11) + (22 * oneMinute) + 33, "11:22:33" )
    , ( oneHour * 100, "100:00:00" )
    ]


suite : Test
suite =
    describe "Duration display"
        (List.map
            (\( input, expected ) ->
                test (String.fromInt input ++ " == " ++ expected) <|
                    \_ -> durationDisplay input |> Expect.equal expected
            )
            cases
        )
