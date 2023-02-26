module Main exposing (..)

import Browser exposing (Document, UrlRequest)
import Browser.Navigation as Nav
import Css exposing (Color, alignItems, backgroundColor, center, color, column, displayFlex, flexDirection, flexGrow, flexShrink, fontFamily, fontSize, height, hex, hidden, hover, int, justifyContent, margin, none, overflow, padding, pct, px, row, sansSerif, textDecoration, underline, width)
import Css.Global
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import JSPlayer
import Page.Album as AlbumPage
import Page.Artist as ArtistPage
import Page.LikedTracks as LikedTracksPage
import Page.Search as SearchPage
import PlayerBar
import Queue
import QueueView exposing (queueView)
import RemoteData exposing (RemoteData(..))
import Route exposing (Route)
import String
import TrackQueue exposing (ActiveTrack, State(..), TrackQueue)
import Types.TrackInfo exposing (Track)
import Url exposing (Url)
import Utilities.AlbumUrl exposing (albumImageUrl, albumUrl)
import Utilities.ArtistUrl exposing (artistUrl)
import Utilities.CssExtensions exposing (gap)



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
    { title = getDocumentTitle model.page (TrackQueue.getActiveTrack model.queue)
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
                    , gap (px 10)
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
                    , gap (px 10)
                    ]
                ]
                [ queueView
                    model.queue
                ]
            ]
        , div [ css [ backgroundColor (hex "#333"), padding (px 10) ] ]
            [ Html.Styled.map (\msg -> PlayerBar msg)
                (PlayerBar.view
                    model.controls
                    (TrackQueue.getActiveTrack model.queue)
                )
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



-- MODEL


type OnPrevious
    = PlayPreviousTrack
    | RestartCurrent


type alias Model =
    { route : Route
    , page : Page
    , navKey : Nav.Key
    , queue : TrackQueue
    , controls : PlayerBar.Model
    , onPreviousBehaviour : OnPrevious
    }


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
            , controls = PlayerBar.init
            , onPreviousBehaviour = PlayPreviousTrack
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
    | PlayerBar PlayerBar.Msg


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
                                    TrackQueue.next model.queue model.controls.repeat

                                cmd =
                                    (case model.controls.repeat of
                                        TrackQueue.RepeatOne ->
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

        ( PlayerBar playerMsg, _ ) ->
            let
                updatedModel =
                    PlayerBar.update model.controls playerMsg
            in
            case playerMsg of
                PlayerBar.OnDragProgressSliderEnd ->
                    let
                        cmd : Cmd Msg
                        cmd =
                            case model.controls.progressSlider of
                                PlayerBar.InteractiveSlider time ->
                                    JSPlayer.seek { timestamp = time }

                                _ ->
                                    Cmd.none
                    in
                    ( { model | controls = updatedModel }, cmd )

                PlayerBar.OnDragVolumeSlider volume ->
                    ( { model | controls = updatedModel }, JSPlayer.setVolume volume )

                PlayerBar.PlayNext ->
                    playNext model

                PlayerBar.PlayPrevious ->
                    playPrevious model

                PlayerBar.Pause ->
                    ( { model | controls = updatedModel }, JSPlayer.pause () )

                PlayerBar.Play ->
                    ( { model | controls = updatedModel }, JSPlayer.play () )

                _ ->
                    ( { model | controls = updatedModel }, Cmd.none )



-- HELPER FUNCTIONS


playNext : Model -> ( Model, Cmd msg )
playNext model =
    let
        updatedQueue =
            TrackQueue.next model.queue model.controls.repeat

        cmd =
            TrackQueue.getActiveTrack updatedQueue
                |> Maybe.map (\{ track } -> JSPlayer.playTrack track.id)
                |> Maybe.withDefault (JSPlayer.pause ())
    in
    ( { model | queue = updatedQueue }, cmd )


{-|

    Plays previous if progress on current track
    is less than threshold, otherwise it restarts the current track
    and if pressed again, it jumps to the previous track

-}
playPrevious : Model -> ( Model, Cmd Msg )
playPrevious model =
    let
        prev : ( Model, Cmd Msg )
        prev =
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
            ( { model | queue = updatedQueue, onPreviousBehaviour = RestartCurrent }, cmd )
    in
    case model.onPreviousBehaviour of
        PlayPreviousTrack ->
            prev

        RestartCurrent ->
            case Queue.getCurrent model.queue |> Maybe.map (\{ progress } -> progress > 5) of
                Just True ->
                    prev

                Just False ->
                    ( { model | onPreviousBehaviour = PlayPreviousTrack }, JSPlayer.seek { timestamp = 0 } )

                _ ->
                    ( model, Cmd.none )


getCurrentlyPlayingTrackInfo : Track -> String
getCurrentlyPlayingTrackInfo track =
    track.title
        ++ " - "
        ++ (track.artists
                |> List.map .name
                |> String.join ", "
           )


getDocumentTitle : Page -> Maybe ActiveTrack -> String
getDocumentTitle page maybeActiveTrack =
    let
        trackIsPlaying =
            maybeActiveTrack
                |> Maybe.map (\{ state } -> state == Playing)
                |> Maybe.withDefault False
    in
    (if trackIsPlaying then
        Maybe.map (\{ track } -> "► " ++ getCurrentlyPlayingTrackInfo track) maybeActiveTrack

     else
        case page of
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

            SearchPage { search } ->
                Just ("Search: " ++ search)
    )
        |> Maybe.withDefault "Orkester"
