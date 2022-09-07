module QueueTests exposing (..)

import Expect
import Queue exposing (Queue, Repeat(..))
import Test exposing (..)


queue : Queue number
queue =
    Queue.init
        { current = Just 1
        , future = Just [ 2, 3, 4 ]
        }


suite : Test
suite =
    describe "Queue"
        [ test "Current"
            (\_ ->
                Expect.equal (Just 1) (Queue.getCurrent queue)
            )
        , test "Previous"
            (\_ ->
                Expect.equal
                    (Just 1)
                    (queue
                        |> Queue.previous
                        |> Queue.getCurrent
                    )
            )
        , test "Next"
            (\_ ->
                Expect.equal
                    (Just 2)
                    (queue
                        |> (\q ->
                                Queue.next q RepeatOff
                                    |> Queue.getCurrent
                           )
                    )
            )
        , test "Next, then previous"
            (\_ ->
                Expect.equal
                    (Just 1)
                    (queue
                        |> (\q ->
                                Queue.next q RepeatOff
                                    |> Queue.previous
                                    |> Queue.getCurrent
                           )
                    )
            )
        , test "RepeatAll queue loops around"
            (\_ ->
                let
                    next : Queue a -> Queue a
                    next q =
                        Queue.next q RepeatAll
                in
                Expect.equal
                    (Just 1)
                    (queue
                        |> next
                        |> next
                        |> next
                        |> next
                        |> Queue.getCurrent
                    )
            )
        ]
