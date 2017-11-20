package faker

import "math/rand"

func getRandomIndex(data []string) int {
	if len(data) > 0 {
		return rand.Intn(len(data))
	}
	return 0
}
