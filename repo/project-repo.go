package repo

import (
	"log"

	"github.com/Desgue/Tasker-Cli/domain"
	"github.com/Desgue/Tasker-Cli/repo/db"
)

type ProjectRepository interface {
	CreateProject(domain.ProjectRequest) (domain.ProjectRequest, error)
	GetProjects() ([]domain.ProjectRequest, error)
}

type projectRepository struct {
	sql *db.SqliteDB
}

func NewProjectRepository(db *db.SqliteDB) ProjectRepository {
	return &projectRepository{sql: db}
}

func (r *projectRepository) CreateProject(p domain.ProjectRequest) (domain.ProjectRequest, error) {
	log.Println("Hello from repo createproject")
	result, err := r.sql.DB.Exec("INSERT INTO Projects (title, description, priority) VALUES (?, ?, ?)", p.Title, p.Description, p.Priority)
	if err != nil {
		log.Println("Error from repo createproject", err)
		return domain.ProjectRequest{}, err
	}
	createdId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last created id from inserted project", err)
		return domain.ProjectRequest{}, err
	}
	Id := int(createdId)

	err = r.sql.DB.QueryRow("SELECT * FROM Projects WHERE id = ?", Id).Scan(&p.Id, &p.Title, &p.Description, &p.Priority)
	if err != nil {
		log.Println("Error getting last created project", err)
		return domain.ProjectRequest{}, err
	}
	log.Println("Project created with id:", p.Id)

	return p, nil
}
func (r *projectRepository) GetProjects() ([]domain.ProjectRequest, error) {
	log.Println("Hello from repo getprojects")
	rows, err := r.sql.DB.Query("SELECT * FROM Projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []domain.ProjectRequest
	for rows.Next() {
		var p domain.ProjectRequest
		err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.Priority)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}
