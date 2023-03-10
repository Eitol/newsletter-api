package newsletter

import (
	"github.com/google/uuid"
)

type Interest string

var (
	InterestTech     Interest = "tech"
	InterestPolitics Interest = "politics"
	InterestSports   Interest = "sports"
)

type Subscription struct {
	UserID    uuid.UUID
	BlogID    uuid.UUID
	Interests []Interest
}
