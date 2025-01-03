package crypto

import "testing"
import "github.com/stretchr/testify/assert"

func TestHashDataToStringShouldProduceCorrectOutput(t *testing.T) {
	result, err := HashDataToString([]byte("testing 123"))
	assert.NoError(t, err)

	expected := "3XKaGNNw9eqMbJX9ZfCCw8ux7xsy476Kz1PDR6sh9zPc6wAqWZQcM6iLb3LReXbGt4UwtTT6qzSUoSqZbL6oG3mb"
	assert.Equal(t, expected, result)
}

func TestHashFile(t *testing.T) {
	result, err := HashFile("../test/data/abc.txt")
	assert.NoError(t, err)

	expected := "2T6ES1EQMLiCMJ9vdg8Cz6Hd5FMqKGiAvTubjRPsyJjcsMVAM3bcKeuwYRh5ExF7VrztiXwWXz1mdyrQM7PA5npE"
	assert.Equal(t, expected, result)
}
