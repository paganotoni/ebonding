package ebonding

const (
	// ClientTypeATT AT&T ebonding client
	ClientTypeATT = iota + 1

	// ClientTypeVerizon Verizon ebonding client
	ClientTypeVerizon

	// ClientTypeSprint Sprint ebonding client
	ClientTypeSprint

	// ClientTypeTmobile TMobile ebonding client
	ClientTypeTmobile

	// Here we would add other types depending on how many clients do we add.
)

//Check for ATTClient
var _ Client = (*ATTClient)(nil)

//Check for VerizonClient
var _ Client = (*VerizonClient)(nil)

// Client interface will allow
type Client interface {
	// FindDevices return all the devices available for that particular client.
	FindDevices() ([]Device, error)

	// FindAccessories returns accessories in that particular client catalog.
	FindAccessories() ([]Accessory, error)
}
