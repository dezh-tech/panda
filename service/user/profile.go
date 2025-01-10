package user

import "errors"

type ProfileRequest struct {
	Pubkey string
}

type ProfileResponse struct {
	Name string `json:"name"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	user, err := s.repo.GetByPubkey(req.Pubkey)
	if err != nil {
		return ProfileResponse{}, errors.New("can't get the profile") // todo::: move to errors.go
	}

	return ProfileResponse{Name: user.Name}, nil
}
