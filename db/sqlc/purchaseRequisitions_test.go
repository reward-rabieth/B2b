package users

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreatePurchaseRequisitions(t *testing.T) {
	params := CreatePurchaseRequisitionParams{
		Requisitionid: 44,
		Requesterid:   88,
		Description:   "vioo",
		Status:        "approved",
	}
	prqs, err := testRepo.CreatePurchaseRequisition(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, prqs)
	require.Equal(t, params.Requisitionid, prqs.Requisitionid)
	require.Equal(t, params.Requesterid, prqs.Requesterid)
	require.Equal(t, params.Description, prqs.Description)
}
