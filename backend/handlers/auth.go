package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tchskk/ecom-RTGP/backend/db"
	"github.com/tchskk/ecom-RTGP/backend/models"
	"github.com/tchskk/ecom-RTGP/backend/utils"
)

func Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", req.Email).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Insert user
	var user models.User
	err = db.DB.QueryRow(
		`INSERT INTO users (email, password, first_name, last_name, role) 
		 VALUES ($1, $2, $3, $4, $5) 
		 RETURNING id, email, first_name, last_name, role, created_at, updated_at`,
		req.Email, hashedPassword, req.FirstName, req.LastName, "user",
	).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, models.LoginResponse{
		Token: token,
		User:  user,
	})
}

func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user
	var user models.User
	var hashedPassword string
	err := db.DB.QueryRow(
		`SELECT id, email, password, first_name, last_name, role, created_at, updated_at 
		 FROM users WHERE email = $1`,
		req.Email,
	).Scan(&user.ID, &user.Email, &hashedPassword, &user.FirstName, &user.LastName, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, hashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		Token: token,
		User:  user,
	})
}

func GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user models.User
	err := db.DB.QueryRow(
		`SELECT id, email, first_name, last_name, role, created_at, updated_at 
		 FROM users WHERE id = $1`,
		userID,
	).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	rows, err := db.DB.Query(
		`SELECT id, email, first_name, last_name, role, created_at, updated_at 
		 FROM users ORDER BY created_at DESC`,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
