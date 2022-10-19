package player

import "fmt"

// ReadyState ...
type ReadyState struct {
	Player *Player
}

// ClickLock ...
func (rs *ReadyState) ClickLock() {
	fmt.Println("player locked")
	rs.Player.setState(rs.Player.Locked)
}
// ClickPlay ...
func (rs *ReadyState) ClickPlay() {
	fmt.Println("playing")
	rs.Player.setState(rs.Player.Playing)
}
// ClickNext ...
func (rs *ReadyState) ClickNext() {
	fmt.Println("switch to next song")
}
// ClickPrevious ...
func (rs *ReadyState) ClickPrevious() {
	fmt.Println("switch to next song")
}

