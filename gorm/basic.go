package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User represents a user entity
type User struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Email   string  `gorm:"unique"`
	Profile Profile `gorm:"constraint:OnDelete:CASCADE;"`
	Posts   []Post  `gorm:"constraint:OnDelete:CASCADE;"`
	Groups  []Group `gorm:"many2many:user_groups;"`
}

// Profile represents a one-to-one relationship
type Profile struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"unique"`
	Bio    string
}

// Post represents a one-to-many relationship
type Post struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Title  string
	Body   string
}

// Group represents a many-to-many relationship
type Group struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique"`
	Users []User `gorm:"many2many:user_groups;"`
}

func main() {
	fmt.Println("Start")
	// Open an in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto migrate the schema
	db.AutoMigrate(&User{}, &Profile{}, &Post{}, &Group{})

	// Create users
	alice := User{Name: "Alice", Email: "alice@example.com", Profile: Profile{Bio: "Software Developer"}}
	bob := User{Name: "Bob", Email: "bob@example.com", Profile: Profile{Bio: "Data Scientist"}}
	db.Create(&alice)
	db.Create(&bob)

	// Create posts
	db.Create(&Post{UserID: alice.ID, Title: "Golang 101", Body: "Introduction to Golang."})
	db.Create(&Post{UserID: alice.ID, Title: "GORM Guide", Body: "How to use GORM."})
	db.Create(&Post{UserID: bob.ID, Title: "Data Science Basics", Body: "Understanding machine learning."})

	// Create groups
	group1 := Group{Name: "Developers"}
	group2 := Group{Name: "Data Scientists"}
	db.Create(&group1)
	db.Create(&group2)

	// Associate users with groups
	db.Model(&alice).Association("Groups").Append(&group1)
	db.Model(&bob).Association("Groups").Append(&group2)

	// Read user with profile and posts
	var user User
	db.Preload("Profile").Preload("Posts").First(&user, "email = ?", "alice@example.com")
	fmt.Println("User:", user.Name, "Profile Bio:", user.Profile.Bio, "Posts:", len(user.Posts))

	// Read user groups
	db.Preload("Groups").First(&user, "email = ?", "alice@example.com")
	fmt.Println("User Groups:", user.Groups)
}
