module QueueTests exposing (..)

import Expect
import Queue
import Test exposing (..)


queue : Queue.Queue number
queue =
    Queue.init (Just 1) (Just [ 2, 3, 4 ])


suite : Test
suite =
    describe "Queue"
        [ test "Current"
            (\_ ->
                Expect.equal (Just 1) (Queue.getCurrent queue)
            )
        , test "Previous"
            (\_ ->
                Expect.equal (Just 1) (queue |> Queue.previous |> Queue.getCurrent)
            )
        , test "Next"
            (\_ ->
                Expect.equal (Just 2) (queue |> Queue.next |> Queue.getCurrent)
            )
        , test "Next, then previous"
            (\_ ->
                Expect.equal (Just 1) (queue |> Queue.next |> Queue.previous |> Queue.getCurrent)
            )
        ]
