# Architecture

## Technologies & Frameworks
Golang, SQLite, Svelte

## The basic idea
Most of the actual work should be done in the backend more precisely the go app. Ideally the client does pretty much zero work apart from displaying information and sending user inputs to the server.

The general idea is to build the workflow around games even more so the actual events during a game like "start", "goal" or "end". 

When a user submits a game after playing, the client will just send a list of events that happened during the game, for example:
* game started at 15:30:00, max (off) and moritz(def) vs hinz(off) and kunz(def), 5 goals to win
* max scores at 15:31:00
* hinz scores a föteli at 15:32:02
* ...
* moritz scores an own goal at 15:38:00
* game ended at 15:42:00 by goals scored

First, these events will be added to a game log table containing all events ever recorded. Each event is obviously assigned to a unique game id. Afterwards some kind of "evaluation engine" should ingest the events and update all the stats (mainly players).

At any time we should be able to feed the "evaluation engine" any amount of log events and extract player and game stats based on the given events. This allows for some interesting options later on:
* remove invalid game results from cheating attempts
* try a different rating system on historical log data
* time-based stats (last month, last year, all-time)
* more involved achievements or stats can be added later on (average playing time, win/lose streaks, scored five unanswered goals)

Since we don't expect massive amounts of data (for now ;) an sqlite table should suffice as event log.

## Data model
see files under /datamodel
