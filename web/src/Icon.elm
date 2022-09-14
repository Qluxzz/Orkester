module Icon exposing (IconType(..), iconUrl)


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


iconUrl : IconType -> String
iconUrl type_ =
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
