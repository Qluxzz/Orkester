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


nextRepeatOff : Queue a -> Queue a
nextRepeatOff q =
    Queue.next q RepeatOff


nextRepeatAll : Queue a -> Queue a
nextRepeatAll q =
    Queue.next q RepeatAll


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
                        |> nextRepeatOff
                        |> Queue.getCurrent
                    )
            )
        , test "Next, then previous"
            (\_ ->
                Expect.equal
                    (Just 1)
                    (queue
                        |> nextRepeatOff
                        |> Queue.previous
                        |> Queue.getCurrent
                    )
            )
        , test "RepeatAll queue loops around"
            (\_ ->
                Expect.equal
                    (Just 1)
                    (queue
                        |> nextRepeatAll
                        |> nextRepeatAll
                        |> nextRepeatAll
                        |> nextRepeatAll
                        |> Queue.getCurrent
                    )
            )
        ]
