module Effect exposing
    ( Effect
    , none, batch
    , sendCmd, sendMsg
    , pushRoute, replaceRoute
    , pushRoutePath, replaceRoutePath
    , loadExternalUrl, back
    , map, toCmd
    , focusElement, likeTrack, pause, play, playNextTrack, playPreviousTrack, playTrack, playTracks, restartTrack, seek, sendApiRequest, setRepeatMode, setVolume, startPlayback, unlikeTrack
    )

{-|

@docs Effect

@docs none, batch
@docs sendCmd, sendMsg

@docs pushRoute, replaceRoute
@docs pushRoutePath, replaceRoutePath
@docs loadExternalUrl, back

@docs map, toCmd

-}

import Browser.Dom
import Browser.Navigation
import Dict exposing (Dict)
import Http
import JSPlayer
import Json.Decode
import Route exposing (Route)
import Route.Path
import Shared.Model
import Shared.Msg
import Task
import Types.TrackId
import Types.TrackInfo
import Types.TrackQueue
import Url exposing (Url)


type Effect msg
    = -- BASICS
      None
    | Batch (List (Effect msg))
    | SendCmd (Cmd msg)
      -- ROUTING
    | PushUrl String
    | ReplaceUrl String
    | LoadExternalUrl String
    | Back
      -- SHARED
    | SendSharedMsg Shared.Msg.Msg
    | SendApiRequest
        { endpoint : String
        , decoder : Json.Decode.Decoder msg
        , onHttpError : Http.Error -> msg
        }
    | FocusElement String Shared.Msg.Msg
      -- PLAYER CONTROLS
    | StartPlayback Types.TrackId.TrackId
    | RestartTrack
    | Play
    | Pause
    | Seek Int
    | SetVolume Int
    | SetRepeatMode Types.TrackQueue.Repeat
    | LikeTrack Types.TrackId.TrackId
    | UnlikeTrack Types.TrackId.TrackId



-- BASICS


{-| Don't send any effect.
-}
none : Effect msg
none =
    None


{-| Send multiple effects at once.
-}
batch : List (Effect msg) -> Effect msg
batch =
    Batch


{-| Send a normal `Cmd msg` as an effect, something like `Http.get` or `Random.generate`.
-}
sendCmd : Cmd msg -> Effect msg
sendCmd =
    SendCmd


{-| Send a message as an effect. Useful when emitting events from UI components.
-}
sendMsg : msg -> Effect msg
sendMsg msg =
    Task.succeed msg
        |> Task.perform identity
        |> SendCmd



-- ROUTING


{-| Set the new route, and make the back button go back to the current route.
-}
pushRoute :
    { path : Route.Path.Path
    , query : Dict String String
    , hash : Maybe String
    }
    -> Effect msg
pushRoute route =
    PushUrl (Route.toString route)


{-| Same as `Effect.pushRoute`, but without `query` or `hash` support
-}
pushRoutePath : Route.Path.Path -> Effect msg
pushRoutePath path =
    PushUrl (Route.Path.toString path)


{-| Set the new route, but replace the previous one, so clicking the back
button **won't** go back to the previous route.
-}
replaceRoute :
    { path : Route.Path.Path
    , query : Dict String String
    , hash : Maybe String
    }
    -> Effect msg
replaceRoute route =
    ReplaceUrl (Route.toString route)


{-| Same as `Effect.replaceRoute`, but without `query` or `hash` support
-}
replaceRoutePath : Route.Path.Path -> Effect msg
replaceRoutePath path =
    ReplaceUrl (Route.Path.toString path)


{-| Redirect users to a new URL, somewhere external to your web application.
-}
loadExternalUrl : String -> Effect msg
loadExternalUrl =
    LoadExternalUrl


{-| Navigate back one page
-}
back : Effect msg
back =
    Back



-- CUSTOM


sendApiRequest :
    { endpoint : String
    , decoder : Json.Decode.Decoder value
    , onResponse : Result Http.Error value -> msg
    }
    -> Effect msg
sendApiRequest options =
    let
        decoder : Json.Decode.Decoder msg
        decoder =
            options.decoder
                |> Json.Decode.map Ok
                |> Json.Decode.map options.onResponse

        onHttpError : Http.Error -> msg
        onHttpError httpError =
            options.onResponse (Err httpError)
    in
    SendApiRequest
        { endpoint = options.endpoint
        , decoder = decoder
        , onHttpError = onHttpError
        }


focusElement : String -> Effect msg
focusElement elementId =
    FocusElement elementId Shared.Msg.NoOp


startPlayback : Types.TrackId.TrackId -> Effect msg
startPlayback id =
    StartPlayback id


playTrack : Types.TrackInfo.Track -> Effect msg
playTrack track =
    playTracks (List.singleton track)


playTracks : List Types.TrackInfo.Track -> Effect msg
playTracks tracks =
    case List.head tracks of
        Just t ->
            Batch
                [ SendSharedMsg (Shared.Msg.PlayTracks tracks)
                , StartPlayback t.id
                ]

        Nothing ->
            None


playNextTrack =
    SendSharedMsg Shared.Msg.PlayNext


playPreviousTrack =
    SendSharedMsg Shared.Msg.PlayPrevious


restartTrack : Effect msg
restartTrack =
    RestartTrack


