package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/abilsabili50/middleware-with-go-fiber/app/dto"
	"github.com/abilsabili50/middleware-with-go-fiber/cmd/factory"
	"github.com/abilsabili50/middleware-with-go-fiber/pkg/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
	config     *config.Config
	repository *factory.Repository
}

func NewMiddleware(config *config.Config, repository *factory.Repository) *Middleware {
	return &Middleware{
		config:     config,
		repository: repository,
	}
}

// JWT Protected func for specify protected route with JWT authentication
func (m *Middleware) JWTProtected() fiber.Handler {
	// create config for JWT authentication middleware
	jwtWareConfig := jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(m.config.App.JWTAccTokenSecretKey)},
		KeyFunc:        m.jwtValidateToken,
		SuccessHandler: m.verifyToken,
		ErrorHandler:   m.jwtError,
	}

	return jwtware.New(jwtWareConfig)
}

func (m *Middleware) IsRegisteredUser(c *fiber.Ctx) error {
	// extract userId
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId, ok := claims["id"].(string)
	if !ok {
		log.Printf("[ERROR] - error occured while extract userId")
		res := dto.Response[any]{
			Status:  "error",
			Message: "forbidden",
			Code:    http.StatusForbidden,
			Data:    nil,
		}

		return c.Status(fiber.StatusForbidden).JSON(res)
	}

	// check user in DB
	_, err := m.repository.User.FindById(userId)
	if err != nil {
		log.Printf("[ERROR] - %v", err)
		res := dto.Response[any]{
			Status:  "error",
			Message: "Unauthorized",
			Code:    http.StatusUnauthorized,
			Data:    nil,
		}

		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	// Store userId in context for later use
	c.Locals("userId", userId)

	return c.Next()
}

func (m *Middleware) jwtValidateToken(t *jwt.Token) (interface{}, error) {
	// validate algorithm
	if t.Method.Alg() != jwtware.HS256 {
		log.Printf("[ERROR] - wrong algorithm")
		return nil, errors.New("wrong algorithm used")
	}

	return []byte(m.config.App.JWTAccTokenSecretKey), nil
}

func (m *Middleware) verifyToken(c *fiber.Ctx) error {
	// extract token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	expires := int64(claims["exp"].(float64))

	// check token expires
	if time.Now().Unix() > expires {
		log.Printf("[ERROR] - expired token")
		return m.jwtError(c, errors.New("token expired"))
	}

	return c.Next()
}

func (m *Middleware) jwtError(c *fiber.Ctx, err error) error {
	// declare status and message error
	status := fiber.StatusUnauthorized
	message := "Unauthorized"

	// check if err != nil
	if err != nil && err.Error() == "Missing or malformed JWT" {
		status = fiber.StatusBadRequest
		message = err.Error()
	}

	// declare response
	res := dto.Response[any]{
		Status:  "error",
		Message: message,
		Code:    status,
		Data:    nil,
	}

	log.Printf("[ERROR] - %v", message)
	return c.Status(status).JSON(res)
}
