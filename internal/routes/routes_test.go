package routes

import (
	"github.com/steinfletcher/apitest"
	"net/http"
	"testing"
)

func TestGetAllMovies(t *testing.T) {
	apitest.New().
		Get("/cinema-service/v1/api/movies").
		Expect(t).
		Status(http.StatusOK).
		End()
}
