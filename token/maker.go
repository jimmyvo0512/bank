package token

import "time"

type Maker interface {
	CreateToken(accountId int64, expiredIn time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}