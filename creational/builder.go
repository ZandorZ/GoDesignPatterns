package creational

// BuildProcess ...
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// ManufacturingDirector Director ...
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// VehicleProduct ...
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (c *BikeBuilder) SetWheels() {
	c.v.Wheels = 2
}

func (c *BikeBuilder) SetSeats() {
	c.v.Seats = 2
}

func (c *BikeBuilder) SetStructure() {
	c.v.Structure = "Bike"
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BusBuilder struct {
	v VehicleProduct
}

func (c *BusBuilder) SetWheels() {
	c.v.Wheels = 4
}

func (c *BusBuilder) SetSeats() {
	c.v.Seats = 12
}

func (c *BusBuilder) SetStructure() {
	c.v.Structure = "Bus"
}

func (c *BusBuilder) GetVehicle() VehicleProduct {
	return c.v
}
