module Route exposing (Route(..), parseUrl)

import Url exposing (Url)
import Url.Parser exposing (..)


type Route
    = NotFound
    | HomePage
    | AlbumWithId Int
    | AlbumWithIdAndUrlName Int String
    | ArtistWithId Int
    | ArtistWithIdAndUrlName Int String


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
        , map AlbumWithIdAndUrlName (s "album" </> int </> string)
        , map AlbumWithId (s "album" </> int)
        , map ArtistWithIdAndUrlName (s "artist" </> int </> string)
        , map ArtistWithId (s "artist" </> int)
        ]
