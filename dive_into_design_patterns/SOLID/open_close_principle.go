package solid

import "time"

type Shipping interface {
	getCost(order *Order) float32
	getDate() time.Time
}

type Order struct {
	lineItems []float32
	shipping  Shipping
}

func (order *Order) getTotal() float32 {
	return 50.0 // solo para mostrar
}
func (order *Order) getTotalWeight() float32 {
	return 25 // solo para mostrar
}

func (order *Order) setShippingType(st Shipping) {
	order.shipping = st
}

func (order *Order) getShippingCost() float32 {
	return order.shipping.getCost(order)
}

func (order *Order) getShippingDatel() time.Time {
	return order.shipping.getDate()
}

// Implement interface on other types of shipping
type GroundShipping struct{}

var _ Shipping = (*GroundShipping)(nil)

func (g *GroundShipping) getCost(order *Order) float32 {

	if order.getTotal() > 100 {
		return 0
	}
	// $1.5 per kilogram, but $10 minimum
	return max(10, order.getTotalWeight()*1.5)

}

func (g *GroundShipping) getDate() time.Time {
	return time.Now()
}

type AirShipping struct{}

var _ Shipping = (*AirShipping)(nil)

func (air *AirShipping) getCost(order *Order) float32 {

	return max(20, order.getTotalWeight()*3)

}

func (g *AirShipping) getDate() time.Time {
	return time.Now()
}
