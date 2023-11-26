package xclone

import (
	"testing"
)

func TestRegisterInput_Sanitize(t *testing.T) {
	input := RegisterInput{
		Username:        " xxx ",
		Email:           " Xxx@gmail.com  ",
		Password:        "password",
		ConfirmPassword: "password",
	}

	want := RegisterInput{
		Username:        "xxx",
		Email:           "xxx@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	input.Sanitize()

	if input != want {
		t.Errorf("Sanitize() = %v, want %v", input, want)
	}
}

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input RegisterInput
		err   error
	}{
		{
			name: "valid",
			input: RegisterInput{
				Username:        "xxx",
				Email:           "xxx@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: nil,
		},
		// Other test cases...

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()

			if tc.err != nil && err == nil {
				t.Errorf("Validate() expected error %v, got nil", tc.err)
			} else if tc.err == nil && err != nil {
				t.Errorf("Validate() expected no error, got %v", err)
			}
		})
	}
}

// Similar changes for other test functions...
