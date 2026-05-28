package uuid

import "github.com/google/uuid"

func GetUuId() string {
	return uuid.New().String()
}
