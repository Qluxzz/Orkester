module Pages.Artist.Id_ exposing (Model, Msg, page)

import Api.Artist
import Effect exposing (Effect)
import Html
import Layouts
import Page exposing (Page)
import RemoteData
import Route exposing (Route)
import Route.Path as Path
import Shared
import View exposing (View)


page : Shared.Model -> Route { id : String } -> Page Model Msg
page shared route =
    Page.new
        { init = \_ -> init route.params.id
        , update = update
        , subscriptions = subscriptions
        , view = view
        }
        |> Page.withLayout toLayout


toLayout : Model -> Layouts.Layout Msg
toLayout _ =
    Layouts.Default {}



-- INIT


type alias Model =
    { artist : RemoteData.WebData Api.Artist.Artist }


init : String -> ( Model, Effect Msg )
init artistId =
    ( Model RemoteData.Loading
    , Effect.sendApiRequest
        { endpoint = "/api/v1/artist/" ++ artistId
        , decoder = Api.Artist.artistDecoder
        , onResponse = RemoteData.fromResult >> GotArtist
        }
    )



-- UPDATE


type Msg
    = GotArtist (RemoteData.WebData Api.Artist.Artist)


update : Msg -> Model -> ( Model, Effect Msg )
update msg model =
    case msg of
        GotArtist artist ->
            case artist of
                RemoteData.Success a ->
                    ( model
                    , Effect.replaceRoutePath (Path.Artist_Id__Name_ { id = String.fromInt a.id, name = a.urlName })
                    )

                _ ->
                    ( model, Effect.none )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Model -> View Msg
view model =
    { title = "Artist"
    , body = [ Html.text "Loading" ]
    }
