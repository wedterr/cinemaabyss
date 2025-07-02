package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wedterr/cinemaabyss/models"
)

// CreateMovieEvent - Создание события фильма
func (c *Container) CreateMovieEvent(ctx echo.Context) error {
	req := NewMovieEventMessage()
	event := new(MovieEventSchema)
	if err := ctx.Bind(event); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	req.Payload = *event
	msg, err := req.toBrokerMessage()
	if err != nil {
		log.Println(err)
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = c.KafkaCtrl.broker.Publish(context.Background(), MovieEventChannelPath, msg)
	if err != nil {
		log.Println(err)
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	// log.Println(req)
	// err := c.KafkaCtrl.SendAsMovieEventOperation(context.Background(), req)
	// if err != nil {
	// 	log.Println(err)
	// 	return ctx.String(http.StatusBadRequest, err.Error())
	// }

	return ctx.JSON(http.StatusCreated, models.EventResponse{
		Status: "success",
	})
}

// CreatePaymentEvent - Создание события платежа
func (c *Container) CreatePaymentEvent(ctx echo.Context) error {
	req := NewPaymentEventMessage()
	event := new(PaymentEventSchema)
	if err := ctx.Bind(event); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	req.Payload = *event
	msg, err := req.toBrokerMessage()
	if err != nil {
		log.Println(err)
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = c.KafkaCtrl.broker.Publish(context.Background(), PaymentsChannelPath, msg)
	if err != nil {
		log.Println(err)
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	// err := c.KafkaCtrl.SendAsPaymentEventOperation(context.Background(), req)
	// if err != nil {
	// 	log.Println(err)
	// 	return ctx.String(http.StatusBadRequest, err.Error())
	// }

	return ctx.JSON(http.StatusCreated, models.EventResponse{
		Status: "success",
	})
}

// CreateUserEvent - Создание события пользователя
func (c *Container) CreateUserEvent(ctx echo.Context) error {
	req := NewUserEventMessage()
	event := new(UserEventSchema)
	if err := ctx.Bind(event); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	req.Payload = *event
	msg, err := req.toBrokerMessage()
	if err != nil {
		log.Println(err)
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = c.KafkaCtrl.broker.Publish(context.Background(), UsersChannelPath, msg)
	if err != nil {
		log.Println(err)
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	// err := c.KafkaCtrl.SendAsUserEventOperation(context.Background(), req)
	// if err != nil {
	// 	log.Println(err)
	// 	return ctx.String(http.StatusBadRequest, err.Error())
	// }

	return ctx.JSON(http.StatusCreated, models.EventResponse{
		Status: "success",
	})
}
