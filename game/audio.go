package game

import (
	"bytes"
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

//go:embed static/sounds/*.mp3
var staticFiles embed.FS

// AudioContextWrapper wraps an audio context for use in Ebiten-based games.
type AudioContextWrapper struct {
	ctx *audio.Context
}

// initAudio initializes the audio context and loads the necessary sound files for the game.
// It loads the following sounds:
// - dead.mp3 — sound for player death
// - pop.mp3 — sound for pop effect
// - turn.mp3 — sound for the turn end
// If there is an error loading any sound file, an error message is logged.
func (g *Game) initAudio() {
	sampleRate := 44100
	g.audioContext = &AudioContextWrapper{ctx: audio.NewContext(sampleRate)}

	deadPlayer, err := loadSound(g.audioContext.ctx, "static/sounds/dead.mp3")
	if err != nil {
		log.Println("Error loading dead.mp3:", err)
	}
	g.soundDead = deadPlayer

	popPlayer, err := loadSound(g.audioContext.ctx, "static/sounds/pop.mp3")
	if err != nil {
		log.Println("Error loading pop.mp3:", err)
	}
	g.soundPop = popPlayer

	turnPlayer, err := loadSound(g.audioContext.ctx, "static/sounds/turn.mp3")
	if err != nil {
		log.Println("Error loading turn.mp3:", err)
	}
	g.soundTurn = turnPlayer
}

// SoundPlayer wraps an Ebiten audio player to handle playback of sound files.
type SoundPlayer struct {
	player *audio.Player
	ctx    *audio.Context
}

// loadSound loads an MP3 sound file from the embedded file system and creates a sound player.
// It returns a SoundPlayer or an error if the file cannot be loaded or decoded.
func loadSound(ctx *audio.Context, path string) (*SoundPlayer, error) {
	// Read the file from the embedded file system
	data, err := staticFiles.ReadFile(path)
	if err != nil {
		return nil, err
	}

	memReader := bytes.NewReader(data)

	// Decode the MP3 data into an Ebiten audio decoder
	d, err := mp3.Decode(ctx, memReader)
	if err != nil {
		return nil, err
	}

	// Create an audio player for the decoded sound
	b, err := audio.NewPlayer(ctx, d)
	if err != nil {
		return nil, err
	}

	return &SoundPlayer{
		player: b,
		ctx:    ctx,
	}, nil
}

// Play plays the sound from the SoundPlayer. If the player is nil or uninitialized, it does nothing.
func (sp *SoundPlayer) Play() {
	if sp == nil || sp.player == nil {
		return
	}
	sp.player.Rewind()
	sp.player.Play()
}

// playDeadSound plays the sound for the player's death.
func (g *Game) playDeadSound() {
	g.soundDead.Play()
}

// playPopSound plays the pop effect sound.
func (g *Game) playPopSound() {
	g.soundPop.Play()
}

// playTurnSound plays the sound for the turn end.
func (g *Game) playTurnSound() {
	g.soundTurn.Play()
}
