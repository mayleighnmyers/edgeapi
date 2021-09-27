package routes

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/redhatinsights/edge-api/pkg/dependencies"
	"github.com/redhatinsights/edge-api/pkg/services/mock_services"
)

func TestListAllImageSets(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := req.Context()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImagesetService := mock_services.NewMockImageSetsServiceInterface(ctrl)
	mockImagesetService.EXPECT().ListAllImageSets(gomock.Any(), gomock.Any()).Return(nil)
	ctx = context.WithValue(ctx, dependencies.Key, &dependencies.EdgeAPIServices{
		ImageSetService: mockImagesetService,
	})

	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListAllImageSets)

	handler.ServeHTTP(rr, req.WithContext(ctx))

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v",
			status, http.StatusOK)

	}
}

func TestGetImageSetByID(t *testing.T) {
	// var id int = 0
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := req.Context()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockImagesetService := mock_services.NewMockImageSetsServiceInterface(ctrl)
	mockImagesetService.EXPECT().GetImageSetsByID(gomock.Any()).Return(nil, nil)
	ctx = context.WithValue(ctx, dependencies.Key, &dependencies.EdgeAPIServices{
		ImageSetService: mockImagesetService,
	})

	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetImageSetsByID)

	handler.ServeHTTP(rr, req)
	fmt.Printf(":: RRR ::: %v\n", rr)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v",
			status, http.StatusOK)

	}
}
