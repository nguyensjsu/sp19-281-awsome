package main

import "golang.org/x/crypto/bcrypt"

func compareHash(inputPassword string, dbPasswordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPasswordHash), []byte(inputPassword))
	return err == nil
}

func getPasswordHash(inputPassword string) string {
	passwordBytes := []byte(inputPassword)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func compareBytes(a, b []byte) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}