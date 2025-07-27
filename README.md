# README

## Game States

### Intro 

This is the "splash" screen to display when the game process first starts

### Menu

Presents the player with an initial menu (single player, multiplayer, etc)

### Setup

Players enter names and other required info

### Playing

The game is in an active playing state.  Snakes are moving.

### Paused

The game is frozen

### Game Over

The game has been won or lost

-----------

## Design

game opens, splash screen is visible for a few sec. this is mostly for style, but there could potentially be some init here also.

### the server
the role of the game server is to own instances of games that one or more "clients" connect to.

the server owns the game state.  it knows where the dots are, where the snakes are, and who the winner is.

the server's job is to notify clients continuously with updates to the game state.

a server instance can be local or remote.  it can serve a single player gamne or a multiplayer game.  regardless, the player client uses the same interface to the server

#### server states
* initializing
* online

### the client
the game client controls rendering, handles input, and sends individual player updates to the game server.

outside the game context, a client can interface with the server by creating a game, joining an existing game, or leaving an existing game.

within the game context, a client sends updates to the server on a snake's movements.

the client listens for updates from the server on dots consumed, collisions, snake growth, and the position of the relevant items within the game grid.

#### client states
* initializing
* menu
* connecting
* in game

### the screen
to normalize the experience for all players within a game, the server creates a grid that matches the lowest common denominator for screen or window sizes.

### the game instance
the game instance includes dots, snakes, a score tracker, grid boundaries, and a leveling system. future instances could include powerups, dynamic dot respawning, etc.

players have points and lives.

a player accumulates points by eating dots, and points continue to accumulate until a player has exhausted its lives.

a player loses a life by colliding with a grid boundary, another snake, or itself.

the object of the game is to consume dots without crashing and dying.  the more dots a snake consumes, the larger it grows.

if all the dots on the grid are consumed, the game goes to the next level.  with each level, the game moves faster.

if a snake collides into another snake, the surviving snake gets a "kill"

the game is over in any of the following events:
- a player has exhausted its lives
- all the dots on the highest level are consumed

there isn't a clear system of winning or losing.  it's up to players to decide how to measure that.  the following metrics might be interpreted as a win or loss:
- points accumulated at end of game
- number of lives left at end of game
- number of kills

#### game states

* initializing
* waiting for players
* running
* paused
* complete