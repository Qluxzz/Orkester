module Main exposing (..)

import Browser
import Css exposing (..)
import Css.Global exposing (Snippet)
import Html.Events exposing (..)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, type_)


type alias Model =
    {}


init : Model
init =
    {}


globalStyle : Html msg
globalStyle =
    Css.Global.global
        [ Css.Global.html
            [ height (pct 100)
            ]
        , Css.Global.body
            [ height (pct 100)
            , color (hex "#FFF")
            ]
        ]


view : Model -> Html Msg
view model =
    div [ css [ height (pct 100), displayFlex, flexDirection column ] ]
        [ globalStyle
        , div
            [ css [ displayFlex, flexDirection row, backgroundColor (hex "#222"), height (pct 100) ] ]
            [ aside [ css [ padding (px 10), backgroundColor (hex "#333"), width (px 200) ] ] [ text "Sidebar" ]
            , section [ css [ displayFlex, flexDirection column, padding (px 20), flexGrow (int 1) ] ]
                [ div [ css [ marginBottom (px 20), displayFlex ] ]
                    [ input [ css [ flexGrow (int 1) ], type_ "text" ] []
                    ]
                , div [] [ text "Main content" ]
                ]
            ]
        , div [ css [ backgroundColor (hex "#333"), padding (px 20) ] ]
            [ text "Player"
            ]
        ]


type Msg
    = Nothing


update : Msg -> Model -> Model
update message model =
    case message of
        _ ->
            model


main : Program () Model Msg
main =
    Browser.sandbox
        { init = init
        , view = \model -> toUnstyled (view model)
        , update = update
        }
