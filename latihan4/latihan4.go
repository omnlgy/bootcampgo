package latihan4

import (
	"fmt"
	"math/rand/v2"
)

type NotificationChannel interface {
	Send(userId, message string) error
	GetChannelName() string
}

type EmailNotifier struct {
	EmailAddress string
}

func NewEmailNotifier(emailAddress string) *EmailNotifier {
	return &EmailNotifier{
		EmailAddress: emailAddress,
	}
}

func (e *EmailNotifier) Send(userId, message string) error {
	if userId == "" {
		return fmt.Errorf("user ID cannot be empty")
	}

	if rand.IntN(10) < 5 {
		return fmt.Errorf("failed to send email")
	}

	fmt.Printf("Sending email to user %s: %s from %s\n", userId, message, e.EmailAddress)
	return nil
}

func (e *EmailNotifier) GetChannelName() string {
	return "Email"
}

type SmsNotifier struct {
	PhoneNumber string
}

func NewSmsNotifier(phoneNumber string) *SmsNotifier {
	return &SmsNotifier{
		PhoneNumber: phoneNumber,
	}
}

func (s *SmsNotifier) Send(userId, message string) error {
	if userId == "" {
		return fmt.Errorf("user ID cannot be empty")
	}

	if rand.IntN(10) < 5 {
		return fmt.Errorf("failed to send sms notification")
	}

	fmt.Printf("Sending SMS to user %s: %s from %s\n", userId, message, s.PhoneNumber)
	return nil
}

func (s *SmsNotifier) GetChannelName() string {
	return "SMS"
}

type PushNotifier struct {
	DeviceToken string
}

func NewPushNotifier(deviceToken string) *PushNotifier {
	return &PushNotifier{
		DeviceToken: deviceToken,
	}
}

func (p *PushNotifier) Send(userId, message string) error {
	if userId == "" {
		return fmt.Errorf("user ID cannot be empty")
	}

	if rand.IntN(10) > 5 {
		return fmt.Errorf("failed to send push notification")
	}

	fmt.Printf("Sending push notification to user %s: %s from %s\n", userId, message, p.DeviceToken)
	return nil
}

func (p *PushNotifier) GetChannelName() string {
	return "Push"
}

type NotificationDispatcher struct {
	Channels []NotificationChannel
}

func NewNotificationDispatcher() *NotificationDispatcher {
	return &NotificationDispatcher{
		Channels: []NotificationChannel{},
	}
}

func (n *NotificationDispatcher) AddChannel(channel NotificationChannel) {
	n.Channels = append(n.Channels, channel)
}

func (n *NotificationDispatcher) BroadcastNotification(userId, message string) {
	for _, channel := range n.Channels {
		if err := channel.Send(userId, message); err != nil {
			fmt.Printf("Error sending via %s: %v\n", channel.GetChannelName(), err)
		}
	}
}

func Main() {
	fmt.Println("Latihan 4")

	dispatcher := NewNotificationDispatcher()

	dispatcher.AddChannel(NewEmailNotifier("user@example.com"))
	dispatcher.AddChannel(NewSmsNotifier("+628123456789"))
	dispatcher.AddChannel(NewPushNotifier("device123"))
	dispatcher.AddChannel(NewEmailNotifier("user2@example.com"))
	dispatcher.AddChannel(NewSmsNotifier("+628987654321"))
	dispatcher.AddChannel(NewPushNotifier("device987"))
	dispatcher.AddChannel(NewSmsNotifier("+62854321987"))
	dispatcher.AddChannel(NewPushNotifier("device654"))
	dispatcher.AddChannel(NewEmailNotifier("user3@example.com"))
	dispatcher.AddChannel(NewSmsNotifier("+62876543210"))

	fmt.Println(len(dispatcher.Channels))

	dispatcher.BroadcastNotification("user123", "Hello World")
}
