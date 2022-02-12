module Main exposing (..)

import Browser
import Css exposing (..)
import Css.Global
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, type_, value)
import Html.Styled.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, Error(..))


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


type alias SearchResult =
    { albums : List Album
    , artists : List Artist
    , tracks : List Track
    }


type alias Model =
    { searchResult : SearchResult
    , searchPhrase : String
    }


init : () -> ( Model, Cmd Msg )
init _ =
    ( { searchResult =
            { albums = []
            , artists = []
            , tracks = []
            }
      , searchPhrase = ""
      }
    , Cmd.none
    )


globalStyle : Html msg
globalStyle =
    Css.Global.global
        [ Css.Global.html
            [ height (pct 100)
            ]
        , Css.Global.body
            [ height (pct 100)
            , color (hex "#FFF")
            , fontFamily sansSerif
            , overflow hidden
            ]
        , Css.Global.h1
            [ margin (px 0)
            , fontSize (px 32)
            ]
        ]


view : Model -> Html Msg
view model =
    let
        curriedSearchList =
            searchResultList model.searchPhrase
    in
    div [ css [ height (pct 100), displayFlex, flexDirection column ] ]
        [ globalStyle
        , div
            [ css
                [ displayFlex
                , flexDirection row
                , backgroundColor (hex "#222")
                , height (pct 100)
                , overflow hidden
                ]
            ]
            [ aside
                [ css
                    [ padding (px 10)
                    , backgroundColor (hex "#333")
                    , width (px 200)
                    , flexShrink (int 0)
                    ]
                ]
                [ text "Sidebar" ]
            , section
                [ css
                    [ displayFlex
                    , flexDirection column
                    , padding (px 20)
                    , flexGrow (int 1)
                    ]
                ]
                [ div
                    [ css
                        [ displayFlex
                        , flexDirection column
                        , overflow auto
                        ]
                    ]
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
                        [ curriedSearchList "Tracks" (List.map (\x -> { id = x.id, name = x.title, urlName = "" }) model.searchResult.tracks)
                        , curriedSearchList "Albums" model.searchResult.albums
                        , curriedSearchList "Artists" model.searchResult.artists
                        ]
                    ]
                , div [] [ text "Main content" ]
                ]
            ]
        , div [ css [ backgroundColor (hex "#333"), padding (px 10) ] ]
            [ text "Nothing is currently playing..."
            ]
        ]


filter : String -> { a | name : String } -> Bool
filter searchPhrase entry =
    if String.isEmpty searchPhrase then
        True

    else
        String.contains (String.toLower searchPhrase) (String.toLower entry.name)


searchResultList : String -> String -> List { a | id : Int, name : String, urlName : String } -> Html Msg
searchResultList phrase title entries =
    let
        filteredEntries =
            List.filter (filter phrase) entries

        result =
            if List.isEmpty filteredEntries then
                [ li [] [ text "No entry matched the prhase" ] ]

            else
                List.map searchResultEntry filteredEntries
    in
    div
        [ css
            [ flexGrow (int 1)
            , flexShrink (int 1)
            , flexBasis (px 0)
            , maxWidth (px 300)
            ]
        ]
        [ h1 [] [ text title ]
        , ul
            [ css
                [ listStyle none
                , padding (px 0)
                , margin (px 0)
                ]
            ]
            result
        ]


searchResultEntry : { a | id : Int, name : String, urlName : String } -> Html Msg
searchResultEntry entry =
    li [ css [ margin2 (px 5) (px 0), padding2 (px 5) (px 0), textDecoration underline ] ] [ text entry.name ]


albumDecoder : Decoder Album
albumDecoder =
    Json.Decode.map3 IdNameAndUrlName
        (Json.Decode.field "id" Json.Decode.int)
        (Json.Decode.field "name" Json.Decode.string)
        (Json.Decode.field "urlName" Json.Decode.string)


artistDecoder : Decoder Artist
artistDecoder =
    Json.Decode.map3 IdNameAndUrlName
        (Json.Decode.field "id" Json.Decode.int)
        (Json.Decode.field "name" Json.Decode.string)
        (Json.Decode.field "urlName" Json.Decode.string)


trackDecoder : Decoder Track
trackDecoder =
    Json.Decode.map2 Track
        (Json.Decode.field "id" Json.Decode.int)
        (Json.Decode.field "title" Json.Decode.string)


searchResultDecoder : Decoder SearchResult
searchResultDecoder =
    Json.Decode.map3 SearchResult
        (Json.Decode.field "albums" (Json.Decode.list albumDecoder))
        (Json.Decode.field "artists" (Json.Decode.list artistDecoder))
        (Json.Decode.field "tracks" (Json.Decode.list trackDecoder))


type Msg
    = UpdateSearchPhrase String
    | DataRecieved (Result Http.Error SearchResult)


getSearchResult : String -> Cmd Msg
getSearchResult query =
    Http.get
        { url = "http://localhost:42000/api/v1/search/" ++ query
        , expect = Http.expectJson DataRecieved searchResultDecoder
        }


update : Msg -> Model -> ( Model, Cmd Msg )
update message model =
    case message of
        UpdateSearchPhrase phrase ->
            ( { model | searchPhrase = phrase }, getSearchResult phrase )

        DataRecieved (Ok searchResult) ->
            ( { model | searchResult = searchResult }, Cmd.none )

        DataRecieved (Err _) ->
            ( model, Cmd.none )


main : Program () Model Msg
main =
    Browser.element
        { init = init
        , view = \model -> toUnstyled (view model)
        , update = update
        , subscriptions = \_ -> Sub.none
        }
