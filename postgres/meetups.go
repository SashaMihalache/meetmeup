package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/sashamihalache/meetmeup/models"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (m *MeetupsRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := m.DB.Model(&meetups).Select()

	if err != nil {
		return nil, err
	}

	return meetups, nil
}
