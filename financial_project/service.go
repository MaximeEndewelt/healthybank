package financial_project

import "fmt"

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) Create(fp FinancialProject) (*FinancialProject, error) {
	return nil, nil
}

func (s *service) Get(id string) *FinancialProject {
	return nil
}

func (s *service) GetAll() []*FinancialProject {
	fmt.Println("Oui Maxime, nous sommes dans le service")
	return nil
}
func (s *service) Update(fp FinancialProject) *FinancialProject {
	return nil
}

func (s *service) Delete(id string) error {
	return nil
}
