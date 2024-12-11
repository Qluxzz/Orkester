module Pages.LikedTracks exposing (Model, Msg, page)

import Api.LikedTracks
import Components.Table
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
import Types.TrackInfo
import Utilities.AlbumUrl
import Utilities.Date
import Utilities.DurationDisplay
import Utilities.ErrorMessage
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
    { likedTracks : RemoteData.WebData (List Api.LikedTracks.Track) }


init : () -> ( Model, Effect Msg )
init () =
    ( Model RemoteData.Loading
    , Effect.sendApiRequest
        { endpoint = "/api/v1/playlist/liked"
        , decoder = Api.LikedTracks.tracksDecoder
        , onResponse = RemoteData.fromResult >> GotLikedTracks
        }
    )



-- UPDATE


type Msg
    = GotLikedTracks (RemoteData.WebData (List Api.LikedTracks.Track))
    | PlayTrack Api.LikedTracks.Track


update : Msg -> Model -> ( Model, Effect Msg )
update msg model =
    case msg of
        GotLikedTracks result ->
            ( { model | likedTracks = result }
            , Effect.none
            )

        PlayTrack track ->
            ( model, Effect.playTrack (toTrack track) )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Model -> View Msg
view model =
    { title = "Liked tracks"
    , body =
        [ Html.h1 [] [ Html.text "Liked tracks" ]
        , case model.likedTracks of
            RemoteData.Success tracks ->
                if List.isEmpty tracks then
                    Html.text "You haven't liked any tracks yet"

                else
                    likedTracksTable tracks

            RemoteData.Loading ->
                Html.text "Loading..."

            RemoteData.Failure err ->
                Utilities.ErrorMessage.errorMessage "Failed to load liked tracks" err

            RemoteData.NotAsked ->
                Html.text ""
        ]
    }


likedTracksTable tracks =
    Components.Table.table
        [ Components.Table.clickableColumn "#" (Tuple.first >> (+) 1 >> String.fromInt >> Html.text) (\( _, t ) -> PlayTrack t)
        , Components.Table.defaultColumn "Title"
            (\( _, t ) ->
                Html.div [ Html.Attributes.class "track-title" ]
                    [ Html.Extra.picture [ Html.Attributes.class "album-cover" ]
                        [ Html.img [ Html.Attributes.src (Utilities.AlbumUrl.albumImageUrl t.album) ] []
                        ]
                    , Html.div []
                        [ Html.div
                            [ Html.Attributes.class "name" ]
                            [ Html.p [] [ Html.text t.title ] ]
                        , Html.div [ Html.Attributes.class "artists" ] (formatTrackArtists t.artists)
                        ]
                    ]
            )
        , Components.Table.linkColumn "Album" (\( _, t ) -> { path = Route.Path.Album_Id__Name_ { id = String.fromInt t.album.id, name = t.album.urlName }, title = t.album.name })
        , Components.Table.textColumn "Date added" (Tuple.second >> .dateAdded >> formatDate)
        , Components.Table.textColumn "Duration" (\( _, t ) -> Utilities.DurationDisplay.durationDisplay t.length)
        ]
        (List.indexedMap Tuple.pair tracks)


formatTrackArtists : List Api.LikedTracks.Artist -> List (Html.Html msg)
formatTrackArtists artists =
    artists
        |> List.map (\artist -> Html.a [ Route.Path.href (Route.Path.Artist_Id__Name_ { id = String.fromInt artist.id, name = artist.urlName }) ] [ Html.text artist.name ])
        |> List.intersperse (Html.span [] [ Html.text ", " ])


toTrack : Api.LikedTracks.Track -> Types.TrackInfo.Track
toTrack track =
    { id = track.id
    , title = track.title
    , length = track.length
    , liked = track.liked
    , album = track.album
    , artists = track.artists
    }


formatDate : Utilities.Date.Date -> String
formatDate { year, month, date, hour, minute } =
    year ++ "-" ++ month ++ "-" ++ date ++ " " ++ hour ++ ":" ++ minute


toPaddedString : Int -> String
toPaddedString =
    String.fromInt >> String.padLeft 2 '0'
