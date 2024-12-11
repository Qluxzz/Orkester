module Components.Table exposing (Align(..), Column, alignData, alignHeader, clickableColumn, defaultColumn, grow, linkColumn, table, textColumn)

import Html
import Html.Attributes
import Html.Events


type Align
    = Left
    | Center
    | Right


type alias Column a msg =
    { title : String
    , data : a -> Html.Html msg
    , hidden : Bool
    , onClick : Maybe (a -> msg)
    , headerStyle : List (Html.Attribute msg)
    , dataStyle : List (Html.Attribute msg)
    }


defaultColumn : String -> (a -> Html.Html msg) -> Column a msg
defaultColumn title data =
    Column title data False Nothing [] []


textColumn : String -> (a -> String) -> Column a msg
textColumn title data =
    Column title (data >> Html.text) False Nothing [] []


linkColumn : String -> (a -> { url : String, title : String }) -> Column a msg
linkColumn title data =
    defaultColumn title (data >> toLink)


toLink : { url : String, title : String } -> Html.Html msg
toLink { url, title } =
    Html.a [ Html.Attributes.href url ] [ Html.text title ]


clickableColumn : String -> (a -> Html.Html msg) -> (a -> msg) -> Column a msg
clickableColumn title data onClick =
    Column title data False (Just onClick) [] []


alignHeader : Align -> Column a msg -> Column a msg
alignHeader alignment c =
    { c | headerStyle = alignmentToStyle alignment :: c.headerStyle }


alignData : Align -> Column a msg -> Column a msg
alignData alignment c =
    { c | dataStyle = alignmentToStyle alignment :: c.dataStyle }


alignmentToStyle : Align -> Html.Attribute msg
alignmentToStyle alignment =
    case alignment of
        Left ->
            Html.Attributes.style "text-align" "left"

        Center ->
            Html.Attributes.style "text-align" "center"

        Right ->
            Html.Attributes.style "text-align" "right"


grow : Column a msg -> Column a msg
grow c =
    { c | headerStyle = Html.Attributes.style "width" "100%" :: c.headerStyle, dataStyle = Html.Attributes.style "width" "100%" :: c.dataStyle }


table : List (Column a msg) -> List a -> Html.Html msg
table columns data =
    let
        visibleColumns : List (Column a msg)
        visibleColumns =
            columns |> List.filter (not << .hidden)
    in
    Html.table [ Html.Attributes.style "border-collapse" "collapse" ]
        [ Html.thead []
            [ Html.tr []
                (List.map column visibleColumns)
            ]
        , Html.tbody
            []
            (List.map
                (\r ->
                    Html.tr
                        []
                        (List.map
                            (\c ->
                                let
                                    attributes : List (Html.Attribute msg)
                                    attributes =
                                        case c.onClick of
                                            Just onClick ->
                                                [ Html.Events.onClick (onClick r), Html.Attributes.style "cursor" "pointer" ]

                                            Nothing ->
                                                []
                                in
                                Html.td
                                    (attributes ++ c.dataStyle)
                                    [ c.data r ]
                            )
                            visibleColumns
                        )
                )
                data
            )
        ]


column : Column a msg -> Html.Html msg
column { title, headerStyle } =
    Html.th
        headerStyle
        [ Html.text title ]
