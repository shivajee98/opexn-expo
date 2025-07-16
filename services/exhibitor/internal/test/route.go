package routes

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/shivajee98/opexn-exhibitors/internal/handlers"
// 	"github.com/shivajee98/opexn-exhibitors/internal/model"
// 	"github.com/shivajee98/opexn-exhibitors/internal/repository"
// 	"github.com/shivajee98/opexn-exhibitors/internal/services"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// // setupTestApp initializes a Fiber app with dependencies for testing
// func setupTestApp(t *testing.T) (*fiber.App, *gorm.DB) {
// 	// Initialize in-memory SQLite database
// 	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("Failed to connect to test database: %v", err)
// 	}

// 	// Run migrations
// 	err = migrate.Migrate(db)
// 	if err != nil {
// 		t.Fatalf("Failed to run migrations: %v", err)
// 	}

// 	// Initialize dependencies
// 	startupRepo := repository.InitStartupRepository(db)
// 	startupService := services.InitStartupService(startupRepo)
// 	startupHandler := handlers.InitStartupHandler(startupService)

// 	// Create Fiber app
// 	app := fiber.New()

// 	// Setup routes
// 	SetupStartupRoutes(app, startupHandler)

// 	return app, db
// }

// // TestStartupRoutes tests all startup endpoints
// func TestStartupRoutes(t *testing.T) {
// 	app, db := setupTestApp(t)

// 	// Seed test data
// 	startup := model.Startup{
// 		Name:            "Test Startup",
// 		DPIITCertNumber: "DPIIT123",
// 		PitchDeckURL:    "https://cloudinary.com/test.pdf",
// 		Address:         model.Address{Street: "123 Test St", City: "Test City", State: "Test State", Pincode: "123456"},
// 		Product:         model.Product{Description: "Test Product", Problem: "Test Problem", Stage: "MVP"},
// 		RevenueInfo:     model.RevenueInfo{RevenueBracket: "₹0–₹5L", UserImpact: 100},
// 		FundingInfo:     model.FundingInfo{Type: "Angel"},
// 		EventIntent:     model.EventIntent{WhyParticipate: "Networking", Expectation: "Growth", ConsentToPay: true},
// 		SPOC:            model.SPOC{Name: "John Doe", Email: "john@example.com", Phone: "1234567890", Position: "Manager"},
// 		Director:        model.Director{Name: "Jane Doe", Email: "jane@example.com"},
// 	}
// 	if err := db.Create(&startup).Error; err != nil {
// 		t.Fatalf("Failed to seed test data: %v", err)
// 	}

