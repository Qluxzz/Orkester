module Pages.Artist.Id_.Name_ exposing (Model, Msg, page)

import Api.Artist
import Effect exposing (Effect)
import Html
import Html.Attributes
import Html.Extra
import Layouts
import Page exposing (Page)
import RemoteData
import Route exposing (Route)
import Route.Path
import Shared
import Types.ReleaseDate exposing (ReleaseDate(..))
import Utilities.AlbumUrl
import Utilities.ErrorMessage
import View exposing (View)


page : Shared.Model -> Route { id : String, name : String } -> Page Model Msg
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
            ( { model | artist = artist }
            , Effect.none
            )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Model -> View Msg
view model =
    case model.artist of
        RemoteData.Success a ->
            { title = a.name
            , body = [ artistView a ]
            }

        RemoteData.Loading ->
            { title = "Loading artist", body = [] }

        RemoteData.Failure err ->
            { title = "Failed to load artist", body = [ Utilities.ErrorMessage.errorMessage "Failed to load album" err ] }

        RemoteData.NotAsked ->
            { title = "Loading artist", body = [] }


artistView : Api.Artist.Artist -> Html.Html msg
artistView artist =
    Html.section [ Html.Attributes.class "artist-page" ]
        [ Html.h1 [] [ Html.text artist.name ]
        , Html.div [ Html.Attributes.class "artist-albums" ] (List.map artistAlbumView artist.albums)
        ]


artistAlbumView : Api.Artist.Album -> Html.Html msg
artistAlbumView album =
    Html.a
        [ Route.Path.href (Route.Path.Album_Id__Name_ { id = String.fromInt album.id, name = album.urlName })
        ]
        [ Html.div [ Html.Attributes.class "artist-album" ]
            [ Html.Extra.picture []
                [ Html.img [ Html.Attributes.src (Utilities.AlbumUrl.albumImageUrl album) ] []
                ]
            , Html.p [] [ Html.text album.name ]
            , Html.p [] [ Html.text (releaseYear album.released) ]
            ]
        ]


releaseYear : ReleaseDate -> String
releaseYear releaseDate =
    String.fromInt
        (case releaseDate of
            Year year ->
                year

            Month { year } ->
                year

            Date { year } ->
                year
        )
