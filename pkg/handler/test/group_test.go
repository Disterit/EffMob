package test

import (
	"EffMob/models"
	"EffMob/pkg/handler"
	"EffMob/pkg/service"
	mock_service "EffMob/pkg/service/mocks"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_CreateGroup(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockGroup, group models.Group)

	testTable := []struct {
		name                 string
		userInput            models.Group
		bodyInput            string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "ok",
			userInput: models.Group{
				GroupName: "Muse",
			},
			bodyInput: `{"name":"Muse"}`,
			mockBehaviour: func(s *mock_service.MockGroup, group models.Group) {
				s.EXPECT().CreateGroup(group.GroupName).Return(1, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			createGroup := mock_service.NewMockGroup(c)
			testCase.mockBehaviour(createGroup, testCase.userInput)

			services := &service.Service{Group: createGroup}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.POST("/group", handlers.CreateGroup)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/group", strings.NewReader(testCase.bodyInput))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_GetAllLibrary(t *testing.T) {
	type mockBehavior func(s *mock_service.MockGroup)

	testTable := []struct {
		name                 string
		mockBehaviour        mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "ok",
			mockBehaviour: func(s *mock_service.MockGroup) {
				s.EXPECT().GetAllLibrary().Return(map[string][]models.Song{
					"test1": {
						{Id: 1, GroupId: 1, SongName: sql.NullString{String: "Song 1", Valid: true}},
						{Id: 2, GroupId: 1, SongName: sql.NullString{String: "Song 2", Valid: true}},
					},
					"test2": {
						{Id: 3, GroupId: 2, SongName: sql.NullString{String: "Song 3", Valid: true}},
					},
				}, nil)

			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"test1":[{"id":1,"group_id":1,"name":"Song 1","text":null,"link":null,"release_date":null},{"id":2,"group_id":1,"name":"Song 2","text":null,"link":null,"release_date":null}],"test2":[{"id":3,"group_id":2,"name":"Song 3","text":null,"link":null,"release_date":null}]}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			GetAllLibrary := mock_service.NewMockGroup(c)
			testCase.mockBehaviour(GetAllLibrary)

			services := &service.Service{Group: GetAllLibrary}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.GET("/group", handlers.GetAllLibrary)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/group", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_GetAllSongGroupById(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockGroup, id int)

	testTable := []struct {
		name                 string
		userInput            int
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "ok",
			mockBehaviour: func(s *mock_service.MockGroup, id int) {
				s.EXPECT().GetAllSongGroupById(id).Return(map[string][]models.Song{
					"test1": {
						{Id: 1, GroupId: 1, SongName: sql.NullString{String: "Song 1", Valid: true}},
						{Id: 2, GroupId: 1, SongName: sql.NullString{String: "Song 2", Valid: true}},
					},
				}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"test1":[{"id":1,"group_id":1,"name":"Song 1","text":null,"link":null,"release_date":null},{"id":2,"group_id":1,"name":"Song 2","text":null,"link":null,"release_date":null}]}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			GetAllSongGroupById := mock_service.NewMockGroup(c)
			testCase.mockBehaviour(GetAllSongGroupById, testCase.userInput)

			services := &service.Service{Group: GetAllSongGroupById}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.GET("/group/:id", handlers.GetAllSongGroupById)

			url := fmt.Sprintf("/group/%d", testCase.userInput)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, url, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_UpdateGroup(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockGroup, id int, input models.Group)

	testTable := []struct {
		name                 string
		userInputId          int
		userInput            models.Group
		bodyInput            string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "ok",
			userInputId: 1,
			userInput: models.Group{
				GroupName: "Muse",
			},
			bodyInput: `{"name":"Muse"}`,
			mockBehaviour: func(s *mock_service.MockGroup, id int, input models.Group) {
				s.EXPECT().UpdateGroup(id, gomock.Any()).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":"ok"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			updateGroup := mock_service.NewMockGroup(c)
			testCase.mockBehaviour(updateGroup, testCase.userInputId, testCase.userInput)

			services := &service.Service{Group: updateGroup}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.PATCH("/group/:id", handlers.UpdateGroup)

			url := fmt.Sprintf("/group/%d", testCase.userInputId)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPatch, url, strings.NewReader(testCase.bodyInput))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_DeleteGroup(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockGroup, id int)

	testTable := []struct {
		name                 string
		userInputId          int
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "ok",
			userInputId: 1,
			mockBehaviour: func(s *mock_service.MockGroup, id int) {
				s.EXPECT().DeleteGroup(id).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":"ok"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			deleteGroup := mock_service.NewMockGroup(c)
			testCase.mockBehaviour(deleteGroup, testCase.userInputId)

			services := &service.Service{Group: deleteGroup}
			handlers := handler.NewHandler(services)

			r := gin.New()
			r.DELETE("/group/:id", handlers.DeleteGroup)

			url := fmt.Sprintf("/group/%d", testCase.userInputId)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, url, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
