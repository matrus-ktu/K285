package station

type Station struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	Address             string    `json:"address"`
	Latitude            string    `json:"latitude"`
	Longitude           string    `json:"longitude"`
	Status              string    `json:"status"`
	ConType             Connector `json:"connector"`
	ParkingRestrictions string    `json:"restrictions"`
	Price               Tariff    `json:"price"`
	Img                 Image     `json:"img"`
	ParkingCost         string    `json:"cost"`
	Operator            string    `json:"operator"`
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
