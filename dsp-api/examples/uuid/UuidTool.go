package uuid

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {

	s := uuid.New().String()
	fmt.Println(s)
}
