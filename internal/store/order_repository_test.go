package store_test

import (
	"testing"

	"github.com/goocarry/wb-internship/internal/model"
	"github.com/goocarry/wb-internship/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestOrderRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, dsn)
	defer teardown("order")

	o, err := s.Order().Create(&model.Order{
		OrderUID: "testorder321",
	})
	assert.NoError(t, err)
	assert.NotNil(t, o)
}
