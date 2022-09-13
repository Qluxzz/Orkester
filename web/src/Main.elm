module Main exposing (..)

import AlbumUrl exposing (albumImageUrl, albumUrl)
import ArtistUrl exposing (artistUrl)
import Browser exposing (Document, UrlRequest)
import Browser.Navigation as Nav
import Css exposing (Color, Style, alignItems, backgroundColor, border, center, color, column, displayFlex, flexBasis, flexDirection, flexGrow, flexShrink, fontFamily, fontSize, height, hex, hidden, hover, int, justifyContent, margin, none, overflow, padding, pct, property, px, row, sansSerif, textDecoration, transparent, underline, width)
import Css.Global
import DurationDisplay exposing (durationDisplay)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src, type_, value)
import Html.Styled.Events exposing (onClick, onInput, onMouseUp)
import JSPlayer
import Page.Album as AlbumPage exposing (formatTrackArtists)
import Page.Artist as ArtistPage
import Page.LikedTracks as LikedTracksPage
import Page.Search as SearchPage
import Queue
import QueueView exposing (queueView)
import RemoteData exposing (RemoteData(..))
import Route exposing (Route)
import String exposing (toInt)
import TrackInfo exposing (Track)
import TrackQueue exposing (ActiveTrack, Repeat(..), State(..), TrackQueue)
import Url exposing (Url)



-- MAIN


main : Program () Model Msg
main =
    Browser.application
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        , onUrlRequest = LinkClicked
        , onUrlChange = UrlChanged
        }



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.batch
        [ Sub.map JSPlayer (JSPlayer.playbackFailed JSPlayer.PlaybackFailed)
        , Sub.map JSPlayer (JSPlayer.progressUpdated JSPlayer.ProgressUpdated)
        , Sub.map JSPlayer (JSPlayer.stateChange JSPlayer.ExternalStateChange)
        ]



-- GLOBAL STYLES


textColor : Color
textColor =
    hex "#FFF"


globalStyle : Html msg
globalStyle =
    Css.Global.global
        [ Css.Global.html
            [ height (pct 100)
            ]
        , Css.Global.body
            [ height (pct 100)
            , color textColor
            , fontFamily sansSerif
            , overflow hidden
            ]
        , Css.Global.h1
            [ margin (px 0)
            , fontSize (px 32)
            ]
        , Css.Global.h2
            [ margin (px 0) ]
        , Css.Global.a
            [ color textColor
            , textDecoration none
            , hover [ textDecoration underline ]
            ]
        , Css.Global.p
            [ margin (px 0)
            ]
        ]



-- VIEW


view : Model -> Document Msg
view model =
    { title = Maybe.withDefault "Orkester" (getDocumentTitle model.page (TrackQueue.getActiveTrack model.queue))
    , body = [ baseView model (currentView model.page) |> toUnstyled ]
    }


baseView : Model -> Html Msg -> Html Msg
baseView model mainContent =
    div [ css [ height (pct 100), displayFlex, flexDirection column ] ]
        [ globalStyle
        , div
            [ css
                [ displayFlex
                , flexDirection row
                , backgroundColor (hex "#222")
                , height (pct 100)
                , overflow hidden
                ]
            ]
            [ aside
                [ css
                    [ backgroundColor (hex "#444")
                    , width (px 250)
                    , flexShrink (int 0)
                    , displayFlex
                    , flexDirection column
                    , property "gap" "10px"
                    ]
                ]
                [ div
                    [ css
                        [ padding (px 10)
                        , flexGrow (int 1)
                        , displayFlex
                        , flexDirection column
                        ]
                    ]
                    [ a [ href "/liked-tracks" ] [ text "Liked Tracks" ]
                    , a [ href "/search" ] [ text "Search" ]
                    ]
                , case TrackQueue.getActiveTrack model.queue of
                    Just { track } ->
                        a
                            [ css [ displayFlex ]
                            , href (albumUrl track.album)
                            ]
                            [ img
                                [ css [ width (pct 100) ]
                                , src (albumImageUrl track.album)
                                ]
                                []
                            ]

                    _ ->
                        Html.Styled.text ""
                ]
            , section
                [ css
                    [ displayFlex
                    , flexDirection column
                    , padding (px 20)
                    , flexGrow (int 1)
                    , overflow hidden
                    ]
                ]
                [ mainContent
                ]
            , aside
                [ css
                    [ padding (px 10)
                    , backgroundColor (hex "#444")
                    , width (px 200)
                    , flexShrink (int 0)
                    , displayFlex
                    , flexDirection column
                    , property "gap" "10px"
                    ]
                ]
                [ queueView
                    model.queue
                ]
            ]
        , div [ css [ backgroundColor (hex "#333"), padding (px 10) ] ]
            [ playerView model
            ]
        ]


