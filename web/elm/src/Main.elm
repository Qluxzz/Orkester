module Main exposing (..)

import Browser
import Css exposing (..)
import Css.Global
import Html.Events exposing (..)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, type_)


type alias IdNameAndUrlName =
    { id : Int
    , name : String
    , urlName : String
    }


type alias Album =
    IdNameAndUrlName


type alias Artist =
    IdNameAndUrlName


type alias Track =
    { id : Int
    , title : String
    }


type alias SearchResult =
    { albums : List Album
    , artists : List Artist
    , tracks : List Track
    }


type alias Model =
    { searchResult : SearchResult
    }


init : Model
init =
    { searchResult =
        { albums = [ { id = 1, name = "Maniac", urlName = "maniac" } ]
        , artists = [ { id = 1, name = "Carpenter Brut", urlName = "carpenter-brut" } ]
        , tracks = []
        }
    }


globalStyle : Html msg
globalStyle =
    Css.Global.global
        [ Css.Global.html
            [ height (pct 100)
            ]
        , Css.Global.body
            [ height (pct 100)
            , color (hex "#FFF")
            , fontFamily sansSerif
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
                [ div [ css [ marginBottom (px 20), displayFlex, flexDirection column ] ]
                    [ input [ css [ flexGrow (int 1) ], type_ "text" ] []
                    , div [ css [ displayFlex ] ]
                        [ ul [ css [ flexGrow (int 1) ] ] (List.map viewSearchResultEntry model.searchResult.albums)
                        , ul [ css [ flexGrow (int 1) ] ] (List.map viewSearchResultEntry model.searchResult.artists)
                        ]
                    ]
                , div [] [ text "Main content" ]
                ]
            ]
        , div [ css [ backgroundColor (hex "#333"), padding (px 10) ] ]
            [ text "Nothing is currently playing..."
            ]
        ]


viewSearchResultEntry : { a | id : Int, name : String, urlName : String } -> Html Msg
viewSearchResultEntry entry =
    li [] [ text entry.name ]


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