play : Effect msg
play =
    Batch [ SendSharedMsg Shared.Msg.Play, Play ]


pause : Effect msg
pause =
    Batch [ SendSharedMsg Shared.Msg.Pause, Pause ]


seek : Int -> Effect msg
seek ms =
    Seek ms


setRepeatMode : Types.TrackQueue.Repeat -> Effect msg
setRepeatMode repeat =
    Batch
        [ SendSharedMsg (Shared.Msg.SetRepeatMode repeat)
        , SetRepeatMode repeat
        ]


setVolume : Int -> Effect msg
setVolume volume =
    Batch
        [ SendSharedMsg (Shared.Msg.SetVolume volume)
        , SetVolume volume
        ]


likeTrack : Types.TrackId.TrackId -> Effect msg
likeTrack =
    LikeTrack


unlikeTrack : Types.TrackId.TrackId -> Effect msg
unlikeTrack =
    UnlikeTrack



-- INTERNALS


{-| Elm Land depends on this function to connect pages and layouts
together into the overall app.
-}
map : (msg1 -> msg2) -> Effect msg1 -> Effect msg2
map fn effect =
    case effect of
        None ->
            None

        Batch list ->
            Batch (List.map (map fn) list)

        SendCmd cmd ->
            SendCmd (Cmd.map fn cmd)

        PushUrl url ->
            PushUrl url

        ReplaceUrl url ->
            ReplaceUrl url

        Back ->
            Back

        LoadExternalUrl url ->
            LoadExternalUrl url

        SendSharedMsg sharedMsg ->
            SendSharedMsg sharedMsg

        SendApiRequest data ->
            SendApiRequest
                { endpoint = data.endpoint
                , decoder = Json.Decode.map fn data.decoder
                , onHttpError = \err -> fn (data.onHttpError err)
                }

        FocusElement elementId msg_ ->
            FocusElement elementId msg_

        -- PLAYER CONTROLS
        StartPlayback trackId ->
            StartPlayback trackId

        RestartTrack ->
            RestartTrack

        Play ->
            Play

        Pause ->
            Pause

        Seek ms ->
            Seek ms

        SetVolume volume ->
            SetVolume volume

        SetRepeatMode mode ->
            SetRepeatMode mode

        LikeTrack trackId ->
            LikeTrack trackId

        UnlikeTrack trackId ->
            UnlikeTrack trackId


{-| Elm Land depends on this function to perform your effects.
-}
toCmd :
    { key : Browser.Navigation.Key
    , url : Url
    , shared : Shared.Model.Model
    , fromSharedMsg : Shared.Msg.Msg -> msg
    , batch : List msg -> msg
    , toCmd : msg -> Cmd msg
    }
    -> Effect msg
    -> Cmd msg
toCmd options effect =
    case effect of
        None ->
            Cmd.none

        Batch list ->
            Cmd.batch (List.map (toCmd options) list)

        SendCmd cmd ->
            cmd

        PushUrl url ->
            Browser.Navigation.pushUrl options.key url

        ReplaceUrl url ->
            Browser.Navigation.replaceUrl options.key url

        Back ->
            Browser.Navigation.back options.key 1

        LoadExternalUrl url ->
            Browser.Navigation.load url

        SendSharedMsg sharedMsg ->
            Task.succeed sharedMsg
                |> Task.perform options.fromSharedMsg

        SendApiRequest data ->
            Http.request
                { method = "GET"
                , url = data.endpoint
                , headers = []
                , body = Http.emptyBody
                , expect =
                    Http.expectJson
                        (\httpResult ->
                            case httpResult of
                                Ok msg ->
                                    msg

                                Err httpError ->
                                    data.onHttpError httpError
                        )
                        data.decoder
                , timeout = Just 15000
                , tracker = Nothing
                }

        FocusElement elementId sharedMsg ->
            Task.attempt (\_ -> options.fromSharedMsg sharedMsg) (Browser.Dom.focus elementId)

        -- PLAYER CONTROLS
        StartPlayback trackId ->
            JSPlayer.playTrack (Types.TrackId.toString trackId)

        RestartTrack ->
            JSPlayer.seek 0

        Play ->
            JSPlayer.play ()

        Pause ->
            JSPlayer.pause ()

        Seek ms ->
            JSPlayer.seek ms

        SetVolume volume ->
            JSPlayer.setVolume volume

        SetRepeatMode mode ->
            JSPlayer.setRepeatMode (Types.TrackQueue.toString mode)

        LikeTrack trackId ->
            Http.request
                { method = "PUT"
                , headers = []
                , url = "/api/v1/track/" ++ Types.TrackId.toString trackId ++ "/like"
                , body = Http.emptyBody
                , expect = Http.expectWhatever (\_ -> options.fromSharedMsg Shared.Msg.NoOp)
                , timeout = Nothing
                , tracker = Nothing
                }

        UnlikeTrack trackId ->
            Http.request
                { method = "DELETE"
                , headers = []
                , url = "/api/v1/track/" ++ Types.TrackId.toString trackId ++ "/like"
                , body = Http.emptyBody
                , expect = Http.expectWhatever (\_ -> options.fromSharedMsg Shared.Msg.NoOp)
                , timeout = Nothing
                , tracker = Nothing
                }
