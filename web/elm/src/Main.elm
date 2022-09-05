module Main exposing (..)

import ApiBaseUrl exposing (apiBaseUrl)
import Browser exposing (Document, UrlRequest)
import Browser.Navigation as Nav
import Css exposing (Color, Style, alignItems, backgroundColor, center, color, column, displayFlex, flexDirection, flexGrow, flexShrink, fontFamily, fontSize, height, hex, hidden, hover, int, justifyContent, margin, marginLeft, none, overflow, padding, padding2, paddingLeft, paddingRight, pct, property, px, row, sansSerif, textDecoration, underline, width)
import Css.Global
import DurationDisplay exposing (durationDisplay)
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src, type_, value)
import Html.Styled.Events exposing (onClick, onInput, onMouseUp)
import Http
import JSPlayer
import Page.Album as AlbumPage exposing (albumUrl, formatTrackArtists)
import Page.Artist as ArtistPage exposing (artistUrl)
import Page.LikedTracks as LikedTracksPage
import Page.Search as SearchPage
import RemoteData exposing (RemoteData(..), WebData)
import Route exposing (Route)
import String exposing (toInt)
import TrackId exposing (TrackId)
import TrackInfo exposing (Track, trackInfoDecoder)
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
        , Sub.map JSPlayer (JSPlayer.stateChange JSPlayer.StateChange)
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
    { title = Maybe.withDefault "Orkester" (getDocumentTitle model.page model.player)
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
                    [ padding (px 10)
                    , backgroundColor (hex "#333")
                    , width (px 200)
                    , flexShrink (int 0)
                    , displayFlex
                    , flexDirection column
                    , property "gap" "10px"
                    ]
                ]
                [ a [ href "/liked-tracks" ] [ text "Liked Tracks" ]
                , a [ href "/search" ] [ text "Search" ]
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
    div [ css [ displayFlex, width (pct 100), height (pct 100), alignItems center, justifyContent center ] ]
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
        [ onClick (JSPlayer JSPlayer.Play), css [ playPauseButtonStyle ] ]
        [ text "Play" ]


pauseButton : Html Msg
pauseButton =
    button
        [ onClick (JSPlayer JSPlayer.Pause), css [ playPauseButtonStyle ] ]
        [ text "Pause" ]


playerView : Model -> Html Msg
playerView model =
    div [ css [ displayFlex, flexDirection column ] ]
        (case model.player of
            Just { track, progress, state } ->
                case track of
                    RemoteData.Success t ->
                        let
                            sliderValue =
                                model.player
                                    |> Maybe.andThen (\p -> p.slider)
                                    |> Maybe.withDefault progress
                        in
                        [ div [ css [ displayFlex ] ]
                            [ a [ href ("/album/" ++ String.fromInt t.album.id ++ "/" ++ t.album.urlName) ]
                                [ img [ css [ width (px 128), height (px 128) ], src (apiBaseUrl ++ "/api/v1/album/" ++ String.fromInt t.album.id ++ "/image") ] []
                                ]
                            , div [ css [ marginLeft (px 10), overflow hidden ] ]
                                [ h1 [] [ text t.title ]
                                , h2 []
                                    (formatTrackArtists t.artists
                                        ++ [ span [] [ text " - " ]
                                           , a [ href ("/album/" ++ String.fromInt t.album.id ++ "/" ++ t.album.urlName) ] [ text t.album.name ]
                                           ]
                                    )
                                ]
                            ]
                        , div [ css [ displayFlex, alignItems center, flexGrow (int 1), property "gap" "10px" ] ]
                            [ div []
                                [ case state of
                                    Playing ->
                                        pauseButton

                                    Paused ->
                                        playButton
                                ]
                            , div [ css [ paddingRight (px 10) ] ] [ text (durationDisplay sliderValue) ]
                            , input
                                [ css [ width (pct 100) ]
                                , type_ "range"
                                , Html.Styled.Attributes.min "0"
                                , Html.Styled.Attributes.max (String.fromInt t.length)
                                , value (sliderValue |> String.fromInt)
                                , onInput (\value -> OnDragSlider (Maybe.withDefault 0 (value |> toInt)))
                                , onMouseUp OnDragSliderEnd
                                ]
                                []
                            , div [ css [ paddingLeft (px 10) ] ] [ text (durationDisplay t.length) ]
                            ]
                        ]

                    _ ->
                        [ text "Nothing is playing right now" ]

            _ ->
                [ text "Nothing is playing right now" ]
        )



-- MODEL


type alias Model =
    { route : Route
    , page : Page
    , navKey : Nav.Key
    , player : Maybe Player
    }


type alias Player =
    { track : WebData Track
    , progress : Int
    , slider : Maybe Int
    , state : PlayerState
    }


type PlayerState
    = Playing
    | Paused


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
            , player = Nothing
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
    | TrackInfoRecieved (WebData Track)
      -- Player slider interaction
    | OnDragSlider Int
    | OnDragSliderEnd


loadTrackInfo : Int -> Cmd Msg
loadTrackInfo trackId =
    Cmd.batch
        [ JSPlayer.playTrack trackId
        , getTrackInfo trackId
        ]


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case ( msg, model.page ) of
        ( AlbumPageMsg albumMsg, AlbumPage pageModel ) ->
            let
                ( updatedPageModel, updatedCmd ) =
                    AlbumPage.update albumMsg pageModel
            in
            case albumMsg of
                AlbumPage.PlayTrack trackId ->
                    ( { model
                        | page = AlbumPage updatedPageModel
                      }
                    , loadTrackInfo trackId
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

        ( LikedTracksPageMsg likedTracksMsg, LikedTracksPage likedTracksModel ) ->
            let
                ( updatedModel, updatedCmd ) =
                    LikedTracksPage.update likedTracksMsg likedTracksModel
            in
            ( { model | page = LikedTracksPage updatedModel }
            , Cmd.map LikedTracksPageMsg updatedCmd
            )

        ( SearchPageMsg searchPageMsg, SearchPage searchPageModel ) ->
            let
                ( updatedModel, updatedCmd ) =
                    SearchPage.update searchPageMsg searchPageModel
            in
            case searchPageMsg of
                SearchPage.PlayTrack trackId ->
                    ( { model | page = SearchPage updatedModel }
                    , loadTrackInfo trackId
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
                    Debug.todo "Handle error"

                JSPlayer.ProgressUpdated updatedProgress ->
                    ( { model | player = updateProgress updatedProgress model.player }, Cmd.none )

                JSPlayer.PlayTrack _ ->
                    ( { model | player = playPlayer model.player }, Cmd.none )

                JSPlayer.Seek _ ->
                    ( model, Cmd.none )

                JSPlayer.Pause ->
                    ( { model | player = pausePlayer model.player }, JSPlayer.pause () )

                JSPlayer.Play ->
                    ( { model | player = playPlayer model.player }, JSPlayer.play () )

                JSPlayer.StateChange state ->
                    case state of
                        "play" ->
                            ( { model | player = playPlayer model.player }, JSPlayer.play () )

                        "pause" ->
                            ( { model | player = pausePlayer model.player }, JSPlayer.pause () )

                        _ ->
                            Debug.todo ("unknown state change " ++ state)

        ( TrackInfoRecieved trackInfo, _ ) ->
            ( { model | player = Just { track = trackInfo, progress = 0, slider = Nothing, state = Playing } }, Cmd.none )

        ( OnDragSlider time, _ ) ->
            ( { model | player = updateSliderValue time model.player }, Cmd.none )

        ( OnDragSliderEnd, _ ) ->
            let
                sliderValue =
                    Maybe.andThen (\s -> s.slider) model.player

                cmd : Cmd Msg
                cmd =
                    case ( model.player, sliderValue ) of
                        ( Just { track }, Just time ) ->
                            case track of
                                RemoteData.Success _ ->
                                    JSPlayer.seek { timestamp = time }

                                _ ->
                                    Cmd.none

                        _ ->
                            Cmd.none
            in
            ( { model | player = clearSliderValue model.player }, cmd )

        ( _, _ ) ->
            Debug.todo "Should never happen!"



-- HELPER FUNCTIONS


getCurrentlyPlayingTrackInfo : Maybe Player -> Maybe String
getCurrentlyPlayingTrackInfo player =
    case player of
        Just { track, state } ->
            if state == Paused then
                Nothing

            else
                track
                    |> RemoteData.toMaybe
                    |> Maybe.map
                        (\t ->
                            t.title
                                ++ " - "
                                ++ (t.artists
                                        |> List.map (\a -> a.name)
                                        |> String.join ", "
                                   )
                        )

        _ ->
            Nothing


pausePlayer : Maybe Player -> Maybe Player
pausePlayer player =
    Maybe.map (\p -> { p | state = Paused }) player


playPlayer : Maybe Player -> Maybe Player
playPlayer player =
    Maybe.map (\p -> { p | state = Playing }) player


clearSliderValue : Maybe Player -> Maybe Player
clearSliderValue player =
    Maybe.map (\p -> { p | slider = Nothing }) player


updateSliderValue : Int -> Maybe Player -> Maybe Player
updateSliderValue value player =
    Maybe.map (\p -> { p | slider = Just value }) player


updateProgress : Int -> Maybe Player -> Maybe Player
updateProgress progress player =
    Maybe.map (\p -> { p | progress = progress }) player


getTrackInfo : TrackId -> Cmd Msg
getTrackInfo trackId =
    Http.get
        { url = apiBaseUrl ++ "/api/v1/track/" ++ String.fromInt trackId
        , expect =
            trackInfoDecoder
                |> Http.expectJson (RemoteData.fromResult >> TrackInfoRecieved)
        }


getDocumentTitle : Page -> Maybe Player -> Maybe String
getDocumentTitle page player =
    case getCurrentlyPlayingTrackInfo player of
        Just title ->
            Just ("► " ++ title)

        _ ->
            case page of
                ArtistPage { artist } ->
                    artist |> RemoteData.toMaybe |> Maybe.map (\a -> a.name)

                AlbumPage { album } ->
                    album |> RemoteData.toMaybe |> Maybe.map (\a -> a.name)

                NotFoundPage ->
                    Just "Not Found"

                IndexPage ->
                    Nothing

                LikedTracksPage _ ->
                    Just "Liked Tracks"

                SearchPage { searchPhrase } ->
                    Just (Maybe.withDefault "Search" searchPhrase)
