package player

import "fmt"

// LockedState ...
type LockedState struct {
	Player *Player
}

// ClickLock ...
func (ls *LockedState) ClickLock() {
	fmt.Println("player is ready")
	ls.Player.setState(ls.Player.Ready)
}
// ClickPlay ...
func (ls *LockedState) ClickPlay() {
	fmt.Println("can't play, player is locked")
}
// ClickNext ...
func (ls *LockedState) ClickNext() {
	fmt.Println("can't play next, player is locked")
}
// ClickPrevious ...
func (ls *LockedState) ClickPrevious() {
	fmt.Println("can't play previous, player is locked")
}
