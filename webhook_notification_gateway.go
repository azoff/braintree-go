package braintree

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
)

type WebhookNotificationGateway struct {
	*Braintree
}

func (w *WebhookNotificationGateway) Parse(signature, payload string) (*WebhookNotification, error) {
	hmacer := newHmacer(w.Braintree)
	if verified, err := hmacer.verifySignature(signature, payload); err != nil {
		return nil, err
	} else if !verified {
		return nil, SignatureError{}
	}

	xmlNotification, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	var n WebhookNotification
	err = xml.Unmarshal(xmlNotification, &n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}


func (w *WebhookNotificationGateway) Verify(challenge string) (response string, err error) {
	if response, err = newHmacer(w.Braintree).hmac(challenge); err == nil {
		response = fmt.Sprintf("%s|%s", w.Braintree.PublicKey, response)
	}
	return
}
