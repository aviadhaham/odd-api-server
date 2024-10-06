package utils

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GetRandomOddNumber() int {
	return rand.Intn(10)*2 + 1
}

func TestNumberInRangeAndOdd(num string) error {
	converted, err := strconv.Atoi(num)
	if err != nil {
		return err
	}
	if converted < 1 || converted > 20 {
		return fmt.Errorf("Number %d is not between 1 and 20", converted)
	}

	if converted%2 == 0 {
		return fmt.Errorf("Number %d is not odd", converted)
	}
	return nil
}
