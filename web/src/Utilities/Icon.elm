module Utilities.Icon exposing (IconType(..), url)


type IconType
    = Play
    | Pause
    | Next
    | Previous
    | RepeatOff
    | RepeatOne
    | RepeatAll


iconBaseUrl : String
iconBaseUrl =
    "/assets/icons/"


iconExtension : String
iconExtension =
    ".svg"


url : IconType -> String
url type_ =
    iconBaseUrl
        ++ (case type_ of
                Play ->
                    "play"

                Pause ->
                    "pause"

                Next ->
                    "next"

                Previous ->
                    "previous"

                RepeatOff ->
                    "repeat-off"

                RepeatOne ->
                    "repeat-one"

                RepeatAll ->
                    "repeat-all"
           )
        ++ iconExtension
