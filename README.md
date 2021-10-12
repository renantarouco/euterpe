# EUTERPE SERVER

The muse of music and lyric poetry in Greek mythology.

This is the a webserver implemented in Go that provides the interaction between
players using ["apollo"](https://github.com/renantarouco/apollo), originally
forked from the ["in-c"](https://github.com/teropa/in-c) project.

The players will interact using the
["apollo"](https://github.com/renantarouco/apollo) interface that connects to
this server for handling a group session.

The server exposes a websocket endpoint that initially registers the user
in a group, gives it an ID and start receiving commands for incremmenting the
user's current counter. That counter represents which fragment the interface
should be playing. Everytime a user decides to to move to the next fragment
every connected player in the same group session will receive a notification,
which will allow the interface to play the record accordingly. The coordination
allows a group of players to compose an entire play based on predefined
ordered fragments like the
[Terry Riley](https://en.wikipedia.org/wiki/Terry_Riley)'s 1964
composition ["In C"](https://en.wikipedia.org/wiki/In_C).

The server also allows to be configured with some rules like:

- maximum players allowed in a session, e.g.: 5 players;
- maximum distance between players fragments, e.g.: a player cannot be more than
  3 fragments ahead from any other player in the session.