package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_I_Get_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func Test_If_I_Get_An_Error_If_Price_Is_Less_Than_Zero(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "invalid price")
}

func Test_If_I_Get_An_Error_If_Tax_Is_Less_Than_Zero(t *testing.T) {
	order := Order{ID: "123", Price: 10}
	assert.Error(t, order.Validate(), "invalid tax")
}

func Test_If_I_Get_No_Error_If_All_The_Values_Are_Valid(t *testing.T) {
	order := Order{ID: "123", Price: 10, Tax: 2}
	assert.NoError(t, order.Validate())
	order.CalculateFinalPrice()
	assert.Equal(t, 12.0, order.FinalPrice)

}
