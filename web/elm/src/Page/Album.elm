module Page.Album exposing (Model, Msg, init, update, view)

import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (href, src)
import Http
import Json.Decode as Decode exposing (Decoder, bool, int, list, string)
import Json.Decode.Pipeline exposing (required)
import RemoteData exposing (WebData)
import Time


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
        |> required "id" int
        |> required "name" string
        |> required "urlName" string
        |> required "tracks" (list trackDecoder)
        |> required "released" string
        |> required "artist" artistDecoder


trackDecoder : Decoder Track
trackDecoder =
    Decode.succeed Track
        |> required "id" int
        |> required "trackNumber" int
        |> required "title" string
        |> required "length" int
        |> required "liked" bool


artistDecoder : Decoder Artist
artistDecoder =
    Decode.succeed Artist
        |> required "id" int
        |> required "name" string
        |> required "urlName" string


type alias Model =
    { album : WebData Album
    }


type Msg
    = FetchAlbum
    | AlbumReceived (WebData Album)


init : Int -> ( Model, Cmd Msg )
init albumId =
    ( { album = RemoteData.Loading }, httpCommand albumId )


httpCommand : Int -> Cmd Msg
httpCommand albumId =
    Http.get
        { url = "http://localhost:42000/api/v1/album/" ++ String.fromInt albumId
        , expect =
            albumDecoder
                |> Http.expectJson (RemoteData.fromResult >> AlbumReceived)
        }



-- UPDATE


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        FetchAlbum ->
            ( { model | album = RemoteData.Loading }, httpCommand 1 )

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

        RemoteData.Failure httpError ->
            text "Failed to load album"


albumView : Album -> Html msg
albumView album =
    section []
        [ img [ src ("http://localhost:42000/api/v1/album/" ++ String.fromInt album.id ++ "/image") ] []
        , div []
            [ h1 [] [ text album.name ]
            , p []
                [ a [ href ("/artist/" ++ String.fromInt album.artist.id ++ "/" ++ album.artist.urlName) ] [ text album.artist.name ]
                , span [] [ text (formatReleaseDate album.released) ]
                , span [] [ text (formatTracksDisplay album.tracks) ]
                , span [] [ text (calculateAlbumLength album.tracks) ]
                ]
            ]
        , table []
            [ thead []
                [ th []
                    [ td [] [ text "#" ]
                    , td [] [ text "TITLE" ]
                    , td [] [ text "DURATION" ]
                    ]
                ]
            , tbody []
                (List.map
                    albumTrackTableRow
                    album.tracks
                )
            ]
        ]


albumTrackTableRow : Track -> Html msg
albumTrackTableRow track =
    tr []
        [ td [] [ text (String.fromInt track.trackNumber) ]
        , td [] [ text track.title ]
        , td [] [ text (String.fromInt track.length) ]
        ]


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


calculateAlbumLength : List Track -> String
calculateAlbumLength tracks =
    let
        totalSeconds =
            List.foldl (\track acc -> acc + track.length) 0 tracks

        minutes =
            totalSeconds // 60

        seconds =
            totalSeconds - minutes * 60
    in
    String.fromInt minutes ++ " min " ++ String.fromInt seconds ++ " sec"


{-| Format release date

Removes time part from date time string

-}
formatReleaseDate : String -> String
formatReleaseDate date =
    date |> String.split "T" |> List.head |> Maybe.withDefault "Unknown release date"
