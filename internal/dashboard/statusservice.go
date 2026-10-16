package dashboard

import "errors"

type StatusService struct{}

func (s *StatusService) GetTotalCount() (int, error) {

	return 1, nil
}

func (s *StatusService) GetEnabledCount() (int, error) {
	return 0, errors.New("WIP")
}
