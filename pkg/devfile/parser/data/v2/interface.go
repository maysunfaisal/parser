package v2

// DevfileDataV2 is an interface that defines functions each devfile spec must implement
type DevfileDataV2 interface {

	// mock new func that each spec must implement
	GetCustomType() string
}
