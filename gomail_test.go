package gomail

import (
	"testing"
)

func TestGoMail(t *testing.T) {
	goMail := NewGoMail()
	goMail.SendNotification("This is the test message that is very long will it be longer than 159 I dont know but lets make it long so we can find out. This is a test of the vtext length and should cut")
}
