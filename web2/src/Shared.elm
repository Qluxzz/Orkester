module Shared exposing
    ( Flags, decoder
    , Model, Msg
    , init, update, subscriptions
    )

{-|

@docs Flags, decoder
@docs Model, Msg
@docs init, update, subscriptions

-}

import Effect exposing (Effect)
import JSPlayer
import Json.Decode
import Json.Decode.Pipeline
import Route exposing (Route)
import Route.Path
import Shared.Model
import Shared.Msg
import Types.TrackQueue as TrackQueue



-- FLAGS


type alias Flags =
    { volume : Int, repeat : TrackQueue.Repeat }


decoder : Json.Decode.Decoder Flags
decoder =
    Json.Decode.map2 Flags
        (Json.Decode.field "volume" (Json.Decode.maybe Json.Decode.int) |> Json.Decode.map (Maybe.withDefault 100))
        (Json.Decode.field "repeat" (Json.Decode.maybe Json.Decode.string) |> Json.Decode.map repeatDecoder)


repeatDecoder : Maybe String -> TrackQueue.Repeat
repeatDecoder option =
    case option of
        Just "RepeatAll" ->
            TrackQueue.RepeatAll

        Just "RepeatOff" ->
            TrackQueue.RepeatOff

        Just "RepeatOne" ->
            TrackQueue.RepeatOne

        _ ->
            TrackQueue.RepeatOff



-- INIT


type alias Model =
    Shared.Model.Model


init : Result Json.Decode.Error Flags -> Route () -> ( Model, Effect Msg )
init flagsResult route =
    let
        baseModel =
            { queue = TrackQueue.empty, volume = 100, repeat = TrackQueue.RepeatOff, onPreviousBehavior = Shared.Model.PlayPreviousTrack }
    in
    ( case flagsResult of
        Ok f ->
            { baseModel
                | volume = f.volume
                , repeat = f.repeat
            }

        Err _ ->
            baseModel
    , Effect.none
    )



-- UPDATE


type alias Msg =
    Shared.Msg.Msg


update : Route () -> Msg -> Model -> ( Model, Effect Msg )
update route msg model =
    case msg of
        Shared.Msg.NoOp ->
            ( model
            , Effect.none
            )

        Shared.Msg.JSPlayer msg_ ->
            case msg_ of
                JSPlayer.PlaybackFailed error ->
                    Debug.todo ("Playback failed " ++ error)

                JSPlayer.ProgressUpdated progress ->
                    ( model, Effect.none )

                JSPlayer.Seek _ ->
                    ( model, Effect.none )

                JSPlayer.ExternalStateChange state ->
                    case state of
                        "play" ->
                            ( { model | queue = TrackQueue.updateActiveTrackState model.queue TrackQueue.Playing }, Effect.none )

                        "pause" ->
                            ( { model | queue = TrackQueue.updateActiveTrackState model.queue TrackQueue.Paused }, Effect.none )

                        "ended" ->
                            let
                                updatedQueue =
                                    TrackQueue.next model.queue model.repeat

                                effect =
                                    (case model.repeat of
                                        TrackQueue.RepeatOne ->
                                            Just Effect.play

                                        _ ->
                                            TrackQueue.getActiveTrack updatedQueue
                                                |> Maybe.map (\{ track } -> Effect.playTrack track.id)
                                    )
                                        |> Maybe.withDefault Effect.none
                            in
                            ( { model | queue = updatedQueue }, effect )

                        "nexttrack" ->
                            playNext model

                        "previoustrack" ->
                            playPrevious model

                        _ ->
                            Debug.todo ("unknown state change " ++ state)



-- SUBSCRIPTIONS


subscriptions : Route () -> Model -> Sub Msg
subscriptions route model =
    Sub.batch
        [ Sub.map Shared.Msg.JSPlayer (JSPlayer.playbackFailed JSPlayer.PlaybackFailed)
        , Sub.map Shared.Msg.JSPlayer (JSPlayer.progressUpdated JSPlayer.ProgressUpdated)
        , Sub.map Shared.Msg.JSPlayer (JSPlayer.stateChange JSPlayer.ExternalStateChange)
        ]



-- HELPER FUNCTIONS


playNext : Model -> ( Model, Effect msg )
playNext model =
    let
        updatedQueue =
            TrackQueue.next model.queue model.repeat

        cmd =
            TrackQueue.getActiveTrack updatedQueue
                |> Maybe.map (\{ track } -> Effect.playTrack track.id)
                |> Maybe.withDefault Effect.pause
    in
    ( { model | queue = updatedQueue }, cmd )


{-|

    Plays previous if progress on current track
    is less than threshold, otherwise it restarts the current track
    and if pressed again, it jumps to the previous track

-}
playPrevious : Model -> ( Model, Effect Msg )
playPrevious model =
    let
        prev : ( Model, Effect Msg )
        prev =
            let
                updatedQueue =
                    TrackQueue.previous model.queue

                current =
                    TrackQueue.getActiveTrack updatedQueue

                cmd : Effect Msg
                cmd =
                    current
                        |> Maybe.map (\{ track } -> Effect.playTrack track.id)
                        |> Maybe.withDefault Effect.none
            in
            ( { model | queue = updatedQueue, onPreviousBehavior = Shared.Model.RestartCurrent }, cmd )
    in
    case model.onPreviousBehavior of
        Shared.Model.PlayPreviousTrack ->
            prev

        Shared.Model.RestartCurrent ->
            case TrackQueue.getActiveTrack model.queue |> Maybe.map (\{ progress } -> progress > 5) of
                Just True ->
                    prev

                Just False ->
                    ( { model | onPreviousBehavior = Shared.Model.PlayPreviousTrack }, Effect.restartTrack )

                _ ->
                    ( model, Effect.none )
