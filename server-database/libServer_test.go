package main

import (
	"database/sql"
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
			name:   "successfully get a book",
			method: http.MethodGet,
			id:     1,
			expected: `{"id":1,"author":"vaidehi","title":"meow"}
`,
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

	db, err := sql.Open("mysql", "root:password@tcp/sample")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	store := NewStore(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := "/book/" + strconv.Itoa(tc.id)
			req := httptest.NewRequest(tc.method, url, nil)
			w := httptest.NewRecorder()

			store.GetBook(w, req)

			resp := w.Result()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Failed to read response body: %v", err)
				return
			}

			if string(b) != tc.expected {
				t.Errorf("Expected body %q, but got %q", tc.expected, string(b))
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
			name:   "successfully added a book",
			method: http.MethodPost,
			body:   `{"author": "vaidehi", "title": "cat"}`,
			expected: `{"id":23,"author":"vaidehi","title":"cat"}
`,
			expStatus: http.StatusOK,
		},

		{
			name:      "invalid request body",
			method:    http.MethodPost,
			body:      `{"author":"Vaidehi","title":1234}`,
			expected:  "Invalid request body\n",
			expStatus: http.StatusBadRequest,
		},

		//{
		//	name:      "empty request body",
		//	method:    http.MethodPost,
		//	body:      ``,
		//	expected:  "Failed to marshal book\n",
		//	expStatus: http.StatusInternalServerError,
		//},

	}
	db, err := sql.Open("mysql", "root:password@tcp/sample")
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	}
	store := NewStore(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/book", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			store.AddBook(w, req)

			resp := w.Result()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Failed to read response body: %v", err)
			}

			if string(b) != tc.expected {
				t.Errorf("Expected body %q, but got %q", tc.expected, string(b))
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
		{name: "successfully updated a book",
			method:    http.MethodPut,
			id:        1,
			body:      `{"author": "updated author","title":"updated title"}`,
			expected:  `{"id":25,"author":"updated author","title":"updated title"}` + "\n",
			expStatus: http.StatusOK,
		},
	}

	db, err := sql.Open("mysql", "root:password@tcp/sample")
	if err != nil {
		t.Errorf("database is not connected %v", err)
	}
	store := NewStore(db)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, "/book", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			store.AddBook(w, req)

			resp := w.Result()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if string(b) != tc.expected {
				t.Errorf("Expected body %q, but got %q", tc.expected, string(b))
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
			id:        17,
			expected:  "Book with ID 17 deleted",
			expStatus: http.StatusOK,
		},
	}

	db, err := sql.Open("mysql", "root:password@tcp/sample")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	store := NewStore(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := "/book/" + strconv.Itoa(tc.id)
			req := httptest.NewRequest(tc.method, url, nil)
			w := httptest.NewRecorder()

			store.DeleteBook(w, req)

			resp := w.Result()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Failed to read response body: %v", err)
				return
			}

			if string(b) != tc.expected {
				t.Errorf("Expected body %q, but got %q", tc.expected, string(b))
			}
		})
	}
}
