package api_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"user_registry/internal/entity"
	"user_registry/internal/handler/http/api"
	mocks "user_registry/internal/handler/http/api/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

type mockBehavior func(r *mocks.MockUseCase, user entity.TextKey)

var (
	signHMACCases = []struct {
		name               string
		inputBody          string
		inputTK            entity.TextKey
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "Ok",
			inputBody: `{"text": "username", "key": "TestName"}`,
			inputTK: entity.TextKey{
				Text: "username",
				Key:  "TestName",
			},
			mockBehavior: func(r *mocks.MockUseCase, tk entity.TextKey) {
				r.EXPECT().Sign(context.Background(), &tk).Return("hexcode", nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:               "Wrong Input",
			inputBody:          `{"username": "username"}`,
			inputTK:            entity.TextKey{},
			mockBehavior:       func(_ *mocks.MockUseCase, _ entity.TextKey) {},
			expectedStatusCode: 400,
		},
		{
			name:      "Service Error",
			inputBody: `{"text": "username", "key": "TestName"}`,
			inputTK: entity.TextKey{
				Text: "username",
				Key:  "TestName",
			},
			mockBehavior: func(r *mocks.MockUseCase, tk entity.TextKey) {
				r.EXPECT().Sign(context.Background(), &tk).Return("", errors.New("something went wrong"))
			},
			expectedStatusCode: 500,
		},
	}
)

func TestHandler_SignHMAC(t *testing.T) {
	t.Parallel()

	signPath := "/sign/hmacsha512"

	for _, test := range signHMACCases {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			usecase := mocks.NewMockUseCase(c)
			test.mockBehavior(usecase, test.inputTK)

			handler := api.NewHandler(usecase)

			// Init Endpoint
			r := gin.New()
			r.POST(signPath, handler.SignHMAC)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, signPath,
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}

// func NewHandler(usecase *mocks.MockUseCase) {
// 	panic("unimplemented")
// }
