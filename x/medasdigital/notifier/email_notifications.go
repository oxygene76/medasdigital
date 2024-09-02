package notifier 

import (
    "fmt"
    "gopkg.in/gomail.v2"
    "context"
    "github.com/cosmos/cosmos-sdk/types"
    "cosmossdk.io/log"
)

// E-Mail-Server Konfiguration (Platzhalter-Werte, bitte ersetzen)
const (
    EmailServer   = "mail.medas-digital.io" // E-Mail-Server-Adresse (Gmail Beispiel)
    EmailPort     = 587              // SMTP Port für TLS
    EmailUsername = "blockchain@medas-digital.io" // Dein E-Mail-Benutzername
    EmailPassword = "TestTest#2024!"  // Dein E-Mail-Passwort
)

type EmailNotifier struct {
    Logger log.Logger
}

func NewEmailNotifier(logger log.Logger) *EmailNotifier {
    return &EmailNotifier{
        Logger: logger,
    }
}

func (e *EmailNotifier) NotifyOnDeposit(ctx context.Context, address string, amount types.Coins) {
    if e.isTargetAddress(address) {
        e.sendEmailNotification(address, amount)
    }
}

func (e *EmailNotifier) isTargetAddress(address string) bool {
    // Ersetze dies durch die Adresse, die Benachrichtigungen empfangen soll
    return address == "medas1cg8u9nk6amf3fphtu9kppq7dgrlx9zqx4wezu3"
}

func (e *EmailNotifier) sendEmailNotification(address string, amount types.Coins) {
    m := gomail.NewMessage()
    m.SetHeader("From", EmailUsername)
    m.SetHeader("To", "info@medas-digital.io") // Empfängeradresse
    m.SetHeader("Subject", "Deposit Received")
    m.SetBody("text/plain", fmt.Sprintf("A deposit of %s has been made to address %s.", amount.String(), address))

    d := gomail.NewDialer(EmailServer, EmailPort, EmailUsername, EmailPassword)

    if err := d.DialAndSend(m); err != nil {
        e.Logger.Error("Failed to send email notification", "error", err)
    }
}

