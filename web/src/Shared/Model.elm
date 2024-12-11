module Shared.Model exposing (Model, OnPrevious(..))

import Types.TrackQueue as TrackQueue


{-| Normally, this value would live in "Shared.elm"
but that would lead to a circular dependency import cycle.

For that reason, both `Shared.Model` and `Shared.Msg` are in their
own file, so they can be imported by `Effect.elm`

-}
type OnPrevious
    = PlayPreviousTrack
    | RestartCurrent


type alias Model =
    { queue : TrackQueue.TrackQueue
    , repeat : TrackQueue.Repeat
    , volume : Int
    , onPreviousBehavior : OnPrevious
    }
