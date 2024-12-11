module Pages.Search.Query_ exposing (Model, Msg, page)

import Api.Album exposing (Album)
import Api.Search
import Effect exposing (Effect)
import Html
import Html.Attributes
import Html.Events
import Http
import Layout exposing (Layout)
import Layouts
import Page exposing (Page)
import RemoteData exposing (WebData)
import Route exposing (Route)
import Route.Path exposing (Path(..))
import Shared
import Types.TrackId
import Types.TrackInfo
import Url exposing (Url)
import View exposing (View)


page : Shared.Model -> Route { query : String } -> Page Model Msg
page shared route =
    Page.new
        { init = \_ -> init route.params.query
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
    { searchResult : WebData Api.Search.SearchResult
    , search : String
    }


init : String -> ( Model, Effect Msg )
init search =
    if String.isEmpty search then
        ( { searchResult = RemoteData.NotAsked, search = "" }
        , Effect.focusElement "search-field"
        )

    else
        ( { searchResult = RemoteData.NotAsked, search = search }
        , Effect.batch
            [ Effect.sendApiRequest
                { endpoint = "/api/v1/search/" ++ Url.percentEncode search
                , decoder = Api.Search.searchResultDecoder
                , onResponse = RemoteData.fromResult >> SearchResultsReceived
                }
            , Effect.focusElement "search-field"
            ]
        )



-- UPDATE


type Msg
    = SearchResultsReceived (WebData Api.Search.SearchResult)
    | UpdateSearchPhrase String
    | PlayTrack Types.TrackInfo.Track


update : Msg -> Model -> ( Model, Effect Msg )
update msg model =
    case msg of
        SearchResultsReceived data ->
            ( { model | searchResult = data }, Effect.none )

        UpdateSearchPhrase phrase ->
            if String.isEmpty phrase then
                ( model
                , -- Go to search overview
                  Effect.pushRoutePath Route.Path.Search
                )

            else
                ( { model | search = phrase }
                , Effect.batch
                    [ Effect.sendApiRequest
                        { endpoint = "/api/v1/search/" ++ Url.percentEncode phrase
                        , decoder = Api.Search.searchResultDecoder
                        , onResponse = RemoteData.fromResult >> SearchResultsReceived
                        }
                    , Effect.replaceRoutePath (Search_Query_ { query = phrase })
                    ]
                )

        PlayTrack track ->
            ( model, Effect.playTrack track )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Model -> View Msg
view model =
    { title = "Search " ++ model.search
    , body =
        [ Html.div [ Html.Attributes.class "search-results-page" ]
            [ Html.input [ Html.Attributes.type_ "text", Html.Attributes.value model.search, Html.Events.onInput UpdateSearchPhrase, Html.Attributes.id "search-field" ] []
            , case model.searchResult of
                RemoteData.Success data ->
                    Html.div [ Html.Attributes.class "search-results" ]
                        [ Html.div []
                            [ Html.ul []
                                (Html.h1 [] [ Html.text "Tracks" ]
                                    :: (if List.isEmpty data.tracks then
                                            [ Html.li [] [ Html.text "No tracks found!" ] ]

                                        else
                                            List.map (\t -> Html.li [ Html.Events.onClick (PlayTrack t), Html.Attributes.class "track-title" ] [ Html.text t.title ]) data.tracks
                                       )
                                )
                            ]
                        , Html.div []
                            [ Html.ul []
                                (Html.h1 [] [ Html.text "Albums" ]
                                    :: (if List.isEmpty data.albums then
                                            [ Html.li [] [ Html.text "No albums found!" ] ]

                                        else
                                            List.map (\album -> Html.li [] [ Html.a [ Html.Attributes.href ("/album/" ++ String.fromInt album.id ++ "/" ++ album.urlName) ] [ Html.text album.name ] ]) data.albums
                                       )
                                )
                            ]
                        , Html.div []
                            [ Html.ul []
                                (Html.h1 [] [ Html.text "Artists" ]
                                    :: (if List.isEmpty data.artists then
                                            [ Html.li [] [ Html.text "No artists found!" ] ]

                                        else
                                            List.map (\artist -> Html.li [] [ Html.a [ Html.Attributes.href ("/artist/" ++ String.fromInt artist.id ++ "/" ++ artist.urlName) ] [ Html.text artist.name ] ]) data.artists
                                       )
                                )
                            ]
                        ]

                _ ->
                    Html.text ""
            ]
        ]
    }
