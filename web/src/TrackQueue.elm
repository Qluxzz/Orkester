module TrackQueue exposing (ActiveTrack, Repeat(..), State(..), TrackQueue, getActiveTrack, init, next, previous, queueLast, queueNext, replaceQueue, updateActiveTrackProgress, updateActiveTrackState)

import Queue exposing (Queue)
import TrackInfo exposing (Track)


type Repeat
    = RepeatOff
      -- Apple Music: Song loops around after it has ended, pressing next starts looping the next track instead
      -- Spotify: Songs loops around after it has ended, pressing next changes to repeat all instead of one
      -- Orkester: ?
    | RepeatOne
    | RepeatAll


type State
    = Playing
    | Paused


type alias ActiveTrack =
    { track : Track
    , progress : Int
    , state : State
    }


type alias TrackQueue =
    Queue Track ActiveTrack


init : { current : Maybe Track, future : Maybe (List Track) } -> TrackQueue
init { current, future } =
    { history = []
    , current = Maybe.map toActiveTrack current
    , future = Maybe.withDefault [] future
    }


updateActiveTrackProgress : TrackQueue -> Int -> TrackQueue
updateActiveTrackProgress ({ current } as queue) progress =
    { queue
        | current = Maybe.map (\c -> { c | progress = progress }) current
    }


updateActiveTrackState : TrackQueue -> State -> TrackQueue
updateActiveTrackState ({ current } as queue) state =
    { queue
        | current = Maybe.map (\c -> { c | state = state }) current
    }


previous : TrackQueue -> TrackQueue
previous ({ history, current, future } as queue) =
    case history of
        [] ->
            queue

        _ ->
            let
                historyTail : Maybe Track
                historyTail =
                    history
                        |> List.reverse
                        |> List.head

                restOfHistory =
                    List.take (List.length history - 1) history
            in
            case historyTail of
                Just tail ->
                    case current of
                        Nothing ->
                            { queue
                                | history = restOfHistory
                                , current = Just (toActiveTrack tail)
                            }

                        Just c ->
                            { queue
                                | history = restOfHistory
                                , current = Just (toActiveTrack tail)
                                , future = toTrack c :: future
                            }

                Nothing ->
                    queue


next : TrackQueue -> Repeat -> TrackQueue
next ({ history, current, future } as queue) repeat =
    let
        updatedHistory =
            case current of
                Nothing ->
                    history

                Just c ->
                    history ++ [ toTrack c ]
    in
    case future of
        first :: rest ->
            { queue
                | history = updatedHistory
                , current = Just (toActiveTrack first)
                , future = rest
            }

        [] ->
            case repeat of
                RepeatOff ->
                    { queue
                        | history = updatedHistory
                        , current = Nothing
                    }

                RepeatOne ->
                    queue

                RepeatAll ->
                    case history of
                        first :: rest ->
                            { queue
                                | history = updatedHistory
                                , current = Just (toActiveTrack first)
                                , future = rest
                            }

                        [] ->
                            { queue
                                | history = updatedHistory
                                , current = Nothing
                            }


getActiveTrack : TrackQueue -> Maybe ActiveTrack
getActiveTrack =
    Queue.getCurrent


toTrack : ActiveTrack -> Track
toTrack { track } =
    track


toActiveTrack : Track -> ActiveTrack
toActiveTrack track =
    { track = track
    , progress = 0
    , state = Paused
    }


{-| Add the tracks to the top of the queue
-}
queueNext : List Track -> TrackQueue -> TrackQueue
queueNext entities ({ current, future } as queue) =
    case current of
        Just _ ->
            { queue | future = entities ++ future }

        Nothing ->
            case entities of
                first :: rest ->
                    { queue
                        | current = Just (toActiveTrack first)
                        , future = rest ++ future
                    }

                [] ->
                    queue


{-| Add the tracks to the bottom of the queue
-}
queueLast : List Track -> TrackQueue -> TrackQueue
queueLast =
    Queue.queueLast


{-| Replaces the currently queued tracks with a new list of tracks
-}
replaceQueue : List Track -> TrackQueue -> TrackQueue
replaceQueue entities ({ history, current } as queue) =
    case entities of
        first :: rest ->
            let
                updatedHistory =
                    case current of
                        Just c ->
                            history ++ [ toTrack c ]

                        Nothing ->
                            history
            in
            { queue | current = Just (toActiveTrack first), future = rest, history = updatedHistory }

        [] ->
            queue
