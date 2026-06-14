package store

import (
	"database/sql"
	"fmt"

	pb "github.com/DmitriiS-dev/basic-backend/proto"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(host, port, user, password, dbname string) (*PostgresStore, error) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (p *PostgresStore) AddTask(t *pb.Task) error {
	query := `
		INSERT INTO tasks (title, description, status)
		VALUES ($1, $2, $3)
		RETURNING id;
	`
	// can add this later on &t.CompletionStatus
	return p.db.QueryRow(query, t.Title, t.Description).
		Scan(&t.Id)
}

func (p *PostgresStore) ListTasks() ([]*pb.Task, error) {
	rows, err := p.db.Query(`SELECT id, title, description, status FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*pb.Task

	for rows.Next() {
		t := &pb.Task{}
		err := rows.Scan(&t.Id, &t.Title, &t.Description) // can add this later on &t.CompletionStatus
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (p *PostgresStore) GetTask(id int32) (*pb.Task, error) {

	t := &pb.Task{}

	// can add this later on &t.CompletionStatus
	err := p.db.QueryRow(
		`SELECT id, title, description, status FROM tasks WHERE id=$1`,
		id,
	).Scan(&t.Id, &t.Title, &t.Description)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (p *PostgresStore) DeleteTask(id int32) (*pb.Task, error) {
	t := &pb.Task{}

	query := `DELETE FROM tasks WHERE id=$1 RETURNING id, title, description`

	err := p.db.QueryRow(query, id).Scan(&t.Id, &t.Title, &t.Description)

	if err != nil {
		// If no row was deleted, Postgres returns sql.ErrNoRows
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Task with Id %d not found", id)
		}
		// Return any other database errors (e.g., connection lost)
		return nil, err
	}

	return t, nil
}
