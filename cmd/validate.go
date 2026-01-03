package cmd

import "github.com/go-playground/validator/v10"

func (flags *PersistentFlags) Validate() error {
	validatorInstance := validator.New()
	err := validatorInstance.Struct(flags)

	if err != nil {
		return err
	}

	return nil
}
