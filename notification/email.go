package notification

type EmailNotification interface {
}

type emailNotification struct {
}

func NewEmailNotification() *emailNotification {
	return &emailNotification{}
}

func (n *emailNotification) Send(to string, subject string, body string) error {
	return nil
}
