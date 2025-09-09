package events

import (
	"errors"

	"example.com/events/common/database"
)

func (event Event) save() (*Event, error) {
	query := `INSERT INTO events (event_id, name, description, start_date, end_date, location, created_at, updated_at, user_id) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := database.DB.Exec(query,
		event.EventID, event.Name, event.Description, event.StartDate,
		event.EndDate, event.Location, event.CreatedAt, event.UpdatedAt, event.UserID)

	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	return &event, nil
}

func find() ([]Event, error) {
	query := `SELECT * FROM events WHERE is_deleted = false`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.EventID, &event.Name, &event.Description, &event.StartDate, &event.EndDate, &event.Location, &event.CreatedAt, &event.UpdatedAt, &event.IsDeleted, &event.UserID)
		if err != nil {
			return nil, errors.New("Something went wrong")
		}
		events = append(events, event)
	}

	return events, nil
}

func findByID(id string) (*Event, error) {
	query := `SELECT * FROM events WHERE event_id = ? AND is_deleted = false`

	row := database.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.EventID, &event.Name, &event.Description, &event.StartDate, &event.EndDate, &event.Location, &event.CreatedAt, &event.UpdatedAt, &event.IsDeleted, &event.UserID)
	if err != nil {
		return nil, errors.New("Event not found")
	}

	return &event, nil
}

func updateOne(event *Event) (*Event, error) {
	query := `UPDATE events SET name = ?, description = ?, start_date = ?, end_date = ?, location = ?, updated_at = ? WHERE event_id = ? AND is_deleted = false`

	_, err := database.DB.Exec(query, event.Name, event.Description, event.StartDate, event.EndDate, event.Location, event.UpdatedAt, event.EventID, event.IsDeleted, event.UserID)
	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	return event, nil
}

func deleteOne(id string) error {
	query := `UPDATE events SET is_deleted = true WHERE event_id = ?`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return errors.New("Something went wrong")
	}
	return nil
}
