module Page.Album exposing (Model, Msg(..), formatTrackArtists, init, update, view)

import Components.Like as Like
import Components.Table exposing (Align(..), clickableColumn, defaultColumn, textColumn)
import Components.Unlike as Unlike
import Css exposing (Style, absolute, alignItems, auto, backgroundColor, border, borderRadius, center, column, cursor, displayFlex, ellipsis, end, flexDirection, flexGrow, flexShrink, height, hex, hidden, hover, int, justifyContent, marginTop, noWrap, nthChild, overflow, overflowX, overflowY, padding, pct, pointer, position, property, px, rgba, right, row, sticky, textAlign, textOverflow, top, transparent, whiteSpace, width)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Html.Styled.Events exposing (onClick)
import Http
import Json.Decode as Decode exposing (Decoder, bool, list, string)
import Json.Decode.Pipeline exposing (required)
import RemoteData exposing (WebData)
import Types.ReleaseDate exposing (ReleaseDate, formatReleaseDate, releaseDateDecoder)
import Types.TrackId exposing (TrackId)
import Types.TrackInfo
import Utilities.AlbumUrl exposing (albumImageUrl)
import Utilities.ApiBaseUrl exposing (apiBaseUrl)
import Utilities.ArtistUrl exposing (artistUrl)
import Utilities.CssExtensions exposing (gap)
import Utilities.DelayedLoader
import Utilities.DurationDisplay exposing (durationDisplay)
import Utilities.ErrorMessage exposing (errorMessage)
import Utilities.Icon as Icon


type alias Album =
    { id : Int
    , name : String
    , urlName : String
    , tracks : List Track
    , released : ReleaseDate
    , artist : Artist
    }


type alias Seconds =
    Int


