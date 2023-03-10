package types

import (
	"errors"
	"time"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

func NewGenesisState(epochs []EpochInfo) *GenesisState {
	return &GenesisState{Epochs: epochs}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	epochs := []EpochInfo{
		{
			Identifier:            "month",
			StartTime:             time.Time{},
			Duration:              time.Hour * 24 * 30,
			CurrentEpoch:          0,
			CurrentEpochStartTime: time.Time{},
			EpochCountingStarted:  false,
			CurrentEpochEnded:     true,
		},
		{
			Identifier:            "week",
			StartTime:             time.Time{},
			Duration:              time.Hour * 24 * 7,
			CurrentEpoch:          0,
			CurrentEpochStartTime: time.Time{},
			EpochCountingStarted:  false,
			CurrentEpochEnded:     true,
		},
		{
			Identifier:            "day",
			StartTime:             time.Time{},
			Duration:              time.Hour * 24,
			CurrentEpoch:          0,
			CurrentEpochStartTime: time.Time{},
			EpochCountingStarted:  false,
			CurrentEpochEnded:     true,
		},
		{
			Identifier:            "hour",
			StartTime:             time.Time{},
			Duration:              time.Hour,
			CurrentEpoch:          0,
			CurrentEpochStartTime: time.Time{},
			EpochCountingStarted:  false,
			CurrentEpochEnded:     true,
		},
		{
			Identifier:            "minute",
			StartTime:             time.Time{},
			Duration:              time.Minute,
			CurrentEpoch:          0,
			CurrentEpochStartTime: time.Time{},
			EpochCountingStarted:  false,
			CurrentEpochEnded:     true,
		},
	}
	return NewGenesisState(epochs)
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// TODO: Epochs identifiers should be unique
	for _, epoch := range gs.Epochs {
		if epoch.Identifier == "" {
			return errors.New("epoch identifier should NOT be empty")
		}
		if epoch.Duration == 0 {
			return errors.New("epoch duration should NOT be 0")
		}
	}
	return nil
}
