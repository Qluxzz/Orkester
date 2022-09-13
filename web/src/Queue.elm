module Queue exposing (Queue, empty, getCurrent, getFuture, getHistory, init, queueLast)


type alias Queue a b =
    { history : List a
    , current : Maybe b
    , future : List a
    }


init : Queue a b
init =
    empty


empty : Queue a b
empty =
    { history = []
    , current = Nothing
    , future = []
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


queueLast : List a -> Queue a b -> Queue a b
queueLast entities ({ future } as queue) =
    { queue | future = future ++ entities }
