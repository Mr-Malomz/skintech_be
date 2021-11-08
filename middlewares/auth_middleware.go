package middlewares

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Token interface {
	CreateToken(firstname string, lastname string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

func newPayLoad(firstname string, lastname string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.New()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Firstname: firstname,
		Lastname:  lastname,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

type jwtMaker struct {
	secretKey string
}

func AuthJWTMiddleware() {
	// return gin.
}
