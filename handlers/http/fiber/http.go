package fiber

import (
	"app/models"
	"app/services/address"
	"app/services/user"
	"context"

	"github.com/gofiber/fiber/v3"
)

type fiberCTX string

func StartHTTPFiber(userSVC user.UserSVC, addresSVC address.AddressSVC) {
	app := fiber.New()
	v1 := app.Group("/v1")

	v1.Post("/user", func(ctx fiber.Ctx) (err error) {
		payload := models.UserReq{}
		err = ctx.Bind().Body(&payload)
		if err != nil {
			return
		}
		out, err := userSVC.Create(context.WithValue(ctx.Context(), fiberCTX("--"), ""), payload)
		if err != nil {
			errOut := models.UnWrapperError(err)
			return ctx.Status(errOut.HTTPCode).JSON(errOut)
		}
		newOut := models.UserResp{
			ID:       out.ID,
			Name:     out.Name,
			LastName: out.LastName,
			Age:      out.Age,
			Address:  nil,
		}
		return ctx.JSON(newOut, "application/json")
	})

	v1.Get("/user/:id", func(ctx fiber.Ctx) (err error) {

		out, err := userSVC.Get(context.WithValue(ctx.Context(), fiberCTX("--"), ""), ctx.Params("id"))
		if err != nil {
			errOut := models.UnWrapperError(err)
			return ctx.Status(errOut.HTTPCode).JSON(errOut)
		}
		newOut := models.UserResp{
			ID:       out.ID,
			Name:     out.Name,
			LastName: out.LastName,
			Age:      out.Age,
			Address:  nil,
		}

		address, err := addresSVC.Get(context.WithValue(ctx.Context(), fiberCTX("--"), ""), ctx.Params("id"), "")
		if err != nil {
			errOut := models.UnWrapperError(err)
			return ctx.Status(errOut.HTTPCode).JSON(errOut)
		}
		newOut.Address = make([]models.AddressResp, len(address))
		for i := 0; i < len(address); i++ {
			newOut.Address[i] = models.AddressResp(address[i])
		}

		return ctx.JSON(newOut, "application/json")
	})

	v1.Delete("/user/:id", func(ctx fiber.Ctx) (err error) {

		err = userSVC.Delete(context.WithValue(ctx.Context(), fiberCTX("--"), ""), ctx.Params("id"))
		if err != nil {
			errOut := models.UnWrapperError(err)
			return ctx.Status(errOut.HTTPCode).JSON(errOut)
		}
		return nil // 200 ok in fiber
	})

	v1.Patch("/user/:id", func(ctx fiber.Ctx) (err error) {
		payload := models.UserReq{}
		err = ctx.Bind().Body(&payload)
		if err != nil {
			return
		}
		err = userSVC.Patch(context.WithValue(ctx.Context(), fiberCTX("--"), ""), payload, ctx.Params("id"))
		if err != nil {
			errOut := models.UnWrapperError(err)
			return ctx.Status(errOut.HTTPCode).JSON(errOut)
		}
		return // 200 ok in fiber
	})

	v1.Post("/user/:id/address", func(ctx fiber.Ctx) (err error) {
		payload := models.AddressReq{}
		err = ctx.Bind().Body(&payload)
		if err != nil {
			return
		}
		payload.UserID = ctx.Params("id")
		out, err := addresSVC.Create(context.WithValue(ctx.Context(), fiberCTX("--"), ""), payload)
		if err != nil {
			errOut := models.UnWrapperError(err)
			return ctx.Status(errOut.HTTPCode).JSON(errOut)
		}
		return ctx.JSON(models.AddressResp(out), "application/json")
	})

	v1.Get("/user/:id/address?address_id", func(ctx fiber.Ctx) (err error) {

		out, err := addresSVC.Get(context.WithValue(ctx.Context(), fiberCTX("--"), ""), ctx.Params("id"), ctx.Query("address_id"))
		if err != nil {
			errOut := models.UnWrapperError(err)
			return ctx.Status(errOut.HTTPCode).JSON(errOut)
		}
		newOut := make([]models.AddressResp, len(out))
		for i := 0; i < len(out); i++ {
			newOut[i] = models.AddressResp(out[i])
		}

		return ctx.JSON(newOut, "application/json")
	})

	v1.Delete("/user/:id/address/:address_id", func(ctx fiber.Ctx) (err error) {

		err = addresSVC.Delete(context.WithValue(ctx.Context(), fiberCTX("--"), ""), ctx.Params("id"), ctx.Params("address_id"))
		if err != nil {
			errOut := models.UnWrapperError(err)
			return ctx.Status(errOut.HTTPCode).JSON(errOut)
		}
		return nil
	})

	v1.Patch("/user/:id/address/:address_id", func(ctx fiber.Ctx) (err error) {
		payload := models.AddressReq{}
		err = ctx.Bind().Body(&payload)
		if err != nil {
			return
		}
		err = addresSVC.Patch(context.WithValue(ctx.Context(), fiberCTX("--"), ""), ctx.Params("id"), ctx.Params("address_id"), payload)
		if err != nil {
			errOut := models.UnWrapperError(err)
			return ctx.Status(errOut.HTTPCode).JSON(errOut)
		}
		return nil
	})
}
