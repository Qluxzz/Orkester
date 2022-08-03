module ReleaseDate exposing (ReleaseDate, formatReleaseDate, releaseDateDecoder)

import Json.Decode as Decode exposing (Decoder, int, string)
import Json.Decode.Pipeline exposing (required)


type ReleaseDate
    = Year Int
    | Month { year : Int, month : Int }
    | Date { year : Int, month : Int, date : Int }


{-| Format release date
Format release date depending on available precision
-}
formatReleaseDate : ReleaseDate -> String
formatReleaseDate d =
    case d of
        Year year ->
            String.padLeft 4 '0' (String.fromInt year)

        Month { year, month } ->
            String.padLeft 4 '0' (String.fromInt year)
                ++ "-"
                ++ String.padLeft 2 '0' (String.fromInt month)

        Date { year, month, date } ->
            String.padLeft 4 '0' (String.fromInt year)
                ++ "-"
                ++ String.padLeft 2 '0' (String.fromInt month)
                ++ "-"
                ++ String.padLeft 2 '0' (String.fromInt date)


monthConstructor : Int -> Int -> ReleaseDate
monthConstructor year month =
    Month { year = year, month = month }


dateConstructor : Int -> Int -> Int -> ReleaseDate
dateConstructor year month date =
    Date { year = year, month = month, date = date }


releaseDateDecoder : Decoder ReleaseDate
releaseDateDecoder =
    Decode.field "precision" string
        |> Decode.andThen
            (\precision ->
                case precision of
                    "year" ->
                        Decode.succeed Year
                            |> required "year" int

                    "month" ->
                        Decode.succeed monthConstructor
                            |> required "year" int
                            |> required "month" int

                    "date" ->
                        Decode.succeed dateConstructor
                            |> required "year" int
                            |> required "month" int
                            |> required "date" int

                    _ ->
                        Decode.fail "Unknown release date precision"
            )
