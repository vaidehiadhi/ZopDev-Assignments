package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func Test_GetBook(t *testing.T) {
	testCases := []struct {
		name      string
		method    string
		id        int
		expected  string
		expStatus int
	}{
		{
			name:      "successfully get a book",
			method:    http.MethodGet,
			id:        1,
			expected:  `{"id":1,"author":"Vaidehi","title":"Vai"}`,
			expStatus: http.StatusOK,
		},
		{
			name:      "book not found",
			method:    http.MethodGet,
			id:        999,
			expected:  "Book not found\n",
			expStatus: http.StatusNotFound,
		},
		{
			name:      "invalid book ID",
			method:    http.MethodGet,
			id:        -1,
			expected:  "Invalid book ID\n",
			expStatus: http.StatusBadRequest,
		},
	}
	Books[1] = Book{Id: 1, Author: "Vaidehi", Title: "Vai"}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := "/book/" + strconv.Itoa(tc.id)
			req := httptest.NewRequest(tc.method, url, nil)
			w := httptest.NewRecorder()

			GetBook(w, req)

			resp := w.Result()
			b, _ := io.ReadAll(resp.Body)

			if string(b) != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, string(b))
			}

			if w.Code != tc.expStatus {
				t.Errorf("Expected %d, but got %d", tc.expStatus, w.Code)
			}
		})
	}
}
func Test_AddBook(t *testing.T) {
	testCases := []struct {
		name      string
		method    string
		body      string
		expected  string
		expStatus int
	}{
		{
			name:      "successfully add a book",
			method:    http.MethodPost,
			body:      `{"author":"Vaidehi","title":"Vai"}`,
			expected:  `{"id":1,"author":"Vaidehi","title":"Vai"}`,
			expStatus: http.StatusOK,
		},
		{
			name:      "invalid request body",
			method:    http.MethodPost,
			body:      `{"author":"Vaidehi","title":1234}`,
			expected:  "Invalid request body\n",
			expStatus: http.StatusBadRequest,
		},
		{
			name:      "empty request body",
			method:    http.MethodPost,
			body:      ``,
			expected:  "Failed to marshal book\n",
			expStatus: http.StatusInternalServerError,
		},
	}

	Books = make(map[int]Book)
	NextID = 1

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/book", strings.NewReader(tc.body))
			w := httptest.NewRecorder()

			AddBook(w, req)

			resp := w.Result()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("failed %v", err)
			}

			if string(b) != tc.expected {
				t.Errorf("expected %s, but got %s", tc.expected, string(b))
			}

			if w.Code != tc.expStatus {
				t.Errorf("expected %d, but got %d", tc.expStatus, w.Code)
			}
		})
	}
}

func Test_UpdateBook(t *testing.T) {
	testCases := []struct {
		name      string
		method    string
		id        int
		body      string
		expected  string
		expStatus int
	}{
		{
			name:      "successfully update a book",
			method:    http.MethodPut,
			id:        1,
			body:      `{"author":"Updated Author","title":"Updated Title"}`,
			expected:  `{"id":1,"author":"Updated Author","title":"Updated Title"}`,
			expStatus: http.StatusOK,
		},
		{
			name:      "update only title",
			method:    http.MethodPut,
			id:        1,
			body:      `{"title":"Only Title Updated"}`,
			expected:  `{"id":1,"author":"Original Author","title":"Only Title Updated"}`,
			expStatus: http.StatusOK,
		},
		{
			name:      "update only author",
			method:    http.MethodPut,
			id:        1,
			body:      `{"author":"Only Author Updated"}`,
			expected:  `{"id":1,"author":"Only Author Updated","title":"Original Title"}`,
			expStatus: http.StatusOK,
		},
		{
			name:      "book not found",
			method:    http.MethodPut,
			id:        999,
			body:      `{"author":"Nonexistent Author","title":"Nonexistent Title"}`,
			expected:  "book not found\n",
			expStatus: http.StatusNotFound,
		},
		{
			name:      "invalid book ID",
			method:    http.MethodPut,
			id:        -1,
			body:      `{"author":"Invalid Author","title":"Invalid Title"}`,
			expected:  "Invalid book ID\n",
			expStatus: http.StatusBadRequest,
		},
		{
			name:      "empty request body",
			method:    http.MethodPut,
			id:        1,
			body:      ``,
			expected:  "failed to read request body\n",
			expStatus: http.StatusInternalServerError,
		},
	}

	Books = make(map[int]Book)
	Books[1] = Book{Id: 1, Author: "Original Author", Title: "Original Title"}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := "/book/" + strconv.Itoa(tc.id)
			req := httptest.NewRequest(tc.method, url, strings.NewReader(tc.body))
			w := httptest.NewRecorder()

			UpdateBook(w, req)

			resp := w.Result()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed: %v", err)
			}

			if string(b) != tc.expected {
				t.Errorf("expected %s, but got %s", tc.expected, string(b))
			}

			if w.Code != tc.expStatus {
				t.Errorf("expected %d, but got %d", tc.expStatus, w.Code)
			}
		})
	}
}

func Test_DeleteBook(t *testing.T) {
	testCases := []struct {
		name      string
		method    string
		id        int
		expected  string
		expStatus int
	}{
		{
			name:      "successfully delete a book",
			method:    http.MethodDelete,
			id:        1,
			expected:  "Book with ID 1 deleted",
			expStatus: http.StatusOK,
		},
		{
			name:      "book not found",
			method:    http.MethodDelete,
			id:        999,
			expected:  "Book not found\n",
			expStatus: http.StatusNotFound,
		},
		{
			name:      "invalid book ID",
			method:    http.MethodDelete,
			id:        -1,
			expected:  "Invalid book ID\n",
			expStatus: http.StatusBadRequest,
		},
	}

	Books = make(map[int]Book)
	Books[1] = Book{Id: 1, Author: "test Author", Title: "test Title"}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := "/book/" + strconv.Itoa(tc.id)
			req := httptest.NewRequest(tc.method, url, nil)
			w := httptest.NewRecorder()

			DeleteBook(w, req)

			resp := w.Result()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed: %v", err)
			}

			if string(b) != tc.expected {
				t.Errorf("expected %s, but got %s", tc.expected, string(b))
			}

			if w.Code != tc.expStatus {
				t.Errorf("expected %d, but got %d", tc.expStatus, w.Code)
			}

			if tc.expStatus == http.StatusOK {
				if _, exists := Books[tc.id]; exists {
					t.Errorf("book with ID %d was not deleted", tc.id)
				}
			}
		})
	}
}
