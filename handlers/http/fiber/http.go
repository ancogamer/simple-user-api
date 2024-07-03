package fiber

import (
	"app/models"
	"app/services/address"
	"app/services/user"
	"context"

	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type fiberCTX string

func StartHTTPFiber(userSVC user.UserSVC, addresSVC address.AddressSVC) {
	app := fiber.New()
	v1 := app.Group("/v1")
	v1.Post("/account", func(c fiber.Ctx) (err error) {

		// chama o serviço que cria uma account, nele, gera um int aleatorio com rand.Int(rand.Reader, big.NewInt(90)) da lib crypto/rand,
		// para servir de salt junto de um hash256 da senha, salva o mesmo junto na account, no caso usuário o serviço também transforma a string de "user" para versão
		// miniscula e salva no banco (aonde a mesma é unica).

		// Create the Claims
		claims := jwt.MapClaims{
			"name":       "John Doe",
			"admin":      true,
			"exp":        time.Now().Add(time.Hour * 72).Unix(),
			"account_id": "",
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.SendStatus(500)
		}

		return c.Status(200).JSON(fiber.Map{"token": t})
	})
	v1.Post("/login", login)

	// JWT Middleware
	// daqui para baixo tudo roda protegida
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	// TODO novo middleware puxa a account do token e checa se ela existe em um 2 middleware e quais as permissões dela

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

		return ctx.Status(200).JSON(newOut)

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

func login(c fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")
	// TODO serviço de login, aonde a senha é comparada junto do salt que está salvo no banco para o user e puxa o id da account para ser colocado junto do token
	if user != "" || pass != "" { // parte de autenticação
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
		// todo add account ID
	}
	//

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
