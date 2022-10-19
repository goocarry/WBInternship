package player

import "fmt"

// PlayingState ...
type PlayingState struct {
	Player *Player
}
// ClickLock ...
func (ps *PlayingState) ClickLock() {
	fmt.Println("player locked")
	ps.Player.setState(ps.Player.Locked)
}
// ClickPlay ...
func (ps *PlayingState) ClickPlay() {
	fmt.Println("player stopped")
	ps.Player.setState(ps.Player.Ready)
}
// ClickNext ...
func (ps *PlayingState) ClickNext() {
	fmt.Println("playing next song")
}
// ClickPrevious ...
func (ps *PlayingState) ClickPrevious() {
	fmt.Println("playing previous song")
}

