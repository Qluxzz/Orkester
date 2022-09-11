module QueueView exposing (..)

import Css exposing (auto, bottom, listStyle, none, overflow, padding, position, px, sticky, top)
import Html.Styled exposing (Html, div, h1, li, text, ul)
import Html.Styled.Attributes exposing (css)
import Queue exposing (Queue, getCurrent, getFuture, getHistory)
import TrackInfo exposing (Track)


styledList : List (Html msg) -> Html msg
styledList =
    ul [ css [ listStyle none, padding (px 0) ] ]


queueView : Queue Track -> Html msg
queueView queue =
    div [ css [ overflow auto ] ]
        [ h1 [ css [ position sticky, top (px 0) ] ] [ text "History" ]
        , styledList (List.map (\{ title } -> li [] [ text title ]) (getHistory queue))
        , h1
            [ css [ position sticky, top (px 0), bottom (px 0) ] ]
            [ text "Now playing" ]
        , styledList [ li [] [ text (Maybe.withDefault "" (queue |> getCurrent |> Maybe.map (\{ title } -> title))) ] ]
        , h1
            [ css [ position sticky, top (px 0), bottom (px 0) ] ]
            [ text "Coming next" ]
        , styledList (List.map (\{ title } -> li [] [ text title ]) (getFuture queue))
        ]
