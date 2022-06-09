package service_test

import (
	"gobasic/errs"
	"gobasic/repository"
	"gobasic/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCustomers(t *testing.T) {

	custRepo := repository.NewCustomerRepositoryMock()

	custService := service.NewCustomerService(custRepo)

	custResp, _ := custService.GetCustomers()
	expected := []service.CustomerResponse{
		{ID: 1000, Name: "Boat"},
		{ID: 1001, Name: "Big"},
		{ID: 1002, Name: "Beam"},
	}

	assert.Equal(t, expected, custResp)

}

func TestGetCustomer(t *testing.T) {

	type testCase struct {
		id       int
		expected service.CustomerResponse
		name     string
	}

	customerResponses := []service.CustomerResponse{
		{ID: 1000, Name: "Boat"},
		{ID: 1001, Name: "Big"},
		{ID: 1002, Name: "Beam"},
	}

	cases := []testCase{
		{id: customerResponses[0].ID, expected: customerResponses[0], name: "id 1000"},
		{id: customerResponses[1].ID, expected: customerResponses[1], name: "id 1001"},
		{id: customerResponses[2].ID, expected: customerResponses[2], name: "id 1002"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			custRepo := repository.NewCustomerRepositoryMock()

			custService := service.NewCustomerService(custRepo)

			custResp, _ := custService.GetCustomer(c.id)

			expected := c.expected

			assert.Equal(t, &expected, custResp)
		})
	}

	t.Run("test not found", func(t *testing.T) {
		custRepo := repository.NewCustomerRepositoryMock()

		custService := service.NewCustomerService(custRepo)

		_, err := custService.GetCustomer(23)

		expected := errs.ErrCustNotFound

		assert.Equal(t, expected, err)
	})
}
