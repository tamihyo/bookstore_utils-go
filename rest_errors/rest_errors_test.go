package rest_errors

import (
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	// err := NewInternalServerError("this is the message", errors.New("database error"))
	// assert.NotNil(t, err)
	// assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	// assert.EqualValues(t, "this is the message", err.Message)
	// assert.EqualValues(t, "internal_server_error", err.Error)
	// assert.NotNil(t, err.Causes)
	// assert.EqualValues(t, 1, len(err.Causes))
	// assert.EqualValues(t, 1, err.Causes[0])

	// errBytes, _ := json.Marshal(err)

	// fmt.Println(string(errBytes))
}

func TestNewbadRequestError(t *testing.T) {
	//TODO : Test
}

func TestNewNotFoundError(t *testing.T) {

}

func TestNewError(t *testing.T) {

}
