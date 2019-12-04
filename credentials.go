package ebonding

// Credentials stores auth settings for the ebonding client
type Credentials struct {
	Username      string
	Password      string
	ExtraSettings map[string]interface{}
}
