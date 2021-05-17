package api

import (
	"context"

	internalclient "github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/chanioxaris/go-datagovgr/types"
)

// Education holds required info to consume Education related endpoints.
type Education struct {
	client *internalclient.Client
}

// NewEducation creates a new instance.
func NewEducation(client *internalclient.Client) *Education {
	return &Education{client: client}
}

// UniversityTeachingStaff retrieves data for the number of teaching staff by school and ranking.
func (e *Education) UniversityTeachingStaff(ctx context.Context) ([]*types.UniversityTeachingStaff, error) {
	req, err := e.client.NewRequestGET(ctx, "minedu_dep")
	if err != nil {
		return nil, err
	}

	response := make([]*types.UniversityTeachingStaff, 0)
	if err := e.client.MakeRequest(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// StudentsBySchool retrieves data for the number of students by school and gender.
func (e *Education) StudentsBySchool(ctx context.Context) ([]*types.StudentsBySchool, error) {
	req, err := e.client.NewRequestGET(ctx, "minedu_students_school")
	if err != nil {
		return nil, err
	}

	response := make([]*types.StudentsBySchool, 0)
	if err := e.client.MakeRequest(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// AtlasInternshipStatistics retrieves annual statistics of internships through the Atlas system.
func (e *Education) AtlasInternshipStatistics(ctx context.Context) ([]*types.AtlasInternshipStatistics, error) {
	req, err := e.client.NewRequestGET(ctx, "grnet_atlas")
	if err != nil {
		return nil, err
	}

	response := make([]*types.AtlasInternshipStatistics, 0)
	if err := e.client.MakeRequest(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// EudoksosRequestsAndDeliveries retrieves statistics of requests and deliveries for books made through the Eudoxus system.
func (e *Education) EudoksosRequestsAndDeliveries(ctx context.Context) ([]*types.EudoksosRequestsAndDeliveries, error) {
	req, err := e.client.NewRequestGET(ctx, "grnet_eudoxus")
	if err != nil {
		return nil, err
	}

	response := make([]*types.EudoksosRequestsAndDeliveries, 0)
	if err := e.client.MakeRequest(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
