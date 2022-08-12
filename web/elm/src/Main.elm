module Main exposing (..)

import ApiBaseUrl exposing (apiBaseUrl)
import Browser exposing (Document, UrlRequest)
import Browser.Navigation as Nav
import Css exposing (Color, alignItems, backgroundColor, center, color, column, displayFlex, flexDirection, flexGrow, flexShrink, fontFamily, fontSize, height, hex, hidden, hover, int, justifyContent, margin, marginLeft, none, overflow, padding, pct, property, px, row, sansSerif, textDecoration, underline, width)
import Css.Global
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Http
import Page.Album as AlbumPage exposing (albumUrl, formatTrackArtists)
import Page.Artist as ArtistPage exposing (artistUrl)
import Page.LikedTracks as LikedTracksPage
import Page.Search as SearchPage
import Player
import RemoteData exposing (RemoteData(..), WebData)
import Route exposing (Route)
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
    Sub.map Player (Player.playbackFailed Player.PlaybackFailed)



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
    { title = Maybe.withDefault "Orkester" (getDocumentTitle model.page)
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
            [ div [ css [ displayFlex ] ]
                (case model.currentlyPlaying of
                    RemoteData.Success track ->
                        [ a [ href ("/album/" ++ String.fromInt track.album.id ++ "/" ++ track.album.urlName) ]
                            [ img [ css [ width (px 128), height (px 128) ], src (apiBaseUrl ++ "/api/v1/album/" ++ String.fromInt track.album.id ++ "/image") ] []
                            ]
                        , div [ css [ marginLeft (px 10), overflow hidden ] ]
                            [ h1 [] [ text track.title ]
                            , h2 []
                                (formatTrackArtists track.artists
                                    ++ [ span [] [ text " - " ]
                                       , a [ href ("/album/" ++ String.fromInt track.album.id ++ "/" ++ track.album.urlName) ] [ text track.album.name ]
                                       ]
                                )
                            ]
                        ]

                    _ ->
                        [ text "Nothing is playing right now..." ]
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
    div [ css [ displayFlex, width (pct 100), height (pct 100), alignItems center, justifyContent center ] ]
        [ h1 [] [ text text_ ]
        ]


indexView : Html Msg
indexView =
    centeredView "Welcome!"


notFoundView : Html Msg
notFoundView =
    centeredView "Page was not found"



-- MODEL


type alias Model =
    { route : Route
    , page : Page
    , navKey : Nav.Key
    , currentlyPlaying : WebData Track
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
            , currentlyPlaying = NotAsked
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
    = AlbumPageMsg AlbumPage.Msg
    | ArtistPageMsg ArtistPage.Msg
    | LikedTracksPageMsg LikedTracksPage.Msg
    | SearchPageMsg SearchPage.Msg
    | LinkClicked UrlRequest
    | UrlChanged Url
    | PlaybackFailed String
    | Player Player.Msg
    | TrackInfoRecieved (WebData Track)


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case ( msg, model.page ) of
        ( AlbumPageMsg albumMsg, AlbumPage pageModel ) ->
            let
                ( updatedPageModel, updatedCmd ) =
                    AlbumPage.update albumMsg pageModel
            in
            case albumMsg of
                AlbumPage.Player (Player.PlayTrack track) ->
                    ( { model
                        | page = AlbumPage updatedPageModel
                      }
                    , Cmd.batch
                        [ Player.playTrack track
                        , getTrackInfo track.id
                        ]
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
                SearchPage.Player (Player.PlayTrack track) ->
                    ( { model
                        | page = SearchPage updatedModel
                      }
                    , Cmd.batch
                        [ Player.playTrack track
                        , getTrackInfo track.id
                        ]
                    )

                _ ->
                    ( { model | page = SearchPage updatedModel }
                    , Cmd.batch
                        (Cmd.map
                            SearchPageMsg
                            updatedCmd
                            :: (if searchPageModel.searchPhrase /= updatedModel.searchPhrase then
                                    [ Nav.replaceUrl model.navKey ("/search/" ++ updatedModel.searchPhrase) ]

                                else
                                    []
                               )
                        )
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

        -- ( Player (Player.PlaybackFailed _), _ ) ->
        --     ( { model | currentlyPlaying = NotAsked }, Cmd.none )
        ( TrackInfoRecieved trackInfo, _ ) ->
            let
                cmd : Cmd Msg
                cmd =
                    case trackInfo of
                        RemoteData.Success t ->
                            Player.playTrack { id = t.id, timestamp = 0 }

                        _ ->
                            Cmd.none
            in
            ( { model | currentlyPlaying = trackInfo }, cmd )

        ( _, _ ) ->
            ( model, Cmd.none )



-- HELPER FUNCTIONS


getTrackInfo : TrackId -> Cmd Msg
getTrackInfo trackId =
    Http.get
        { url = apiBaseUrl ++ "/api/v1/track/" ++ String.fromInt trackId
        , expect =
            trackInfoDecoder
                |> Http.expectJson (RemoteData.fromResult >> TrackInfoRecieved)
        }


getDocumentTitle : Page -> Maybe String
getDocumentTitle page =
    case page of
        ArtistPage { artist } ->
            case artist of
                RemoteData.Success a ->
                    Just a.name

                _ ->
                    Nothing

        AlbumPage { album } ->
            case album of
                RemoteData.Success a ->
                    Just a.name

                _ ->
                    Nothing

        NotFoundPage ->
            Just "Not Found"

        IndexPage ->
            Nothing

        LikedTracksPage _ ->
            Just "Liked Tracks"

        SearchPage { searchPhrase } ->
            Just searchPhrase
