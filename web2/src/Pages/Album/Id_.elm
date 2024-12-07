module Pages.Album.Id_ exposing (Model, Msg, page)

import Api.Album
import Components.Table
import Effect exposing (Effect)
import Html
import Html.Attributes
import Layouts
import Page exposing (Page)
import RemoteData
import Route exposing (Route)
import Shared
import Types.ReleaseDate exposing (formatReleaseDate)
import Utilities.DurationDisplay
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
toLayout model =
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
            ( { model | album = album }
            , Effect.none
            )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Model -> View Msg
view model =
    { title = "Pages.Album.Id_"
    , body =
        case model.album of
            RemoteData.Success album ->
                [ albumView album ]

            _ ->
                [ Html.text "Loading" ]
    }


albumView : Api.Album.Album -> Html.Html Msg
albumView album =
    Html.section [ Html.Attributes.class "album" ]
        [ Html.div [ Html.Attributes.class "album-info"]
            [ picture []
                [ Html.img [ Html.Attributes.src (albumImageUrl album.id) ] []
                ]
            , Html.div []
                [ Html.h1 [] [ Html.text album.name ]
                , Html.div [ Html.Attributes.class "info"]
                    [ Html.div [] [ Html.text (formatReleaseDate album.released) ]
                    , Html.div [] [ Html.text (formatTracksDisplay album.tracks) ]
                    , Html.div [] [ Html.text (formatAlbumLength album.tracks) ]
                    ]
                ]
            ]
        , Html.div []
            [ Components.Table.table
                [ Components.Table.textColumn "#" (.trackNumber >> String.fromInt)
                , Components.Table.defaultColumn "Title"
                    (\t ->
                        Html.div []
                            [ Html.div [ Html.Attributes.class "name" ] [ Html.p [] [ Html.text t.title ] ]
                            , Html.div [ Html.Attributes.class "artists" ] (formatTrackArtists t.artists)
                            ]
                    )
                    |> Components.Table.grow
                    |> Components.Table.alignHeader Components.Table.Left
                , Components.Table.textColumn ""
                    (\t ->
                        if t.liked then
                            "Liked"

                        else
                            "Like"
                    )
                    |> Components.Table.alignHeader Components.Table.Center
                    |> Components.Table.alignData Components.Table.Center
                , Components.Table.textColumn "Duration" (.length >> Utilities.DurationDisplay.durationDisplay)
                    |> Components.Table.alignHeader Components.Table.Center
                    |> Components.Table.alignData Components.Table.Center
                ]
                album.tracks
            ]
        ]


picture : List (Html.Attribute msg) -> List (Html.Html msg) -> Html.Html msg
picture =
    Html.node "picture"


albumImageUrl : Int -> String
albumImageUrl id =
    "/api/v1/album/" ++ String.fromInt id ++ "/image"


formatTrackArtists : List Api.Album.Artist -> List (Html.Html msg)
formatTrackArtists artists =
    artists
        |> List.map (\artist -> Html.a [ Html.Attributes.href (artistUrl artist) ] [ Html.text artist.name ])
        |> List.intersperse (Html.span [] [ Html.text ", " ])


formatTracksDisplay : List Api.Album.Track -> String
formatTracksDisplay tracks =
    let
        amountOfTracks =
            tracks |> List.length

        suffix =
            if amountOfTracks /= 1 then
                "s"

            else
                ""
    in
    String.fromInt amountOfTracks ++ " track" ++ suffix


formatAlbumLength : List Api.Album.Track -> String
formatAlbumLength tracks =
    tracks
        |> List.map .length
        |> List.foldl (+) 0
        |> Utilities.DurationDisplay.durationDisplay


artistUrl : { r | id : Int, urlName : String } -> String
artistUrl artist =
    "/artist/" ++ String.fromInt artist.id ++ "/" ++ artist.urlName
