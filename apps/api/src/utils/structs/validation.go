package structs

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}
