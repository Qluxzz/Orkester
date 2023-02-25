module Utilities.DelayedLoader exposing (custom, default)

import Css exposing (Duration, animationDelay, animationName, int)
import Css.Animations exposing (keyframes)
import Html.Styled exposing (Html, div)
import Html.Styled.Attributes exposing (css)


{-| Use this to avoid rendering your loading state if the thing you're waiting for happens before the specified delay
-}
custom : Duration ms -> List (Html msg) -> Html msg
custom delay =
    div
        [ css
            [ animationDelay delay
            , Css.opacity (int 0)
            , animationName
                (keyframes
                    [ ( 0, [ Css.Animations.opacity (int 0) ] )
                    , ( 100, [ Css.Animations.opacity (int 1) ] )
                    ]
                )
            ]
        ]


default : Duration ms -> Html msg
default delay =
    custom delay [ Html.Styled.text "Loading..." ]
