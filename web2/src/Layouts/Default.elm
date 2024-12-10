module Layouts.Default exposing (Model, Msg, Props, layout)

import Effect exposing (Effect)
import Html exposing (Html)
import Html.Attributes exposing (class)
import Layout exposing (Layout)
import Route exposing (Route)
import Shared
import Types.TrackQueue
import View exposing (View)


type alias Props =
    {}


layout : Props -> Shared.Model -> Route () -> Layout () Model Msg contentMsg
layout props shared route =
    Layout.new
        { init = init
        , update = update
        , view = view shared
        , subscriptions = subscriptions
        }



-- MODEL


type alias Model =
    {}


init : () -> ( Model, Effect Msg )
init _ =
    ( {}
    , Effect.none
    )



-- UPDATE


type Msg
    = ReplaceMe


update : Msg -> Model -> ( Model, Effect Msg )
update msg model =
    case msg of
        ReplaceMe ->
            ( model
            , Effect.none
            )


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none



-- VIEW


view : Shared.Model -> { toContentMsg : Msg -> contentMsg, content : View contentMsg, model : Model } -> View contentMsg
view shared { toContentMsg, model, content } =
    let
        trackTitle =
            Types.TrackQueue.getActiveTrack shared.queue |> Maybe.map (.track >> .title) |> Maybe.withDefault "Nothing is playing!"
    in
    { title = content.title
    , body =
        [ Html.aside [ Html.Attributes.class "sidebar" ]
            [ Html.ul [] [ Html.li [] [ Html.a [ Html.Attributes.href "/search" ] [ Html.text "Search" ] ] ]
            ]
        , Html.main_ [] content.body
        , Html.aside [ Html.Attributes.class "queue" ] [ Html.text "Queue" ]
        , Html.div [ Html.Attributes.class "player-bar" ] [ Html.text trackTitle ]
        ]
    }
