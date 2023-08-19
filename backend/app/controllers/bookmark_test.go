package controllers

import (
	"api/db/crud"
	"api/db/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var (
	testUser2 *models.User
	bookmark1 *models.Bookmark
)

func testBookmarkController(t *testing.T) {
	// prepare test user
	testUser2, _ = crud.AddUser(&models.User{
		Name:      "testerBookmark",
		Authority: models.AuthorityNone,
	}, nil)

	bookmark1, _ = crud.AddBookmark(&models.Bookmark{
		UserID: testUser2.ID,
		Title:  "Bookmark Test Title 1",
		Link:   "https://cheesecat47.github.io/",
	})

	t.Run("testAddBookmarkHandler", testAddBookmarkHandler)
	t.Run("testGetBookmarkByIdHandler", testGetBookmarkByIdHandler)
	t.Run("testGetAllBookmarksHandler", testGetAllBookmarksHandler)
}

// testAddBookmarkHandler : [controllers.AddBookmarkHandler] 테스트
//
// # Test Cases
//   - Case 1: 정상적으로 북마크 추가
//   - Case 2: 유저 아이디 없이 북마크 추가
//   - Case 3: 제목 없이 북마크 추가
//   - Case 4: 링크 없이 북마크 추가
func testAddBookmarkHandler(t *testing.T) {
	testCases := []struct {
		name            string
		method          string
		route           string
		body            models.AddBookmarkRequest
		expectedError   bool
		expectedCode    int
		expectedMessage string
	}{
		{
			name:   "Add bookmark -> success",
			method: "POST",
			route:  "/api/v1/bookmark/",
			body: models.AddBookmarkRequest{
				UserID: testUser2.ID,
				Title:  "bookmark_test case 1 title",
				Link:   "https://cheesecat47.github.io/bookmark_test/case1/link",
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
		{
			name:   "Add bookmark without userId -> fail",
			method: "POST",
			route:  "/api/v1/bookmark/",
			body: models.AddBookmarkRequest{
				Title: "bookmark_test case 2 title",
				Link:  "https://cheesecat47.github.io/bookmark_test/case2/link",
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
		{
			name:   "Add bookmark without title -> success",
			method: "POST",
			route:  "/api/v1/bookmark/",
			body: models.AddBookmarkRequest{
				UserID: testUser2.ID,
				Link:   "https://cheesecat47.github.io/bookmark_test/case3/link",
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
		{
			name:   "Add bookmark without link -> fail",
			method: "POST",
			route:  "/api/v1/bookmark/",
			body: models.AddBookmarkRequest{
				UserID: testUser2.ID,
				Title:  "bookmark_test case 4 title",
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			err := json.NewEncoder(buf).Encode(tt.body)
			require.NoError(t, err)

			req := httptest.NewRequest(tt.method, tt.route, buf)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("UserID", strconv.Itoa(int(tt.body.UserID)))
			t.Log("req: ", req)

			resp, err := app.Test(req, -1)
			t.Log("resp: ", resp)
			require.NoError(t, err)

			result := &models.AddBookmarkResponse{}
			err = json.NewDecoder(resp.Body).Decode(result)
			require.NoError(t, err)

			require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
			require.Equal(t, tt.expectedMessage, result.Message, result.Message)
		})

	}
}

// testGetBookmarkByIdHandler : [controllers.GetBookmarkByIdHandler] 테스트
//
// # Test Cases:
//   - Case 1: bookmark1의 아이디를 사용해 조회
//   - Case 2: 존재하지 않는 북마크 아이디를 사용해 조회
func testGetBookmarkByIdHandler(t *testing.T) {
	testCases := []struct {
		name            string
		method          string
		route           string
		body            models.GetBookmarkByIdRequest
		expectedError   bool
		expectedCode    int
		expectedMessage string
	}{
		{
			name:   "Get bookmark which index is 1 -> success",
			method: "GET",
			route:  "/api/v1/bookmark/",
			body: models.GetBookmarkByIdRequest{
				ID: bookmark1.ID,
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
		{
			name:   "Get bookmark which index not exists -> fail",
			method: "GET",
			route:  "/api/v1/bookmark/",
			body: models.GetBookmarkByIdRequest{
				ID: 0,
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "record not found",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, fmt.Sprintf("%s%d", tt.route, tt.body.ID), nil)
			req.Header.Set("Content-Type", "application/json")
			t.Log("req: ", req)

			resp, err := app.Test(req, -1)
			t.Log("resp: ", resp)
			require.NoError(t, err)

			result := &models.GetBookmarkByIdResponse{}
			err = json.NewDecoder(resp.Body).Decode(result)
			require.NoError(t, err)

			require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
			require.Equal(t, tt.expectedMessage, result.Message, result.Message)
		})
	}
}

// testGetAllBookmarksHandler : [controllers.GetAllBookmarksHandler] 테스트
//
// # Test Cases:
//   - Case 1: 쿼리 파라미터 없이 전체 북마크 조회
//   - Case 2: offset 없이 북마크 조회
//   - Case 3: offset 파라미터 값으로 음수를 사용하는 경우
//   - Case 4: limit 파라미터 값으로 음수를 사용하는 경우
//   - Case 5: offset, limit 값을 설정해 일부 범위의 유저 조회
func testGetAllBookmarksHandler(t *testing.T) {
	testCases := []struct {
		name            string
		method          string
		route           string
		body            models.GetAllBookmarksRequest
		expectedError   bool
		expectedCode    int
		expectedMessage string
	}{
		{
			name:            "Get all bookmark -> success",
			method:          "GET",
			route:           "/api/v1/bookmark/",
			body:            models.GetAllBookmarksRequest{},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
		{
			name:   "Get all bookmarks with no offset -> success",
			method: "GET",
			route:  "/api/v1/bookmark/",
			body: models.GetAllBookmarksRequest{
				Limit: 1,
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
		{
			name:   "Get all bookmarks using negative offset value -> fail",
			method: "GET",
			route:  "/api/v1/bookmark/",
			body: models.GetAllBookmarksRequest{
				Offset: -1,
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
		{
			name:   "Get all bookmarks using negative limit value -> fail",
			method: "GET",
			route:  "/api/v1/bookmark/",
			body: models.GetAllBookmarksRequest{
				Limit: -1,
			},
			expectedError:   true,
			expectedCode:    http.StatusBadRequest,
			expectedMessage: "Validation failed",
		},
		{
			name:   "Get all bookmarks using limit and offset -> success",
			method: "GET",
			route:  "/api/v1/bookmark/",
			body: models.GetAllBookmarksRequest{
				Offset: int(bookmark1.ID),
				Limit:  1,
			},
			expectedError:   false,
			expectedCode:    http.StatusOK,
			expectedMessage: "Success",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				tt.method,
				fmt.Sprintf("%s?offset=%d&limit=%d", tt.route, tt.body.Offset, tt.body.Limit),
				nil)
			req.Header.Set("Content-Type", "application/json")
			t.Log("req: ", req)

			resp, err := app.Test(req, -1)
			t.Log("resp: ", resp)
			require.NoError(t, err)

			result := &models.GetAllBookmarksResponse{}
			err = json.NewDecoder(resp.Body).Decode(result)
			require.NoError(t, err, result.Message)

			require.Equal(t, tt.expectedCode, resp.StatusCode, result.Message)
			require.Equal(t, tt.expectedMessage, result.Message, result.Message)
		})
	}
}
