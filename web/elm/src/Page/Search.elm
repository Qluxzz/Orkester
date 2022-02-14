module Page.Search exposing (..)


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
            if String.isEmpty phrase then
                ( { model | searchPhrase = phrase }, Cmd.none )

            else
                ( { model | searchPhrase = phrase }, getSearchResult phrase )

        DataRecieved (Ok searchResult) ->
            ( { model | searchResult = searchResult }, Cmd.none )

        DataRecieved (Err _) ->
            ( model, Cmd.none )


searchView : Model -> Html Msg
searchView model =
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
