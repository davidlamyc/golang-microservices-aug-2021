package rest

import (
	"os"
	"testing"
	"net/http"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL: "http://localhost:8080/users/login",
		HTTPMethod: http.MethodPost,
		ReqBody: `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: -1,
		RespBody: `{}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)

}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	
}

func TestLoginUserNoError(t *testing.T) {
	
}