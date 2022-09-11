module Queue exposing (Queue, Repeat(..), empty, getCurrent, getFuture, getHistory, init, next, previous, queueLast, queueNext, replaceQueue)


type Repeat
    = RepeatOff
      -- Apple Music: Song loops around after it has ended, pressing next starts looping the next track instead
      -- Spotify: Songs loops around after it has ended, pressing next changes to repeat all instead of one
      -- Orkester: ?
    | RepeatOne
    | RepeatAll


type alias Queue a =
    { history : List a
    , current : Maybe a
    , future : List a
    }


empty : Queue a
empty =
    { history = []
    , current = Nothing
    , future = []
    }


init : { current : Maybe a, future : Maybe (List a) } -> Queue a
init { current, future } =
    { history = []
    , current = current
    , future = Maybe.withDefault [] future
    }


getCurrent : Queue a -> Maybe a
getCurrent { current } =
    current


getHistory : Queue a -> List a
getHistory { history } =
    history


getFuture : Queue a -> List a
getFuture { future } =
    future


previous : Queue a -> Queue a
previous ({ history, current, future } as queue) =
    case history of
        [] ->
            queue

        _ ->
            let
                historyTail : Maybe a
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
                                , current = Just tail
                            }

                        Just c ->
                            { queue
                                | history = restOfHistory
                                , current = Just tail
                                , future = c :: future
                            }

                Nothing ->
                    queue


next : Queue a -> Repeat -> Queue a
next ({ history, current, future } as queue) repeat =
    let
        updatedHistory =
            case current of
                Nothing ->
                    history

                Just c ->
                    history ++ [ c ]
    in
    case future of
        first :: rest ->
            { queue
                | history = updatedHistory
                , current = Just first
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
                                , current = Just first
                                , future = rest
                            }

                        [] ->
                            { queue
                                | history = updatedHistory
                                , current = Nothing
                            }


queueNext : List a -> Queue a -> Queue a
queueNext entities ({ current, future } as queue) =
    case current of
        Just _ ->
            { queue | future = entities ++ future }

        Nothing ->
            case entities of
                first :: rest ->
                    { queue
                        | current = Just first
                        , future = rest ++ future
                    }

                [] ->
                    queue


replaceQueue : List a -> Queue a -> Queue a
replaceQueue entities ({ history, current } as queue) =
    case entities of
        first :: rest ->
            let
                updatedHistory =
                    case current of
                        Just c ->
                            history ++ [ c ]

                        Nothing ->
                            history
            in
            { queue | current = Just first, future = rest, history = updatedHistory }

        [] ->
            queue


queueLast : List a -> Queue a -> Queue a
queueLast entities ({ future } as queue) =
    { queue | future = future ++ entities }
