module Pages.Album.Id_ exposing (Model, Msg, page)

import Api.Album
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
page _ route =
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
    { album : RemoteData.WebData Api.Album.Album }


init : String -> ( Model, Effect Msg )
init id =
    ( Model RemoteData.Loading
    , Effect.sendApiRequest
        { endpoint = "/api/v1/album/" ++ id
        , decoder = Api.Album.albumDecoder
        , onResponse = RemoteData.fromResult >> GotAlbum
        }
    )



-- UPDATE


type Msg
    = GotAlbum (RemoteData.WebData Api.Album.Album)


update : Msg -> Model -> ( Model, Effect Msg )
update msg model =
    case msg of
        GotAlbum album ->
            case album of
                RemoteData.Success a ->
                    ( { model | album = album }
                    , Effect.replaceRoutePath (Path.Album_Id__Name_ { id = String.fromInt a.id, name = a.urlName })
                    )

                _ ->
                    ( { model | album = album }
                    , Effect.none
                    )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none



-- VIEW


view : Model -> View Msg
view _ =
    { title = "Pages.Album.Id_"
    , body = [ Html.text "Loading..." ]
    }
