package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func main() {

	redisdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		result, err := redisdb.Incr("counter").Result()
		if err != nil {
			panic(err)
		}
		return c.SendString(fmt.Sprintf("Hello, visitor number %d!", result))
	})

	log.Fatal(app.Listen(os.Getenv("SERVICE_ADDR")))

}
