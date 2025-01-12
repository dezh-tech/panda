package domainService

import (
	"context"
	"fmt"

	"github.com/dezh-tech/panda/pkg/validator"
	schema "github.com/dezh-tech/panda/schemas"
)

func (s DomainService) GetAll(ctx context.Context, filter interface{}) (*[]schema.Domain, *validator.Varror) {
	domains, err := s.repo.GetAll(ctx, filter)
	if err != nil {
		return nil, &validator.Varror{Error: err.Error()}
	}

	return domains, nil
}
