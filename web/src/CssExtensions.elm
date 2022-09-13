module CssExtensions exposing (..)

import Css exposing (LengthOrNumber, property)


gap : LengthOrNumber compatible -> Css.Style
gap { value } =
    property "gap" value