// 	// Define test cases
// 	tests := []struct {
// 		name           string
// 		method         string
// 		path           string
// 		body           interface{}
// 		headers        map[string]string
// 		expectedStatus int
// 		expectedBody   string
// 		setupContext   func(c *fiber.Ctx)
// 	}{
// 		{
// 			name:           "Check DPIITCertNumber exists",
// 			method:         "GET",
// 			path:           "/api/startup/check/DPIIT123",
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   `{"exists":true}`,
// 		},
// 		{
// 			name:           "Check DPIITCertNumber does not exist",
// 			method:         "GET",
// 			path:           "/api/startup/check/DPIIT999",
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   `{"exists":false}`,
// 		},
// 		{
// 			name:           "Get Startup by ID",
// 			method:         "GET",
// 			path:           "/api/startup/1",
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   `"id":1,"name":"Test Startup","dpiitCertNumber":"DPIIT123"`,
// 		},
// 		{
// 			name:           "Get Startup by ID - Not Found",
// 			method:         "GET",
// 			path:           "/api/startup/999",
// 			expectedStatus: http.StatusNotFound,
// 			expectedBody:   `"Startup not found"`,
// 		},
// 		{
// 			name:           "Get All Startups",
// 			method:         "GET",
// 			path:           "/api/startup",
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   `[{"id":1,"name":"Test Startup","dpiitCertNumber":"DPIIT123"`,
// 		},
// 		{
// 			name:   "Register Startup - Success",
// 			method: "POST",
// 			path:   "/api/startup/register",
// 			body: model.Startup{
// 				Name:            "New Startup",
// 				DPIITCertNumber: "DPIIT456",
// 				PitchDeckURL:    "https://cloudinary.com/new.pdf",
// 				Address:         model.Address{Street: "456 New St", City: "New City", State: "New State", Pincode: "654321"},
// 				Product:         model.Product{Description: "New Product", Problem: "New Problem", Stage: "Prototype"},
// 				RevenueInfo:     model.RevenueInfo{RevenueBracket: "₹5–₹25L", UserImpact: 200},
// 				FundingInfo:     model.FundingInfo{Type: "VC"},
// 				EventIntent:     model.EventIntent{WhyParticipate: "Exposure", Expectation: "Funding", ConsentToPay: false},
// 				SPOC:            model.SPOC{Name: "Alice Smith", Email: "alice@example.com", Phone: "0987654321", Position: "CEO"},
// 				Director:        model.Director{Name: "Bob Johnson", Email: "bob@example.com"},
// 			},
// 			headers:        map[string]string{"Content-Type": "application/json"},
// 			setupContext:   func(c *fiber.Ctx) { c.Locals("clerk_id", "test_clerk_id") },
// 			expectedStatus: http.StatusCreated,
// 			expectedBody:   `"message":"Startup registered successfully","startup":{"id":2,"name":"New Startup","dpiitCertNumber":"DPIIT456"}`,
// 		},
// 		{
// 			name:           "Register Startup - Missing Clerk ID",
// 			method:         "POST",
// 			path:           "/api/startup/register",
// 			body:           model.Startup{Name: "Invalid Startup", DPIITCertNumber: "DPIIT789", PitchDeckURL: "https://cloudinary.com/invalid.pdf"},
// 			headers:        map[string]string{"Content-Type": "application/json"},
// 			expectedStatus: http.StatusUnauthorized,
// 			expectedBody:   `"Unauthorized"`,
// 		},
// 		{
// 			name:           "Register Startup - Duplicate DPIITCertNumber",
// 			method:         "POST",
// 			path:           "/api/startup/register",
// 			body:           model.Startup{Name: "Duplicate Startup", DPIITCertNumber: "DPIIT123", PitchDeckURL: "https://cloudinary.com/duplicate.pdf"},
// 			headers:        map[string]string{"Content-Type": "application/json"},
// 			setupContext:   func(c *fiber.Ctx) { c.Locals("clerk_id", "test_clerk_id") },
// 			expectedStatus: http.StatusConflict,
// 			expectedBody:   `"Startup with this DPIIT certificate number already exists"`,
// 		},
// 		{
// 			name:   "Update Startup - Success",
// 			method: "PUT",
// 			path:   "/api/startup/update",
// 			body: model.Startup{
// 				ID:              1,
// 				Name:            "Updated Startup",
// 				DPIITCertNumber: "DPIIT123",
// 				PitchDeckURL:    "https://cloudinary.com/updated.pdf",
// 				Address:         model.Address{Street: "789 Update St", City: "Update City", State: "Update State", Pincode: "987654"},
// 			},
// 			headers:        map[string]string{"Content-Type": "application/json"},
// 			setupContext:   func(c *fiber.Ctx) { c.Locals("clerk_id", "test_clerk_id") },
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   `"message":"Startup updated successfully","startup":{"id":1,"name":"Updated Startup","dpiitCertNumber":"DPIIT123"}`,
// 		},
// 		{
// 			name:           "Update Startup - Not Found",
// 			method:         "PUT",
// 			path:           "/api/startup/update",
// 			body:           model.Startup{ID: 999, Name: "Nonexistent Startup", DPIITCertNumber: "DPIIT999", PitchDeckURL: "https://cloudinary.com/nonexistent.pdf"},
// 			headers:        map[string]string{"Content-Type": "application/json"},
// 			setupContext:   func(c *fiber.Ctx) { c.Locals("clerk_id", "test_clerk_id") },
// 			expectedStatus: http.StatusNotFound,
// 			expectedBody:   `"Startup not found"`,
// 		},
// 		{
// 			name:           "Delete Startup - Success",
// 			method:         "DELETE",
// 			path:           "/api/startup/1",
// 			setupContext:   func(c *fiber.Ctx) { c.Locals("clerk_id", "test_clerk_id") },
// 			expectedStatus: http.StatusOK,
// 			expectedBody:   `"message":"Startup deleted successfully"`,
// 		},
// 		{
// 			name:           "Delete Startup - Not Found",
// 			method:         "DELETE",
// 			path:           "/api/startup/999",
// 			setupContext:   func(c *fiber.Ctx) { c.Locals("clerk_id", "test_clerk_id") },
// 			expectedStatus: http.StatusNotFound,
// 			expectedBody:   `"Startup not found"`,
// 		},
// 	}

// 	// Run tests
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create request
// 			var body []byte
// 			if tt.body != nil {
// 				body, _ = json.Marshal(tt.body)
// 			}
// 			req := httptest.NewRequest(tt.method, tt.path, bytes.NewReader(body))
// 			for k, v := range tt.headers {
// 				req.Header.Set(k, v)
// 			}

// 			// Setup Fiber context for middleware simulation
// 			app.TestMiddleware(func(c *fiber.Ctx) error {
// 				if tt.setupContext != nil {
// 					tt.setupContext(c)
// 				}
// 				return c.Next()
// 			})

// 			// Execute request
// 			resp, err := app.Test(req, -1)
// 			if err != nil {
// 				t.Fatalf("Failed to execute request: %v", err)
// 			}

// 			// Check status code
// 			if resp.StatusCode != tt.expectedStatus {
// 				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
// 			}

// 			// Check response body
// 			bodyBytes := new(bytes.Buffer)
// 			_, err = bodyBytes.ReadFrom(resp.Body)
// 			if err != nil {
// 				t.Fatalf("Failed to read response body: %v", err)
// 			}
// 			bodyStr := bodyBytes.String()
// 			if !strings.Contains(bodyStr, tt.expectedBody) {
// 				t.Errorf("Expected body to contain %q, got %q", tt.expectedBody, bodyStr)
// 			}
// 		})
// 	}
// }
