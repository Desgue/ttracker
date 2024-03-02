package svc

import (
	"log"

	"github.com/Desgue/Tasker-Cli/domain"
	"github.com/Desgue/Tasker-Cli/repo"
)

type ProjectService interface {
	AddProject(domain.ProjectRequest) (domain.ProjectItem, error)
	GetProjects() ([]domain.ProjectItem, error)
}

type projectService struct {
	repo repo.ProjectRepository
}

func NewProjectService(repo repo.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) AddProject(p domain.ProjectRequest) (domain.ProjectItem, error) {
	log.Println("Hello from service addproject")
	projectRes, err := s.repo.CreateProject(p)
	if err != nil {
		return domain.ProjectItem{}, err
	}
	projectItem := domain.NewProjectItem(projectRes)
	return projectItem, nil
}

func (s *projectService) GetProjects() ([]domain.ProjectItem, error) {
	log.Println("Hello from service getprojects")
	res, err := s.repo.GetProjects()
	if err != nil {
		return nil, err
	}
	var projects []domain.ProjectItem
	for _, p := range res {
		projects = append(projects, domain.NewProjectItem(p))
	}
	return projects, nil
}
