package domain

const (
	GamerCreatedEvent     = "V1.Gamers.GamerCreated"
	GamerApprovedEvent    = "V1.Gamers.GamerApproved"
	GamerDisapprovedEvent = "V1.Gamers.GamerDisapproved"
)

type GamerCreated struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsApproved bool   `json:"is_approved"`
}

func (GamerCreated) Key() string { return GamerCreatedEvent }

type GamerApproved struct{}

func (GamerApproved) Key() string { return GamerApprovedEvent }

type GamerDisapproved struct{}

func (GamerDisapproved) Key() string { return GamerDisapprovedEvent }
