package financial_project

type Service interface {
	Create(fp FinancialProject) (*FinancialProject, error)
	Get(id string) *FinancialProject
	GetAll() []*FinancialProject
	Update(fp FinancialProject) *FinancialProject
	Delete(id string) error
}
