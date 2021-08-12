package postgres

import (
	"context"
	"database/sql"
	"diffme.dev/diffme-api/core"
	domain "diffme.dev/diffme-api/modules/events"
)

type postgresEventRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewPostgresEventRepository(Conn *sql.DB) domain.EventRepository {
	return &postgresEventRepository{Conn}
}

func (m *postgresEventRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Event, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			println(errRow)
		}
	}()

	result = make([]domain.Event, 0)

	for rows.Next() {
		t := domain.Event{}

		err = rows.Scan(
			&t.ID,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *postgresEventRepository) GetByID(ctx context.Context, id string) (res domain.Event, err error) {
	query := `SELECT id,updated_at, created_at
  						FROM article WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Event{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, core.ErrNotFound
	}

	return
}
