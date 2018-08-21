package telnet

import (
	"errors"
	"strings"

	"github.com/aybabtme/rgbterm"
	"github.com/soupstore/coda/config"
	"github.com/soupstore/coda/simulation/model"
	"github.com/soupstore/go-core/logging"
)

// stateWorld is the scene used interacting with the world
type stateWorld struct {
	config      *config.Config
	conn        *connection
	characterID model.CharacterID
}

// onEnter is called when the scene is first loaded
func (s *stateWorld) onEnter() error {
	s.conn.writelnString("You are in the world!\n\r")

	s.characterID = CharacterIDFromContext(s.conn.ctx)
	events, err := s.conn.sim.WakeUpCharacter(s.characterID)
	if err != nil {
		return err
	}

	go renderEvents(s.conn, events)

	return nil
}

func (s *stateWorld) onExit() error {
	logging.Debug("Disconnecting from world server")
	// err := s.cc.Close()
	logging.Debug("Disconnected from world server")
	return nil
}

// handleInput parses input from the client and performs any appropriate command
func (s *stateWorld) handleInput(input string) error {
	tokens := strings.Split(input, " ")
	commandText := strings.ToLower(tokens[0])

	command, ok := worldCommands[commandText]
	if !ok {
		echo := rgbterm.String("Huh?", 255, 100, 100, 0, 0, 0)
		s.conn.writelnString(echo)
		s.conn.writePrompt()
		return nil
	}

	characterID := CharacterIDFromContext(s.conn.ctx)
	err := command(characterID, s.conn.sim, tokens[1:])
	if err != nil {
		return err
	}

	// TODO: I dont like this - need to fix it
	if commandText == "quit" {
		s.conn.close()
		return errors.New("closed")
	}

	return nil
}