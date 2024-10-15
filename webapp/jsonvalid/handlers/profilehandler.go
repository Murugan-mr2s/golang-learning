package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type Profile struct {
	Id     uint    `json:"id" validate:"required`
	Name   string  `json:"name" validate:"required,min=4`
	Email  string  `json:"email" validate:"required,email" `
	Age    uint    `json:"age" validate:"required,gte=18,lte=100" `
	Salary float64 `json:"salary" validate:"required"`
}

func NewProfileHandler() *Profile {
	return &Profile{}
}

func (p *Profile) GetAllProfile(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Get profiles"))
}

func (p *Profile) AddNewProfile(rw http.ResponseWriter, r *http.Request) {
	profile := r.Context().Value("profile").(Profile)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, profile)

	//rw.Write([]byte(string(profile.Id) + ":" + profile.Name + ":" + profile.Email + ":" + string(profile.Age) + ":" + profile.Salary) )
}

var validate = validator.New()

func ProfileValidator(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var profile = new(Profile)

		if err := json.NewDecoder(r.Body).Decode(profile); err != nil {
			http.Error(w, "unable to json marshalliing", http.StatusBadRequest)
			return
		}

		if err := validate.Struct(profile); err != nil {
			validationError := err.(validator.ValidationErrors)
			http.Error(w, formatValidationErrors(validationError), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "profile", *profile)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func formatValidationErrors(errors validator.ValidationErrors) string {
	var message string
	for _, err := range errors {
		message += err.Field() + " is " + err.Tag() + ". "
	}
	return message
}
