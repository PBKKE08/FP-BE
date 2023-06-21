package handler

import (
	"github.com/PBKKE08/FP-BE/api/command/buat_booking"
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

	apiGroup.POST("/book", h.BookPartner)
}

func (h *BookingHandler) BookPartner(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*customClaims)

	type TD struct {
		PenggunaID  string `json:"-"`
		PartnerID   string `json:"partner_id"`
		BookingDate string `json:"booking_date"`
		TimeStart   string `json:"time_start"`
		TimeEnd     string `json:"time_end"`
		PaymentType string `json:"payment_type"`
		Msg         string `json:"message"`
	}

	var td TD

	if err := c.Bind(&td); err != nil {
		return c.JSON(422, Response(422, err.Error()))
	}

	parsedDate, err := time.Parse("2006-01-02", td.BookingDate)
	if err != nil {
		return c.JSON(422, Response(422, err.Error()))
	}

	req := buat_booking.Request{
		PenggunaID:  "",
		PartnerID:   td.PartnerID,
		BookingDate: parsedDate,
		TimeStart:   td.TimeStart,
		TimeEnd:     td.TimeEnd,
		PaymentType: td.PaymentType,
		Msg:         td.Msg,
	}

	err = h.bookingUsecase.CreateBooking(c.Request().Context(), claims.ID, req)
	if err != nil {
		return c.JSON(400, Response(400, err.Error()))
	}

	return nil
}
