package config

type Config struct{
	lnt Lenta
	pr Proxy
	ctg Categories
	del Delivery
}

type Lenta struct {
	URL string ``
	APIURL string ``
	Timeout int ``
}

type Proxy struct {
	URL string ``
	Type string ``
	Enabled bool ``
}

type Categories struct {
	ID string ``
	Name string ``
	Path string ``
}

type Delivery struct {
	Address string ``
	Latitude float64 ``
	Longitude float64 ``
}