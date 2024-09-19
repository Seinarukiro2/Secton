package unit

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/Seinarukiro2/Secton/backend/controllers"
	"github.com/Seinarukiro2/Secton/backend/tests"
)

func TestCreateUser_Success(t *testing.T) {
	// Инициализация тестовой базы данных
	tests.InitTestDB()

	// Создаем Echo и запрос/ответ
	e := echo.New()
	userJSON := `{"telegram_id": "123456789"}`
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(userJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Вызываем контроллер
	err := controllers.CreateUser(c)

	// Проверка результата
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var user map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, "123456789", user["telegram_id"])
}

func TestCreateUser_BadRequest(t *testing.T) {
	// Инициализация тестовой базы данных
	tests.InitTestDB()

	// Создаем Echo и запрос/ответ с неверными данными
	e := echo.New()
	invalidJSON := `{"telegram_id": 123}` // Поле должно быть строкой, а не числом
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(invalidJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Вызываем контроллер
	err := controllers.CreateUser(c)

	// Проверка результата
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid request", response["error"])
}
