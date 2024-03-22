# Example
You have an e-commerce application with an Order class that calculates shipping costs and all shipping methods are hardcoded inside the class. If you need to add a new shipping method, you have to change the code of the Order class and risk breaking it.

```go
type Order struct {
    lineItems any
    shipping any
}

func (order *Order) getTotal() {}
func (order *Order) getTotalWeight() {}
func (order *Order) setShippingType(st any) {}
func (order *Order) getShippingCost() any {
    if shipping == "ground" {
        if order.getTotal() > 100 {
            return 0
        }
        // $1.5 per kilogram, but $10 minimum
        return max(10, order.getTotalWeight() * 1.5)
    }

    if shipping == "air" {
        // $3 per kilogram, but $20 minimum
        return max(20, order.getTotalWeight() * 3)
    }
}
func (order *Order) getShippingDate() {}
```

You can solve the problem by applying the *Strategy* pattern. Start by extracting shipping methods into separate classes with a common interface.

```go

type Shipping interface {
    getCost(order)
    getDate()
}

type Order struct {
    lineItems any
    shipping any
}

func (order *Order) getTotal() {}
func (order *Order) getTotalWeight() {}
func (order *Order) setShippingType(st Shipping) {}
func (order *Order) getShippingCost() any {
    return order.shipping.getCost()
}
func (order *Order) getShippingDatel() {}

// Implement interface on other types of shipping
type Ground struct{}

func (g *Ground) getCost(order Order) {
    if shipping == "ground" {
        if order.getTotal() > 100 {
            return 0
        }
        // $1.5 per kilogram, but $10 minimum
        return max(10, order.getTotalWeight() * 1.5)
    }
}
type Air struct{}
```

Now when you need to implement a new shipping method, you can derive a new class from the Shipping interface without touching any of the Order class’ code. The client code of the **Order** class will link orders with a shipping object of the new class whenever the user selects this shipping methods in the UI.

As a bonus, this solution let you move the delivery time calculation to more relevant classes, according to the *single responsibility principle*.