package reader

import "math/rand"

func pickRandom(list []string) string {

	if len(list) == 0 {
		return ""
	}

	index := rand.Intn(len(list))
	return list[index]
}
