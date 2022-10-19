package player

import (
	"github.com/goocarry/wb-internship/patterns/state/state"
)

// Player ...
type Player struct {
	Playing state.State
	Ready state.State
	Locked state.State

	CurrentState state.State

	Songs []string
}


// NewPlayer ...
func NewPlayer(songs []string) *Player {
    p := &Player{
        Songs: songs,
    }

    Playing := &PlayingState{
        Player: p,
    }
    Locked := &LockedState{
        Player: p,
    }
    Ready := &ReadyState{
        Player: p,
    }
    

    p.setState(Ready)
    p.Ready = Ready
    p.Locked = Locked
    p.Playing = Playing
    return p
}
// ClickPlay ...
func (p *Player) ClickPlay() {
    p.CurrentState.ClickPlay()
}

// ClickLock ...
func (p *Player) ClickLock() {
    p.CurrentState.ClickLock()
}

// ClickNext ...
func (p *Player) ClickNext() {
    p.CurrentState.ClickNext()
}

// ClickPrevious ...
func (p *Player) ClickPrevious() {
    p.CurrentState.ClickPrevious()
}

func (p *Player) setState(s state.State) {
    p.CurrentState = s
}
