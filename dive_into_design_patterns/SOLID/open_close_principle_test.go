package solid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenClose(t *testing.T) {
	var airShipping = &AirShipping{}
	var groundShipping = &GroundShipping{}

	var order = &Order{
		lineItems: []float32{12.45, 5.55, 2},
		shipping:  airShipping,
	}

	t.Run("AirShipping", func(t *testing.T) {
		// t.Log(order.getShippingCost())
		assert.Equal(t, order.getShippingCost(), float32(75))
	})

	t.Run("GroundShipping", func(t *testing.T) {
		order.setShippingType(groundShipping)
		// t.Log(order.getShippingCost())
		assert.Equal(t, order.getShippingCost(), float32(37.5))
	})
}
