package gym

type BasicMember struct {
	Name  string
	Email string
}

type PremiumMember struct {
	Name        string
	Email       string
	BonusPoints int
}

type Member interface {
	GetName() string
	GetEmail() string
	GetPoints() int
}

type Gym struct {
	Members map[uint64]Member
}

func (basicMember BasicMember) GetName() string {
	return basicMember.Name
}

func (basicMember BasicMember) GetEmail() string {
	return basicMember.Email
}

func (basicMember BasicMember) GetPoints() int {
	return 0
}

func (premiumMember PremiumMember) GetName() string {
	return premiumMember.Name
}

func (premiumMember PremiumMember) GetEmail() string {
	return premiumMember.Email
}

func (premiumMember PremiumMember) GetPoints() int {
	return premiumMember.BonusPoints
}
