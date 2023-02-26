module Page.Artist exposing (Model, Msg(..), init, update, view)

import AlbumUrl exposing (albumImageUrl, albumUrl)
import ApiBaseUrl exposing (apiBaseUrl)
import Css exposing (Style, absolute, auto, backgroundColor, block, bold, column, display, displayFlex, ellipsis, flex3, flexDirection, fontWeight, height, hex, hidden, left, lineHeight, marginTop, noWrap, overflow, padding, padding4, paddingTop, pct, position, property, px, relative, textOverflow, top, whiteSpace, width)
import ErrorMessage exposing (errorMessage)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Http
import Json.Decode as Decode exposing (Decoder, int, list, string)
import Json.Decode.Pipeline exposing (required)
import ReleaseDate exposing (ReleaseDate(..), releaseDateDecoder)
import RemoteData exposing (WebData)
import Utilities.CssExtensions exposing (gap)
import Utilities.DelayedLoader


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
    , released : ReleaseDate
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
        |> required "released" releaseDateDecoder


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
        { url = apiBaseUrl ++ "/api/v1/artist/" ++ String.fromInt artistId
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
            Utilities.DelayedLoader.default (Css.ms 500)

        RemoteData.Success artist ->
            artistView artist

        RemoteData.Failure error ->
            errorMessage "Failed to load artist" error


artistView : Artist -> Html Msg
artistView artist =
    section [ css [ displayFlex, flexDirection column, overflow hidden ] ]
        [ h1 [] [ text artist.name ]
        , div
            [ css
                [ property "display" "grid"
                , gap (px 24)
                , property "grid-template-columns" "repeat(auto-fill, minmax(256px, 1fr))"
                , property "grid-template-rows" "1fr"
                , overflow auto
                , marginTop (px 10)
                ]
            ]
            (List.map
                albumView
                artist.albums
            )
        ]


pStyle : Style
pStyle =
    Css.batch
        [ whiteSpace noWrap
        , overflow hidden
        , textOverflow ellipsis
        , fontWeight bold
        , padding4 (px 10) (px 0) (px 5) (px 0)
        , lineHeight (px 10)
        ]


albumView : Album -> Html Msg
albumView album =
    a
        [ href (albumUrl album)
        ]
        [ div
            [ css
                [ displayFlex
                , flexDirection column
                , flex3 (Css.int 1) (Css.int 1) (Css.int 0)
                , backgroundColor (hex "#333")
                , padding (px 10)
                , overflow hidden
                ]
            ]
            [ node "picture"
                [ css
                    [ position relative
                    , overflow hidden
                    , paddingTop (pct 100)
                    , height (px 0)
                    ]
                ]
                [ img
                    [ css
                        [ display block
                        , position absolute
                        , top (px 0)
                        , left (px 0)
                        , width (pct 100)
                        , height (pct 100)
                        ]
                    , src (albumImageUrl album)
                    ]
                    []
                ]
            , p [ css [ pStyle ] ] [ text album.name ]
            , p [ css [ pStyle ] ] [ text (releaseYear album.released) ]
            ]
        ]



-- HELPER FUNCTIONS


releaseYear : ReleaseDate -> String
releaseYear releaseDate =
    String.fromInt
        (case releaseDate of
            Year year ->
                year

            Month { year } ->
                year

            Date { year } ->
                year
        )
