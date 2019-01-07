package main

import (
    "fmt"
    // "math" // for math.Min(a, b) etc.
)

func main() {
    // TODO
    // Const outside of main?
    const maxNumberOfCards uint = 5

    // TODO
    // Different log levels
    //  None - Winner
    //  Info - Played Card, Yell
    //  Verbose - Given Cards, Played Card, Yell
    //  Debug - Given Cards, Hard Cards, Played Card, Yell
    // Const outside of main?

    // TODO
    // Create Players
    //  "Alice", "Bob", "Carol", "Dave", "Erin", "Frank", "Grace", "Heidi"
    // Different rulesets
    //  Always Highest Card
    //  Highest Card matching Colour
    //  Always Lowest Card
    //  Always Lowest Card matching Colour

    // TODO
    // Make number of players adjustable
    var players uint = 4

    // TODO
    // Proper print
    fmt.Println("Number of players", players, "and cards", maxNumberOfCards)

    // TODO
    // Give cards
    // Create cards stack
    // long for-loop with the players playing until one player has 0 cards
}

func giveCards() {

}

func playTurn(int playerNumber) {
    // TODO
    // Previous Card
    
    // TODO
    // Possible Cards

    // TODO
    // Different logic

    // TODO
    // Action, log Yell
}


/*
  TODO
  - Apply Names
  - Create Available Cards Array/List
  - Create Card class
  - Create static rules (previous Card, player choice, action)
  - Allow you decide which card to play!!!
*/