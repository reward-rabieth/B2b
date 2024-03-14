-- name: CreateQuotationRequest :one
INSERT INTO quotation_requests (
proposal_request_id, requisition_id, supplier_id, status
) VALUES ($1,$2,$3,$4)
RETURNING *;