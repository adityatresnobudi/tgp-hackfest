package receipt_pg

const GET_ALL_RECEIPTS = `
	SELECT r.name, r.category, r.total as total_bill, u.id, u.name, i.id, i.product, i.price, i.quantity, i.discount, i.tax, i.service, i.total
	FROM items i
	LEFT JOIN receipts r on i.receipt_id = r.id
	LEFT JOIN users u on i.user_id = u.id;
`

const GET_ONE_RECEIPT_BY_USER_ID = `
	SELECT r.name, r.category, r.total as total_bill, u.id, u.name, i.id, i.product, i.price, i.quantity, i.discount, i.tax, i.service, i.total
	FROM items i
	LEFT JOIN receipts r on i.receipt_id = r.id
	LEFT JOIN users u on i.user_id = u.id
	WHERE r.id = $1 AND u.id = $2;
`