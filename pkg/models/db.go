package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sqlx.DB

func InitDB(conn_string string) error {
	conn, err := sqlx.Connect("postgres", conn_string)
	if err != nil {
		return err
	}
	if err := conn.Ping(); err != nil {
		return err
	}
	db = conn
	if err := MigrateTables(); err != nil {
		return fmt.Errorf("error while migrating tables: %v", err)
	}
	return nil
}

func MigrateTables() error {
	if _, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS projects (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	); 
	CREATE INDEX IF NOT EXISTS idx_projects_id ON projects (id);
	`); err != nil {
		return fmt.Errorf("error while creating table projects: %v", err)
	}

	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM projects WHERE name = $1", "первая запись")
	if err != nil {
		return fmt.Errorf("error while checking if project exists: %v", err)
	}
	if count == 0 {
		_, err := db.Exec("INSERT INTO projects (name) VALUES ($1)", "первая запись")
		if err != nil {
			return fmt.Errorf("error while inserting in table projects: %v", err)
		}
	}

	if _, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS goods (
		id SERIAL PRIMARY KEY NOT NULL,
		project_id INTEGER REFERENCES projects(id) NOT NULL,
		name TEXT NOT NULL,
		description TEXT,
		priority INTEGER NOT NULL,
		removed BOOLEAN NOT NULL DEFAULT false,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
	CREATE INDEX IF NOT EXISTS idx_goods_id ON goods (id);
	CREATE INDEX IF NOT EXISTS idx_goods_project_id ON goods (project_id);
	CREATE INDEX IF NOT EXISTS idx_goods_name ON goods (name);
	`); err != nil {
		return fmt.Errorf("error while creating goods: %v", err)
	}

	return nil
}

func GetGoodsFromDB(limit int, offset int)(*[]Good,*Meta,error){
	var goods []Good
	var total int
	var removed_count int
	query_goods := `SELECT * FROM goods LIMIT $1 OFFSET $2`
	query_total := `SELECT COUNT(*) FROM goods`
	query_removed := `SELECT COUNT(*) FROM goods WHERE removed = true`
	err := db.Select(&goods,query_goods,limit,offset)
	if err != nil{
		return nil,nil,err
	}
	err = db.Get(&total,query_total)	
	if err != nil{
		return nil,nil,err
	}
	err = db.Get(&removed_count,query_removed)	
	if err != nil{
		return nil,nil,err
	}
	meta := Meta{
		Limit: limit,
		Offset: offset,
		Removed: removed_count,
		Total: total,
	}
	return &goods,&meta,err
}

	
