package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyEip191Sig(t *testing.T) {
	address := "0x40fcc42c5a25945c02b19204d082a67591d30cf6"
	originMessage := "Hello World!"
	signature := "0xff0a604b4400dbc23d2a8ed7a728c552246cd59bcd6a795a7e212622142e9b814f1da8e8af26e03205131b323cb1076486755abb1fbed5f852879257cb4e60c01b"

	errorAddress := "0x3946d96a4b46657ca95CBE85d8a60b822186Ad1f"

	res := VerifyEip191Sig(address, originMessage, signature)
	errorRes := VerifyEip191Sig(errorAddress, originMessage, signature)

	assert.Equal(t, res, true)
	assert.Equal(t, errorRes, false)
}
