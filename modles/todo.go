package modles

import "time"

type Todo struct{
	TodoId string
	Task string
	CreatedAt time.Time
	CompletedAt time.Time
	Completed bool
	UserId string

}