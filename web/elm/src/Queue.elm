module Queue exposing (Queue, empty, getCurrent, init, next, previous, queueLast, queueNext)


type Queue a
    = Queue
        { history : List a
        , current : Maybe a
        , future : List a
        }


empty : Queue a
empty =
    Queue
        { history = []
        , current = Nothing
        , future = []
        }


init : Maybe a -> Maybe (List a) -> Queue a
init current future =
    Queue
        { history = []
        , current = current
        , future = Maybe.withDefault [] future
        }


getCurrent : Queue a -> Maybe a
getCurrent (Queue { current }) =
    current


previous : Queue a -> Queue a
previous (Queue { history, current, future }) =
    -- If history is nothing, return existing
    -- If history is something:
    -- If current is something, take last from history, set as current, move current to future
    -- If current is nothing, take last from history, set as current
    case history of
        [] ->
            Queue
                { history = []
                , current = current
                , future = future
                }

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
                            Queue
                                { history = restOfHistory
                                , current = Just tail
                                , future = future
                                }

                        Just c ->
                            Queue
                                { history = restOfHistory
                                , current = Just tail
                                , future = c :: future
                                }

                Nothing ->
                    Queue
                        { history = history
                        , current = current
                        , future = future
                        }


next : Queue a -> Queue a
next (Queue { history, current, future }) =
    case current of
        Just c ->
            case future of
                first :: rest ->
                    Queue
                        { history = history ++ [ c ]
                        , current = Just first
                        , future = rest
                        }

                [] ->
                    Queue
                        { history = history ++ [ c ]
                        , current = Nothing
                        , future = []
                        }

        _ ->
            Queue
                { history = history
                , current = current
                , future = future
                }


queueNext : List a -> Queue a -> Queue a
queueNext entities (Queue { history, current, future }) =
    Queue
        { history = history
        , current = current
        , future = entities ++ future
        }


queueLast : List a -> Queue a -> Queue a
queueLast entities (Queue { history, current, future }) =
    Queue
        { history = history
        , current = current
        , future = future ++ entities
        }
