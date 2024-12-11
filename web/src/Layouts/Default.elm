module Layouts.Default exposing (Model, Msg, Props, layout)

import Api.Artist
import Effect exposing (Effect)
import Html exposing (Html)
import Html.Attributes exposing (class)
import Html.Events
import Layout exposing (Layout)
import Route exposing (Route)
import Route.Path
import Shared
import Shared.Msg
import Types.Queue
import Types.TrackInfo
import Types.TrackQueue
import Utilities.AlbumUrl
import Utilities.DurationDisplay
import Utilities.Icon as Icon
import View exposing (View)


type alias Props =
    {}


layout : Props -> Shared.Model -> Route () -> Layout () Model Msg contentMsg
layout props shared route =
    Layout.new
        { init = init
        , update = update
        , view = view shared
        , subscriptions = subscriptions
        }



-- MODEL


type Slider
    = NonInteractiveSlider
    | InteractiveSlider Int


type alias Model =
    { progressSlider : Slider
    }


init : () -> ( Model, Effect Msg )
init _ =
    ( { progressSlider = NonInteractiveSlider }
    , Effect.none
    )



-- UPDATE


type Msg
    = PlayPrevious
    | PlayNext
    | Pause
    | Play
    | OnDragProgressSlider Int
    | OnDragProgressSliderEnd
    | OnDragVolumeSlider Int
    | OnRepeatChange Types.TrackQueue.Repeat


update : Msg -> Model -> ( Model, Effect Msg )
update msg model =
    case msg of
        PlayPrevious ->
            ( model
            , Effect.playPreviousTrack
            )

        PlayNext ->
            ( model, Effect.playNextTrack )

        Pause ->
            ( model, Effect.pause )

        Play ->
            ( model, Effect.play )

        OnDragProgressSlider time ->
            ( { model | progressSlider = InteractiveSlider time }, Effect.none )

        OnDragProgressSliderEnd ->
            case model.progressSlider of
                InteractiveSlider time ->
                    ( { model | progressSlider = NonInteractiveSlider }, Effect.seek time )

                NonInteractiveSlider ->
                    ( model, Effect.none )

        OnDragVolumeSlider volume ->
            ( model, Effect.setVolume volume )

        OnRepeatChange repeat ->
            ( model, Effect.setRepeatMode repeat )


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Shared.Model -> { toContentMsg : Msg -> contentMsg, content : View contentMsg, model : Model } -> View contentMsg
view shared { toContentMsg, model, content } =
    let
        currentlyPlayingTrack =
            Types.TrackQueue.getActiveTrack shared.queue
    in
    { title = currentlyPlayingTrack |> Maybe.map (.track >> getCurrentlyPlayingTrackInfo) |> Maybe.withDefault content.title
    , body =
        [ sidebarView currentlyPlayingTrack
        , Html.main_ [] content.body
        , Html.aside [ Html.Attributes.class "queue" ] [ queueView shared.queue ]
        , Html.div [ Html.Attributes.class "player-bar" ] (playerBarView shared.volume model.progressSlider currentlyPlayingTrack shared.repeat)
            |> Html.map toContentMsg
        ]
    }


getCurrentlyPlayingTrackInfo : { r | title : String, artists : List { y | name : String } } -> String
getCurrentlyPlayingTrackInfo track =
    "► "
        ++ track.title
        ++ " - "
        ++ (track.artists
                |> List.map .name
                |> String.join ", "
           )


sidebarView : Maybe Types.TrackQueue.ActiveTrack -> Html msg
sidebarView activeTrack =
    Html.aside [ Html.Attributes.class "sidebar" ]
        [ Html.ul []
            [ Html.li [] [ Html.a [ Route.Path.href Route.Path.Search ] [ Html.text "Search" ] ]
            , Html.li [] [ Html.a [ Route.Path.href Route.Path.LikedTracks ] [ Html.text "Liked tracks" ] ]
            ]
        , case activeTrack of
            Just { track } ->
                Html.a
                    [ Route.Path.href (Route.Path.Album_Id__Name_ { id = String.fromInt track.album.id, name = track.album.urlName })
                    ]
                    [ Html.img
                        [ Html.Attributes.src (Utilities.AlbumUrl.albumImageUrl track.album)
                        ]
                        []
                    ]

            _ ->
                Html.text ""
        ]



-- PLAYER BAR


playerBarView : Int -> Slider -> Maybe Types.TrackQueue.ActiveTrack -> Types.TrackQueue.Repeat -> List (Html.Html Msg)
playerBarView volume progressSlider activeTrack repeat =
    case activeTrack of
        Just t ->
            [ currentlyPlayingView t.track
            , controls volume progressSlider t repeat
            ]

        Nothing ->
            [ Html.text "Nothing is playing right now" ]


