package coupon

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/lucas-deschamps/marketing-bot/internal/pkg/redis"

	. "github.com/lucas-deschamps/marketing-bot/internal/config"
)

func (s *couponService) Track(c *fiber.Ctx) error {
	ctx := context.Background()

	var redeemedMsg string
	var status bool = false

	couponCode := strings.ToUpper(CouponID)

	rdb := redis.SetupRedisClient()

	// Increment the counter for the session ID
	if err := rdb.Incr(ctx, fmt.Sprintf("counter:%s", SessionID)).Err(); err != nil {
		fmt.Println("Error incrementing counter:", err)
		return err
	}

	// Increment the total coupon clicks counter
	if err := rdb.Incr(ctx, "coupon_clicks").Err(); err != nil {
		fmt.Println("Error incrementing total coupon clicks:", err)
		return err
	}

	fmt.Println("\nCounter incremented successfully for session ID:", SessionID)

	// Get the count for the session ID
	count, ctErr := rdb.Get(ctx, fmt.Sprintf("counter:%s", SessionID)).Int()

	if ctErr != nil {
		fmt.Println("Error getting count:", ctErr)
		return ctErr
	}

	fmt.Println("Count for session ID:", SessionID, "is", count)

	// Get total coupon clicks
	total, err := rdb.Get(ctx, "coupon_clicks").Int()

	if err != nil {
		fmt.Println("Error getting total coupon clicks:", err)
		return err
	}

	fmt.Println("Total coupon clicks is: ", total)

	if count > 1 {
		status = true
	}

	if status == true {
		redeemedMsg = fmt.Sprintf(
			"Coupon redeemed! Coupon code: #%s.\nSession ID: %s\nPreviously redeemed: Yes.\nTotal coupon clicks in application: %s\n",
			couponCode,
			SessionID,
			strconv.Itoa(total),
		)
	} else {
		redeemedMsg = fmt.Sprintf(
			"Coupon redeemed! Coupon code: #%s.\nSession ID: %s\nPreviously redeemed: No.\nTotal coupon clicks in application: %s\n",
			couponCode,
			SessionID,
			strconv.Itoa(total),
		)
	}

	return c.SendString(redeemedMsg)
}
