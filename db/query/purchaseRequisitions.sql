-- name: CreatePurchaseRequisition :one
INSERT INTO purchase_requisitions(
requisition_id,
requester_id,
description,
title,
status
)VALUES ($1, $2, $3, $4,$5
         ) RETURNING *;

-- name: GetPurchaseRequisitionByID :one
SELECT * FROM purchase_requisitions
WHERE requisition_id = $1 LIMIT 1;



-- name: ListPurchaseRequisitions :many
SELECT * FROM purchase_requisitions
ORDER BY date_submitted;


-- name: UpdatePurchaseRequisition :exec
UPDATE purchase_requisitions
set status= $1
WHERE requisition_id = $2
RETURNING *;



-- name: DeletePurchaseRequisition :exec
DELETE FROM purchase_requisitions
WHERE requisition_id = $1;

