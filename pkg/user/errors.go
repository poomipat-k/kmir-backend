package user

type UsernameRequiredError struct{}

func (e UsernameRequiredError) Error() string {
	return "username is required"
}

type PasswordTooShortError struct{}

func (e PasswordTooShortError) Error() string {
	return "password minimum length are 8 characters"
}

type PasswordRequiredError struct{}

func (e PasswordRequiredError) Error() string {
	return "password is required"
}

type PasswordTooLongError struct{}

func (e PasswordTooLongError) Error() string {
	return "password maximum length are 60 characters"
}

type InvalidCredentialError struct{}

func (e InvalidCredentialError) Error() string {
	return "invalid credentials"
}
