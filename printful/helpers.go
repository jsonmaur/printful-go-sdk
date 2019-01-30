package printful

import (
	"encoding/json"
	"errors"
)

func unmarshal(data []byte, pointer interface{}) error {
	if err := json.Unmarshal(data, &pointer); err != nil {
		return errors.New("Could not parse response JSON")
	}

	return nil
}
