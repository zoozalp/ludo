package state

import "github.com/libretro/go-playthemall/libretro"

// State is a type for the global state of the app
type State struct {
	Core        libretro.Core // Current libretro core
	FrameTimeCb libretro.FrameTimeCallback
	AudioCb     libretro.AudioCallback
	CoreRunning bool
	MenuActive  bool // When set to true, will display the menu layer
	Verbose     bool
	CorePath    string // Path of the current libretro core
	GamePath    string // Path of the current game
}

// Global state
var Global State
