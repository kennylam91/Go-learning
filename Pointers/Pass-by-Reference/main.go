package Pass_by_Reference

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(analyticsPointer *Analytics, msg Message) {
	(*analyticsPointer).MessagesTotal++
	if msg.Success {
		(*analyticsPointer).MessagesSucceeded++
	} else {
		(*analyticsPointer).MessagesFailed++
	}
}
