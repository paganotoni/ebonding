package ebonding

// NewClient returns a client for the passed clientType
func NewClient(clientType int, credentials Credentials) (*Client, error) {
	var client Client

	switch clientType {
	case ClientTypeATT:
		client = ATTClient{
			credentials: credentials,
		}
	case ClientTypeVerizon:
		client = VerizonClient{
			credentials: credentials,
		}
	}

	return &client, nil
}
