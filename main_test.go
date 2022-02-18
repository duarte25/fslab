package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	t.Run("sucesso da soma de 1 + 2 deve ser 3", func(t *testing.T) {
		out := soma(1, 2)
		expected := 3.0
		assert.Equal(t, expected, out)
	})

	t.Run("sucesso da soma de 5 + 8 deve ser 13", func(t *testing.T) {
		out := soma(5, 8)
		expected := 13.0
		assert.Equal(t, expected, out)
	})

}

func TestMA(t *testing.T) {
	t.Run("sucesso da media aritmetica de 1 e 2 deve ser 1.5", func(t *testing.T) {
		out := mediaAritmetica(1, 2)
		expected := 1.5
		assert.Equal(t, expected, out)
	})

	t.Run("sucesso da media aritmetica de 5 e 8 deve ser 6.5", func(t *testing.T) {
		out := mediaAritmetica(5, 8)
		expected := 6.5
		assert.Equal(t, expected, out)
	})

}

func TestMP(t *testing.T) {
	t.Run("sucesso da media ponderada de 2 e 2 deve ser 2.0", func(t *testing.T) {
		out := mediaPonderada(2, 2)
		expected := 2.0
		assert.Equal(t, expected, out)
	})

	t.Run("sucesso da media ponderada de 5 e 8 deve ser 7", func(t *testing.T) {
		out := mediaPonderada(5, 8)
		expected := 7.0
		assert.Equal(t, expected, out)
	})

}
