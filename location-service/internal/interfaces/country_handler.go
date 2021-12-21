package interfaces

import (
	"github.com/labstack/echo/v4"
	"grpc-demo/location-service/internal/usecase/country"
	"net/http"
	"strconv"
)

type CountryHandler struct {
	countryService country.CountryServiceIFace
}

func NewCountryHandler(countryService country.CountryServiceIFace) *CountryHandler {
	return &CountryHandler{
		countryService: countryService,
	}
}

func (h *CountryHandler) GetAll(ctx echo.Context) (err error) {
	resp, err := h.countryService.GetAll()
	if err != nil {
		panic(err)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (h *CountryHandler) GetById(ctx echo.Context) (err error) {
	countryId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	resp, err := h.countryService.GetById(countryId)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(http.StatusOK, resp)
}
