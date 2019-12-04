package ebonding

// VerizonClient is the implementation of the client for A&T
type VerizonClient struct {
	credentials Credentials
}

// FindDevices implementation for AT&T
func (att VerizonClient) FindDevices() ([]Device, error) {
	return []Device{}, nil
}

// FindAccessories implementation for AT&T
func (att VerizonClient) FindAccessories() ([]Accessory, error) {
	return []Accessory{}, nil
}
