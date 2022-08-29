module Page.Search exposing (Model, Msg(..), init, update, view)

import ApiBaseUrl exposing (apiBaseUrl)
import Browser.Dom
import Css exposing (auto, column, displayFlex, flexBasis, flexDirection, flexGrow, flexShrink, hidden, int, listStyle, margin, margin2, marginTop, maxWidth, none, overflow, padding, padding2, padding4, paddingLeft, paddingRight, px, textDecoration, underline)
import ErrorMessage exposing (errorMessage)
import Html.Styled exposing (Html, a, div, h1, input, li, text, ul)
import Html.Styled.Attributes exposing (css, href, id, type_, value)
import Html.Styled.Events exposing (onClick, onInput)
import Http
import JSPlayer
import Json.Decode as Decode exposing (Decoder, list)
import Json.Decode.Pipeline exposing (required)
import RemoteData exposing (WebData)
import Task


type alias IdNameAndUrlName =
    { id : Int
    , name : String
    , urlName : String
    }


type alias Album =
    IdNameAndUrlName


type alias Artist =
    IdNameAndUrlName


type alias Track =
    { id : Int
    , title : String
    }


type Type
    = AlbumLink
    | ArtistLink
    | TrackLink


type alias SearchResult =
    { albums : List Album
    , artists : List Artist
    , tracks : List Track
    }


type alias Model =
    { searchResult : WebData SearchResult
    , searchPhrase : Maybe String
    }


init : Maybe String -> ( Model, Cmd Msg )
init phrase =
    let
        p : String
        p =
            Maybe.withDefault "" phrase
    in
    if String.isEmpty p then
        ( { searchResult = RemoteData.NotAsked
          , searchPhrase = Nothing
          }
        , focusSearchField
        )

    else
        ( { searchResult = RemoteData.Loading
          , searchPhrase = Just p
          }
        , Cmd.batch
            [ getSearchResult p
            , focusSearchField
            ]
        )


focusSearchField : Cmd Msg
focusSearchField =
    Task.attempt (\_ -> FocusedSearchField) (Browser.Dom.focus "search-field")


searchResultList : Type -> List { a | id : Int, name : String, urlName : String } -> Html Msg
searchResultList type_ entries =
    let
        result =
            if List.isEmpty entries then
                [ li [] [ text "No entry matched the phrase" ] ]

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

                    TrackLink ->
                        "Tracks"

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

                TrackLink ->
                    "track"

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
            (case type_ of
                TrackLink ->
                    [ onClick (Player (JSPlayer.PlayTrack { id = entry.id }))
                    , href ""
                    ]

                _ ->
                    [ href link ]
            )
            [ text entry.name ]
        ]


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


trackDecoder : Decoder Track
trackDecoder =
    Decode.succeed Track
        |> required "id" Decode.int
        |> required "title" Decode.string


searchResultDecoder : Decoder SearchResult
searchResultDecoder =
    Decode.succeed SearchResult
        |> required "albums" (list albumDecoder)
        |> required "artists" (list artistDecoder)
        |> required "tracks" (list trackDecoder)


type Msg
    = UpdateSearchPhrase String
    | SearchResultsRecieved (WebData SearchResult)
    | FocusedSearchField
    | Player JSPlayer.Msg


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
                ( { model
                    | searchPhrase = Nothing
                    , searchResult = RemoteData.NotAsked
                  }
                , Cmd.none
                )

            else
                ( { model | searchPhrase = Just phrase, searchResult = RemoteData.Loading }, getSearchResult phrase )

        SearchResultsRecieved searchResult ->
            ( { model | searchResult = searchResult }, Cmd.none )

        FocusedSearchField ->
            ( model, Cmd.none )

        Player _ ->
            ( model, Cmd.none )


view : Model -> Html Msg
view model =
    div [ css [ overflow hidden, displayFlex, flexDirection column ] ]
        [ input
            [ css
                [ flexGrow (int 1)
                ]
            , type_ "text"
            , value (Maybe.withDefault "" model.searchPhrase)
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
                    [ searchResultList TrackLink (List.map (\x -> { id = x.id, name = x.title, urlName = "" }) data.tracks)
                    , searchResultList AlbumLink data.albums
                    , searchResultList ArtistLink data.artists
                    ]

                RemoteData.Failure error ->
                    [ errorMessage "Search failed" error ]

                RemoteData.NotAsked ->
                    [ h1 [] [ text "Start typing to search..." ] ]

                _ ->
                    []
            )
        ]
