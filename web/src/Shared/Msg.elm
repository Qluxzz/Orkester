module Shared.Msg exposing (Msg(..))

import JSPlayer
import Types.TrackInfo
import Types.TrackQueue


{-| Normally, this value would live in "Shared.elm"
but that would lead to a circular dependency import cycle.

For that reason, both `Shared.Model` and `Shared.Msg` are in their
own file, so they can be imported by `Effect.elm`

-}
type Msg
    = NoOp
    | JSPlayer JSPlayer.Msg
    | PlayTrack Types.TrackInfo.Track
    | PlayTracks (List Types.TrackInfo.Track)
    | PlayPrevious
    | PlayNext
    | Pause
    | Play
    | SetRepeatMode Types.TrackQueue.Repeat
    | SetVolume Int
