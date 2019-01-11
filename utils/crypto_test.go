package utils

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateSecretKey(t *testing.T){
	assert := assert.New(t)
	log.Println(createSecretKey(16))
	assert.Equal(createSecretKey(16),1)
}

func TestAesEncrypt(t *testing.T){
	assert := assert.New(t)
	aes, _ :=aesEncrypt("123",createSecretKey(16))
	assert.Equal("3dLUln+QYI6+PdeW1R5byQ==",aes)
}