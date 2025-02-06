package model

import (
	"database/sql/driver"

	"errors"
	"time"
)

type CustomTime time.Time

func (t *CustomTime) Scan(value any) error {
	if v, ok := value.(string); ok {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = CustomTime(parsedTime)
		return nil
	}
	return errors.New("failed to scan")
}

func (t *CustomTime) Value() (driver.Value, error) {
	if time.Time(*t).IsZero() {
		return nil, nil
	}
	return time.Time(*t).Format("2006-01-02 15:04:05"), nil
}
