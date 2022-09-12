module QueueView exposing (..)

import Css exposing (auto, bold, fontSize, fontWeight, listStyle, none, overflow, padding, px)
import Html.Styled exposing (Html, div, li, p, text, ul)
import Html.Styled.Attributes exposing (css)
import Queue exposing (ActiveTrack, Queue, getCurrent, getFuture, getHistory)
import TrackInfo exposing (Track)


styledList : List (Html msg) -> Html msg
styledList =
    ul [ css [ listStyle none, padding (px 0) ] ]


styledP : List (Html msg) -> Html msg
styledP =
    p [ css [ fontSize (px 20), fontWeight bold ] ]


queueView : Queue Track ActiveTrack -> Html msg
queueView queue =
    div [ css [ overflow auto ] ]
        [ styledP [ text "History" ]
        , styledList (List.map (\{ title } -> li [] [ text title ]) (getHistory queue))
        , styledP
            [ text "Now playing" ]
        , styledList [ li [] [ text (Maybe.withDefault "" (queue |> getCurrent |> Maybe.map (\{ track } -> track.title))) ] ]
        , styledP
            [ text "Coming next" ]
        , styledList (List.map (\{ title } -> li [] [ text title ]) (getFuture queue))
        ]
