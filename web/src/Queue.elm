module Queue exposing (ActiveTrack, Queue, Repeat(..), State(..), empty, getCurrent, getFuture, getHistory, init, next, previous, queueLast, queueNext, replaceQueue, updateActiveTrackProgress, updateActiveTrackState)

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


type alias Queue a b =
    { history : List a
    , current : Maybe b
    , future : List a
    }


type alias TrackQueue =
    Queue Track ActiveTrack


empty : TrackQueue
empty =
    { history = []
    , current = Nothing
    , future = []
    }


init : { current : Maybe Track, future : Maybe (List Track) } -> TrackQueue
init { current, future } =
    { history = []
    , current = Maybe.map toActiveTrack current
    , future = Maybe.withDefault [] future
    }


getCurrent : Queue a b -> Maybe b
getCurrent { current } =
    current


getHistory : Queue a b -> List a
getHistory { history } =
    history


getFuture : Queue a b -> List a
getFuture { future } =
    future


updateActiveTrackProgress : Queue Track ActiveTrack -> Int -> Queue Track ActiveTrack
updateActiveTrackProgress ({ current } as queue) progress =
    { queue
        | current = Maybe.map (\c -> { c | progress = progress }) current
    }


updateActiveTrackState : Queue Track ActiveTrack -> State -> Queue Track ActiveTrack
updateActiveTrackState ({ current } as queue) state =
    { queue
        | current = Maybe.map (\c -> { c | state = state }) current
    }


previous : Queue Track ActiveTrack -> Queue Track ActiveTrack
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


next : Queue Track ActiveTrack -> Repeat -> Queue Track ActiveTrack
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


toTrack : ActiveTrack -> Track
toTrack { track } =
    track


toActiveTrack : Track -> ActiveTrack
toActiveTrack track =
    { track = track
    , progress = 0
    , state = Paused
    }


queueNext : List Track -> Queue Track ActiveTrack -> Queue Track ActiveTrack
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


replaceQueue : List Track -> Queue Track ActiveTrack -> Queue Track ActiveTrack
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


queueLast : List Track -> Queue Track ActiveTrack -> Queue Track ActiveTrack
queueLast entities ({ future } as queue) =
    { queue | future = future ++ entities }
