module Main exposing (..)

import Browser exposing (Document, UrlRequest)
import Browser.Navigation as Nav
import Css exposing (..)
import Css.Global
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css)
import Page.Album as AlbumPage exposing (getAlbumUrl)
import Page.Artist as ArtistPage exposing (getArtistUrl)
import RemoteData
import Route exposing (Route)
import Url exposing (Url)


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
        , Css.Global.a
            [ color textColor
            , textDecoration none
            ]
        , Css.Global.p
            [ margin (px 0)
            ]
        ]


baseView : Html Msg -> Html Msg
baseView mainContent =
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
                    ]
                ]
                [ text "Sidebar" ]
            , section
                [ css
                    [ displayFlex
                    , flexDirection column
                    , padding (px 20)
                    , flexGrow (int 1)
                    ]
                ]
                [ mainContent
                ]
            ]
        , div [ css [ backgroundColor (hex "#333"), padding (px 10) ] ]
            [ text "Nothing is currently playing..."
            ]
        ]


type alias Model =
    { route : Route
    , page : Page
    , navKey : Nav.Key
    }


type Page
    = NotFoundPage
    | IndexPage
    | AlbumPage AlbumPage.Model
    | ArtistPage ArtistPage.Model


type Msg
    = AlbumPageMsg AlbumPage.Msg
    | ArtistPageMsg ArtistPage.Msg
    | LinkClicked UrlRequest
    | UrlChanged Url


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case ( msg, model.page ) of
        ( AlbumPageMsg subMsg, AlbumPage pageModel ) ->
            let
                ( updatedPageModel, updatedCmd ) =
                    AlbumPage.update subMsg pageModel
            in
            case subMsg of
                AlbumPage.AlbumReceived (RemoteData.Success album) ->
                    ( { model | page = AlbumPage updatedPageModel }
                    , Cmd.batch
                        [ Cmd.map AlbumPageMsg updatedCmd
                        , Nav.replaceUrl model.navKey (getAlbumUrl album)
                        ]
                    )

                _ ->
                    ( { model | page = AlbumPage updatedPageModel }
                    , Cmd.map AlbumPageMsg updatedCmd
                    )

        ( ArtistPageMsg subMsg, ArtistPage pageModel ) ->
            let
                ( updatedPageModel, updatedCmd ) =
                    ArtistPage.update subMsg pageModel
            in
            case subMsg of
                ArtistPage.ArtistRecieved (RemoteData.Success artist) ->
                    ( { model | page = ArtistPage updatedPageModel }
                    , Cmd.batch
                        [ Cmd.map ArtistPageMsg updatedCmd
                        , Nav.replaceUrl model.navKey (getArtistUrl artist)
                        ]
                    )

                _ ->
                    ( { model | page = ArtistPage updatedPageModel }
                    , Cmd.map ArtistPageMsg updatedCmd
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

        ( _, _ ) ->
            ( model, Cmd.none )


init : () -> Url -> Nav.Key -> ( Model, Cmd Msg )
init _ url navKey =
    let
        model =
            { route = Route.parseUrl url
            , page = NotFoundPage
            , navKey = navKey
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
    in
    ( { model | page = currentPage }
    , Cmd.batch [ existingCmds, mappedPageCmds ]
    )


view : Model -> Document Msg
view model =
    { title = Maybe.withDefault "Orkester" (getDocumentTitle model.page)
    , body = [ baseView (currentView model) |> toUnstyled ]
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

        _ ->
            Nothing


currentView : Model -> Html Msg
currentView model =
    case model.page of
        NotFoundPage ->
            notFoundView

        AlbumPage albumModel ->
            AlbumPage.view albumModel
                |> map AlbumPageMsg

        ArtistPage artistModel ->
            ArtistPage.view artistModel
                |> map ArtistPageMsg

        IndexPage ->
            indexView


indexView : Html Msg
indexView =
    div [ css [ displayFlex, width (pct 100), height (pct 100), alignItems center, justifyContent center ] ]
        [ h1 [] [ text "Welcome!" ]
        ]


notFoundView : Html Msg
notFoundView =
    div [ css [ displayFlex, width (pct 100), height (pct 100), alignItems center, justifyContent center ] ]
        [ h1 [] [ text "Page was not found" ]
        ]


main : Program () Model Msg
main =
    Browser.application
        { init = init
        , view = view
        , update = update
        , subscriptions = \_ -> Sub.none
        , onUrlRequest = LinkClicked
        , onUrlChange = UrlChanged
        }
