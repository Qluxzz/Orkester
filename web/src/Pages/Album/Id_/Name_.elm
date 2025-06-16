module Pages.Album.Id_.Name_ exposing (Model, Msg, page)

import Api.Album
import Components.Table
import Effect exposing (Effect)
import Html
import Html.Attributes
import Html.Events
import Html.Extra
import Layouts
import Page exposing (Page)
import RemoteData exposing (WebData)
import Route exposing (Route)
import Route.Path
import Shared
import Types.ReleaseDate
import Types.TrackId
import Types.TrackInfo
import Utilities.DurationDisplay
import Utilities.ErrorMessage
import Utilities.Icon as Icon
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
toLayout model =
    Layouts.Default {}



-- INIT


type alias Model =
    { album : WebData Api.Album.Album }


init : String -> ( Model, Effect Msg )
init albumId =
    ( Model RemoteData.Loading
    , Effect.sendApiRequest
        { endpoint = "/api/v1/album/" ++ albumId
        , decoder = Api.Album.albumDecoder
        , onResponse = RemoteData.fromResult >> GotAlbum
        }
    )



-- UPDATE


type Msg
    = GotAlbum (RemoteData.WebData Api.Album.Album)
    | PlayTrack Types.TrackInfo.Track
    | PlayTracks (List Types.TrackInfo.Track)
    | UnlikeTrack Types.TrackId.TrackId
    | LikeTrack Types.TrackId.TrackId


update : Msg -> Model -> ( Model, Effect Msg )
update msg model =
    case msg of
        GotAlbum album ->
            ( { model | album = album }
            , Effect.none
            )

        PlayTrack track ->
            ( model, Effect.playTrack track )

        PlayTracks tracks ->
            ( model, Effect.playTracks tracks )

        LikeTrack trackId ->
            ( { model
                | album =
                    RemoteData.map
                        (\x ->
                            { x
                                | tracks =
                                    List.map
                                        (\t ->
                                            if t.id == trackId then
                                                { t | liked = True }

                                            else
                                                t
                                        )
                                        x.tracks
                            }
                        )
                        model.album
              }
            , Effect.likeTrack trackId
            )

        UnlikeTrack trackId ->
            ( { model
                | album =
                    RemoteData.map
                        (\x ->
                            { x
                                | tracks =
                                    List.map
                                        (\t ->
                                            if t.id == trackId then
                                                { t | liked = False }

                                            else
                                                t
                                        )
                                        x.tracks
                            }
                        )
                        model.album
              }
            , Effect.unlikeTrack trackId
            )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Model -> View Msg
view model =
    case model.album of
        RemoteData.Success a ->
            { title = a.name ++ " • " ++ a.artist.name
            , body = [ albumView a ]
            }

        RemoteData.Loading ->
            { title = "Loading album", body = [ Html.text "Loading..." ] }

        RemoteData.Failure err ->
            { title = "Failed to load album", body = [ Utilities.ErrorMessage.errorMessage "Failed to load album" err ] }

        RemoteData.NotAsked ->
            { title = "Loading album", body = [ Html.text "Loading..." ] }


albumView : Api.Album.Album -> Html.Html Msg
albumView album =
    Html.section [ Html.Attributes.class "album-page" ]
        [ Html.div [ Html.Attributes.class "album-info" ]
            [ Html.Extra.picture []
                [ Html.img [ Html.Attributes.src (albumImageUrl album.id) ] []
                , playButton (PlayTracks (List.map (mapAlbumTrackToTrack album) album.tracks))
                ]
            , Html.div []
                [ Html.h1 [] [ Html.text album.name ]
                , Html.div [ Html.Attributes.class "info" ]
                    [ Html.a [ Route.Path.href (Route.Path.Artist_Id__Name_ { id = String.fromInt album.artist.id, name = album.artist.urlName }) ] [ Html.text album.artist.name ]
                    , Html.div [] [ Html.text (Types.ReleaseDate.formatReleaseDate album.released) ]
                    , Html.div [] [ Html.text (formatTracksDisplay album.tracks) ]
                    , Html.div [] [ Html.text (formatAlbumLength album.tracks) ]
                    ]
                ]
            ]
        , Html.div [ Html.Attributes.class "track-list" ]
            [ Components.Table.table
                [ Components.Table.clickableColumn "#" (.trackNumber >> String.fromInt >> Html.text) (mapAlbumTrackToTrack album >> PlayTrack)
                , Components.Table.defaultColumn "Title"
                    (\t ->
                        Html.div []
                            [ Html.div [ Html.Attributes.class "name" ] [ Html.p [] [ Html.text t.title ] ]
                            , Html.div [ Html.Attributes.class "artists" ] (formatTrackArtists t.artists)
                            ]
                    )
                    |> Components.Table.grow
                , Components.Table.clickableColumn ""
                    (\t ->
                        Html.text
                            (if t.liked then
                                "Liked"

                             else
                                "Like"
                            )
                    )
                    (\t ->
                        (if t.liked then
                            UnlikeTrack

                         else
                            LikeTrack
                        )
                            t.id
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


playButton : Msg -> Html.Html Msg
playButton msg =
    Html.button [ Html.Events.onClick msg, Html.Attributes.class "play-button" ]
        [ Html.img [ Html.Attributes.src (Icon.url Icon.Play) ] []
        ]


albumImageUrl : Int -> String
albumImageUrl id =
    "/api/v1/album/" ++ String.fromInt id ++ "/image"


formatTrackArtists : List Api.Album.Artist -> List (Html.Html msg)
formatTrackArtists artists =
    artists
        |> List.map (\artist -> Html.a [ Route.Path.href (Route.Path.Artist_Id__Name_ { id = String.fromInt artist.id, name = artist.urlName }) ] [ Html.text artist.name ])
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


mapAlbumTrackToTrack : Api.Album.Album -> Api.Album.Track -> Types.TrackInfo.Track
mapAlbumTrackToTrack album track =
    { id = track.id
    , title = track.title
    , length = track.length
    , liked = track.liked
    , album =
        { id = album.id
        , name = album.name
        , urlName = album.urlName
        }
    , artists = track.artists
    }
