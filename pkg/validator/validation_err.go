package validator

type Varror struct {
	Error            string             `json:"error"`
	ValidationErrors []*ValidationError `json:"validation_errors,omitempty"`
}
