module Page.Album exposing (Model, Msg, formatReleaseDate, init, update, view)

import BaseUrl exposing (baseUrl)
import Css exposing (displayFlex, flexGrow, int, paddingTop, px, width)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Http
import Json.Decode as Decode exposing (Decoder, bool, list, string)
import Json.Decode.Pipeline exposing (required)
import RemoteData exposing (WebData)


type alias Album =
    { id : Int
    , name : String
    , urlName : String
    , tracks : List Track
    , released : String
    , artist : Artist
    }


type alias Track =
    { id : Int
    , trackNumber : Int
    , title : String
    , length : Int
    , liked : Bool
    }


type alias Artist =
    { id : Int
    , name : String
    , urlName : String
    }


albumDecoder : Decoder Album
albumDecoder =
    Decode.succeed Album
        |> required "id" Decode.int
        |> required "name" string
        |> required "urlName" string
        |> required "tracks" (list trackDecoder)
        |> required "released" string
        |> required "artist" artistDecoder


trackDecoder : Decoder Track
trackDecoder =
    Decode.succeed Track
        |> required "id" Decode.int
        |> required "trackNumber" Decode.int
        |> required "title" string
        |> required "length" Decode.int
        |> required "liked" bool


artistDecoder : Decoder Artist
artistDecoder =
    Decode.succeed Artist
        |> required "id" Decode.int
        |> required "name" string
        |> required "urlName" string


type alias Model =
    { album : WebData Album
    }


type Msg
    = AlbumReceived (WebData Album)


init : Int -> ( Model, Cmd Msg )
init albumId =
    ( { album = RemoteData.Loading }, getAlbumById albumId )


getAlbumById : Int -> Cmd Msg
getAlbumById albumId =
    Http.get
        { url = baseUrl ++ "/api/v1/album/" ++ String.fromInt albumId
        , expect =
            albumDecoder
                |> Http.expectJson (RemoteData.fromResult >> AlbumReceived)
        }



-- UPDATE


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        AlbumReceived response ->
            ( { model | album = response }, Cmd.none )



-- VIEWS


view : Model -> Html Msg
view model =
    albumViewOrError model


albumViewOrError : Model -> Html Msg
albumViewOrError model =
    case model.album of
        RemoteData.NotAsked ->
            text ""

        RemoteData.Loading ->
            h3 [] [ text "Loading..." ]

        RemoteData.Success album ->
            albumView album

        RemoteData.Failure _ ->
            text "Failed to load album"


albumView : Album -> Html msg
albumView album =
    section []
        [ img [ src (baseUrl ++ "/api/v1/album/" ++ String.fromInt album.id ++ "/image") ] []
        , div []
            [ h1 [] [ text album.name ]
            , div [ css [ displayFlex ] ]
                [ a [ css [ Css.padding2 (px 5) (px 5) ], href ("/artist/" ++ String.fromInt album.artist.id ++ "/" ++ album.artist.urlName) ] [ text album.artist.name ]
                , div [ css [ Css.padding2 (px 5) (px 5) ] ] [ text (formatReleaseDate album.released) ]
                , div [ css [ Css.padding2 (px 5) (px 5) ] ] [ text (formatTracksDisplay album.tracks) ]
                , div [ css [ Css.padding2 (px 5) (px 5) ] ] [ text (formatAlbumLength album.tracks) ]
                ]
            ]
        , div []
            (table
                album.tracks
            )
        ]


table : List Track -> List (Html msg)
table tracks =
    tableRow "#" "TITLE" "LIKED" "DURATION"
        :: List.map
            (\track -> tableRow (String.fromInt track.trackNumber) track.title (likedDisplay track.liked) (formatTrackLength track.length))
            tracks


tableRow : String -> String -> String -> String -> Html msg
tableRow col1 col2 col3 col4 =
    div [ css [ displayFlex, paddingTop (px 5) ] ]
        [ div [ css [ width (px 50) ] ] [ text col1 ]
        , div [ css [ flexGrow (int 1) ] ] [ text col2 ]
        , div [ css [ width (px 50) ] ] [ text col3 ]
        , div [ css [] ] [ text col4 ]
        ]


likedDisplay : Bool -> String
likedDisplay liked =
    if liked then
        "Liked"

    else
        "Like"


formatTracksDisplay : List Track -> String
formatTracksDisplay tracks =
    let
        amountOfTracks =
            tracks |> List.length |> String.fromInt
    in
    if List.length tracks /= 1 then
        amountOfTracks ++ " tracks"

    else
        amountOfTracks ++ " track"


formatAlbumLength : List Track -> String
formatAlbumLength tracks =
    formatTrackLength <| List.foldl (\track acc -> acc + track.length) 0 tracks


{-| Format track length
Returns track length formatted as x min x sec
-}
formatTrackLength : Int -> String
formatTrackLength length =
    let
        minutes =
            length // 60

        seconds =
            length - minutes * 60
    in
    if minutes > 0 then
        String.fromInt minutes ++ " min " ++ String.fromInt seconds ++ " sec"

    else
        String.fromInt seconds ++ " sec"


{-| Format release date

Removes time part from date time string

-}
formatReleaseDate : String -> String
formatReleaseDate date =
    date |> String.split "T" |> List.head |> Maybe.withDefault "Unknown release date"
