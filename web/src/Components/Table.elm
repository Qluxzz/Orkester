module Components.Table exposing (Align(..), Column, alignData, alignHeader, clickableColumn, defaultColumn, grow, table, textColumn)

import Css exposing (collapse)
import Html.Styled as Html
import Html.Styled.Attributes
import Html.Styled.Events


type Align
    = Left
    | Center
    | Right


type alias Column a msg =
    { title : String
    , data : a -> Html.Html msg
    , hidden : Bool
    , onClick : Maybe (a -> msg)
    , headerStyle : List Css.Style
    , dataStyle : List Css.Style
    }


defaultColumn : String -> (a -> Html.Html msg) -> Column a msg
defaultColumn title data =
    Column title data False Nothing [] []


textColumn : String -> (a -> String) -> Column a msg
textColumn title data =
    Column title (data >> Html.text) False Nothing [] []


clickableColumn : String -> (a -> Html.Html msg) -> (a -> msg) -> Column a msg
clickableColumn title data onClick =
    Column title data False (Just onClick) [] []


alignHeader : Align -> Column a msg -> Column a msg
alignHeader alignment c =
    { c
        | headerStyle =
            (case alignment of
                Left ->
                    Css.textAlign Css.left

                Center ->
                    Css.textAlign Css.center

                Right ->
                    Css.textAlign Css.right
            )
                :: c.headerStyle
    }


alignData : Align -> Column a msg -> Column a msg
alignData alignment c =
    { c
        | dataStyle =
            (case alignment of
                Left ->
                    Css.textAlign Css.left

                Center ->
                    Css.textAlign Css.center

                Right ->
                    Css.textAlign Css.right
            )
                :: c.dataStyle
    }


grow : Column a msg -> Column a msg
grow c =
    { c | headerStyle = Css.width (Css.pct 100) :: c.headerStyle, dataStyle = Css.width (Css.pct 100) :: c.dataStyle }


table : List (Column a msg) -> List a -> Html.Html msg
table columns data =
    let
        visibleColumns : List (Column a msg)
        visibleColumns =
            columns |> List.filter (not << .hidden)
    in
    Html.table [ Html.Styled.Attributes.css [ Css.borderCollapse collapse ] ]
        [ Html.thead []
            [ Html.tr []
                (List.map column visibleColumns)
            ]
        , Html.tbody
            []
            (List.map
                (\r ->
                    Html.tr
                        [ Html.Styled.Attributes.css
                            [ Css.nthChild "even"
                                [ Css.backgroundColor (Css.hex "#333") ]
                            , Css.nthChild "odd"
                                [ Css.backgroundColor (Css.hex "#222") ]
                            ]
                        ]
                        (List.map
                            (\c ->
                                let
                                    attributes : List (Html.Attribute msg)
                                    attributes =
                                        case c.onClick of
                                            Just onClick ->
                                                [ Html.Styled.Events.onClick (onClick r), Html.Styled.Attributes.css [ Css.cursor Css.pointer ] ]

                                            Nothing ->
                                                []
                                in
                                Html.td
                                    (attributes ++ [ Html.Styled.Attributes.css (Css.padding (Css.px 10) :: c.dataStyle) ])
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
        [ Html.Styled.Attributes.css (Css.padding (Css.px 10) :: headerStyle)
        ]
        [ Html.text title ]
