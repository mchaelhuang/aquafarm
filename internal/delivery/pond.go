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

type PondDelivery struct {
	cfg    *config.Cfg
	pondUC internal.PondUC
}

func NewPondDelivery(
	cfg *config.Cfg,
	pondUC internal.PondUC,
) *PondDelivery {
	return &PondDelivery{
		cfg:    cfg,
		pondUC: pondUC,
	}
}

func (f *PondDelivery) GetList(c echo.Context) error {
	request := entity.PondFilterRequest{}
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	filter := entity.PondFilter{
		Pond: entity.Pond{
			FarmID: request.FarmID,
			Label:  request.Label,
			Volume: request.Volume,
		},
	}
	ponds, err := f.pondUC.GetList(c.Request().Context(), filter)
	if err != nil {
		return handlerErrorResponse(err)
	}
	return c.JSON(http.StatusOK, ponds)
}

func (f *PondDelivery) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handlerErrorResponse(constant.ErrNotFound)
	}

	pond, err := f.pondUC.Get(
		c.Request().Context(),
		entity.PondFilter{Pond: entity.Pond{ID: id}},
	)
	if err != nil {
		return handlerErrorResponse(err)
	}
	return c.JSON(http.StatusOK, pond)
}

func (f *PondDelivery) Create(c echo.Context) error {
	request := entity.PondFormRequest{}
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Validation
	if request.FarmID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, constant.ErrIncorrectFarmID.Error())
	}

	pond := entity.Pond{
		FarmID: request.FarmID,
		Label:  request.Label,
		Volume: request.Volume,
	}
	id, err := f.pondUC.Store(c.Request().Context(), pond)
	if err != nil {
		return handlerErrorResponse(err)
	}

	return c.JSON(http.StatusOK, entity.PondFormResponse{ID: id})
}

func (f *PondDelivery) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handlerErrorResponse(constant.ErrNotFound)
	}

	request := entity.PondFormRequest{}
	err = c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	pond := entity.Pond{
		ID:     id,
		Label:  request.Label,
		Volume: request.Volume,
	}
	id, err = f.pondUC.Store(c.Request().Context(), pond)
	if err != nil {
		return handlerErrorResponse(err)
	}

	return c.JSON(http.StatusOK, entity.PondFormResponse{ID: id})
}

func (f *PondDelivery) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return handlerErrorResponse(constant.ErrNotFound)
	}

	err = f.pondUC.Delete(c.Request().Context(), id)
	if err != nil {
		return handlerErrorResponse(err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}
