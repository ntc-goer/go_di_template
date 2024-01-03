package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestXMLDecode(t *testing.T) {
	data, err := os.ReadFile("onixfile.xml")
	assert.Nil(t, err)
	var result XMLProduct
	err = XMLDecode(data, &result)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
