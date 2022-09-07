module Queue exposing (Queue, Repeat(..), empty, getCurrent, getRepeat, init, next, previous, queueLast, queueNext)


type Repeat
    = Off
      -- Apple Music: Song loops around after it has ended, pressing next starts looping the next track instead
      -- Spotify: Songs loops around after it has ended, pressing next changes to repeat all instead of one
      -- Orkester: ?
    | One
    | All


type alias Queue a =
    { history : List a
    , current : Maybe a
    , future : List a
    , repeat : Repeat
    }


empty : Queue a
empty =
    { history = []
    , current = Nothing
    , future = []
    , repeat = Off
    }


init : { current : Maybe a, future : Maybe (List a), repeat : Repeat } -> Queue a
init { current, future, repeat } =
    { history = []
    , current = current
    , future = Maybe.withDefault [] future
    , repeat = repeat
    }


getCurrent : Queue a -> Maybe a
getCurrent { current } =
    current


getRepeat : Queue a -> Repeat
getRepeat { repeat } =
    repeat


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


next : Queue a -> Queue a
next ({ history, current, future, repeat } as queue) =
    case current of
        Nothing ->
            queue

        Just c ->
            case future of
                first :: rest ->
                    { queue
                        | history = history ++ [ c ]
                        , current = Just first
                        , future = rest
                    }

                [] ->
                    case repeat of
                        Off ->
                            { queue
                                | history = history ++ [ c ]
                                , current = Nothing
                            }

                        One ->
                            queue

                        All ->
                            case history of
                                first :: rest ->
                                    { queue
                                        | history = history ++ [ c ]
                                        , current = Just first
                                        , future = rest
                                    }

                                [] ->
                                    { queue
                                        | history = history ++ [ c ]
                                        , current = Nothing
                                    }


queueNext : List a -> Queue a -> Queue a
queueNext entities ({ future } as queue) =
    { queue | future = entities ++ future }


queueLast : List a -> Queue a -> Queue a
queueLast entities ({ future } as queue) =
    { queue | future = future ++ entities }
