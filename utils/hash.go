package utils

import "golang.org/x/crypto/bcrypt"

// Package utils provides utility functions for password hashing and verification.
//
// HashPassword takes a plain text password and returns its bcrypt hash using a cost of 14.
// It returns the hashed password as a string and any error encountered during hashing.
//
// CheckPasswordHash compares a plain text password with a bcrypt hashed password.
// It returns true if the password matches the hash, and false otherwise.

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}