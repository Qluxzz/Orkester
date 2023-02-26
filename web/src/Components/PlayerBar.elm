module Components.PlayerBar exposing (..)

import Components.TrackQueue as TrackQueue
import Css exposing (Style, alignItems, backgroundColor, border, borderRadius, center, displayFlex, flexBasis, flexDirection, flexGrow, height, hex, hidden, hover, int, justifyContent, overflow, padding, pct, px, row, transparent, width)
import Html.Styled exposing (Html, a, button, div, h1, h2, img, input, span, text)
import Html.Styled.Attributes exposing (href, src, type_, value)
import Html.Styled.Events exposing (onClick, onInput, onMouseUp)
import Page.Album exposing (Model, formatTrackArtists)
import String exposing (toInt)
import Svg.Styled.Attributes exposing (css)
import Types.TrackInfo exposing (Track)
import Utilities.AlbumUrl exposing (albumUrl)
import Utilities.CssExtensions exposing (gap)
import Utilities.DurationDisplay exposing (durationDisplay)
import Utilities.Icon as Icon


init : Model
init =
    { progressSlider = NonInteractiveSlider
    , repeat = TrackQueue.RepeatOff
    , volume = 50
    }


type alias Model =
    { progressSlider : Slider
    , repeat : TrackQueue.Repeat
    , volume : Int
    }


type Msg
    = OnDragProgressSlider Int
    | OnDragProgressSliderEnd
    | OnDragVolumeSlider Int
    | OnRepeatChange TrackQueue.Repeat
    | PlayNext
    | PlayPrevious
    | Pause
    | Play


type Slider
    = NonInteractiveSlider
    | InteractiveSlider Int



-- VIEW


view : Model -> Maybe TrackQueue.ActiveTrack -> Html Msg
view model track =
    div [ css [ displayFlex, flexDirection row, gap (px 10) ] ]
        (case track of
            Just activeTrack ->
                [ currentlyPlayingView activeTrack.track
                , controls model activeTrack
                ]

            _ ->
                [ text "Nothing is playing right now" ]
        )


playerButtonStyle : Style
playerButtonStyle =
    Css.batch
        [ width (px 24)
        , height (px 24)
        , border (px 0)
        , backgroundColor transparent
        , borderRadius (pct 50)
        , displayFlex
        , justifyContent center
        , alignItems center
        , padding (px 15)
        , hover [ backgroundColor (hex "111") ]
        ]


playButton : Html Msg
playButton =
    button
        [ onClick Play, css [ playerButtonStyle ] ]
        [ img [ src (Icon.url Icon.Play) ] [] ]


pauseButton : Html Msg
pauseButton =
    button
        [ onClick Pause, css [ playerButtonStyle ] ]
        [ img [ src (Icon.url Icon.Pause) ] [] ]


controls : Model -> TrackQueue.ActiveTrack -> Html Msg
controls { progressSlider, repeat, volume } { track, progress, state } =
    let
        sliderValue =
            case progressSlider of
                NonInteractiveSlider ->
                    progress

                InteractiveSlider x ->
                    x
    in
    div [ css [ displayFlex, flexDirection row, alignItems center, flexGrow (int 1), gap (px 10) ] ]
        [ div [ css [ displayFlex, gap (px 10) ] ]
            [ button [ css [ playerButtonStyle ], onClick PlayPrevious ] [ img [ src (Icon.url Icon.Previous) ] [] ]
            , case state of
                TrackQueue.Playing ->
                    pauseButton

                TrackQueue.Paused ->
                    playButton
            , button [ css [ playerButtonStyle ], onClick PlayNext ] [ img [ src (Icon.url Icon.Next) ] [] ]
            , repeatButton repeat
            ]
        , div [ css [ displayFlex, flexGrow (int 1), alignItems center, gap (px 10) ] ]
            [ div [] [ text (durationDisplay sliderValue) ]
            , input
                [ css [ width (pct 100) ]
                , type_ "range"
                , Html.Styled.Attributes.min "0"
                , Html.Styled.Attributes.max (String.fromInt track.length)
                , value (sliderValue |> String.fromInt)
                , onInput (\value -> OnDragProgressSlider (Maybe.withDefault 0 (value |> toInt)))
                , onMouseUp OnDragProgressSliderEnd
                ]
                []
            , div [] [ text (durationDisplay track.length) ]
            ]
        , div [ css [ displayFlex ] ]
            [ input
                [ css [ width (pct 100) ]
                , type_ "range"
                , Html.Styled.Attributes.min "0"
                , Html.Styled.Attributes.max "100"
                , value (volume |> String.fromInt)
                , onInput (\value -> OnDragVolumeSlider (Maybe.withDefault 0 (value |> toInt)))
                ]
                []
            ]
        ]


currentlyPlayingView : Track -> Html Msg
currentlyPlayingView { title, album, artists } =
    div [ css [ displayFlex, flexBasis (pct 50), flexGrow (int 0) ] ]
        [ div [ css [ overflow hidden ] ]
            [ h1 [] [ text title ]
            , h2 []
                (formatTrackArtists artists
                    ++ [ span [] [ text " - " ]
                       , a [ href (albumUrl album) ] [ text album.name ]
                       ]
                )
            ]
        ]


repeatButton : TrackQueue.Repeat -> Html Msg
repeatButton repeat =
    let
        styledButton : msg -> String -> Html msg
        styledButton click icon =
            button
                [ css
                    [ playerButtonStyle ]
                , onClick click
                ]
                [ img [ src icon ] [] ]
    in
    case repeat of
        TrackQueue.RepeatOff ->
            styledButton (OnRepeatChange TrackQueue.RepeatAll) (Icon.url Icon.RepeatOff)

        TrackQueue.RepeatAll ->
            styledButton (OnRepeatChange TrackQueue.RepeatOne) (Icon.url Icon.RepeatAll)

        TrackQueue.RepeatOne ->
            styledButton (OnRepeatChange TrackQueue.RepeatOff) (Icon.url Icon.RepeatOne)



-- UPDATE


update : Model -> Msg -> Model
update model msg =
    case msg of
        OnDragProgressSlider time ->
            { model | progressSlider = InteractiveSlider time }

        OnDragProgressSliderEnd ->
            { model | progressSlider = NonInteractiveSlider }

        OnDragVolumeSlider volume ->
            { model | volume = volume }

        OnRepeatChange repeat ->
            { model | repeat = repeat }

        PlayNext ->
            model

        PlayPrevious ->
            model

        Pause ->
            model

        Play ->
            model
