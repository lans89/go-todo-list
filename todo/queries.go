package todoheader

import (
	"api-todolist/db"

	"github.com/google/uuid"
)

func SaveHeader(header *ToDoHeaderDto) error {
	uuid, _ := uuid.NewRandom()
	header.UUID = uuid.String()
	query :=
		`INSERT INTO tbl_list_header (id, "name", description, "owner") 
		VALUES ($1, $2, $3, $4)`
	_, err := db.SqlDB.Exec(query, header.UUID, header.Name, header.Description, header.Owner)
	if err != nil {
		return err
	}
	return nil
}

func ListHeaderFrom(owner string) ([]ToDoHeaderDto, error) {
	query :=
		`SELECT id, name, description, active, created_at, finish_at, last_updated_at, owner
		FROM tbl_list_header 
		WHERE "owner" = $1 and status is true`
	rows, err := db.SqlDB.Query(query, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []ToDoHeaderDto
	for rows.Next() {
		var header ToDoHeaderDto
		if err := rows.Scan(&header.UUID, &header.Name, &header.Description, &header.Active, &header.CreationTime, &header.FinishTime, &header.UpdateTime, &header.Owner); err != nil {
			return nil, err
		}
		results = append(results, header)
	}

	return results, nil
}
