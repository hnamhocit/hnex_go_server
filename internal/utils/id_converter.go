package utils

import (
	"strconv"
)

func ConvertID(id string) (uint32, error) {
	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(idUint64), nil
}