currentlyPlayingView : Types.TrackInfo.Track -> Html.Html msg
currentlyPlayingView { title, album, artists } =
    Html.div [ Html.Attributes.class "track-info" ]
        [ Html.h1 [] [ Html.text title ]
        , Html.h2 []
            (formatTrackArtists artists
                ++ [ Html.span [] [ Html.text " - " ]
                   , Html.a [ Route.Path.href (Route.Path.Album_Id__Name_ { id = String.fromInt album.id, name = album.urlName }) ] [ Html.text album.name ]
                   ]
            )
        ]


formatTrackArtists : List { r | id : Int, name : String, urlName : String } -> List (Html.Html msg)
formatTrackArtists artists =
    artists
        |> List.map (\artist -> Html.a [ Route.Path.href (Route.Path.Artist_Id__Name_ { id = String.fromInt artist.id, name = artist.urlName }) ] [ Html.text artist.name ])
        |> List.intersperse (Html.span [] [ Html.text ", " ])


controls : Int -> Slider -> Types.TrackQueue.ActiveTrack -> Types.TrackQueue.Repeat -> Html.Html Msg
controls volume progressSlider { progress, state, track } repeat =
    let
        sliderValue =
            case progressSlider of
                NonInteractiveSlider ->
                    progress

                InteractiveSlider x ->
                    x
    in
    Html.div
        [ Html.Attributes.class "controls"
        ]
        [ Html.div [ Html.Attributes.class "buttons" ]
            [ Html.button [ Html.Attributes.class "player-button", Html.Events.onClick PlayPrevious ] [ Html.img [ Html.Attributes.src (Icon.url Icon.Previous) ] [] ]
            , case state of
                Types.TrackQueue.Playing ->
                    pauseButton

                Types.TrackQueue.Paused ->
                    playButton
            , Html.button [ Html.Attributes.class "player-button", Html.Events.onClick PlayNext ] [ Html.img [ Html.Attributes.src (Icon.url Icon.Next) ] [] ]
            , repeatButton repeat
            ]
        , Html.div [ Html.Attributes.style "display" "flex", Html.Attributes.style "gap" "10px" ]
            [ Html.text (Utilities.DurationDisplay.durationDisplay sliderValue)
            , Html.input
                [ Html.Attributes.type_ "range"
                , Html.Attributes.min "0"
                , Html.Attributes.max (String.fromInt track.length)
                , Html.Attributes.value (sliderValue |> String.fromInt)
                , Html.Events.onInput (\v -> OnDragProgressSlider (Maybe.withDefault 0 (v |> String.toInt)))
                , Html.Events.onMouseUp OnDragProgressSliderEnd
                ]
                []
            , Html.text (Utilities.DurationDisplay.durationDisplay track.length)
            ]
        , Html.div []
            [ Html.input
                [ Html.Attributes.type_ "range"
                , Html.Attributes.min "0"
                , Html.Attributes.max "100"
                , Html.Attributes.value (String.fromInt volume)
                , Html.Events.onInput (\v -> OnDragVolumeSlider (Maybe.withDefault 0 (v |> String.toInt)))
                ]
                []
            ]
        ]


playButton : Html.Html Msg
playButton =
    Html.button [ Html.Attributes.class "player-button", Html.Events.onClick Play ]
        [ Html.img [ Html.Attributes.src (Icon.url Icon.Play) ] []
        ]


pauseButton : Html.Html Msg
pauseButton =
    Html.button [ Html.Attributes.class "player-button", Html.Events.onClick Pause ]
        [ Html.img [ Html.Attributes.src (Icon.url Icon.Pause) ] [] ]


repeatButton : Types.TrackQueue.Repeat -> Html Msg
repeatButton repeat =
    let
        styledButton : msg -> String -> Html msg
        styledButton click icon =
            Html.button
                [ Html.Attributes.class "player-button"
                , Html.Events.onClick click
                ]
                [ Html.img [ Html.Attributes.src icon ] [] ]
    in
    case repeat of
        Types.TrackQueue.RepeatOff ->
            styledButton (OnRepeatChange Types.TrackQueue.RepeatAll) (Icon.url Icon.RepeatOff)

        Types.TrackQueue.RepeatAll ->
            styledButton (OnRepeatChange Types.TrackQueue.RepeatOne) (Icon.url Icon.RepeatAll)

        Types.TrackQueue.RepeatOne ->
            styledButton (OnRepeatChange Types.TrackQueue.RepeatOff) (Icon.url Icon.RepeatOne)



-- QUEUE VIEW


queueView : Types.TrackQueue.TrackQueue -> Html msg
queueView queue =
    Html.div []
        [ Html.p [] [ Html.text "History" ]
        , Html.ul [] (List.map (\{ title } -> Html.li [] [ Html.text title ]) (Types.Queue.getHistory queue))
        , Html.p [] [ Html.text "Now playing" ]
        , Html.ul [] [ Html.li [] [ Html.text (Maybe.withDefault "" (queue |> Types.TrackQueue.getActiveTrack |> Maybe.map (.track >> .title))) ] ]
        , Html.p [] [ Html.text "Coming next" ]
        , Html.ul [] (List.map (\{ title } -> Html.li [] [ Html.text title ]) (Types.Queue.getFuture queue))
        ]
