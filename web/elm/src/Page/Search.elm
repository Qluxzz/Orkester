module Page.Search exposing (Model, Msg, init, update, view)

import BaseUrl exposing (baseUrl)
import Css exposing (auto, displayFlex, flexBasis, flexGrow, flexShrink, int, listStyle, margin, margin2, marginBottom, marginTop, maxWidth, none, overflow, padding, padding2, px, textDecoration, underline)
import Html.Styled exposing (Html, a, div, h1, input, li, text, ul)
import Html.Styled.Attributes exposing (css, href, type_, value)
import Html.Styled.Events exposing (onInput)
import Http
import Json.Decode as Decode exposing (Decoder, list)
import Json.Decode.Pipeline exposing (required)
import RemoteData exposing (WebData)


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
    , searchPhrase : String
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
          , searchPhrase = p
          }
        , Cmd.none
        )

    else
        ( { searchResult = RemoteData.Loading
          , searchPhrase = p
          }
        , getSearchResult p
        )


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
    li [ css [ margin2 (px 5) (px 0), padding2 (px 5) (px 0), textDecoration underline ] ]
        [ a [ href link ] [ text entry.name ] ]


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


getSearchResult : String -> Cmd Msg
getSearchResult query =
    Http.get
        { url = baseUrl ++ "/api/v1/search/" ++ query
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
                    | searchPhrase = phrase
                    , searchResult = RemoteData.NotAsked
                  }
                , Cmd.none
                )

            else
                ( { model | searchPhrase = phrase, searchResult = RemoteData.Loading }, getSearchResult phrase )

        SearchResultsRecieved searchResult ->
            ( { model | searchResult = searchResult }, Cmd.none )


view : Model -> Html Msg
view model =
    div []
        [ input
            [ css
                [ flexGrow (int 1)
                ]
            , type_ "text"
            , value model.searchPhrase
            , onInput UpdateSearchPhrase
            ]
            []
        , div
            [ css
                [ displayFlex
                , overflow auto
                , marginTop (px 20)
                , marginBottom (px 20)
                ]
            ]
            (case model.searchResult of
                RemoteData.Success data ->
                    [ searchResultList TrackLink (List.map (\x -> { id = x.id, name = x.title, urlName = "" }) data.tracks)
                    , searchResultList AlbumLink data.albums
                    , searchResultList ArtistLink data.artists
                    ]

                RemoteData.Failure _ ->
                    [ h1 [] [ text "Search failed" ] ]

                RemoteData.NotAsked ->
                    [ h1 [] [ text "Start typing to search..." ] ]

                _ ->
                    []
            )
        ]
