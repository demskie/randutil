package randutil

import (
	"errors"
	"math/rand"
)

type WeightedObj struct {
	Weight int
	Object interface{}
}

func WeightedChoice(choices []WeightedObj, rnum *rand.Rand) (WeightedObj, error) {
	max := 0
	for _, choice := range choices {
		max += choice.Weight
	}
	num := rnum.Intn(max)
	for _, choice := range choices {
		num -= choice.Weight
		if num < 0 {
			return choice, nil
		}
	}
	return WeightedObj{}, errors.New("undefined behavior")
}
