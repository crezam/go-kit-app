package rpc

type UpperCaseRequest struct {
	OriginalWord string `json:"originalWord"`
}

type UpperCaseResponse struct {
	TransformedWord string `json:"transformedWord"`
	Err string `json:"err,omitempty"`
}
