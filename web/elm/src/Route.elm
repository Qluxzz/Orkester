module Route exposing (Route(..), parseUrl)

import Url exposing (Url)
import Url.Parser exposing (..)


type Route
    = NotFound
    | HomePage
    | Album Int String
    | Artist Int String


parseUrl : Url -> Route
parseUrl url =
    case parse matchRoute url of
        Just route ->
            route

        Nothing ->
            NotFound


matchRoute : Parser (Route -> a) a
matchRoute =
    oneOf
        [ map HomePage top
        , map Album (s "album" </> int </> string)
        , map Artist (s "artist" </> int </> string)
        ]
