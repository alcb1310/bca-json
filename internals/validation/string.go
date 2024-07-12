package validation

import "errors"

func ValidateString(name string, minLength int, required bool) error {
    if !required && name == "" {
        return nil
    }

    if required && name == "" {
        return errors.New("Name is required")
    }

    if len(name) < minLength {
        return errors.New("Invalid name")
    }

    return nil
}
