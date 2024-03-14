-- name: CreateSupplier :many
INSERT INTO suppliers (
                    supplier_id, supplier_name, contact_person, contact_mail, supplier_type
) VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: ListSupplier :many
SELECT * FROM  suppliers
ORDER BY supplier_name;