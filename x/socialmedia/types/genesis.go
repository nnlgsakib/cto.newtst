package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:  DefaultParams(),
		PostMap: []Post{}, ProfileMap: []Profile{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	postIndexMap := make(map[string]struct{})

	for _, elem := range gs.PostMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := postIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for post")
		}
		postIndexMap[index] = struct{}{}
	}
	profileIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProfileMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := profileIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for profile")
		}
		profileIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
