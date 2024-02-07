-- name: CreatePurchaseRequisition :one
INSERT INTO purchaserequisitions (
    requisitionid, requesterid,title, description,status
) VALUES (
             $1, $2,$3,$4,$5
         )
    RETURNING *;

-- name: GetPurchaseRequisition :one
SELECT * FROM purchaserequisitions
WHERE requisitionid = $1 LIMIT 1;

-- name: ListPurchaseRequisitions :many
SELECT * FROM purchaserequisitions
ORDER BY datesubmitted;


-- name: UpdatePurchaseRequisition :exec
UPDATE purchaserequisitions
set requesterid = $2,
    description = $3,
    status= $4,
    datesubmitted=$5,
    dateapproved=$6,
    daterejected=$7

WHERE requisitionid = $1;

-- name: DeletePurchaseRequisition :exec
DELETE FROM purchaserequisitions
WHERE requisitionid = $1;