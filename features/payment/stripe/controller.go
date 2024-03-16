package stripe

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/refund"

	_ "main/docs" // This is required for Swagger to find your documentation
)

// @title Stripe Payment API
// @version 1.0
// @description API for processing payments with Stripe
// @host localhost:8080
// @BasePath /
func Init() {
	stripe.Key = "sk_test_51Ngm3RGhokhcgA0sLs7kgb0RX34VmN8tk8mJCq6oliiMX3Sxng0M4hCemz3Bikbd7K76Palkb9bbFndeTwWVE3lm00zHVsKwg0"
	// stripe.Key = "pk_test_51Ngm3RGhokhcgA0sRWmKDxJDSld4r4je29GB4v1RGKLur8lJFrcLDql0Ahq1glDykEnShRyfvK9Cosi6GselKd5l00eXGHv5M9"

}

// Authentication middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken != "pk_test_51Ngm3RGhokhcgA0sRWmKDxJDSld4r4je29GB4v1RGKLur8lJFrcLDql0Ahq1glDykEnShRyfvK9Cosi6GselKd5l00eXGHv5M9" {
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

// // TestStripeCharge simulates creating a charge using Stripe API
// // @Summary Test Stripe Charge
// // @Description Simulates creating a charge using Stripe API
// // @Accept json
// // @Produce json
// // @Success 200 {object} string "Charge created successfully"
// // @Failure 500 {object} string "Error creating charge"
// // @Router /api/v1/stripe/test_charge [post]
// // @Tags Stripe Payment
// func TestStripeCharge(c *gin.Context) {
// 	// Create a new charge
// 	params := &stripe.ChargeParams{
// 		Amount:   stripe.Int64(2000),
// 		Currency: stripe.String(string(stripe.CurrencyUSD)),
// 		Desc:     stripe.String("Test Charge"),
// 	}

// 	sc := &client.API{}
// 	sc.Init("sk_test_4eC39HqLyjWDarjtT1zdp7dc", nil)
// 	sc.Charges.Get("ch_3Ln3j02eZvKYlo2C0d5IZWuG", params)
// 	params.SetSource("tok_visa") // use a test card token provided by Stripe
// 	ch, err := charge.New(params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check for errors
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating charge"})
// 		return
// 	}

// 	// Print charge details
// 	fmt.Printf("Charge ID: %s\n", ch.ID)
// 	fmt.Printf("Amount: %d\n", ch.Amount)
// 	fmt.Printf("Description: %s\n", ch.Description)

// 	c.JSON(http.StatusOK, gin.H{"message": "Charge created successfully"})
// }

// stripe_custom.go

// SetAPIKey sets the Stripe API key for authentication
func SetAPIKey(secretKey string) {
	stripe.Key = secretKey
	// stripe.Key = publishableKey
}

// CreateCustomer creates a new customer with the given parameters
func CreateCustomerFun(email, accountID string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	if accountID != "" {
		params.SetStripeAccount(accountID)
	}
	c, err := customer.New(params)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// GetCustomer retrieves a customer by ID
func GetCustomer(customerID string) (*stripe.Customer, error) {
	c, err := customer.Get(customerID, nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// func main() {
// 	// apiKey := "sk_test..."
// 	apiKey := "sk_test_51Ngm3RGhokhcgA0sLs7kgb0RX34VmN8tk8mJCq6oliiMX3Sxng0M4hCemz3Bikbd7K76Palkb9bbFndeTwWVE3lm00zHVsKwg0"
// 	SetAPIKey(apiKey)

// 	// Create a customer
// 	customer, err := CreateCustomerFun("jenny.rosen@example.com", "")
// 	if err != nil {
// 		fmt.Println("Error creating customer:", err)
// 		return
// 	}
// 	fmt.Println("Customer ID:", customer.ID)

// 	// Get a customer
// 	retrievedCustomer, err := GetCustomer(customer.ID)
// 	if err != nil {
// 		fmt.Println("Error retrieving customer:", err)
// 		return
// 	}
// 	fmt.Println("Retrieved Customer ID:", retrievedCustomer.ID)
// }

func StripeCustom() {
	publishableKey := "pk_test_51Ngm3RGhokhcgA0sRWmKDxJDSld4r4je29GB4v1RGKLur8lJFrcLDql0Ahq1glDykEnShRyfvK9Cosi6GselKd5l00eXGHv5M9"
	// secretKey := "sk_test_51Ngm3RGhokhcgA0sLs7kgb0RX34VmN8tk8mJCq6oliiMX3Sxng0M4hCemz3Bikbd7K76Palkb9bbFndeTwWVE3lm00zHVsKwg0"

	SetAPIKey(publishableKey)

	// Create a customer
	customer, err := CreateCustomerFun("jenny.rosen@example.com", "")
	if err != nil {
		fmt.Println("Error creating customer:", err)
		return
	}
	fmt.Println("Customer ID:", customer.ID)

	// Get a customer
	retrievedCustomer, err := GetCustomer(customer.ID)
	if err != nil {
		fmt.Println("Error retrieving customer:", err)
		return
	}
	fmt.Println("Retrieved Customer ID:", retrievedCustomer.ID)
}

// @Summary Create a new customer
// @Description Create a new customer in Stripe
// @Accept  json
// @Produce  json
// @Param email formData string true "Email address of the customer"
// @Param account_id formData string false "Account ID (optional)"
// @Success 200 {object} string "customer_id"
// @Failure 500 {object} string "error"
// @Router /create_customer [post]
func CreateCustomerHandler(c *gin.Context) {
	email := c.PostForm("email")
	accountID := c.PostForm("account_id") // If you have account-based authentication

	customer, err := CreateCustomerFun(email, accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"customer_id": customer.ID})
}

// @Summary Retrieve a customer
// @Description Retrieve a customer from Stripe
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID"
// @Success 200 {object} object
// @Failure 500 {object} string "error"
// @Router /get_customer/{id} [get]
func GetCustomerHandler(c *gin.Context) {
	customerID := c.Param("id")

	customer, err := GetCustomer(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// Define Req struct
type Req struct {
	CustomerID string `json:"customer_id"`
	Amount     int64  `json:"amount"`
	Currency   string `json:"currency"`
}

// @Summary Retrieve a customer
// @Description Retrieve a customer from Stripe
// @Accept  json
// @Produce  json
// @Param id path string true "Customer ID"
// @Param Req body Req true "Req object"
// @Success 200 {object} object
// @Failure 500 {object} string "error"
// @Router /charge_customer/{id} [get]
func ChargeCustomer(c *gin.Context) {
	// Retrieve customer ID from Request body
	var req Req
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request payload"})
		return
	}

	// Create a charge
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(req.Currency),
		Customer: stripe.String(req.CustomerID),
	}

	ch, err := charge.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"charge_id": ch.ID})
}
