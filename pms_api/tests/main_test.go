package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"pms_backend/pms_api/internal/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	host := "http://localhost:8080"

	user := model.UserInserted{
		Username:   "pmsuser",
		Password:   "1234",
		IsAdmin:    false,
		FirstName:  "Пользователь",
		MiddleName: "",
		LastName:   "",
	}
	data, err := json.Marshal(user)
	assert.NoError(t, err)

	resp, err := http.Post(host+"/api/v1/users", "application/json", bytes.NewReader(data))
	assert.NoError(t, err)

	userResp := &model.User{}
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(userResp)

	project := model.InsertProject{
		Name:        "Тестовый проект",
		Description: "Создан для тестирования приложения",
	}
	data, err = json.Marshal(project)
	assert.NoError(t, err)

	resp, err = http.Post(host+"/api/v1/projects", "application/json", bytes.NewReader(data))
	assert.NoError(t, err)

	projectResp := &model.Project{}
	decoder = json.NewDecoder(resp.Body)
	decoder.Decode(projectResp)

	task := model.TaskInserted{
		Name:        "Тестовая задача",
		Description: "Создана для теста приложения",
		Status:      "Открыта",
		ProjectID:   projectResp.ID,
		AuthorID:    userResp.ID,
		ExecutorID:  userResp.ID,
		TesterID:    userResp.ID,
	}
	data, err = json.Marshal(task)
	assert.NoError(t, err)

	_, err = http.Post(host+"/api/v1/task", "application/json", bytes.NewReader(data))
	assert.NoError(t, err)
}
