package game

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type AudioContextWrapper struct {
	ctx *audio.Context
}

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

type SoundPlayer struct {
	player *audio.Player
	ctx    *audio.Context
}

func loadSound(ctx *audio.Context, path string) (*SoundPlayer, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	memReader := bytes.NewReader(data)

	d, err := mp3.Decode(ctx, memReader)
	if err != nil {
		return nil, err
	}

	b, err := audio.NewPlayer(ctx, d)
	if err != nil {
		return nil, err
	}

	return &SoundPlayer{
		player: b,
		ctx:    ctx,
	}, nil
}

func (sp *SoundPlayer) Play() {
	if sp == nil || sp.player == nil {
		return
	}
	sp.player.Rewind()
	sp.player.Play()
}

func (g *Game) playDeadSound() {
	g.soundDead.Play()
}

func (g *Game) playPopSound() {
	g.soundPop.Play()
}

func (g *Game) playTurnSound() {
	g.soundTurn.Play()
}
