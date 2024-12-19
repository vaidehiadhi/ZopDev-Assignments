package main

import (
	"bytes"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"gotest.tools/v3/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBook(t *testing.T) {
	type testCase struct {
		name               string
		requestPath        string
		mockQuerySetup     func(mock sqlmock.Sqlmock)
		expectedStatusCode int
		expectedResponse   string
	}

	db, mock, err := sqlmock.New()
	assert.NilError(t, err)

	defer db.Close()

	store := &Store{db: db}
	tests := []testCase{
		{
			name:        "Valid Book ID",
			requestPath: "/book/1",
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id, author, title FROM book WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "author", "title"}).
						AddRow(1, "vaidehi", "meow"))
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"id":1,"author":"vaidehi","title":"meow"}` + "\n",
		},
		{
			name:        "Invalid Book ID",
			requestPath: "/book/abc",
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   "Invalid book ID\n",
		},
		{
			name:        "Book Not Found",
			requestPath: "/book/999",
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id, author, title FROM book WHERE id = ?").
					WithArgs(999).
					WillReturnError(sql.ErrNoRows)
			},
			expectedStatusCode: http.StatusNotFound,
			expectedResponse:   "Book not found\n",
		},
		{
			name:        "Database Error",
			requestPath: "/book/1",
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id, author, title FROM book WHERE id = ?").
					WithArgs(1).
					WillReturnError(sql.ErrConnDone)
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   "Database error :connection done\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockQuerySetup != nil {
				tc.mockQuerySetup(mock)
			}

			req := httptest.NewRequest("GET", tc.requestPath, nil)
			rec := httptest.NewRecorder()

			store.GetBook(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)

			body := rec.Body.String()
			assert.Equal(t, tc.expectedResponse, body)
			assert.NilError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAddBook(t *testing.T) {
	type testCase struct {
		name               string
		requestBody        string
		mockQuerySetup     func(mock sqlmock.Sqlmock)
		expectedStatusCode int
		expectedResponse   string
	}

	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	store := &Store{db: db}
	tests := []testCase{
		{
			name:        "Valid Book",
			requestBody: `{"author": "vaidehi", "title": "book-2"}`,
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO book (author, title) VALUES (?, ?)").
					WithArgs("vaidehi", "book-2").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"id":1,"author":"vaidehi","title":"book-2"}` + "\n",
		},
		{
			name:        "Invalid JSON Body",
			requestBody: `{"author": "vaidehi", "title":}`,
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   "Invalid request body\n",
		},
		{
			name:        "Database Error",
			requestBody: `{"author": "vaidehi", "title": "book-2"}`,
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO book (author, title) VALUES (?, ?)").
					WithArgs("vaidehi", "book-2").
					WillReturnError(errors.New("database error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   "Failed to add book. err: database error\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectationsWereMet()
			if tc.mockQuerySetup != nil {
				tc.mockQuerySetup(mock)
			}

			req := httptest.NewRequest("POST", "/book", bytes.NewBufferString(tc.requestBody))
			rec := httptest.NewRecorder()

			store.AddBook(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)

			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, tc.expectedResponse, string(body))

			assert.NilError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdateBook(t *testing.T) {
	type testCase struct {
		name               string
		requestURL         string
		requestBody        string
		mockQuerySetup     func(mock sqlmock.Sqlmock)
		expectedStatusCode int
		expectedResponse   string
	}

	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	store := &Store{db: db}

	tests := []testCase{
		{
			name:        "Valid Update",
			requestURL:  "/book/1",
			requestBody: `{"author": "updated author", "title": "updated title"}`,
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE book SET author = ?, title = ? WHERE id = ?").
					WithArgs("updated author", "updated title", 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "Book with ID 1 updated",
		},
		{
			name:        "Invalid Book ID",
			requestURL:  "/book/invalid",
			requestBody: `{"author": "updated author", "title": "updated title"}`,
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   "Invalid book ID: strconv.Atoi: parsing \"invalid\": invalid syntax\n",
		},

		{
			name:        "Database Error",
			requestURL:  "/book/1",
			requestBody: `{"author": "updated author", "title": "updated title"}`,
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE book SET author = ?, title = ? WHERE id = ?").
					WithArgs("updated author", "updated title", 1).
					WillReturnError(errors.New("database error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   "Failed to update book: database error\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectationsWereMet()

			if tc.mockQuerySetup != nil {
				tc.mockQuerySetup(mock)
			}

			req := httptest.NewRequest("PUT", tc.requestURL, bytes.NewBufferString(tc.requestBody))
			rec := httptest.NewRecorder()

			store.UpdateBook(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)

			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, tc.expectedResponse, string(body))

			assert.NilError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteBook(t *testing.T) {
	type testCase struct {
		name               string
		requestURL         string
		mockQuerySetup     func(mock sqlmock.Sqlmock)
		expectedStatusCode int
		expectedResponse   string
	}

	db, mock, err := sqlmock.New()
	assert.NilError(t, err)
	defer db.Close()

	store := &Store{db: db}

	tests := []testCase{
		{
			name:       "Valid Deletion",
			requestURL: "/book/1",
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM book WHERE id = ?").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "Book with ID 1 deleted",
		},
		{
			name:       "Invalid Book ID",
			requestURL: "/book/invalid",
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   "Invalid book ID:strconv.Atoi: parsing \"invalid\": invalid syntax\n",
		},
		{
			name:       "Book Not Found",
			requestURL: "/book/999",
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM book WHERE id = ?").
					WithArgs(999).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "Book with ID 999 deleted",
		},
		{
			name:       "Database Error",
			requestURL: "/book/2",
			mockQuerySetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM book WHERE id = ?").
					WithArgs(2).
					WillReturnError(errors.New("database error"))
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   "Failed to delete book:database error\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mock.ExpectationsWereMet()

			if tc.mockQuerySetup != nil {
				tc.mockQuerySetup(mock)
			}

			req := httptest.NewRequest("DELETE", tc.requestURL, nil)
			rec := httptest.NewRecorder()

			store.DeleteBook(rec, req)

			res := rec.Result()
			assert.Equal(t, tc.expectedStatusCode, res.StatusCode)

			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, tc.expectedResponse, string(body))

			assert.NilError(t, mock.ExpectationsWereMet())
		})
	}
}
