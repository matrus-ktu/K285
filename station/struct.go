package station

type Station struct {
	ID                  int
	Name                string
	Address             string
	Latitude            string
	Longitude           string
	Status              string
	ConType             Connector
	ParkingRestrictions string
	Price               Tariff
	Img                 Image
	ParkingCost         string
	Operator            string
}

type Connector struct {
	ID               int
	Standard         string
	Format           string
	PowerType        string
	MaxVoltage       int
	MaxAmperage      int
	MaxElectricPower int
	Power            float32
}

type Tariff struct {
	CountryCode string
	Currency    string
	Type        string
	MinPrice    float32
	MaxPrice    float32
}

type Image struct {
	Url      string
	Category string
	Type     string
	Width    int
	Height   int
}
