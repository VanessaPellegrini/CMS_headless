package user_service

import (
	userRepository "github.com/VanessaPellegrini/CMS_headless/app/repositories"

	m "github.com/VanessaPellegrini/CMS_headless/app/models"
)

func Create(user m.User) error {
	err := userRepository.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func Read() (m.Users, error) {

	users, err := userRepository.Read()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func Update(user m.User, userId string) error {
	err := userRepository.Update(user, userId)

	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {
	err := userRepository.Delete(userId)

	if err != nil {
		return err
	}
	return nil
}
