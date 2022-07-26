module Page.LikedTracks exposing (Model, Msg, init, update, view)

import BaseUrl exposing (baseUrl)
import Css exposing (marginBottom, px)
import ErrorMessage exposing (errorMessage)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css)
import Http
import Json.Decode as Decode exposing (Decoder, bool, list, string)
import Json.Decode.Pipeline exposing (required)
import Page.Album exposing (Msg(..))
import RemoteData exposing (WebData)
import TrackId exposing (TrackId)
import Unlike
import Url exposing (Protocol(..))


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
    { likedTracks : WebData (List Track)
    }


type Msg
    = TracksRecieved (WebData (List Track))
    | UnlikeTrack TrackId
    | Unlike Unlike.Msg


init : ( Model, Cmd Msg )
init =
    ( { likedTracks = RemoteData.Loading }, getLikedTracks )


getLikedTracks : Cmd Msg
getLikedTracks =
    Http.get
        { url = baseUrl ++ "/api/v1/playlist/liked"
        , expect =
            Decode.list trackDecoder
                |> Http.expectJson (RemoteData.fromResult >> TracksRecieved)
        }



-- UPDATE


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        TracksRecieved tracks ->
            ( { model | likedTracks = tracks }, Cmd.none )

        UnlikeTrack trackId ->
            ( model, Cmd.map Unlike (Unlike.unlikeTrackById trackId) )

        Unlike (Unlike.UnlikeTrackResponse trackId state) ->
            case state of
                RemoteData.Success _ ->
                    let
                        ( tracks, cmd ) =
                            RemoteData.update (removeTrackById trackId) model.likedTracks
                    in
                    ( { model | likedTracks = tracks }, cmd )

                _ ->
                    ( model, Cmd.none )


removeTrackById : TrackId -> List Track -> ( List Track, Cmd Msg )
removeTrackById trackId tracks =
    ( List.filter (\{ id } -> id /= trackId) tracks, Cmd.none )



-- VIEW


view : Model -> Html Msg
view model =
    likedTracksViewOrError model


likedTracksViewOrError : Model -> Html Msg
likedTracksViewOrError model =
    case model.likedTracks of
        RemoteData.NotAsked ->
            text ""

        RemoteData.Loading ->
            h3 [] [ text "Loading..." ]

        RemoteData.Success tracks ->
            likedTracksView tracks

        RemoteData.Failure error ->
            errorMessage "Failed to load liked tracks" error


likedTracksView : List Track -> Html Msg
likedTracksView tracks =
    section []
        (h1 [ css [ marginBottom (px 20) ] ] [ text "Liked Tracks" ]
            :: List.map
                (\track -> h1 [] [ text track.title ])
                tracks
        )
