package integration

import (
	"testing"
	"github.com/Seinarukiro2/Secton/backend/models"
	"github.com/Seinarukiro2/Secton/backend/config"
	"github.com/Seinarukiro2/Secton/backend/tests"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserInDB(t *testing.T) {
	// Инициализируем тестовую базу данных
	tests.InitTestDB()

	// Создаем пользователя в базе данных
	user := models.User{TelegramID: "123456789"}
	err := config.DB.Create(&user).Error
	assert.NoError(t, err)

	// Проверяем, что пользователь был создан
	var fetchedUser models.User
	err = config.DB.First(&fetchedUser, "telegram_id = ?", "123456789").Error
	assert.NoError(t, err)
	assert.Equal(t, "123456789", fetchedUser.TelegramID)
}