currentView : Page -> Html Msg
currentView page =
    case page of
        NotFoundPage ->
            notFoundView

        AlbumPage albumModel ->
            AlbumPage.view albumModel
                |> map AlbumPageMsg

        ArtistPage artistModel ->
            ArtistPage.view artistModel
                |> map ArtistPageMsg

        LikedTracksPage likedTracksModel ->
            LikedTracksPage.view likedTracksModel
                |> map LikedTracksPageMsg

        SearchPage searchPageModel ->
            SearchPage.view searchPageModel
                |> map SearchPageMsg

        IndexPage ->
            indexView


centeredView : String -> Html Msg
centeredView text_ =
    div
        [ css
            [ displayFlex
            , width (pct 100)
            , height (pct 100)
            , alignItems center
            , justifyContent center
            ]
        ]
        [ h1 [] [ text text_ ]
        ]


indexView : Html Msg
indexView =
    centeredView "Welcome!"


notFoundView : Html Msg
notFoundView =
    centeredView "Page was not found"


playPauseButtonStyle : Style
playPauseButtonStyle =
    Css.batch
        [ width (px 50)
        ]


playButton : Html Msg
playButton =
    button
        [ onClick Play, css [ playPauseButtonStyle ] ]
        [ text "Play" ]


pauseButton : Html Msg
pauseButton =
    button
        [ onClick Pause, css [ playPauseButtonStyle ] ]
        [ text "Pause" ]


playerView : Model -> Html Msg
playerView model =
    div [ css [ displayFlex, flexDirection row, property "gap" "10px" ] ]
        (case TrackQueue.getActiveTrack model.queue of
            Just activeTrack ->
                [ currentlyPlayingView activeTrack.track
                , controls model activeTrack
                ]

            _ ->
                [ text "Nothing is playing right now" ]
        )


