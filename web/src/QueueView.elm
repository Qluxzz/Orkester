module QueueView exposing (..)

import Css exposing (auto, bottom, listStyle, none, overflow, padding, position, px, sticky, top)
import Html.Styled exposing (Attribute, Html, div, h1, li, text, ul)
import Html.Styled.Attributes exposing (css)
import Queue exposing (Queue, getCurrent, getFuture, getHistory)
import TrackId exposing (TrackId)


styledList : List (Html msg) -> Html msg
styledList =
    ul [ css [ listStyle none, padding (px 0) ] ]


queueView : Queue TrackId -> Html msg
queueView queue =
    div [ css [ overflow auto ] ]
        [ h1 [ css [ position sticky, top (px 0) ] ] [ text "History" ]
        , styledList (List.map (\tId -> li [] [ text (String.fromInt tId) ]) (getHistory queue))
        , h1
            [ css [ position sticky, top (px 0), bottom (px 0) ] ]
            [ text "Now playing" ]
        , styledList [ li [] [ text (Maybe.withDefault "" (queue |> getCurrent |> Maybe.map (\c -> String.fromInt c))) ] ]
        , h1
            [ css [ position sticky, top (px 0), bottom (px 0) ] ]
            [ text "Coming next" ]
        , styledList (List.map (\tId -> li [] [ text (String.fromInt tId) ]) (getFuture queue))
        ]
