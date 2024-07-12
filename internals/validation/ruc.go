package validation

import (
	"errors"
	"log/slog"
	"strconv"
)

var coef = [9]int{4, 3, 2, 7, 6, 5, 4, 3, 2}

func ValidateRuc(ruc string, required bool) error {
	if !required && ruc == "" {
		return nil
	}

	if required && ruc == "" {
		return errors.New("ID is required")
	}
	if _, err := strconv.Atoi(ruc); err != nil {
		return errors.New("Invalid RUC")
	}

	prov, _ := strconv.Atoi(ruc[:2])
	if prov > 24 || prov < 1 {
		return errors.New("Invalid ID")
	}

	slog.Info("ValidateRuc", "ruc", ruc, "length", len(ruc))
	switch len(ruc) {
	case 10:
		val := cedulaValidation(ruc[:10])
		if val != "" {
			return errors.New(val)
		}
	case 13:
		val := rucValidation(ruc)
		if val != "" {
			return errors.New(val)
		}
	default:
		return errors.New("Invalid ID Length")
	}
	return nil
}

func rucValidation(ruc string) string {
	if ruc[10:] != "001" {
		return "Invalid ID"
	}

	if ruc[2] != '9' {
		return cedulaValidation(ruc[:10])
	}

	sum := 0
	for i, val := range coef {
		x, _ := strconv.Atoi(string(ruc[i]))
		sum += x * val
	}

	ver := 11 - (sum % 11)
	if ver == 11 {
		ver = 0
	}
	rucVer, _ := strconv.Atoi(string(ruc[9]))

	if rucVer != ver {
		return "Invalid RUC"
	}

	return ""
}

func cedulaValidation(cedula string) string {
	sum := 0

	for i := 1; i < 10; i++ {
		coef := 1
		if i%2 != 0 {
			coef = 2
		}

		x, _ := strconv.Atoi(string(cedula[i-1]))

		res := x * coef
		if res > 9 {
			res = res%10 + 1
		}

		sum += res
	}

	ver := sum % 10
	if ver != 0 {
		ver = 10 - ver
	}
	cedulaVer, _ := strconv.Atoi(string(cedula[9]))

	if ver != cedulaVer {
		return "Invalid ID"
	}

	return ""
}
