package telegram

type TelegramService interface {
	SendMessage(form map[string]string) error
	GetUserName() string
}
