package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Layla": 20,
			"Izzy":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("returns Layla's score", func(t *testing.T) {
		request := newGetScoreRequest("Layla")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "20"

		assertResponseBody(t, response.Body.String(), want)
		assertResponseCode(t, response.Code, http.StatusOK)
	})
	t.Run("returns Izzy's score", func(t *testing.T) {
		request := newGetScoreRequest("Izzy")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "10"

		assertResponseBody(t, response.Body.String(), want)
		assertResponseCode(t, response.Code, http.StatusOK)
	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Sam")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it returns acceepted on POST", func(t *testing.T) {
		player := "Layla"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("did not store correct winnget got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeage := []Player{
			{"Sam", 100},
			{"Izzy", 50},
			{"Layla", 75},
		}
		store := StubPlayerStore{nil, nil, wantedLeage}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertResponseCode(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeage)
		assertContentType(t, response, jsonContentType)

	})

}

func getLeagueFromResponse(t testing.TB, body io.Reader) []Player {
	t.Helper()
	var league []Player
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return league

}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, contentType string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != contentType {
		t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
	}
}

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertResponseCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}
