package config

import "github.com/go-playground/validator/v10"

// Validator is a shared instance of the validator
var Validator = validator.New()
