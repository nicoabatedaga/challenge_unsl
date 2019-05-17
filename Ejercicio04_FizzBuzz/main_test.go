package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "FizzBuzz"},
		FizzBuzz([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}),
		"La lista que se genero no es la que se esperaba")

	assert.Equal(
		[]string{"Fizz", "Buzz", "Fizz", "Fizz", "FizzBuzz", "1", "Fizz", "2", "Buzz", "7", "7"},
		FizzBuzz([]int{3, 5, 6, 9, 15, 1, 3, 2, 5, 7, 7}),
		"La lista que se genero no es la que se esperaba")

}
