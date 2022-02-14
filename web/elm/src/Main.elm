module Main exposing (..)

import Browser
import Css exposing (..)
import Css.Global
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css)
import Html.Styled.Events exposing (..)
import Page.Album as AlbumPage



-- globalStyle : Html msg
-- globalStyle =
--     Css.Global.global
--         [ Css.Global.html
--             [ height (pct 100)
--             ]
--         , Css.Global.body
--             [ height (pct 100)
--             , color (hex "#FFF")
--             , fontFamily sansSerif
--             , overflow hidden
--             ]
--         , Css.Global.h1
--             [ margin (px 0)
--             , fontSize (px 32)
--             ]
--         ]
-- type Msg
--     = Nothing
-- type alias Model =
--     {}
-- view : Model -> Html Msg
-- view model =
--     div [ css [ height (pct 100), displayFlex, flexDirection column ] ]
--         [ globalStyle
--         , div
--             [ css
--                 [ displayFlex
--                 , flexDirection row
--                 , backgroundColor (hex "#222")
--                 , height (pct 100)
--                 , overflow hidden
--                 ]
--             ]
--             [ aside
--                 [ css
--                     [ padding (px 10)
--                     , backgroundColor (hex "#333")
--                     , width (px 200)
--                     , flexShrink (int 0)
--                     ]
--                 ]
--                 [ text "Sidebar" ]
--             , section
--                 [ css
--                     [ displayFlex
--                     , flexDirection column
--                     , padding (px 20)
--                     , flexGrow (int 1)
--                     ]
--                 ]
--                 [ div [] [ text "Main content" ]
--                 ]
--             ]
--         , div [ css [ backgroundColor (hex "#333"), padding (px 10) ] ]
--             [ text "Nothing is currently playing..."
--             ]
--         ]


main : Program () AlbumPage.Model AlbumPage.Msg
main =
    Browser.element
        { init = AlbumPage.init
        , view = \model -> toUnstyled <| AlbumPage.view model
        , update = AlbumPage.update
        , subscriptions = \_ -> Sub.none
        }
