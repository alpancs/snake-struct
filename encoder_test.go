package snaky

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalStruct(t *testing.T) {
	v := struct {
		ID         uint
		Name       string
		Price      float64
		IsNew      bool
		CategoryID uint
	}{1, "duck", 15000, true, 14}
	data, err := Marshal(v)
	assert.NoError(t, err)
	assert.Equal(t, `{"id":1,"name":"duck","price":15000,"is_new":true,"category_id":14}`, string(data))
}
