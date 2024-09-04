package yaakcli

import (
	"math/rand/v2"
)

var adjectives = []string{
	"young", "youthful", "yellow", "yielding", "yappy",
	"yawning", "yummy", "yucky", "yearly",
}

var nouns = []string{
	"yak", "yarn", "year", "yell", "yoke", "yoga",
	"yam", "yacht", "yodel",
}

func RandomName() string {
	adjective := adjectives[rand.IntN(len(adjectives))]
	noun := nouns[rand.IntN(len(nouns))]
	return adjective + "-" + noun
}
