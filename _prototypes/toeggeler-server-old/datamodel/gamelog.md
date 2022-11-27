# GameLog V0.1

The gamelog table stores every event in a very flat structure, there is certainly room for improvement...

## EVENTS
* GAME_START
    ```json
    {
        "event": "GAME_START",
        "time": 123123123123,
        "setup": {
            "settings": {
                "timeLimit": null,
                "goalLimit": 8
            },
            "side1": {
                "onOffense": 122,
                "onDefense": 123,
                "color": "blue"
            },
            "side2": {
                "onOffense": 222,
                "onDefense": 223,
                "color": "red"
            }
        }
    }
    ```

* GAME_END
    ```json
    {
        "event": "GAME_END",
        "time": 123123123123
    }
    ```

* GAME_TIMEOUT
    ```json
    {
        "event": "GAME_TIMEOUT",
        "time": 123123123123
    }
    ```

* GOAL
    ```json
    {
        "event": "GOAL",
        "time": 1231231232,
        "scoredBy": 122
    }
    ```

* OWN_GOAL
    ```json
    {
        "event": "OWN_GOAL",
        "time": 1231231232,
        "scoredBy": 122
    }
    ```

* COUNTER_GOAL
    ```json
    {
        "event": "COUNTER_GOAL",
        "time": 1231231232,
        "scoredBy": 122
    }
    ```

|Name|Type| |
|----|----|-|
|gamelog_id  |INTEGER| PK
|game_id|INTEGER|
|event|INTEGER|
|time|INTEGER|
|time_limit_ms|INTEGER|
|goal_limit|INTEGER|
|side1_on_offense|INTEGER|
|side1_on_defense|INTEGER|
|side1_color|INTEGER|
|side2_on_offense|INTEGER|
|side2_on_defense|INTEGER|
|side2_color|INTEGER|
|scored_by|INTEGER|

