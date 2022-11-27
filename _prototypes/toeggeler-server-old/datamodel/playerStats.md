# PlayerStats V0.1
Player stats to keep track of a player's stats like ranking, wins and goals scored.

Will be updated  after each submitted game, but could be build or rebuilt based on the game log anytime.

```json
{
    "id": 1,
    "evaluationDate": 12312312312,
    "playerId": 123,
    "fromDate": 123123123123,
    "toDate": 2312312321,
    "elo": 1680, 
    "overall": {
        "wins": 14,
        "losses": 8,
        "goalsFor": 15,
        "goalsAgainst": 16,
        "ownGoals": 2,
        "counterGoals": 2,
        "shutouts": 0,
        "winLossRatio": 0.55
    },
    "offense": {
        "wins": 12,
        "losses": 5,
        "goalsFor": 12,
        "goalsAgainst": 4,
        "ownGoals": 1,
        "counterGoals": 2,
        "winLossRatio": 0.6,
    },
    "defense": {
        "wins": 2,
        "losses": 3,
        "goalsFor": 3,
        "goalsAgainst": 12,
        "ownGoals": 1,
        "counterGoals": 0,
        "winLossRatio": 0.5
    }
}
```

|Name|Type| |
|----|----|-|
|playerstats_id  |INTEGER| PK
|evaluation_date|INTEGER|
|player_id|INTEGER|FK
|from_date|INTEGER|
|to_date|INTEGER
|elo|INTEGER
|ovr_wins|INTEGER
|ovr_losses|INTEGER
|ovr_goals_for|INTEGER
|ovr_goals_against|INTEGER
|ovr_own_goals|INTEGER
|ovr_counter_goals|INTEGER
|ovr_shutouts|INTEGER
|ovr_win_loss_ratio|INTEGER
|off_wins|INTEGER
|off_losses|INTEGER
|off_goals_for|INTEGER
|off_goals_against|INTEGER
|off_own_goals|INTEGER
|off_counter_goals|INTEGER
|off_win_loss_ratio|INTEGER
|def_wins|INTEGER
|def_losses|INTEGER
|def_goals_for|INTEGER
|def_goals_against|INTEGER
|def_own_goals|INTEGER
|def_counter_goals|INTEGER
|def_win_loss_ratio|INTEGER