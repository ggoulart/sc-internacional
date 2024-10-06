package teams

import (
	"bytes"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestController_PostTeam(t *testing.T) {
	tests := []struct {
		name                 string
		setup                func(*serviceMock)
		requestBody          string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "when request is invalid",
			setup:                func(s *serviceMock) {},
			requestBody:          "abcd",
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: "{\"error\":\"invalid character 'a' looking for beginning of value\"}",
		},
		{
			name: "when failed to create a team",
			setup: func(s *serviceMock) {
				receivedTeam := Team{Id: "", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				s.On("createTeam", mock.Anything, receivedTeam).Return(Team{}, errors.New("failed to create team"))
			},
			requestBody:          "{\"fullName\": \"Sport Club Internacional\", \"name\": \"Internacional\", \"foundationDate\": \"1909-04-04T00:00:00.000Z\", \"website\": \"internacional.com.br\"}",
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: "{\"error\":\"failed to create team\"}",
		},
		{
			name: "when successfully creates a team",
			setup: func(s *serviceMock) {
				receivedTeam := Team{Id: "", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				returnTeam := Team{Id: "1", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				s.On("createTeam", mock.Anything, receivedTeam).Return(returnTeam, nil)
			},
			requestBody:          "{\"fullName\": \"Sport Club Internacional\", \"name\": \"Internacional\", \"foundationDate\": \"1909-04-04T00:00:00.000Z\", \"website\": \"internacional.com.br\"}",
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: "{\"id\":\"1\",\"name\":\"Internacional\",\"fullName\":\"Sport Club Internacional\",\"website\":\"internacional.com.br\",\"foundationDate\":\"1909-04-04T00:00:00Z\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceMock{}
			tt.setup(s)

			c := NewController(s)

			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = &http.Request{
				Header: make(http.Header),
				Body:   io.NopCloser(bytes.NewBuffer([]byte(tt.requestBody))),
			}

			c.PostTeam(ctx)

			assert.Equal(t, tt.expectedStatusCode, recorder.Code)
			assert.Equal(t, tt.expectedResponseBody, recorder.Body.String())
		})
	}
}

func TestController_GetTeam(t *testing.T) {
	tests := []struct {
		name               string
		setup              func(*serviceMock)
		id                 string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name: "when failed to get team",
			setup: func(s *serviceMock) {
				s.On("getTeam", mock.Anything, "1").Return(Team{}, errors.New("failed to get team"))
			},
			id:                 "1",
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       "{\"error\":\"failed to get team\"}",
		},
		{
			name: "when team is not found",
			setup: func(s *serviceMock) {
				s.On("getTeam", mock.Anything, "1").Return(Team{}, nil)
			},
			id:                 "1",
			expectedStatusCode: http.StatusNotFound,
			expectedBody:       "{\"error\":\"team not found\"}",
		},
		{
			name: "when successfully get team",
			setup: func(s *serviceMock) {
				team := Team{Id: "1", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				s.On("getTeam", mock.Anything, "1").Return(team, nil)
			},
			id:                 "1",
			expectedStatusCode: http.StatusOK,
			expectedBody:       "{\"id\":\"1\",\"name\":\"Internacional\",\"fullName\":\"Sport Club Internacional\",\"website\":\"internacional.com.br\",\"foundationDate\":\"1909-04-04T00:00:00Z\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceMock{}
			tt.setup(s)

			c := NewController(s)

			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.AddParam("id", tt.id)
			ctx.Request = &http.Request{Header: make(http.Header)}

			c.GetTeam(ctx)

			assert.Equal(t, tt.expectedStatusCode, recorder.Code)
			assert.Equal(t, tt.expectedBody, recorder.Body.String())
		})
	}
}

type serviceMock struct {
	service
	mock.Mock
}

func (m *serviceMock) createTeam(ctx context.Context, team Team) (Team, error) {
	args := m.Called(ctx, team)

	return args.Get(0).(Team), args.Error(1)
}

func (m *serviceMock) getTeam(ctx context.Context, id string) (Team, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(Team), args.Error(1)
}
