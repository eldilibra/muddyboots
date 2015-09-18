
L.D Libra [6:44 PM]
Your challenge: Implement a MUD like game in Go (https://en.wikipedia.org/wiki/MUD)
- There are Rooms, Players, and Items
- A Player has an Inventory composed of Items
- A Room may have multiple Players in it, but only 1 Item at a time. If a Player picks it up, it moves to that Player’s Inventory, and leaves the Room
- The Map in which Players play is composed of Rooms. Each room has 1 - 4 doors to other Rooms, so the map can be sparse.
- The game is multiplayer
- A Player may either:
   - move North
   - move South
   - move East
   - move West
   - examine a Room to see its doors and Item (if it has an item)
   - pick up an Item
   - put down an Item in a room that has no Item
   - quit the game (edited)

L.D Libra [6:45 PM]
Start here, then add Monsters and a combat system.

L.D Libra [6:45 PM]
This will at least keep you busy for the better part of a day or two :~)
