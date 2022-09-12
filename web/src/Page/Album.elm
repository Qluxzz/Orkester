module Page.Album exposing (Model, Msg(..), formatTrackArtists, init, update, view)

import AlbumUrl exposing (albumImageUrl)
import ApiBaseUrl exposing (apiBaseUrl)
import ArtistUrl exposing (artistUrl)
import Css exposing (Style, absolute, alignItems, auto, backgroundColor, center, column, cursor, displayFlex, ellipsis, end, flexDirection, flexGrow, flexShrink, hex, hidden, int, justifyContent, marginTop, noWrap, nthChild, overflow, overflowX, overflowY, padding, pointer, position, property, px, right, row, sticky, textAlign, textOverflow, top, whiteSpace, width)
import DurationDisplay exposing (durationDisplay)
import ErrorMessage exposing (errorMessage)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Html.Styled.Events exposing (onClick)
import Http
import Json.Decode as Decode exposing (Decoder, bool, list, string)
import Json.Decode.Pipeline exposing (required)
import Like
import ReleaseDate exposing (ReleaseDate, formatReleaseDate, releaseDateDecoder)
import RemoteData exposing (WebData)
import TrackId exposing (TrackId)
import TrackInfo
import Unlike


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
    | PlayTrack TrackInfo.Track
    | PlayAlbum (List TrackInfo.Track)


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
            h3 [] [ text "Loading..." ]

        RemoteData.Success album ->
            albumView album

        RemoteData.Failure err ->
            errorMessage "Failed to load album" err


picture : List (Attribute msg) -> List (Html msg) -> Html msg
picture =
    node "picture"


albumView : Album -> Html Msg
albumView album =
    section
        [ css [ displayFlex, flexDirection column, overflow hidden ]
        ]
        [ div [ css [ displayFlex, alignItems end ] ]
            [ picture [ css [ displayFlex, alignItems center, justifyContent center ] ]
                [ img [ css [ property "aspect-ratio" "1/1", width (px 192) ], src (albumImageUrl album) ] []
                , button [ css [ position absolute, padding (px 10) ], onClick (PlayAlbum (List.map (mapAlbumTrackToTrack album) album.tracks)) ] [ text "Play" ]
                ]
            , div [ css [ Css.paddingLeft (px 10), overflow hidden ] ]
                [ h1 [ css [ whiteSpace noWrap, textOverflow ellipsis, overflowX hidden, overflowY auto ] ] [ text album.name ]
                , div [ css [ displayFlex, flexDirection row, property "gap" "10px" ] ]
                    [ a [ css [], href (artistUrl album.artist) ] [ text album.artist.name ]
                    , div [] [ text (formatReleaseDate album.released) ]
                    , div [] [ text (formatTracksDisplay album.tracks) ]
                    , div [] [ text (formatAlbumLength album.tracks) ]
                    ]
                ]
            ]
        , div [ css [ marginTop (px 10), displayFlex, flexDirection column, overflow auto ] ]
            (table
                album
            )
        ]


table : Album -> List (Html Msg)
table album =
    tableHeaderRow "#" "TITLE" "LIKED" "DURATION"
        :: List.map
            (trackRow
                album
            )
            album.tracks


trackNumberColStyle : Style
trackNumberColStyle =
    Css.batch [ width (px 30), cursor pointer, flexShrink (int 0) ]


trackTitleColStyle : Style
trackTitleColStyle =
    Css.batch [ flexGrow (int 1), overflow hidden, textOverflow ellipsis ]


trackLikedColStyle : Style
trackLikedColStyle =
    Css.batch [ width (px 50), flexShrink (int 0) ]


trackDurationColStyle : Style
trackDurationColStyle =
    Css.batch [ width (px 85), textAlign right, flexShrink (int 0) ]


trackRowStyle : Style
trackRowStyle =
    Css.batch
        [ displayFlex
        , property "gap" "10px"
        , padding (px 10)
        , nthChild "even"
            [ backgroundColor (hex "#333") ]
        , nthChild "odd"
            [ backgroundColor (hex "#222") ]
        ]


tableHeaderRow : String -> String -> String -> String -> Html msg
tableHeaderRow col1 col2 _ col4 =
    div [ css [ trackRowStyle, position sticky, top (px 0) ] ]
        [ div [ css [ trackNumberColStyle ] ] [ text col1 ]
        , div [ css [ trackTitleColStyle ] ] [ text col2 ]
        , div [ css [ trackLikedColStyle ] ] []
        , div [ css [ trackDurationColStyle ] ] [ text col4 ]
        ]


trackRow : Album -> Track -> Html Msg
trackRow album track =
    let
        onClickLike =
            if track.liked then
                UnlikeTrack track.id

            else
                LikeTrack track.id

        likedDisplay liked =
            if liked then
                "Liked"

            else
                "Like"
    in
    div [ css [ trackRowStyle ] ]
        [ div [ css [ trackNumberColStyle ], onClick (PlayTrack (mapAlbumTrackToTrack album track)) ] [ text (String.fromInt track.trackNumber) ]
        , div [ css [ trackTitleColStyle, displayFlex, flexDirection column ] ]
            [ div [] [ p [ css [ whiteSpace noWrap, overflow hidden, textOverflow ellipsis ] ] [ text track.title ] ]
            , div [] (formatTrackArtists track.artists)
            ]
        , div [ css [ trackLikedColStyle ], onClick onClickLike ] [ text (likedDisplay track.liked) ]
        , div [ css [ trackDurationColStyle ] ] [ text (durationDisplay track.length) ]
        ]


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


mapAlbumTrackToTrack : Album -> Track -> TrackInfo.Track
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
