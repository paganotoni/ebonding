package ebonding

// ATTClient is the implementation of the client for A&T
type ATTClient struct {
	credentials Credentials
}

// FindDevices implementation for AT&T
func (att ATTClient) FindDevices() ([]Device, error) {
	return []Device{}, nil
}

// FindAccessories implementation for AT&T
func (att ATTClient) FindAccessories() ([]Accessory, error) {
	return []Accessory{}, nil
}
