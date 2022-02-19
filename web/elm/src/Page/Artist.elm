module Page.Artist exposing (Model, Msg(..), getArtistUrl, init, update, view)

import BaseUrl exposing (baseUrl)
import Css exposing (auto, column, displayFlex, flexDirection, flexWrap, hidden, marginTop, overflow, px, width, wrap)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Http
import Json.Decode as Decode exposing (Decoder, int, list, string)
import Json.Decode.Pipeline exposing (required)
import Page.Album exposing (formatReleaseDate)
import RemoteData exposing (WebData)


type alias Artist =
    { id : Int
    , name : String
    , urlName : String
    , albums : List Album
    }


type alias Album =
    { id : Int
    , name : String
    , urlName : String
    , released : String
    }


artistDecoder : Decoder Artist
artistDecoder =
    Decode.succeed Artist
        |> required "id" int
        |> required "name" string
        |> required "urlName" string
        |> required "albums" (list albumDecoder)


albumDecoder : Decoder Album
albumDecoder =
    Decode.succeed Album
        |> required "id" int
        |> required "name" string
        |> required "urlName" string
        |> required "released" string


type alias Model =
    { artist : WebData Artist
    }


type Msg
    = FetchArtist Int
    | ArtistRecieved (WebData Artist)


init : Int -> ( Model, Cmd Msg )
init artistId =
    ( { artist = RemoteData.Loading }, getArtist artistId )


getArtist : Int -> Cmd Msg
getArtist artistId =
    Http.get
        { url = baseUrl ++ "/api/v1/artist/" ++ String.fromInt artistId
        , expect =
            artistDecoder
                |> Http.expectJson (RemoteData.fromResult >> ArtistRecieved)
        }



-- UPDATE


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        FetchArtist artistId ->
            ( { model | artist = RemoteData.Loading }, getArtist artistId )

        ArtistRecieved artist ->
            ( { model | artist = artist }, Cmd.none )



-- VIEWS


view : Model -> Html Msg
view model =
    artistViewOrError model


artistViewOrError : Model -> Html Msg
artistViewOrError model =
    case model.artist of
        RemoteData.NotAsked ->
            text ""

        RemoteData.Loading ->
            h3 [] [ text "Loading..." ]

        RemoteData.Success artist ->
            artistView artist

        RemoteData.Failure _ ->
            text "Failed to load artist"


artistView : Artist -> Html Msg
artistView artist =
    section [ css [ displayFlex, flexDirection column, overflow hidden ] ]
        [ h1 [] [ text artist.name ]
        , div [ css [ displayFlex, flexWrap wrap, overflow auto, marginTop (px 10) ] ]
            (List.map
                albumView
                artist.albums
            )
        ]


albumView : Album -> Html Msg
albumView album =
    div [ css [ width (px 300) ] ]
        [ a [ href ("/album/" ++ String.fromInt album.id ++ "/" ++ album.urlName) ]
            [ img [ src (baseUrl ++ "/api/v1/album/" ++ String.fromInt album.id ++ "/image") ] []
            , h2 [] [ text album.name ]
            , h3 [] [ text (formatReleaseDate album.released) ]
            ]
        ]



-- HELPER FUNCTIONS


getArtistUrl : Artist -> String
getArtistUrl artist =
    "/artist/" ++ String.fromInt artist.id ++ "/" ++ artist.urlName
