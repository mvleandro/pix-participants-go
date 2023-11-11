package server

// Represents a financial institution that is currently participating in Pix Brazilian financial scheme.
type Participant struct {
	// The institution id.
	Id string `json:"id"`

	// The institution name.
	Name string `json:"name"`

	// the institution short name.
	ShortName string `json:"shortName"`

	//The mode of pix participation.
	PixParticipationMode string `json:"pixParticipationMode"`

	// The kind of spi participation.
	SpiParticipationKind string `json:"spiParticipationKind"`

	// Timestamp since the begin of operation.
	BeginOperationTimestamp string `json:"beginOperationTimestamp"`
}
