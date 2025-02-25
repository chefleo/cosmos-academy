package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SystemInfo: SystemInfo{
			NextId:        uint64(DefaultIndex),
			FifoHeadIndex: NoFifoIndex,
			FifoTailIndex: NoFifoIndex,
		},
		StoredGameList: []StoredGame{},
		PlayerInfoList: []PlayerInfo{},
		Leaderboard: Leaderboard{
			Winners: []WinningPlayer{},
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in storedGame
	storedGameIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredGameList {
		index := string(StoredGameKey(elem.Index))
		if _, ok := storedGameIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedGame")
		}
		storedGameIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in playerInfo
	playerInfoIndexMap := make(map[string]struct{})

	for _, elem := range gs.PlayerInfoList {
		index := string(PlayerInfoKey(elem.Index))
		if _, ok := playerInfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for playerInfo")
		}
		playerInfoIndexMap[index] = struct{}{}
	}
	// Validate Leaderboard
	if err := gs.Leaderboard.Validate(); err != nil {
		return err
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
