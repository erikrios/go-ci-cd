package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/erikrios/go-clean-arhictecture/model"
	"github.com/erikrios/go-clean-arhictecture/service/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	mockUserService := mocks.UserService{}
	mockAuthService := mocks.AuthService{}

	dummyResp := []model.GetUserResponse{
		{
			Id:        1,
			Email:     "naruto@gmail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        2,
			Email:     "sasuke@gmail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	t.Run("success scenario", func(t *testing.T) {
		mockUserService.On("List").Return(
			func() []model.GetUserResponse {
				return dummyResp
			},
			func() error {
				return nil
			},
		).Once()

		t.Run("it should return 200 status code with valid response", func(t *testing.T) {
			controller := NewUserController(&mockUserService, &mockAuthService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users")

			if assert.NoError(t, controller.GetAll(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)

				body := rec.Body.String()

				response := make(map[string]interface{})
				if assert.NoError(t, json.Unmarshal([]byte(body), &response)) {
					users := response["data"].([]interface{})
					assert.Equal(t, len(dummyResp), len(users))

					for i, res := range dummyResp {
						gotUser := users[i].(map[string]interface{})
						gotUserId := uint(gotUser["id"].(float64))
						gotUserEmail := gotUser["email"].(string)
						assert.Equal(t, res.Id, gotUserId)
						assert.Equal(t, res.Email, gotUserEmail)
					}
				}
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		mockUserService.On("List").Return(
			func() []model.GetUserResponse {
				return []model.GetUserResponse{}
			},
			func() error {
				return errors.New("something went wrong.")
			},
		).Once()

		t.Run("it should return 500 status code if error happened", func(t *testing.T) {
			controller := NewUserController(&mockUserService, &mockAuthService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users")

			if assert.NoError(t, controller.GetAll(c)) {
				assert.Equal(t, http.StatusInternalServerError, rec.Code)
			}
		})
	})
}

func TestCreate(t *testing.T) {
	mockUserService := mocks.UserService{}
	mockAuthService := mocks.AuthService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyReq := model.CreateUserRequest{
			Email:    "naruto@gmail.com",
			Password: "naruto",
		}
		dummyResp := model.GetUserResponse{
			Id:        1,
			Email:     "naruto@gmail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockUserService.On("Create", mock.AnythingOfType("model.CreateUserRequest")).Return(
			func(request model.CreateUserRequest) model.GetUserResponse {
				return dummyResp
			},
			func(request model.CreateUserRequest) error {
				return nil
			},
		).Once()

		t.Run("it should return 200 status code with valid response", func(t *testing.T) {
			controller := NewUserController(&mockUserService, &mockAuthService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users")

			if assert.NoError(t, controller.Create(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				body := rec.Body.String()

				response := make(map[string]interface{})
				err := json.Unmarshal([]byte(body), &response)

				if assert.NoError(t, err) {
					user := response["data"].(map[string]interface{})
					assert.Equal(t, dummyResp.Id, uint(user["id"].(float64)))
					assert.Equal(t, dummyResp.Email, user["email"])
				}
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		t.Run("it should return 400 status code if request body is invalid", func(t *testing.T) {
			controller := NewUserController(&mockUserService, &mockAuthService)
			invalidJSONreq := `{"email": "naruto@gmail.com", "password": false}`

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(invalidJSONreq))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users")

			if assert.NoError(t, controller.Create(c)) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
				body := rec.Body.String()

				response := make(map[string]interface{})
				err := json.Unmarshal([]byte(body), &response)

				if assert.NoError(t, err) {
					errorMsg := response["error"].(string)
					assert.NotZero(t, errorMsg)
				}
			}
		})

		t.Run("it should return 500 status code if error happened", func(t *testing.T) {
			mockUserService.On("Create", mock.AnythingOfType("model.CreateUserRequest")).Return(
				func(request model.CreateUserRequest) model.GetUserResponse {
					return model.GetUserResponse{}
				},
				func(request model.CreateUserRequest) error {
					return errors.New("something went wrong.")
				},
			).Once()

			controller := NewUserController(&mockUserService, &mockAuthService)
			reqBody := `{"email": "naruto@gmail.com", "password": "naruto"}`

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users")

			if assert.NoError(t, controller.Create(c)) {
				assert.Equal(t, http.StatusInternalServerError, rec.Code)
				body := rec.Body.String()

				response := make(map[string]interface{})
				err := json.Unmarshal([]byte(body), &response)

				if assert.NoError(t, err) {
					errorMsg := response["error"].(string)
					assert.NotZero(t, errorMsg)
				}
			}
		})
	})
}

func TestLogin(t *testing.T) {
	mockUserService := mocks.UserService{}
	mockAuthService := mocks.AuthService{}

	t.Run("success scenario", func(t *testing.T) {
		dummyReq := model.LoginUserRequest{
			Email:    "naruto@gmail.com",
			Password: "naruto",
		}
		dummyResp := model.GetUserResponse{
			Id:        1,
			Email:     "naruto@gmail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		dummyToken := "xyz.xyz.xyz"

		mockUserService.On("Login", mock.AnythingOfType("model.LoginUserRequest")).Return(
			func(request model.LoginUserRequest) model.GetUserResponse {
				return dummyResp
			},
			func(request model.LoginUserRequest) error {
				return nil
			},
		).Once()

		mockAuthService.On("GenerateToken", dummyResp.Id, dummyResp.Email).Return(
			func(id uint, email string) string {
				return dummyToken
			},
			func(id uint, email string) error {
				return nil
			},
		).Once()

		t.Run("it should return 200 status code with valid response", func(t *testing.T) {
			controller := NewUserController(&mockUserService, &mockAuthService)
			requestBody, err := json.Marshal(dummyReq)
			assert.NoError(t, err)

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(requestBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/login")

			if assert.NoError(t, controller.Login(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				body := rec.Body.String()

				response := make(map[string]interface{})
				err := json.Unmarshal([]byte(body), &response)

				if assert.NoError(t, err) {
					token := response["token"].(string)
					assert.Equal(t, dummyToken, token)
				}
			}
		})
	})

	t.Run("failed scenario", func(t *testing.T) {
		t.Run("it should return 400 status code if request body is invalid", func(t *testing.T) {
			controller := NewUserController(&mockUserService, &mockAuthService)
			invalidJSONreq := `{"email": "naruto@gmail.com", "password": false}`

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(invalidJSONreq))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/login")

			if assert.NoError(t, controller.Login(c)) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
				body := rec.Body.String()

				response := make(map[string]interface{})
				err := json.Unmarshal([]byte(body), &response)

				if assert.NoError(t, err) {
					errorMsg := response["error"].(string)
					assert.NotZero(t, errorMsg)
				}
			}
		})

		t.Run("it should return 404 status code if username or password not found", func(t *testing.T) {
			mockUserService.On("Login", mock.AnythingOfType("model.LoginUserRequest")).Return(
				func(request model.LoginUserRequest) model.GetUserResponse {
					return model.GetUserResponse{}
				},
				func(request model.LoginUserRequest) error {
					return errors.New("invalid email or password.")
				},
			).Once()

			controller := NewUserController(&mockUserService, &mockAuthService)
			reqBody := `{"email": "naruto@gmail.com", "password": "naruto"}`

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/login")

			if assert.NoError(t, controller.Login(c)) {
				assert.Equal(t, http.StatusNotFound, rec.Code)
				body := rec.Body.String()

				response := make(map[string]interface{})
				err := json.Unmarshal([]byte(body), &response)

				if assert.NoError(t, err) {
					errorMsg := response["error"].(string)
					assert.NotZero(t, errorMsg)
				}
			}
		})

		t.Run("it should return 500 status code if error happened", func(t *testing.T) {
			dummyResp := model.GetUserResponse{
				Id:        1,
				Email:     "naruto@gmail.com",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			mockUserService.On("Login", mock.AnythingOfType("model.LoginUserRequest")).Return(
				func(request model.LoginUserRequest) model.GetUserResponse {
					return dummyResp
				},
				func(request model.LoginUserRequest) error {
					return nil
				},
			).Once()

			mockAuthService.On("GenerateToken", dummyResp.Id, dummyResp.Email).Return(
				func(id uint, email string) string {
					return ""
				},
				func(id uint, email string) error {
					return errors.New("something went wrong.")
				},
			).Once()

			controller := NewUserController(&mockUserService, &mockAuthService)
			reqBody := `{"email": "naruto@gmail.com", "password": "naruto"}`

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/login")

			if assert.NoError(t, controller.Login(c)) {
				assert.Equal(t, http.StatusInternalServerError, rec.Code)
				body := rec.Body.String()

				response := make(map[string]interface{})
				err := json.Unmarshal([]byte(body), &response)

				if assert.NoError(t, err) {
					errorMsg := response["error"].(string)
					assert.NotZero(t, errorMsg)
				}
			}
		})
	})
}
