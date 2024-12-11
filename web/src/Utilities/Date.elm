module Utilities.Date exposing (Date, dateDecoder)

import Json.Decode as Decode
import Parser exposing ((|.), (|=), Parser, chompUntil, getChompedString, int, succeed, symbol)


type alias Date =
    { year : String
    , month : String
    , date : String
    , hour : String
    , minute : String
    , second : String
    }


date : Parser.Parser Date
date =
    succeed Date
        -- Year
        |= digitParser "-"
        |. symbol "-"
        -- Month
        |= digitParser "-"
        |. symbol "-"
        -- Date
        |= digitParser "T"
        |. symbol "T"
        -- Hour
        |= digitParser ":"
        |. symbol ":"
        -- Minute
        |= digitParser ":"
        |. symbol ":"
        -- Second
        |= digitParser "."


digitParser : String -> Parser.Parser String
digitParser until =
    getChompedString <|
        succeed ()
            |. chompUntil until


dateDecoder : Decode.Decoder Date
dateDecoder =
    Decode.string
        |> Decode.andThen
            (\dateString ->
                case Parser.run date dateString of
                    Ok d ->
                        Decode.succeed d

                    Err err ->
                        Decode.fail ("Failed to decode " ++ dateString ++ " to a date! Error: " ++ Debug.toString err)
            )
