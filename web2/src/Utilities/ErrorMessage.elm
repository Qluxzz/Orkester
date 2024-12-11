module Utilities.ErrorMessage exposing (..)

import Html
import Html.Attributes
import Http


errorMessage : String -> Http.Error -> Html.Html msg
errorMessage userFriendlyMessage error =
    Html.div [ Html.Attributes.class "error-message" ]
        [ Html.h1 [] [ Html.text userFriendlyMessage ]
        , Html.p []
            [ Html.text
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
