module Page.Album exposing (Model, Msg(..), formatReleaseDate, getAlbumUrl, init, update, view)

import BaseUrl exposing (baseUrl)
import Css exposing (column, displayFlex, flexDirection, flexGrow, int, marginRight, paddingTop, px, right, textAlign, width)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Html.Styled.Events exposing (onClick)
import Http
import Json.Decode as Decode exposing (Decoder, bool, list, string)
import Json.Decode.Pipeline exposing (required)
import Page.Artist exposing (getArtistUrl)
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
    | LikeTrack Int
    | LikeTrackResponse Int (WebData ())
    | UnlikeTrack Int
    | UnlikeTrackResponse Int (WebData ())


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


likeTrackById : Int -> Cmd Msg
likeTrackById trackId =
    Http.request
        { method = "PUT"
        , headers = []
        , url = baseUrl ++ "/api/v1/track/" ++ String.fromInt trackId ++ "/like"
        , body = Http.emptyBody
        , expect = Http.expectWhatever (RemoteData.fromResult >> LikeTrackResponse trackId)
        , timeout = Nothing
        , tracker = Nothing
        }


unlikeTrackById : Int -> Cmd Msg
unlikeTrackById trackId =
    Http.request
        { method = "DELETE"
        , headers = []
        , url = baseUrl ++ "/api/v1/track/" ++ String.fromInt trackId ++ "/like"
        , body = Http.emptyBody
        , expect = Http.expectWhatever (RemoteData.fromResult >> UnlikeTrackResponse trackId)
        , timeout = Nothing
        , tracker = Nothing
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



-- UPDATE


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        AlbumReceived response ->
            ( { model | album = response }, Cmd.none )

        LikeTrackResponse trackId (RemoteData.Success _) ->
            let
                ( album, cmd ) =
                    RemoteData.update (setTrackLikeStatus trackId True) model.album
            in
            ( { model | album = album }, cmd )

        LikeTrackResponse _ _ ->
            ( model, Cmd.none )

        LikeTrack trackId ->
            ( model, likeTrackById trackId )

        UnlikeTrack trackId ->
            ( model, unlikeTrackById trackId )

        UnlikeTrackResponse trackId (RemoteData.Success _) ->
            let
                ( album, cmd ) =
                    RemoteData.update (setTrackLikeStatus trackId False) model.album
            in
            ( { model | album = album }, cmd )

        UnlikeTrackResponse _ _ ->
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

        RemoteData.Failure _ ->
            text "Failed to load album"


albumView : Album -> Html Msg
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
