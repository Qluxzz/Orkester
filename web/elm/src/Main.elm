module Main exposing (..)

import Browser exposing (Document, UrlRequest)
import Browser.Navigation as Nav
import Html.Styled exposing (..)
import Page.Album as AlbumPage
import Route exposing (Route)
import Url exposing (Url)



-- globalStyle : Html msg
-- globalStyle =
--     Css.Global.global
--         [ Css.Global.html
--             [ height (pct 100)
--             ]
--         , Css.Global.body
--             [ height (pct 100)
--             , color (hex "#FFF")
--             , fontFamily sansSerif
--             , overflow hidden
--             ]
--         , Css.Global.h1
--             [ margin (px 0)
--             , fontSize (px 32)
--             ]
--         ]
-- type Msg
--     = Nothing
-- type alias Model =
--     {}
-- view : Model -> Html Msg
-- view model =
--     div [ css [ height (pct 100), displayFlex, flexDirection column ] ]
--         [ globalStyle
--         , div
--             [ css
--                 [ displayFlex
--                 , flexDirection row
--                 , backgroundColor (hex "#222")
--                 , height (pct 100)
--                 , overflow hidden
--                 ]
--             ]
--             [ aside
--                 [ css
--                     [ padding (px 10)
--                     , backgroundColor (hex "#333")
--                     , width (px 200)
--                     , flexShrink (int 0)
--                     ]
--                 ]
--                 [ text "Sidebar" ]
--             , section
--                 [ css
--                     [ displayFlex
--                     , flexDirection column
--                     , padding (px 20)
--                     , flexGrow (int 1)
--                     ]
--                 ]
--                 [ div [] [ text "Main content" ]
--                 ]
--             ]
--         , div [ css [ backgroundColor (hex "#333"), padding (px 10) ] ]
--             [ text "Nothing is currently playing..."
--             ]
--         ]


type alias Model =
    { route : Route
    , page : Page
    , navKey : Nav.Key
    }


type Page
    = NotFoundPage
    | IndexPage
    | AlbumPage AlbumPage.Model


type Msg
    = AlbumPageMsg AlbumPage.Msg
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
            ( { model | page = AlbumPage updatedPageModel }
            , Cmd.map AlbumPageMsg updatedCmd
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
            ( { model | route = newRoute }, Cmd.none )
                |> initCurrentPage

        ( _, _ ) ->
            ( model, Cmd.none )


init : () -> Url -> Nav.Key -> ( Model, Cmd Msg )
init flags url navKey =
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

                Route.Album id ->
                    let
                        ( pageModel, pageCmds ) =
                            AlbumPage.init id
                    in
                    ( AlbumPage pageModel, Cmd.map AlbumPageMsg pageCmds )

                Route.HomePage ->
                    ( IndexPage, Cmd.none )
    in
    ( { model | page = currentPage }
    , Cmd.batch [ existingCmds, mappedPageCmds ]
    )


view : Model -> Document Msg
view model =
    { title = "Orkester"
    , body = [ currentView model |> toUnstyled ]
    }


currentView : Model -> Html Msg
currentView model =
    case model.page of
        NotFoundPage ->
            notFoundView

        AlbumPage albumModel ->
            AlbumPage.view albumModel
                |> map AlbumPageMsg

        IndexPage ->
            indexView


indexView : Html Msg
indexView =
    h1 [] [ text "Welcome!" ]


notFoundView : Html Msg
notFoundView =
    h3 [] [ text "Page was not found" ]


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
