package users

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreatePurchaseRequisitions(t *testing.T) {
	params := CreatePurchaseRequisitionParams{
		Title:       "ujenzi",
		Description: "vifaaa ",
	}
	prqs, err := testRepo.CreatePurchaseRequisition(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, prqs)
	require.Equal(t, params.Title, prqs.Title)
	require.Equal(t, params.Description, prqs.Description)
}
