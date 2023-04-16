package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/mchaelhuang/aquafarm/internal"
	"github.com/mchaelhuang/aquafarm/internal/config"
	"github.com/mchaelhuang/aquafarm/internal/constant"
	"github.com/mchaelhuang/aquafarm/internal/entity"
)

type FarmDelivery struct {
	cfg    *config.Cfg
	farmUC internal.FarmUC
}

func NewFarmDelivery(
	cfg *config.Cfg,
	farmUC internal.FarmUC,
) *FarmDelivery {
	return &FarmDelivery{
		cfg:    cfg,
		farmUC: farmUC,
	}
}

func (f *FarmDelivery) GetList(c echo.Context) error {
	request := entity.FarmFilterRequest{}
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	filter := entity.FarmFilter{
		Farm: entity.Farm{
			Name:     request.Name,
			Location: request.Location,
		},
	}
	farms, err := f.farmUC.GetList(c.Request().Context(), filter)
	if err != nil {
		return handlerErrorResponse(err)
	}
	return c.JSON(http.StatusOK, farms)
}

func (f *FarmDelivery) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handlerErrorResponse(constant.ErrNotFound)
	}

	farm, err := f.farmUC.Get(
		c.Request().Context(),
		entity.FarmFilter{Farm: entity.Farm{ID: id}},
	)
	if err != nil {
		return handlerErrorResponse(err)
	}
	return c.JSON(http.StatusOK, farm)
}

func (f *FarmDelivery) Create(c echo.Context) error {
	request := entity.FarmFormRequest{}
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Validation
	if request.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name required")
	}
	if request.Location == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "location required")
	}

	farm := entity.Farm{
		Name:     request.Name,
		Location: request.Location,
	}
	id, err := f.farmUC.Store(c.Request().Context(), farm)
	if err != nil {
		return handlerErrorResponse(err)
	}

	return c.JSON(http.StatusOK, entity.FarmFormResponse{ID: id})
}

func (f *FarmDelivery) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handlerErrorResponse(constant.ErrNotFound)
	}

	request := entity.FarmFormRequest{}
	err = c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	farm := entity.Farm{
		ID:       id,
		Name:     request.Name,
		Location: request.Location,
	}
	id, err = f.farmUC.Store(c.Request().Context(), farm)
	if err != nil {
		return handlerErrorResponse(err)
	}

	return c.JSON(http.StatusOK, entity.FarmFormResponse{ID: id})
}

func (f *FarmDelivery) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handlerErrorResponse(constant.ErrNotFound)
	}

	err = f.farmUC.Delete(c.Request().Context(), id)
	if err != nil {
		return handlerErrorResponse(err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}
