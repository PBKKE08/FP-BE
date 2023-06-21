package handler

import (
	"github.com/PBKKE08/FP-BE/api/usecase"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"time"
)

type BookingHandler struct {
	bookingUsecase *usecase.BookingUsecase
}

func NewBookingHandler(bookingUsecase *usecase.BookingUsecase) *BookingHandler {
	return &BookingHandler{bookingUsecase: bookingUsecase}
}

func (h *BookingHandler) Load(e *echo.Echo) {
	apiGroup := e.Group("/booking")
	apiGroup.Use(echojwt.WithConfig(jwtConfig))

	apiGroup.POST("/", h.BookPartner)
}

func (h *BookingHandler) BookPartner(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*customClaims)

	type payload struct {
		PartnerID   string    `json:"partner_id"`
		BookingDate time.Time `json:"booking_date"`
		TimeStart   string    `json:"time_start"`
		TimeEnd     string    `json:"time_end"`
	}

	var td payload

	if err := c.Bind(&td); err != nil {
		return c.JSON(422, Response(422, err.Error()))
	}

	err := h.bookingUsecase.CreateBooking(c.Request().Context(), claims.ID, td.PartnerID, td.BookingDate, td.TimeStart, td.TimeEnd)
	if err != nil {
		return c.JSON(400, Response(400, err.Error()))
	}

	return nil
}
