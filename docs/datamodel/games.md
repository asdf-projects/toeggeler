# Games V0.1

The games table stores summaries of games created based on the gamelog table. Could be useful if we want to show the last x games played by a user or something like that.

Will be updated  after each submitted game, but could be build or rebuilt based on the game log anytime

```json
{
    "id": 123, 
    "setup": {
        "settings": {
            "timeLimit": null,
            "goalLimit": 8
        },
        "side1": {
            "offensePlayer": 122,
            "defensePlayer": 123,
            "color": "blue"
        },
        "side2": {
            "offensePlayer": 222,
            "defensePlayer": 223,
            "color": "red"
        }
    },
    "score": {
        "side1": 5,
        "side2": 3
    },
    "startTime": 12312312312321,
    "endTime": 12312312312321,
}
```
# Games V0.1



|Name|Type| |
|----|----|-|
|game_id|INTEGER| PK
|time_limit|INTEGER|
|goal_limit|INTEGER| 
|side1_off_player|INTEGER|
|side1_def_player|INTEGER|
|side1_goals|INTEGER
|side1_color|TEXT|
|side2_off_player|INTEGER
|side2_def_player|INTEGER|
|side2_goals|INTEGER
|side2_color|TEXT|
|start_time|INTEGER
|end_time|INTEGER