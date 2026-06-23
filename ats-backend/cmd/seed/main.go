package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=1234 dbname=talentflow port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 1. Seed Tenant
	tenantID := "00000000-0000-0000-0000-000000000001"
	db.Exec(`
		INSERT INTO tenants (id, name, domain, created_at, updated_at)
		VALUES (?, 'Default Tenant', 'default.talentflow.com', NOW(), NOW())
		ON CONFLICT (id) DO NOTHING
	`, tenantID)

	// 2. Seed Role
	roleID := "11111111-1111-1111-1111-111111111111"
	db.Exec(`
		INSERT INTO roles (id, name, description, created_at, updated_at)
		VALUES (?, 'Admin', 'Super Administrator', NOW(), NOW())
		ON CONFLICT (id) DO NOTHING
	`, roleID)

	// 3. Hash password dengan bcrypt yang benar
	password := "password123"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// 4. Seed User (upsert berdasarkan email)
	userID := "22222222-2222-2222-2222-222222222222"
	result := db.Exec(`
		INSERT INTO users (id, tenant_id, role_id, email, password, first_name, last_name, is_active, created_at, updated_at)
		VALUES (?, ?, ?, 'admin@talentflow.com', ?, 'Super', 'Admin', true, NOW(), NOW())
		ON CONFLICT (email) DO UPDATE SET password = EXCLUDED.password
	`, userID, tenantID, roleID, string(hash))

	if result.Error != nil {
		log.Fatalf("Failed to seed user: %v", result.Error)
	}

	fmt.Println("✅ Seeding berhasil!")
	fmt.Println("   Email   : admin@talentflow.com")
	fmt.Println("   Password: password123")
}
