module Html.Extra exposing (..)

import Html


picture : List (Html.Attribute msg) -> List (Html.Html msg) -> Html.Html msg
picture =
    Html.node "picture"
