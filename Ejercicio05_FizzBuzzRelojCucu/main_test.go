package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(
		"Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu",
		RunCucu("12:00"),
		"No coincide la respuesta")
	assert.Equal(
		"Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu",
		RunCucu("00:00"),
		"No coincide la respuesta")
	assert.Equal(
		"Cucuu",
		RunCucu("01:00"),
		"No coincide la respuesta")
	assert.Equal(
		"Cucuu",
		RunCucu("13:00"),
		"No coincide la respuesta")
	assert.Equal(
		"Cucuu Cucuu",
		RunCucu("14:00"),
		"No coincide la respuesta")
	assert.Equal(
		"Fizz Buzz",
		RunCucu("11:15"),
		"No coincide la respuesta")
	assert.Equal(
		"Buzz",
		RunCucu("11:55"),
		"No coincide la respuesta")
	assert.Equal(
		"Fizz",
		RunCucu("12:03"),
		"No coincide la respuesta")
	assert.Equal(
		"Cucuu",
		RunCucu("12:30"),
		"No coincide la respuesta")
}
