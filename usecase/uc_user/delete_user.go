package ucuser

func (uu *userUsecase) DeleteUser(userId uint) error {
	err := uu.userRepository.DeleteUser(userId)
	return err
}
