package usecases

// import (
// 	"context"
// 	"testing"

// 	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
// 	"github.com/AlexanderMorozov1919/mobileapp/internal/services/websocket"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// )

// func TestOneCWebhookUsecase_HandleReceptionsUpdate(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockCacheRepo := mocks.NewMockOneCCacheRepository(ctrl)
// 	mockHub := &websocket.Hub{} // или мок, если нужно

// 	usecase := NewOneCWebhookUsecase(mockCacheRepo, mockHub)

// 	update := models.OneCReceptionsUpdate{
// 		CallID: 123,
// 		Receptions: []models.Reception{
// 			{ID: 1, Status: "Новая"},
// 		},
// 	}

// 	mockCacheRepo.EXPECT().
// 		SaveReceptions(gomock.Any(), update.CallID, update.Receptions).
// 		Return(nil)

// 	err := usecase.HandleReceptionsUpdate(context.Background(), update)
// 	assert.NoError(t, err)
// }
