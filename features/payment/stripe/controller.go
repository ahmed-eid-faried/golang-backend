package stripe

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/refund"

	_ "main/docs" // This is required for Swagger to find your documentation
)

// @title Stripe Payment API
// @version 1.0
// @description API for processing payments with Stripe
// @host localhost:8080
// @BasePath /
func Init() {
	stripe.Key = "pk_test_51Ngm3RGhokhcgA0sRWmKDxJDSld4r4je29GB4v1RGKLur8lJFrcLDql0Ahq1glDykEnShRyfvK9Cosi6GselKd5l00eXGHv5M9"
}

// Authentication middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken != "sk_test_51Ngm3RGhokhcgA0sLs7kgb0RX34VmN8tk8mJCq6oliiMX3Sxng0M4hCemz3Bikbd7K76Palkb9bbFndeTwWVE3lm00zHVsKwg0" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Customer routes
// @Summary Create a new customer
// @Description Create a new customer in Stripe
// @Accept  json
// @Produce  json
// @Success 200 {object} string "customer_id"
// @Failure 500 {object} string "error"
// @Router /stripe/create_customer [post]
// @Tags Stripe Payment
func CreateCustomer(c *gin.Context) {
	params := &stripe.CustomerParams{}
	cust, err := customer.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"customer_id": cust.ID})
}

// @Summary Retrieve customer
// @Description Retrieve a customer from Stripe
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID"
// @Success 200 {object} object
// @Failure 500 {object} string "error"
// @Router /stripe/retrieve_customer/{id} [get]
// @Tags Stripe Payment
func RetrieveCustomer(c *gin.Context) {
	id := c.Param("id")
	cust, err := customer.Get(id, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cust)
}

// Payment routes
// @Summary Create payment intent
// @Description Create a payment intent in Stripe
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Failure 500 {object} string "error"
// @Router /stripe/create_payment_intent [post]
// @Tags Stripe Payment
func CreatePaymentIntent(c *gin.Context) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	intent, err := paymentintent.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, intent)
}

// @Summary Capture payment
// @Description Capture a payment in Stripe
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Failure 500 {object} string "error"
// @Router /stripe/capture_payment [post]
// @Tags Stripe Payment
func CapturePayment(c *gin.Context) {
	// You need to retrieve payment intent ID from request payload
	// Assuming it's provided in JSON format
	var req struct {
		PaymentIntentID string `json:"payment_intent_id"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	intent, err := paymentintent.Capture(req.PaymentIntentID, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, intent)
}

// @Summary Refund payment
// @Description Refund a payment in Stripe
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Failure 500 {object} string "error"
// @Router /stripe/refund_payment [post]
// @Tags Stripe Payment
func RefundPayment(c *gin.Context) {
	// You need to retrieve payment intent ID from request payload
	// Assuming it's provided in JSON format
	var req struct {
		PaymentIntentID string `json:"payment_intent_id"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	refundParams := &stripe.RefundParams{
		PaymentIntent: stripe.String(req.PaymentIntentID),
	}
	ref, err := refund.New(refundParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ref)
}

// r := setupRouter()
// // Swagger documentation routes
// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
// r.Run(":8080")

// func setupRouter() *gin.Engine {
// 	r := gin.Default()
// 	// payment:=
// 	// Authentication middleware
// 	r.Use(authMiddleware())
// 	r.POST("/create_customer", createCustomer)
// 	r.GET("/retrieve_customer/:id", retrieveCustomer)
// 	r.POST("/create_payment_intent", createPaymentIntent)
// 	r.POST("/capture_payment", capturePayment)
// 	r.POST("/refund_payment", refundPayment)
// 	return r
// }
