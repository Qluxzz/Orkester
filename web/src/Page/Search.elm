module Page.Search exposing (Model, Msg(..), init, update, view)

import Browser.Dom
import Css exposing (auto, column, cursor, displayFlex, flexBasis, flexDirection, flexGrow, flexShrink, hidden, int, listStyle, margin, marginTop, maxWidth, none, overflow, padding, paddingLeft, paddingRight, pointer, px, textDecoration, underline)
import Html.Styled exposing (Html, a, div, h1, input, li, text, ul)
import Html.Styled.Attributes exposing (css, href, id, type_, value)
import Html.Styled.Events exposing (onClick, onInput)
import Http
import Json.Decode as Decode exposing (Decoder, list)
import Json.Decode.Pipeline exposing (required)
import Process
import RemoteData exposing (WebData)
import Task
import Types.TrackInfo
import Utilities.ApiBaseUrl exposing (apiBaseUrl)
import Utilities.DelayedLoader
import Utilities.ErrorMessage exposing (errorMessage)


type alias IdNameAndUrlName =
    { id : Int
    , name : String
    , urlName : String
    }


type alias Album =
    IdNameAndUrlName


type alias Artist =
    IdNameAndUrlName


type Type
    = AlbumLink
    | ArtistLink


type alias SearchResult =
    { albums : List Album
    , artists : List Artist
    , tracks : List Types.TrackInfo.Track
    }


type alias Model =
    { searchResult : WebData SearchResult
    , search : String
    }


init : Maybe String -> ( Model, Cmd Msg )
init search =
    let
        p : String
        p =
            Maybe.withDefault "" search
                |> String.replace "%20" " "
    in
    if String.isEmpty p then
        ( { searchResult = RemoteData.NotAsked
          , search = ""
          }
        , focusSearchField
        )

    else
        ( { searchResult = RemoteData.NotAsked
          , search = p
          }
        , Cmd.batch
            [ debounceSearch p
            , focusSearchField
            ]
        )


debounceSearch : String -> Cmd Msg
debounceSearch search =
    Process.sleep 250 |> Task.perform (\_ -> DebouncedSearch search)


focusSearchField : Cmd Msg
focusSearchField =
    Task.attempt (\_ -> FocusedSearchField) (Browser.Dom.focus "search-field")


albumDecoder : Decoder Album
albumDecoder =
    Decode.succeed IdNameAndUrlName
        |> required "id" Decode.int
        |> required "name" Decode.string
        |> required "urlName" Decode.string


artistDecoder : Decoder Artist
artistDecoder =
    Decode.succeed IdNameAndUrlName
        |> required "id" Decode.int
        |> required "name" Decode.string
        |> required "urlName" Decode.string


searchResultDecoder : Decoder SearchResult
searchResultDecoder =
    Decode.succeed SearchResult
        |> required "albums" (list albumDecoder)
        |> required "artists" (list artistDecoder)
        |> required "tracks" (list Types.TrackInfo.trackInfoDecoder)


type Msg
    = UpdateSearchPhrase String
    | DebouncedSearch String
    | SearchResultsRecieved (WebData SearchResult)
    | FocusedSearchField
    | PlayTrack Types.TrackInfo.Track


getSearchResult : String -> Cmd Msg
getSearchResult query =
    Http.get
        { url = apiBaseUrl ++ "/api/v1/search/" ++ query
        , expect =
            searchResultDecoder
                |> Http.expectJson (RemoteData.fromResult >> SearchResultsRecieved)
        }


update : Msg -> Model -> ( Model, Cmd Msg )
update message model =
    case message of
        UpdateSearchPhrase phrase ->
            if String.isEmpty phrase then
                ( model
                , Cmd.none
                )

            else
                ( { model | search = phrase }
                , debounceSearch phrase
                )

        DebouncedSearch search ->
            if model.search == search then
                ( { model | searchResult = RemoteData.Loading }, getSearchResult search )

            else
                ( model, Cmd.none )

        SearchResultsRecieved searchResult ->
            ( { model | searchResult = searchResult }, Cmd.none )

        FocusedSearchField ->
            ( model, Cmd.none )

        {- this case is handled by the update method in Main.elm -}
        PlayTrack _ ->
            ( model, Cmd.none )


view : Model -> Html Msg
view model =
    div [ css [ overflow hidden, displayFlex, flexDirection column ] ]
        [ input
            [ css
                [ flexGrow (int 1)
                ]
            , type_ "text"
            , value model.search
            , onInput UpdateSearchPhrase
            , id "search-field"
            ]
            []
        , div
            [ css
                [ displayFlex
                , overflow auto
                , marginTop (px 20)
                ]
            ]
            (case model.searchResult of
                RemoteData.Success data ->
                    [ div
                        [ css
                            [ flexGrow (int 1)
                            , flexShrink (int 1)
                            , flexBasis (px 0)
                            , maxWidth (px 300)
                            , paddingLeft (px 5)
                            , paddingRight (px 5)
                            ]
                        ]
                        [ h1 []
                            [ text "Tracks"
                            ]
                        , ul
                            [ css
                                [ listStyle none
                                , padding (px 0)
                                , margin (px 0)
                                ]
                            ]
                            (if List.isEmpty data.tracks then
                                [ li [ css [ marginTop (px 10) ] ] [ text "No entry matched the phrase" ] ]

                             else
                                List.map
                                    (\t ->
                                        li
                                            [ css
                                                [ marginTop (px 10)
                                                , textDecoration underline
                                                , cursor pointer
                                                ]
                                            , onClick (PlayTrack t)
                                            ]
                                            [ text t.title ]
                                    )
                                    data.tracks
                            )
                        ]
                    , searchResultList AlbumLink data.albums
                    , searchResultList ArtistLink data.artists
                    ]

                RemoteData.Failure error ->
                    [ errorMessage "Search failed" error ]

                RemoteData.NotAsked ->
                    [ h1 [] [ text "Start typing to search..." ] ]

                RemoteData.Loading ->
                    [ Utilities.DelayedLoader.default (Css.ms 500) ]
            )
        ]


searchResultList : Type -> List { a | id : Int, name : String, urlName : String } -> Html Msg
searchResultList type_ entries =
    let
        result =
            if List.isEmpty entries then
                [ li [ css [ marginTop (px 10) ] ] [ text "No entry matched the phrase" ] ]

            else
                List.map (searchResultEntry type_) entries
    in
    div
        [ css
            [ flexGrow (int 1)
            , flexShrink (int 1)
            , flexBasis (px 0)
            , maxWidth (px 300)
            , paddingLeft (px 5)
            , paddingRight (px 5)
            ]
        ]
        [ h1 []
            [ text
                (case type_ of
                    AlbumLink ->
                        "Albums"

                    ArtistLink ->
                        "Artists"
                )
            ]
        , ul
            [ css
                [ listStyle none
                , padding (px 0)
                , margin (px 0)
                ]
            ]
            result
        ]


searchResultEntry : Type -> { a | id : Int, name : String, urlName : String } -> Html Msg
searchResultEntry type_ entry =
    let
        linkType : String
        linkType =
            case type_ of
                ArtistLink ->
                    "artist"

                AlbumLink ->
                    "album"

        link : String
        link =
            "/"
                ++ linkType
                ++ "/"
                ++ String.fromInt entry.id
                ++ "/"
                ++ entry.urlName
    in
    li [ css [ marginTop (px 10), textDecoration underline ] ]
        [ a
            [ href link ]
            [ text entry.name ]
        ]
