package maps

import (
	"errors"
)

type Dict map[string]string

func (d Dict) Search(key string) (string, error) {
	definition, exist := d[key]
	if !exist {
		return "", errors.New("could not find a word")
	}

	return definition, nil
}
