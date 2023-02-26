module Utilities.ErrorMessage exposing (..)

import Css exposing (border3, hex, marginTop, padding, px, solid)
import Html.Styled exposing (Html, div, h1, p, text)
import Html.Styled.Attributes exposing (css)
import Http


errorMessage : String -> Http.Error -> Html msg
errorMessage userFriendlyMessage error =
    div [ css [ padding (px 10), border3 (px 2) solid (hex "#eee") ] ]
        [ h1 [] [ text userFriendlyMessage ]
        , p [ css [ marginTop (px 20) ] ]
            [ text
                (case error of
                    Http.BadStatus status ->
                        "HTTP Error Code: " ++ String.fromInt status

                    Http.BadBody body ->
                        body

                    Http.BadUrl url ->
                        "Invalid URL " ++ url

                    Http.Timeout ->
                        "Request timed out"

                    Http.NetworkError ->
                        "Network connection lost"
                )
            ]
        ]