type alias Track =
    { id : TrackId
    , trackNumber : Int
    , title : String
    , length : Seconds
    , liked : Bool
    , artists : List Artist
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
        |> required "released" releaseDateDecoder
        |> required "artist" artistDecoder


trackDecoder : Decoder Track
trackDecoder =
    Decode.succeed Track
        |> required "id" Decode.int
        |> required "trackNumber" Decode.int
        |> required "title" string
        |> required "length" Decode.int
        |> required "liked" bool
        |> required "artists" (list artistDecoder)


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
    | LikeTrack TrackId
    | UnlikeTrack TrackId
    | Like Like.Msg
    | Unlike Unlike.Msg
    | PlayTrack Types.TrackInfo.Track
    | PlayAlbum (List Types.TrackInfo.Track)


init : Int -> ( Model, Cmd Msg )
init albumId =
    ( { album = RemoteData.Loading }, getAlbumById albumId )


getAlbumById : Int -> Cmd Msg
getAlbumById albumId =
    Http.get
        { url = apiBaseUrl ++ "/api/v1/album/" ++ String.fromInt albumId
        , expect =
            albumDecoder
                |> Http.expectJson (RemoteData.fromResult >> AlbumReceived)
        }


setTrackLikeStatus : Int -> Bool -> Album -> ( Album, Cmd msg )
setTrackLikeStatus trackId liked album =
    let
        updatedTracks =
            List.map
                (\track ->
                    if track.id == trackId then
                        { track | liked = liked }

                    else
                        track
                )
                album.tracks
    in
    ( { album | tracks = updatedTracks }, Cmd.none )


unlikeTrack : TrackId -> Album -> ( Album, Cmd msg )
unlikeTrack trackId album =
    setTrackLikeStatus trackId False album


likeTrack : TrackId -> Album -> ( Album, Cmd msg )
likeTrack trackId album =
    setTrackLikeStatus trackId True album



-- UPDATE


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        AlbumReceived response ->
            ( { model | album = response }, Cmd.none )

        Like (Like.LikeTrackResponse trackId state) ->
            case state of
                RemoteData.Success _ ->
                    let
                        ( album, cmd ) =
                            RemoteData.update (likeTrack trackId) model.album
                    in
                    ( { model | album = album }, cmd )

                _ ->
                    ( model, Cmd.none )

        Unlike (Unlike.UnlikeTrackResponse trackId state) ->
            case state of
                RemoteData.Success _ ->
                    let
                        ( album, cmd ) =
                            RemoteData.update (unlikeTrack trackId) model.album
                    in
                    ( { model | album = album }, cmd )

                _ ->
                    ( model, Cmd.none )

        LikeTrack trackId ->
            ( model, Cmd.map Like (Like.likeTrackById trackId) )

        UnlikeTrack trackId ->
            ( model, Cmd.map Unlike (Unlike.unlikeTrackById trackId) )

        {- these cases are handled by the update method in Main.elm -}
        PlayTrack _ ->
            ( model, Cmd.none )

        PlayAlbum _ ->
            ( model, Cmd.none )



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
            Utilities.DelayedLoader.default (Css.ms 500)

        RemoteData.Success album ->
            albumView album

        RemoteData.Failure err ->
            errorMessage "Failed to load album" err


picture : List (Attribute msg) -> List (Html msg) -> Html msg
picture =
    node "picture"


playButton : Msg -> Html Msg
playButton msg =
    button
        [ onClick msg
        , css
            [ width (px 48)
            , height (px 48)
            , border (px 0)
            , backgroundColor transparent
            , borderRadius (pct 50)
            , displayFlex
            , justifyContent center
            , alignItems center
            , padding (px 15)
            , backgroundColor (rgba 0 0 0 0.8)
            , hover [ backgroundColor (hex "333") ]
            , cursor pointer
            , position absolute
            ]
        ]
        [ img [ src (Icon.url Icon.Play) ] [] ]


albumView : Album -> Html Msg
albumView album =
    section
        [ css [ displayFlex, flexDirection column, overflow hidden ]
        ]
        [ div [ css [ displayFlex, alignItems end ] ]
            [ picture [ css [ displayFlex, alignItems center, justifyContent center ] ]
                [ img [ css [ property "aspect-ratio" "1/1", width (px 192) ], src (albumImageUrl album) ] []
                , playButton (PlayAlbum (List.map (mapAlbumTrackToTrack album) album.tracks))
                ]
            , div [ css [ Css.paddingLeft (px 10), overflow hidden ] ]
                [ h1 [ css [ whiteSpace noWrap, textOverflow ellipsis, overflowX hidden, overflowY auto ] ] [ text album.name ]
                , div [ css [ displayFlex, flexDirection row, gap (px 10) ] ]
                    [ a [ css [], href (artistUrl album.artist) ] [ text album.artist.name ]
                    , div [] [ text (formatReleaseDate album.released) ]
                    , div [] [ text (formatTracksDisplay album.tracks) ]
                    , div [] [ text (formatAlbumLength album.tracks) ]
                    ]
                ]
            ]
        , div [ css [ marginTop (px 10), displayFlex, flexDirection column, overflow auto ] ]
            [ table album
            ]
        ]


table : Album -> Html Msg
table album =
    Components.Table.table
        [ clickableColumn "#" (.trackNumber >> String.fromInt >> Html.Styled.text) (mapAlbumTrackToTrack album >> PlayTrack)
        , defaultColumn "Title"
            (\t ->
                div [ css [ trackTitleColStyle, displayFlex, flexDirection column ] ]
                    [ div [] [ p [ css [ whiteSpace noWrap, overflow hidden, textOverflow ellipsis ] ] [ text t.title ] ]
                    , div [] (formatTrackArtists t.artists)
                    ]
            )
            |> Components.Table.grow
            |> Components.Table.alignHeader Left
        , clickableColumn ""
            (\t ->
                (if t.liked then
                    "Liked"

                 else
                    "Like"
                )
                    |> Html.Styled.text
            )
            (\t ->
                if t.liked then
                    UnlikeTrack t.id

                else
                    LikeTrack t.id
            )
            |> Components.Table.alignHeader Center
            |> Components.Table.alignData Center
        , textColumn "Duration" (.length >> durationDisplay)
            |> Components.Table.alignHeader Center
            |> Components.Table.alignData Center
        ]
        album.tracks


trackTitleColStyle : Style
trackTitleColStyle =
    Css.batch [ overflow hidden, textOverflow ellipsis ]


formatTrackArtists : List Artist -> List (Html msg)
formatTrackArtists artists =
    artists
        |> List.map (\artist -> a [ href (artistUrl artist) ] [ text artist.name ])
        |> List.intersperse (span [] [ text ", " ])


formatTracksDisplay : List Track -> String
formatTracksDisplay tracks =
    let
        amountOfTracks =
            tracks |> List.length

        suffix =
            if amountOfTracks /= 1 then
                "s"

            else
                ""
    in
    String.fromInt amountOfTracks ++ " track" ++ suffix


formatAlbumLength : List Track -> String
formatAlbumLength tracks =
    tracks
        |> List.map .length
        |> List.foldl (+) 0
        |> durationDisplay


mapAlbumTrackToTrack : Album -> Track -> Types.TrackInfo.Track
mapAlbumTrackToTrack album track =
    { id = track.id
    , title = track.title
    , length = track.length
    , liked = track.liked
    , album =
        { id = album.id
        , name = album.name
        , urlName = album.urlName
        }
    , artists = track.artists
    }
