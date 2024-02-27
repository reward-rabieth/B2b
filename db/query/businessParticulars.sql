-- name: CreateBusinessParticular :one
INSERT INTO business_particulars(REG_BUSINESS_NAME,
                                 BRELA_REG_NUMBER,
                                 USER_ID,
                                 PO_BOX,
                                 OCCUPATION_LOCATION,
                                 COUNTRY,
                                 REGION,
                                 TIN,
                                 FIRST_NAME,
                                 LAST_NAME,
                                 CONTACT,
                                 WEBSITE
                                 ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
returning *;


-- name: UpdateBusinessParticulars :exec
UPDATE business_particulars
SET reg_business_name=$1,
    brela_reg_number=$2,
    po_box=$3,
    occupation_location=$4,
    country=$5,
    region=$6,
    tin=$7,
    first_name=$8,
    last_name=$9,
    contact=$10,
    website=$11
where brela_reg_number=$2;