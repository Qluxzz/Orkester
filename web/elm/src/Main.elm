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
        { albums = List.map (\_ -> { id = 1, name = "Maniac", urlName = "maniac" }) (List.range 1 100)
        , artists = List.map (\_ -> { id = 1, name = "Carpenter Brut", urlName = "carpenter-brut" }) (List.range 1 50)
        , tracks = [ { id = 1, title = "Maniac" } ]
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
            , overflow hidden
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
                [ div [ css [ displayFlex, flexDirection column, overflow auto ] ]
                    [ input [ css [ flexGrow (int 1) ], type_ "text" ] []
                    , div [ css [ displayFlex, overflow auto, marginTop (px 20), marginBottom (px 20) ] ]
                        [ searchResultList model.searchResult.albums
                        , searchResultList model.searchResult.artists
                        , searchResultList (List.map (\x -> { id = x.id, name = x.title, urlName = "" }) model.searchResult.tracks)
                        ]
                    ]
                , div [] [ text "Main content" ]
                ]
            ]
        , div [ css [ backgroundColor (hex "#333"), padding (px 10) ] ]
            [ text "Nothing is currently playing..."
            ]
        ]


searchResultList : List { a | id : Int, name : String, urlName : String } -> Html Msg
searchResultList entries =
    ul [ css [ flexGrow (int 1), listStyle none, padding (px 0), margin (px 0) ] ] (List.map searchResultEntry entries)


searchResultEntry : { a | id : Int, name : String, urlName : String } -> Html Msg
searchResultEntry entry =
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
