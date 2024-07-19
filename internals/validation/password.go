package validation

import "errors"

func ValidatePassword(name string, minLength int, required bool) error {
    if !required && name == "" {
        return nil
    }

    if required && name == "" {
        return errors.New("Password is required")
    }

    if len(name) < minLength {
        return errors.New("Invalid password")
    }

    return nil
}
