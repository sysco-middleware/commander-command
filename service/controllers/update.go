package controllers

import (
	"encoding/json"

	"github.com/sysco-middleware/commander"
)

// UpdateModal used during a "update" command
type UpdateModal struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// OnUpdateUser handles a "update" command
func OnUpdateUser(command *commander.Command) {
	var data UpdateModal

	UnmarshalErr := json.Unmarshal(command.Data, &data)
	if UnmarshalErr != nil {
		command.NewError("DataParseError", nil)
		return
	}

	// ...

	res, MarshalErr := json.Marshal(data)

	if MarshalErr != nil {
		command.NewError("ResponseParseError", nil)
		return
	}

	command.NewEvent("Updated", command.Key, res)
}