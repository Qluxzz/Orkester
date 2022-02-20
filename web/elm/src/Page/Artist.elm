module Page.Artist exposing (Model, Msg(..), getArtistUrl, init, update, view)

import BaseUrl exposing (baseUrl)
import Css exposing (alignItems, alignSelf, auto, backgroundColor, center, column, content, displayFlex, ellipsis, flex, flexDirection, flexGrow, flexWrap, hex, hidden, justifyContent, margin, marginTop, noWrap, overflow, padding, px, rgb, start, textOverflow, whiteSpace, width, wrap)
import Css.Transitions exposing (gridGap)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Http
import Json.Decode as Decode exposing (Decoder, int, list, string)
import Json.Decode.Pipeline exposing (required)
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
    a [ href ("/album/" ++ String.fromInt album.id ++ "/" ++ album.urlName) ]
        [ div
            [ css
                [ width (px 128)
                , overflow hidden
                , textOverflow ellipsis
                , padding (px 10)
                , backgroundColor (hex "#333")
                , margin (px 10)
                , flex content
                ]
            ]
            [ div [ css [ displayFlex, justifyContent center ] ]
                [ img [ src (baseUrl ++ "/api/v1/album/" ++ String.fromInt album.id ++ "/image") ] []
                ]
            , p [ css [ overflow hidden, textOverflow ellipsis, whiteSpace noWrap, alignSelf start, marginTop (px 5) ] ] [ text album.name ]
            , p [] [ text (formatReleaseDate album.released) ]
            ]
        ]



-- HELPER FUNCTIONS


formatReleaseDate : String -> String
formatReleaseDate releseDate =
    String.split "-" releseDate |> List.head |> Maybe.withDefault "XXXX"


getArtistUrl : Artist -> String
getArtistUrl artist =
    "/artist/" ++ String.fromInt artist.id ++ "/" ++ artist.urlName
