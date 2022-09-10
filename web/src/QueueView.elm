module QueueView exposing (..)

import Css exposing (auto, overflow, position, px, sticky, top)
import Html.Styled exposing (Html, div, h1, li, text, ul)
import Html.Styled.Attributes exposing (css)
import Queue exposing (Queue, getCurrent, getFuture, getHistory)
import TrackId exposing (TrackId)


queueView : Queue TrackId -> Html msg
queueView queue =
    div [ css [ overflow auto ] ]
        [ h1 [ css [ position sticky, top (px 0) ] ] [ text "History" ]
        , ul [] (List.map (\tId -> li [] [ text (String.fromInt tId) ]) (getHistory queue))
        , h1 [ css [ position sticky, top (px 0) ] ] [ text "Now playing" ]
        , ul [] [ li [] [ text (Maybe.withDefault "" (queue |> getCurrent |> Maybe.map (\c -> String.fromInt c))) ] ]
        , h1 [ css [ position sticky, top (px 0) ] ] [ text "Coming next" ]
        , ul [] (List.map (\tId -> li [] [ text (String.fromInt tId) ]) (getFuture queue))
        ]
