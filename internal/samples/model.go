package samples

type Sample struct {
	ID int `json:"id"`
	SampleName string `json:"sample_name"`
	SampleBody string `json:"sample_body"`
	UserID int `json:"user_id"`
}