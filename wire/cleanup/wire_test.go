package cleanup

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	mission, clean, err := New()
	if err != nil {
		if clean != nil {
			clean()
		}
	}

	fmt.Println(mission.Player.Name)
	fmt.Println(mission.Monster.Name)
	clean()
}
