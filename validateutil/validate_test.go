package validateutil

import "testing"

func TestValidateString(t *testing.T) {
	password1 := "465465454464546"
	password2 := "156+_12.?"

	if err := ValidateString(password1); err != nil {
		t.Errorf("ValidateString(%s) = %v, want nil", password1, err)
	}
	if err := ValidateString(password2); err == nil {
		t.Errorf("ValidateString(%s) = nil, want error", password2)
	}
}
