package domain

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/google/uuid"
)

const (
	userFullNameMinLength = 3
	userFullNameMaxLength = 100
)

var userPhoneNumberRegexp = regexp.MustCompile(`^\+[1-9]\d{9,14}$`)

type User struct {
	ID          uuid.UUID
	Version     int64
	FullName    string
	PhoneNumber *string
}

func CreateUser(fullName string, phoneNumber *string) User {
	return User{
		ID:          uuid.New(),
		Version:     1,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func (u User) Validate() error {
	if length := utf8.RuneCountInString(u.FullName); length < userFullNameMinLength || length > userFullNameMaxLength {
		return fmt.Errorf("`FullName` must be between %d and %d characters long", userFullNameMinLength, userFullNameMaxLength)
	}
	if u.PhoneNumber != nil && !userPhoneNumberRegexp.MatchString(*u.PhoneNumber) {
		return fmt.Errorf("`PhoneNumber` must be a valid phone number in E.164 format")
	}
	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func (p UserPatch) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf("`FullName` can't be null")
	}
	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u
	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}
	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp
	return nil
}
