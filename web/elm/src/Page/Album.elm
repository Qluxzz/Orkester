module Page.Album exposing (Model, Msg(..), formatReleaseDate, getAlbumUrl, init, update, view)

import BaseUrl exposing (baseUrl)
import Css exposing (alignItems, column, displayFlex, end, flexDirection, flexGrow, int, marginLeft, marginRight, marginTop, padding, paddingTop, px, right, textAlign, width)
import ErrorMessage exposing (errorMessage)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Html.Styled.Events exposing (onClick)
import Http
import Json.Decode as Decode exposing (Decoder, bool, list, string)
import Json.Decode.Pipeline exposing (required)
import Like
import Page.Artist exposing (getArtistUrl)
import RemoteData exposing (WebData)
import TrackId exposing (TrackId)
import Unlike


type alias Album =
    { id : Int
    , name : String
    , urlName : String
    , tracks : List Track
    , released : String
    , artist : Artist
    }


type alias Track =
    { id : TrackId
    , trackNumber : Int
    , title : String
    , length : Int
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


albumView : Album -> Html Msg
albumView album =
    section []
        [ div [ css [ displayFlex, alignItems end ] ]
            [ img [ src (baseUrl ++ "/api/v1/album/" ++ String.fromInt album.id ++ "/image") ] []
            , div [ css [ Css.paddingLeft (px 10) ] ]
                [ h1 [] [ text album.name ]
                , div [ css [ displayFlex ] ]
                    [ a [ css [], href ("/artist/" ++ String.fromInt album.artist.id ++ "/" ++ album.artist.urlName) ] [ text album.artist.name ]
                    , div [ css [ marginLeft (px 5) ] ] [ text (formatReleaseDate album.released) ]
                    , div [ css [ marginLeft (px 5) ] ] [ text (formatTracksDisplay album.tracks) ]
                    , div [ css [ marginLeft (px 5) ] ] [ text (formatAlbumLength album.tracks) ]
                    ]
                ]
            ]
        , div [ css [ marginTop (px 10) ] ]
            (table
                album.tracks
            )
        ]


table : List Track -> List (Html Msg)
table tracks =
    tableRow "#" "TITLE" "LIKED" "DURATION"
        :: List.map
            (\track -> trackRow track)
            tracks


trackNumberColStyle : Attribute msg
trackNumberColStyle =
    css [ width (px 50) ]


trackTitleColStyle : Attribute msg
trackTitleColStyle =
    css [ flexGrow (int 1) ]


trackLikedColStyle : Attribute msg
trackLikedColStyle =
    css [ width (px 50) ]


trackDurationColStyle : Attribute msg
trackDurationColStyle =
    css [ width (px 125), textAlign right ]


trackRowStyle : Attribute msg
trackRowStyle =
    css [ displayFlex, paddingTop (px 5) ]


tableRow : String -> String -> String -> String -> Html msg
tableRow col1 col2 _ col4 =
    div [ trackRowStyle ]
        [ div [ trackNumberColStyle ] [ text col1 ]
        , div [ trackTitleColStyle ] [ text col2 ]
        , div [ trackLikedColStyle ] []
        , div [ trackDurationColStyle ] [ text col4 ]
        ]


trackRow : Track -> Html Msg
trackRow track =
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
    div [ trackRowStyle ]
        [ div [ trackNumberColStyle ] [ text (String.fromInt track.trackNumber) ]
        , div [ trackTitleColStyle, css [ displayFlex, flexDirection column ] ]
            [ div [] [ p [] [ text track.title ] ]
            , div [] (formatTrackArtists track.artists)
            ]
        , div [ trackLikedColStyle, onClick onClickLike ] [ text (likedDisplay track.liked) ]
        , div [ trackDurationColStyle ] [ text (formatTrackLength track.length) ]
        ]


formatTrackArtists : List Artist -> List (Html msg)
formatTrackArtists artists =
    let
        amountOfArtists =
            List.length artists

        elements =
            List.concat (List.indexedMap (formatTrackArtist amountOfArtists) artists)
    in
    elements


formatTrackArtist : Int -> Int -> Artist -> List (Html msg)
formatTrackArtist amount index artist =
    let
        spanText =
            if index == amount - 1 then
                ""

            else
                ","
    in
    [ a [ href (getArtistUrl artist) ] [ text artist.name ]
    , span [ css [ marginRight (px 10) ] ] [ text spanText ]
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


getAlbumUrl : Album -> String
getAlbumUrl album =
    "/album/" ++ String.fromInt album.id ++ "/" ++ album.urlName
