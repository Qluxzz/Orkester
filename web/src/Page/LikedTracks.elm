module Page.LikedTracks exposing (Model, Msg, init, update, view)

import Components.Table
import Components.Unlike as Unlike
import Css exposing (..)
import Html.Styled exposing (..)
import Html.Styled.Attributes as HA exposing (css)
import Http
import Json.Decode as Decode exposing (Decoder, bool, list, string)
import Json.Decode.Pipeline
import Page.Album exposing (Msg(..))
import RemoteData exposing (WebData)
import Types.TrackId exposing (TrackId)
import Url exposing (Protocol(..))
import Utilities.AlbumUrl exposing (albumImageUrl)
import Utilities.ApiBaseUrl exposing (apiBaseUrl)
import Utilities.ArtistUrl
import Utilities.CssExtensions exposing (gap)
import Utilities.DelayedLoader
import Utilities.DurationDisplay exposing (durationDisplay)
import Utilities.ErrorMessage exposing (errorMessage)


type alias Track =
    { id : Int
    , trackNumber : Int
    , title : String
    , length : Int
    , liked : Bool
    , artists : List Artist
    , dateAdded : String
    , album : { id : Int }
    }


type alias Artist =
    { id : Int
    , name : String
    , urlName : String
    }


trackDecoder : Decoder Track
trackDecoder =
    Decode.succeed Track
        |> Json.Decode.Pipeline.required "id" Decode.int
        |> Json.Decode.Pipeline.required "trackNumber" Decode.int
        |> Json.Decode.Pipeline.required "title" string
        |> Json.Decode.Pipeline.required "length" Decode.int
        |> Json.Decode.Pipeline.required "liked" bool
        |> Json.Decode.Pipeline.required "artists" (list artistDecoder)
        |> Json.Decode.Pipeline.required "dateAdded" string
        |> Json.Decode.Pipeline.required "album" albumDecoder


artistDecoder : Decoder Artist
artistDecoder =
    Decode.succeed Artist
        |> Json.Decode.Pipeline.required "id" Decode.int
        |> Json.Decode.Pipeline.required "name" string
        |> Json.Decode.Pipeline.required "urlName" string


type alias Album =
    { id : Int }


albumDecoder : Decoder Album
albumDecoder =
    Decode.succeed Album
        |> Json.Decode.Pipeline.required "id" Decode.int


type alias Model =
    { likedTracks : WebData (List Track)
    }


type Msg
    = TracksReceived (WebData (List Track))
    | PlayTrack Track
    | UnlikeTrack TrackId
    | Unlike Unlike.Msg


init : ( Model, Cmd Msg )
init =
    ( { likedTracks = RemoteData.Loading }, getLikedTracks )


getLikedTracks : Cmd Msg
getLikedTracks =
    Http.get
        { url = apiBaseUrl ++ "/api/v1/playlist/liked"
        , expect =
            Decode.list trackDecoder
                |> Http.expectJson (RemoteData.fromResult >> TracksReceived)
        }



-- UPDATE


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        TracksReceived tracks ->
            ( { model | likedTracks = tracks }, Cmd.none )

        UnlikeTrack trackId ->
            ( model, Cmd.map Unlike (Unlike.unlikeTrackById trackId) )

        {- these cases are handled by the update method in Main.elm -}
        PlayTrack _ ->
            ( model, Cmd.none )

        Unlike (Unlike.UnlikeTrackResponse trackId state) ->
            case state of
                RemoteData.Success _ ->
                    let
                        ( tracks, cmd ) =
                            RemoteData.update (\t -> ( removeTrackById trackId t, Cmd.none )) model.likedTracks
                    in
                    ( { model | likedTracks = tracks }, cmd )

                _ ->
                    ( model, Cmd.none )


removeTrackById : TrackId -> List Track -> List Track
removeTrackById trackId tracks =
    List.filter (\{ id } -> id /= trackId) tracks



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
            Utilities.DelayedLoader.default (ms 500)

        RemoteData.Success tracks ->
            likedTracksView tracks

        RemoteData.Failure error ->
            errorMessage "Failed to load liked tracks" error


likedTracksView : List Track -> Html Msg
likedTracksView tracks =
    section [ css [ displayFlex, flexDirection column, overflow hidden, gap (px 10), textOverflow ellipsis ] ]
        (h1 [ css [ marginBottom (px 20) ] ] [ text "Liked Tracks" ]
            :: (if List.length tracks == 0 then
                    [ text "You haven't liked any tracks yet!" ]

                else
                    [ Components.Table.table
                        [ Components.Table.clickableColumn "#" (\( i, _ ) -> String.fromInt (i + 1) |> Html.Styled.text) (\( _, t ) -> PlayTrack t)
                        , Components.Table.defaultColumn "TITLE"
                            (\( _, t ) ->
                                div [ css [ displayFlex, flexDirection row ] ]
                                    [ picture [ css [ displayFlex, alignItems center, justifyContent center ] ]
                                        [ img [ css [ property "aspect-ratio" "1/1", width (px 32) ], HA.src (albumImageUrl t.album) ] []
                                        ]
                                    , div [ css [ displayFlex, flexDirection column ] ]
                                        [ div [] [ p [ css [ whiteSpace noWrap, overflow hidden, textOverflow ellipsis ] ] [ text t.title ] ]
                                        , div [] (formatTrackArtists t.artists)
                                        ]
                                    ]
                            )
                            |> Components.Table.grow
                            |> Components.Table.alignHeader Components.Table.Left
                        , Components.Table.textColumn "Added" (\( _, t ) -> formatDate t.dateAdded)
                            |> Components.Table.alignHeader Components.Table.Center
                            |> Components.Table.alignData Components.Table.Center
                        , Components.Table.clickableColumn "Liked" (\_ -> Html.Styled.text "Liked") (\( _, t ) -> UnlikeTrack t.id)
                        , Components.Table.textColumn "Duration" (Tuple.second >> .length >> durationDisplay)
                            |> Components.Table.alignHeader Components.Table.Center
                            |> Components.Table.alignData Components.Table.Center
                        ]
                        (List.indexedMap Tuple.pair tracks)
                    ]
               )
        )


picture : List (Attribute msg) -> List (Html msg) -> Html msg
picture =
    node "picture"


formatTrackArtists : List Artist -> List (Html msg)
formatTrackArtists artists =
    artists
        |> List.map (\artist -> a [ HA.href (Utilities.ArtistUrl.artistUrl artist) ] [ text artist.name ])
        |> List.intersperse (span [] [ text ", " ])


formatDate : String -> String
formatDate dateString =
    case String.split "T" dateString of
        [ date, _ ] ->
            date

        _ ->
            "Unknown date"
