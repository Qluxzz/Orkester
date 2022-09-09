module QueueView exposing (..)

import Html.Styled exposing (Html, div, h1, li, text, ul)
import Queue exposing (Queue, getCurrent, getFuture, getHistory)
import TrackId exposing (TrackId)


queueView : Queue TrackId -> Html msg
queueView queue =
    div []
        [ h1 [] [ text "History" ]
        , ul [] (List.map (\tId -> li [] [ text (String.fromInt tId) ]) (getHistory queue))
        , h1 [] [ text "Now playing" ]
        , ul [] [ li [] [ text (Maybe.withDefault "" (queue |> getCurrent |> Maybe.map (\c -> String.fromInt c))) ] ]
        , h1 [] [ text "Coming next" ]
        , ul [] (List.map (\tId -> li [] [ text (String.fromInt tId) ]) (getFuture queue))
        ]
