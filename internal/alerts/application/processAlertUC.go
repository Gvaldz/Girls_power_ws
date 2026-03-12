package application

import (
	"fmt"
	"ws-server/internal/alerts/domain/entities"
	domain "ws-server/internal/alerts/domain/repository"
)

type ProcessAlertUseCase struct {
	wsHub domain.WSNotifier
}

func NewProcessAlertUseCase(ws domain.WSNotifier) *ProcessAlertUseCase {
	return &ProcessAlertUseCase{wsHub: ws}
}

func (uc *ProcessAlertUseCase) Execute(senderID int, payload entities.AlertPayload) {
	payload.SenderID = senderID

	networkIDs := make([]int, 0, len(payload.UsersNetwork))
	for _, u := range payload.UsersNetwork {
		networkIDs = append(networkIDs, u.UserID)
	}
	uc.wsHub.NotifyMultiple(networkIDs, "NEARBY_ALERT", map[string]interface{}{
		"sender_id":   senderID,
		"sender_name": payload.SenderName,
		"message":     fmt.Sprintf("¡%s está en peligro y necesita ayuda! Ingresa a la aplicación para obtener más información", payload.SenderName),
	})

	familyIDs := make([]int, 0, len(payload.UsersFamily))
	for _, u := range payload.UsersFamily {
		familyIDs = append(familyIDs, u.UserID)
	}
	uc.wsHub.NotifyMultiple(familyIDs, "FAMILY_ALERT", map[string]interface{}{
		"sender_id":   senderID,
		"sender_name": payload.SenderName,
		"message":     fmt.Sprintf("¡Tu familiar %s está en peligro! Ingresa a la aplicación para obtener más información", payload.SenderName),
	})
}