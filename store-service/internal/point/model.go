package point

type SubmitedPoint struct {
	Amount int `json:"amount"`
}

type Point struct {
	ID     int `json:"id" db:"id"`
	UserID int `json:"user_id" db:"user_id"`
	Amount int `json:"amount" db:"amount"`
}

type TotalPoint struct {
	Point int `json:"point"`
}
