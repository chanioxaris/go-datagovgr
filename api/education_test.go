package api_test

import (
	"context"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/chanioxaris/go-datagovgr/datagovgrtest"
	"github.com/jarcoal/httpmock"
)

func TestEducation_UniversityTeachingStaff_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.UniversityTeachingStaff,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.UniversityTeachingStaff),
	)

	got, err := fixture.API.Education.UniversityTeachingStaff(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.UniversityTeachingStaff) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.UniversityTeachingStaff, got)
	}
}

func TestEducation_UniversityTeachingStaff_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.UniversityTeachingStaff,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Education.UniversityTeachingStaff(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestEducation_StudentsBySchool_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.StudentsBySchool,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.StudentsBySchool),
	)

	got, err := fixture.API.Education.StudentsBySchool(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.StudentsBySchool) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.StudentsBySchool, got)
	}
}

func TestEducation_StudentsBySchool_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.StudentsBySchool,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Education.StudentsBySchool(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestEducation_AtlasInternshipStatistics_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.AtlasInternshipStatistics,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.AtlasInternshipStatistics),
	)

	got, err := fixture.API.Education.AtlasInternshipStatistics(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.AtlasInternshipStatistics) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.AtlasInternshipStatistics, got)
	}
}

func TestEducation_AtlasInternshipStatistics_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.AtlasInternshipStatistics,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Education.AtlasInternshipStatistics(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestEducation_EudoksosRequestsAndDeliveries_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.EudoksosRequestsAndDeliveries,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.EudoksosRequestsAndDeliveries),
	)

	got, err := fixture.API.Education.EudoksosRequestsAndDeliveries(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.EudoksosRequestsAndDeliveries) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.EudoksosRequestsAndDeliveries, got)
	}
}

func TestEducation_EudoksosRequestsAndDeliveries_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.EudoksosRequestsAndDeliveries,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Education.EudoksosRequestsAndDeliveries(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}
