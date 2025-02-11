package Send_messagepackage

// ?

// don't touch below this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func (user User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= user.MessageCharLimit {
		return message, true
	} else {
		return "", false
	}
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}
