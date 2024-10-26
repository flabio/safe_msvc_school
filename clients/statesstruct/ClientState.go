package statesstruct

import (
	"net/http"
	"strconv"

	constants "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
)

func MsvcStateFindById(id uint, c *fiber.Ctx) (StateResponse, string) {
	data, err := http.NewRequest(constants.GET, constants.MSVC_STATES_URL+"/"+strconv.Itoa(int(id)), nil)
	data.Header.Set(constants.AUTHORIZATION, c.Get(constants.AUTHORIZATION))
	resp, msg := DataStateClient(data, err)
	return resp, msg
}
