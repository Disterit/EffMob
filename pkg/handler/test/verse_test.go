package test

import (
	"EffMob/pkg/handler"
	"EffMob/pkg/service"
	mock_service "EffMob/pkg/service/mocks"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetVerses(t *testing.T) {
	type mockBehavior func(s *mock_service.MockVerse, songId, verseId, limit int)

	testTable := []struct {
		name                 string
		userInputId          int
		userInputVerseId     int
		userInputLimit       int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:             "ok",
			userInputId:      1,
			userInputVerseId: 1,
			userInputLimit:   0,
			mockBehavior: func(s *mock_service.MockVerse, songId, verseId, limit int) {
				s.EXPECT().GetVerses(songId, verseId, limit).Return(map[string]string{
					"1": "First verse",
					"2": "Second verse",
				}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"1":"First verse","2":"Second verse"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockVerse := mock_service.NewMockVerse(c)
			testCase.mockBehavior(mockVerse, testCase.userInputId, testCase.userInputVerseId, testCase.userInputLimit)

			services := &service.Service{Verse: mockVerse}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.GET("/verses/:id/:verse", handlers.GetVerses)

			url := fmt.Sprintf("/verses/%d/%d?limit=%d", testCase.userInputId, testCase.userInputVerseId, testCase.userInputLimit)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.JSONEq(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
