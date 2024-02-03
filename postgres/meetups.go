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

func (m *MeetupsRepo) GetById(id string) (*models.Meetup, error) {
	meetup := &models.Meetup{ID: id}

	err := m.DB.Model(meetup).WherePK().Select()

	if err != nil {
		return nil, err
	}

	return meetup, nil
}

func (m *MeetupsRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()

	if err != nil {
		return nil, err
	}

	return meetup, nil
}

func (m *MeetupsRepo) UpdateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).WherePK().Update()

	if err != nil {
		return nil, err
	}

	return meetup, nil
}
