package repository

import (
	"course/entity"
	"database/sql"
)

type CourseMySqlRepository struct {
	Db *sql.DB
}

func (c CourseMySqlRepository) Insert(course entity.Course) error {
	stmt, err := c.Db.Prepare(`INSERT INTO courses(id, name, description, status) values(?,?,?,?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status,
	)

	if err != nil {
		return err
	}

	return nil
}
