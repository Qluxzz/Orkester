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
import Shared
import Types.TrackInfo
import Utilities.AlbumUrl
import Utilities.ArtistUrl
import Utilities.Date
import Utilities.DurationDisplay
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
                likedTracksTable tracks

            RemoteData.Loading ->
                Html.text "Loading..."

            RemoteData.Failure err ->
                Html.text (Debug.toString err)

            RemoteData.NotAsked ->
                Html.text ""
        ]
    }


likedTracksTable tracks =
    Components.Table.table
        [ Components.Table.clickableColumn "#" (.trackNumber >> String.fromInt >> Html.text) PlayTrack
        , Components.Table.defaultColumn "Title"
            (\t ->
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
        , Components.Table.linkColumn "Album" (\t -> { url = "/album" ++ String.fromInt t.album.id ++ "/" ++ t.album.urlName, title = t.album.name })
        , Components.Table.textColumn "Date added" (.dateAdded >> formatDate)
        , Components.Table.textColumn "Duration" (\t -> Utilities.DurationDisplay.durationDisplay t.length)
        ]
        tracks


formatTrackArtists : List Api.LikedTracks.Artist -> List (Html.Html msg)
formatTrackArtists artists =
    artists
        |> List.map (\artist -> Html.a [ Html.Attributes.href (Utilities.ArtistUrl.artistUrl artist) ] [ Html.text artist.name ])
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
