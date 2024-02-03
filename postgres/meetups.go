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

	err := m.DB.Model(&meetups).Order("id").Select()

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

func (m *MeetupsRepo) DeleteMeetup(id string) (bool, error) {
	meetup := &models.Meetup{ID: id}

	_, err := m.DB.Model(meetup).WherePK().Delete()

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *MeetupsRepo) GetMeetupsForUser(obj *models.User) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := m.DB.Model(&meetups).Where("user_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return meetups, nil
}