controls : Model -> ActiveTrack -> Html Msg
controls { progressSlider, repeat, volume } { track, progress, state } =
    let
        sliderValue =
            case progressSlider of
                NonInteractiveSlider ->
                    progress

                InteractiveSlider x ->
                    x
    in
    div [ css [ displayFlex, flexDirection row, alignItems center, flexGrow (int 1), property "gap" "10px" ] ]
        [ div [ css [ displayFlex, property "gap" "10px" ] ]
            [ button [ onClick PlayPrevious ] [ text "⬅️" ]
            , case state of
                Playing ->
                    pauseButton

                Paused ->
                    playButton
            , button [ onClick PlayNext ] [ text "➡️" ]
            , repeatButton repeat
            ]
        , div [ css [ displayFlex, flexGrow (int 1), alignItems center, property "gap" "10px" ] ]
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


repeatButton : Repeat -> Html Msg
repeatButton repeat =
    let
        styledButton : msg -> String -> Html msg
        styledButton click tx =
            button
                [ css
                    [ border (px 0)
                    , padding (px 0)
                    , backgroundColor transparent
                    , fontSize (px 20)
                    ]
                , onClick click
                ]
                [ text tx ]
    in
    case repeat of
        RepeatOff ->
            styledButton (OnRepeatChange RepeatAll) "➡️"

        RepeatAll ->
            styledButton (OnRepeatChange RepeatOne) "🔁"

        RepeatOne ->
            styledButton (OnRepeatChange RepeatOff) "🔂"



-- MODEL


type alias Model =
    { route : Route
    , page : Page
    , navKey : Nav.Key
    , progressSlider : Slider
    , queue : TrackQueue
    , repeat : Repeat
    , volume : Int
    }


type Slider
    = NonInteractiveSlider
    | InteractiveSlider Int


type Page
    = NotFoundPage
    | IndexPage
    | AlbumPage AlbumPage.Model
    | ArtistPage ArtistPage.Model
    | LikedTracksPage LikedTracksPage.Model
    | SearchPage SearchPage.Model


init : () -> Url -> Nav.Key -> ( Model, Cmd Msg )
init _ url navKey =
    let
        model : Model
        model =
            { route = Route.parseUrl url
            , page = NotFoundPage
            , navKey = navKey
            , queue = Queue.empty
            , repeat = RepeatOff
            , volume = 50
            , progressSlider = NonInteractiveSlider
            }
    in
    initCurrentPage ( model, Cmd.none )


initCurrentPage : ( Model, Cmd Msg ) -> ( Model, Cmd Msg )
initCurrentPage ( model, existingCmds ) =
    let
        ( currentPage, mappedPageCmds ) =
            case model.route of
                Route.NotFound ->
                    ( NotFoundPage, Cmd.none )

                Route.AlbumWithIdAndUrlName id _ ->
                    let
                        ( pageModel, pageCmds ) =
                            AlbumPage.init id
                    in
                    ( AlbumPage pageModel, Cmd.map AlbumPageMsg pageCmds )

                Route.AlbumWithId id ->
                    let
                        ( pageModel, pageCmds ) =
                            AlbumPage.init id
                    in
                    ( AlbumPage pageModel, Cmd.map AlbumPageMsg pageCmds )

                Route.ArtistWithId id ->
                    let
                        ( pageModel, pageCmds ) =
                            ArtistPage.init id
                    in
                    ( ArtistPage pageModel, Cmd.map ArtistPageMsg pageCmds )

                Route.ArtistWithIdAndUrlName id _ ->
                    let
                        ( pageModel, pageCmds ) =
                            ArtistPage.init id
                    in
                    ( ArtistPage pageModel, Cmd.map ArtistPageMsg pageCmds )

                Route.HomePage ->
                    ( IndexPage, Cmd.none )

                Route.LikedTracks ->
                    let
                        ( pageModel, pageCmds ) =
                            LikedTracksPage.init
                    in
                    ( LikedTracksPage pageModel, Cmd.map LikedTracksPageMsg pageCmds )

                Route.Search ->
                    let
                        ( pageModel, pageCmds ) =
                            SearchPage.init Nothing
                    in
                    ( SearchPage pageModel, Cmd.map SearchPageMsg pageCmds )

                Route.SearchWithQuery query ->
                    let
                        ( pageModel, pageCmds ) =
                            SearchPage.init (Just query)
                    in
                    ( SearchPage pageModel, Cmd.map SearchPageMsg pageCmds )
    in
    ( { model | page = currentPage }
    , Cmd.batch [ existingCmds, mappedPageCmds ]
    )



-- UPDATE


type Msg
    = -- Pages
      AlbumPageMsg AlbumPage.Msg
    | ArtistPageMsg ArtistPage.Msg
    | LikedTracksPageMsg LikedTracksPage.Msg
    | SearchPageMsg SearchPage.Msg
      -- Navigation
    | LinkClicked UrlRequest
    | UrlChanged Url
      -- JS Player
    | JSPlayer JSPlayer.Msg
      -- Controls
    | OnDragProgressSlider Int
    | OnDragProgressSliderEnd
    | OnDragVolumeSlider Int
    | OnRepeatChange Repeat
    | PlayNext
    | PlayPrevious
    | Pause
    | Play


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case ( msg, model.page ) of
        ( AlbumPageMsg albumMsg, AlbumPage pageModel ) ->
            let
                ( updatedPageModel, updatedCmd ) =
                    AlbumPage.update albumMsg pageModel
            in
            case albumMsg of
                AlbumPage.PlayTrack track ->
                    let
                        updatedQueue =
                            TrackQueue.replaceQueue [ track ] model.queue
                    in
                    ( { model
                        | page = AlbumPage updatedPageModel
                        , queue = updatedQueue
                      }
                    , JSPlayer.playTrack track.id
                    )

                AlbumPage.PlayAlbum tracks ->
                    let
                        updatedQueue =
                            TrackQueue.replaceQueue tracks model.queue

                        track =
                            TrackQueue.getActiveTrack updatedQueue
                    in
                    ( { model
                        | page = AlbumPage updatedPageModel
                        , queue = updatedQueue
                      }
                    , case track of
                        Just t ->
                            JSPlayer.playTrack t.track.id

                        Nothing ->
                            Cmd.none
                    )

                AlbumPage.AlbumReceived (RemoteData.Success album) ->
                    ( { model | page = AlbumPage updatedPageModel }
                    , Cmd.batch
                        [ Cmd.map AlbumPageMsg updatedCmd
                        , Nav.replaceUrl model.navKey (albumUrl album)
                        ]
                    )

                _ ->
                    ( { model | page = AlbumPage updatedPageModel }
                    , Cmd.map AlbumPageMsg updatedCmd
                    )

        ( AlbumPageMsg _, _ ) ->
            ( model, Cmd.none )

        ( ArtistPageMsg artistMsg, ArtistPage pageModel ) ->
            let
                ( updatedPageModel, updatedCmd ) =
                    ArtistPage.update artistMsg pageModel
            in
            case artistMsg of
                ArtistPage.ArtistRecieved (RemoteData.Success artist) ->
                    ( { model | page = ArtistPage updatedPageModel }
                    , Cmd.batch
                        [ Cmd.map ArtistPageMsg updatedCmd
                        , Nav.replaceUrl model.navKey (artistUrl artist)
                        ]
                    )

                _ ->
                    ( { model | page = ArtistPage updatedPageModel }
                    , Cmd.map ArtistPageMsg updatedCmd
                    )

        ( ArtistPageMsg _, _ ) ->
            ( model, Cmd.none )

        ( LikedTracksPageMsg likedTracksMsg, LikedTracksPage likedTracksModel ) ->
            let
                ( updatedModel, updatedCmd ) =
                    LikedTracksPage.update likedTracksMsg likedTracksModel
            in
            ( { model | page = LikedTracksPage updatedModel }
            , Cmd.map LikedTracksPageMsg updatedCmd
            )

        ( LikedTracksPageMsg _, _ ) ->
            ( model, Cmd.none )

        ( SearchPageMsg searchPageMsg, SearchPage searchPageModel ) ->
            let
                ( updatedModel, updatedCmd ) =
                    SearchPage.update searchPageMsg searchPageModel
            in
            case searchPageMsg of
                SearchPage.PlayTrack track ->
                    let
                        updatedQueue =
                            TrackQueue.replaceQueue [ track ] model.queue
                    in
                    ( { model
                        | page = SearchPage updatedModel
                        , queue = updatedQueue
                      }
                    , JSPlayer.playTrack track.id
                    )

                SearchPage.UpdateSearchPhrase phrase ->
                    ( { model | page = SearchPage updatedModel }
                    , Nav.replaceUrl
                        model.navKey
                        ("/search/" ++ phrase)
                    )

                _ ->
                    ( { model | page = SearchPage updatedModel }
                    , Cmd.map
                        SearchPageMsg
                        updatedCmd
                    )

        ( SearchPageMsg _, _ ) ->
            ( model, Cmd.none )

        ( LinkClicked urlRequest, _ ) ->
            case urlRequest of
                Browser.Internal url ->
                    ( model
                    , Nav.pushUrl model.navKey (Url.toString url)
                    )

                Browser.External url ->
                    ( model
                    , Nav.load url
                    )

        ( UrlChanged url, _ ) ->
            let
                newRoute =
                    Route.parseUrl url
            in
            -- This fixes the problem with infinite loops on replaceUrl
            if newRoute /= model.route then
                ( { model | route = newRoute }, Cmd.none )
                    |> initCurrentPage

            else
                ( model, Cmd.none )

        ( JSPlayer playerMsg, _ ) ->
            case playerMsg of
                JSPlayer.PlaybackFailed error ->
                    Debug.todo ("Playback failed " ++ error)

                JSPlayer.ProgressUpdated updatedProgress ->
                    ( { model | queue = TrackQueue.updateActiveTrackProgress model.queue updatedProgress }, Cmd.none )

                JSPlayer.Seek _ ->
                    ( model, Cmd.none )

                JSPlayer.ExternalStateChange state ->
                    case state of
                        "play" ->
                            ( { model | queue = TrackQueue.updateActiveTrackState model.queue Playing }, Cmd.none )

                        "pause" ->
                            ( { model | queue = TrackQueue.updateActiveTrackState model.queue Paused }, Cmd.none )

                        "ended" ->
                            let
                                updatedQueue =
                                    TrackQueue.next model.queue model.repeat

                                cmd =
                                    (case model.repeat of
                                        RepeatOne ->
                                            Just (JSPlayer.play ())

                                        _ ->
                                            TrackQueue.getActiveTrack updatedQueue
                                                |> Maybe.map (\{ track } -> JSPlayer.playTrack track.id)
                                    )
                                        |> Maybe.withDefault Cmd.none
                            in
                            ( { model | queue = updatedQueue }, cmd )

                        "nexttrack" ->
                            playNext model

                        "previoustrack" ->
                            playPrevious model

                        _ ->
                            Debug.todo ("unknown state change " ++ state)

        ( OnDragProgressSlider time, _ ) ->
            ( { model | progressSlider = InteractiveSlider time }, Cmd.none )

        ( OnDragProgressSliderEnd, _ ) ->
            let
                cmd : Cmd Msg
                cmd =
                    case model.progressSlider of
                        InteractiveSlider time ->
                            JSPlayer.seek { timestamp = time }

                        _ ->
                            Cmd.none
            in
            ( { model | progressSlider = NonInteractiveSlider }, cmd )

        ( OnDragVolumeSlider volume, _ ) ->
            ( { model | volume = volume }, JSPlayer.setVolume volume )

        ( OnRepeatChange repeat, _ ) ->
            ( { model | repeat = repeat }, Cmd.none )

        ( PlayNext, _ ) ->
            playNext model

        ( PlayPrevious, _ ) ->
            playPrevious model

        ( Pause, _ ) ->
            ( model, JSPlayer.pause () )

        ( Play, _ ) ->
            ( model, JSPlayer.play () )



-- HELPER FUNCTIONS


playNext : Model -> ( Model, Cmd msg )
playNext model =
    let
        updatedQueue =
            TrackQueue.next model.queue model.repeat

        cmd =
            TrackQueue.getActiveTrack updatedQueue
                |> Maybe.map (\{ track } -> JSPlayer.playTrack track.id)
                |> Maybe.withDefault (JSPlayer.pause ())

        -- Clear player if no new track
    in
    ( { model | queue = updatedQueue }, cmd )


playPrevious : Model -> ( Model, Cmd Msg )
playPrevious model =
    let
        updatedQueue =
            TrackQueue.previous model.queue

        current =
            TrackQueue.getActiveTrack updatedQueue

        cmd : Cmd Msg
        cmd =
            current
                |> Maybe.map (\{ track } -> JSPlayer.playTrack track.id)
                |> Maybe.withDefault Cmd.none
    in
    ( { model | queue = updatedQueue }, cmd )


getCurrentlyPlayingTrackInfo : Track -> String
getCurrentlyPlayingTrackInfo track =
    track.title
        ++ " - "
        ++ (track.artists
                |> List.map .name
                |> String.join ", "
           )


getDocumentTitle : Page -> Maybe ActiveTrack -> Maybe String
getDocumentTitle page maybeActiveTrack =
    maybeActiveTrack
        |> Maybe.map
            (\{ track, state } ->
                if state == Playing then
                    Just ("► " ++ getCurrentlyPlayingTrackInfo track)

                else
                    Nothing
            )
        |> Maybe.withDefault
            (case page of
                ArtistPage { artist } ->
                    artist |> RemoteData.toMaybe |> Maybe.map .name

                AlbumPage { album } ->
                    album |> RemoteData.toMaybe |> Maybe.map .name

                NotFoundPage ->
                    Just "Not Found"

                IndexPage ->
                    Nothing

                LikedTracksPage _ ->
                    Just "Liked Tracks"

                SearchPage { searchPhrase } ->
                    Just ("Search: " ++ Maybe.withDefault "" searchPhrase)
            )
