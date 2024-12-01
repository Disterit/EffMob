package test

import (
	"EffMob/models"
	"EffMob/pkg/handler"
	"EffMob/pkg/service"
	mock_service "EffMob/pkg/service/mocks"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_CreateSong(t *testing.T) {
	type mockBehavior func(s *mock_service.MockSong, song handler.AddSong, songInfo *models.SongInfo)

	testTable := []struct {
		name                 string
		userInput            string
		bodyInput            handler.AddSong
		externalApiResponse  *models.SongInfo
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			userInput: `{"group": "Muse", "song": "Supermassive Black Hole"}`,
			bodyInput: handler.AddSong{
				GroupName: "Muse",
				SongName:  "Supermassive Black Hole",
			},
			externalApiResponse: &models.SongInfo{
				Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
				Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
				ReleaseDate: "16.07.2006",
			},
			mockBehavior: func(s *mock_service.MockSong, song handler.AddSong, songInfo *models.SongInfo) {
				s.EXPECT().CreateSong(song.GroupName, song.SongName, songInfo).Return(1, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			createSong := mock_service.NewMockSong(c)
			testCase.mockBehavior(createSong, testCase.bodyInput, testCase.externalApiResponse)

			services := &service.Service{Song: createSong}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.POST("/song", handlers.CreateSong)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/song", bytes.NewBufferString(testCase.userInput))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_GetAllSong(t *testing.T) {
	type mockBehavior func(s *mock_service.MockSong)

	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "ok",
			mockBehavior: func(s *mock_service.MockSong) {
				s.EXPECT().GetAllSongs().Return([]models.Song{
					{
						1,
						1,
						sql.NullString{},
						sql.NullString{},
						sql.NullString{},
						sql.NullTime{},
					},
				}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":1,"group_id":1,"name":null,"text":null,"link":null,"release_date":null}]`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			getAllSong := mock_service.NewMockSong(c)
			testCase.mockBehavior(getAllSong)

			services := &service.Service{Song: getAllSong}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.GET("/song", handlers.GetAllSong)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/song", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_GetSongById(t *testing.T) {
	type mockBehavior func(s *mock_service.MockSong, id int)

	testTable := []struct {
		name                 string
		userInput            int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			userInput: 1,
			mockBehavior: func(s *mock_service.MockSong, id int) {
				s.EXPECT().GetSongById(1).Return(models.Song{
					1,
					1,
					sql.NullString{},
					sql.NullString{},
					sql.NullString{},
					sql.NullTime{},
				}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"id":1,"group_id":1,"name":null,"text":null,"link":null,"release_date":null}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			GetSongById := mock_service.NewMockSong(c)
			testCase.mockBehavior(GetSongById, testCase.userInput)

			services := &service.Service{Song: GetSongById}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.GET("/song/:id", handlers.GetSongById)

			url := fmt.Sprintf("/song/%d", testCase.userInput)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, url, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func stringToPtr(s string) *string {
	return &s
}

func TestHandler_UpdateSong(t *testing.T) {
	type mockBehavior func(s *mock_service.MockSong, id int, input models.UpdateSong)

	testTable := []struct {
		name                 string
		userInput            int
		bodyInput            models.UpdateSong
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			userInput: 1,
			bodyInput: models.UpdateSong{
				SongName:    stringToPtr("New Song Name"),
				Text:        stringToPtr("New Song Text"),
				Link:        stringToPtr("https://example.com"),
				ReleaseDate: stringToPtr("2024-12-01"),
			},
			mockBehavior: func(s *mock_service.MockSong, id int, input models.UpdateSong) {
				s.EXPECT().UpdateSong(1, input).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":"ok"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			updateSong := mock_service.NewMockSong(c)
			testCase.mockBehavior(updateSong, testCase.userInput, testCase.bodyInput)

			services := &service.Service{Song: updateSong}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.PATCH("/song/:id", handlers.UpdateSong)

			reqBody, _ := json.Marshal(testCase.bodyInput)
			url := fmt.Sprintf("/song/%d", testCase.userInput)
			req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(string(reqBody)))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteSong(t *testing.T) {
	type mockBehavior func(s *mock_service.MockSong, id int)

	testTable := []struct {
		name                 string
		userInput            int
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			userInput: 1,
			mockBehavior: func(s *mock_service.MockSong, id int) {
				s.EXPECT().DeleteSong(1).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":"ok"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			deleteSong := mock_service.NewMockSong(c)
			testCase.mockBehavior(deleteSong, testCase.userInput)

			services := &service.Service{Song: deleteSong}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.DELETE("/song/:id", handlers.DeleteSong)

			url := fmt.Sprintf("/song/%d", testCase.userInput)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, url, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
