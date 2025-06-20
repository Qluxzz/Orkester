module Pages.Search exposing (Model, Msg, page)

import Effect exposing (Effect)
import Html
import Html.Attributes
import Html.Events
import Layouts
import Page exposing (Page)
import Route exposing (Route)
import Route.Path exposing (Path(..))
import Shared
import View exposing (View)


page : Shared.Model -> Route () -> Page Model Msg
page shared route =
    Page.new
        { init = init
        , update = update
        , subscriptions = subscriptions
        , view = view
        }
        |> Page.withLayout toLayout


toLayout : Model -> Layouts.Layout Msg
toLayout model =
    Layouts.Default {}



-- INIT


type alias Model =
    {}


init : () -> ( Model, Effect Msg )
init () =
    ( {}
    , Effect.focusElement "search-field"
    )



-- UPDATE


type Msg
    = UpdateSearchPhrase String


update : Msg -> Model -> ( Model, Effect Msg )
update msg model =
    case msg of
        UpdateSearchPhrase phrase ->
            -- Go to search result page
            ( model, Effect.pushRoutePath (Route.Path.Search_Query_ { query = phrase }) )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Model -> View Msg
view model =
    { title = "Search"
    , body =
        [ Html.div [ Html.Attributes.class "search-start-page" ]
            [ Html.h2 [] [ Html.text "Search" ]
            ]
        ]
    }
