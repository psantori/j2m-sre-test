package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func greet(n int64) string {
	return fmt.Sprintf("Hello, visitor number %d!", n)
}

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
		return c.SendString(greet(result))
	})
	app.Listen(":8080")
	log.Fatal(app.Listen(os.Getenv("SERVICE_ADDR")))

}
