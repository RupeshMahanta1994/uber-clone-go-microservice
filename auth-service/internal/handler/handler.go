package handler

import (
	"context"
	"errors"
	"log"

	"github.com/rupeshmahanta/auth-service/internal/service"
)

// RegisterRequest represents the register request payload
type RegisterRequest struct {
	Email    string
	Password string
	Role     string
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string
	Password string
}

// AuthResponse represents the response with token
type AuthResponse struct {
	Token string
	Error string
}

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService service.IAuthService // Use interface for flexibility
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService service.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(ctx context.Context, req *RegisterRequest) (*AuthResponse, error) {
	// Validate input
	if err := validateRegisterRequest(req); err != nil {
		return &AuthResponse{Error: err.Error()}, nil
	}

	// Call service
	token, err := h.authService.Register(ctx, req.Email, req.Password, req.Role)
	if err != nil {
		log.Printf("Register error: %v", err)
		return &AuthResponse{Error: "Failed to register user"}, nil
	}

	return &AuthResponse{Token: token}, nil
}

// Login handles user login
func (h *AuthHandler) Login(ctx context.Context, req *LoginRequest) (*AuthResponse, error) {
	// Validate input
	if err := validateLoginRequest(req); err != nil {
		return &AuthResponse{Error: err.Error()}, nil
	}

	// Call service
	token, err := h.authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		log.Printf("Login error: %v", err)
		return &AuthResponse{Error: "Invalid credentials"}, nil
	}

	return &AuthResponse{Token: token}, nil
}

// Validation functions
func validateRegisterRequest(req *RegisterRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}
	if req.Password == "" {
		return errors.New("password is required")
	}
	if req.Role == "" {
		return errors.New("role is required")
	}
	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}

func validateLoginRequest(req *LoginRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}
	if req.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
