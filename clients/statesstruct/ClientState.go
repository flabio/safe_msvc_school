package statesstruct

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/insfractruture/utils"
)

func MsvcStateFindById(id uint, c *fiber.Ctx) (StateResponse, string) {
	data, err := http.NewRequest(utils.GET, utils.MSVC_STATES_URL+"/"+strconv.Itoa(int(id)), nil)
	data.Header.Set(utils.AUTHORIZATION, c.Get(utils.AUTHORIZATION))
	resp, msg := DataStateClient(data, err)
	return resp, msg
}
