package main

import (
	"errors"
	"strings"
	"strconv"
)

type PuppySearch struct {
    Breed string
    Gender string
    ZipCode string
    Radius string
}

func (puppySearch PuppySearch) validate() error {
	if puppySearch.Breed == "" {
		return errors.New("Breed cannot be empty")
	}

	if puppySearch.Gender != "" {
		genderLower := strings.ToLower(puppySearch.Gender)
		if genderLower != "male" || genderLower != "female" {
			return errors.New("Gender must be male or female")
		}
	}

	if puppySearch.ZipCode != "" {
		_, err := strconv.Atoi(puppySearch.ZipCode)
		if err != nil {
			return errors.New("ZipCode must be numeric")
		}
		if len(puppySearch.ZipCode) != 5 {
			return errors.New("ZipCode must be 5 numbers")
		}
	}

	if puppySearch.Radius != "" {
		if puppySearch.ZipCode == "" {
			return errors.New("Cannot have a radius without a ZipCode")
		}

		val, err := strconv.Atoi(puppySearch.Radius)
		if err != nil {
			if strings.ToLower(puppySearch.Radius) != "nationwide" {
				return errors.New("Radius must be one of Nationwide, 500, 250, 100, 50, 25")
			}
		} else {
			if val != 500 && val != 250 && val != 100 && val != 50 && val != 25 {
				return errors.New("Radius must be one of Nationwide, 500, 250, 100, 50, 25")
			}
		}
	}

	return nil
}

func (puppySearch PuppySearch) createRequestParams(page int) string {
	result := ""
	token := "?"
	if puppySearch.Gender != "" {
		result = result + token + "gender=" + puppySearch.Gender
		token = "&"
	}
	if puppySearch.ZipCode != "" {
		result = result + token + "location=" + puppySearch.ZipCode
		token = "&"
	}
	if puppySearch.Radius != "" {
		result = result + token + "radius=" + puppySearch.Radius
	}
	result = result + token + "page=" + strconv.Itoa(page)
	return result
}