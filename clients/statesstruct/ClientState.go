package statesstruct

import (
	"net/http"
	"strconv"

	utils "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
)

func MsvcStateFindById(id uint, c *fiber.Ctx) (StateResponse, string) {
	data, err := http.NewRequest(utils.GET, utils.MSVC_STATES_URL+"/"+strconv.Itoa(int(id)), nil)
	data.Header.Set(utils.AUTHORIZATION, c.Get(utils.AUTHORIZATION))
	resp, msg := DataStateClient(data, err)
	return resp, msg
}
