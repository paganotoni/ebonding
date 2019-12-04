package ebonding

//Device Represents a mobile device that's not an accessory for accessories see Accessory
type Device struct {
	SKU                string
	ManufacturerPartID string

	Color    string
	OS       string
	Capacity string

	URL                   string
	CompatibleAccessories []Accessory
}
