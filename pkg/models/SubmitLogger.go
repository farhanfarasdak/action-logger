package models

type SubmitLogger struct {
	RefID     string `json:"ref_id"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}
