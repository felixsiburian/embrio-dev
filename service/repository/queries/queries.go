package queries

const (
	QueryGetUserByEmail = `
		SELECT 
			n.fullname
		FROM 
			nasabahs n
		WHERE
			n.email = $1 AND
			n.is_active = true
	`
)
