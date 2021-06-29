package utils_test

import (
	"github.com/JIeeiroSst/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToInt(t *testing.T){
	var result= 2
	assert.Equal(t,result,utils.StringToInt("2"))
}