package queries

const (
	QueryGetUserByEmail = `
		SELECT 
			n.fullname
		FROM 
			nasabahs n
		WHERE
			(n.email = $1 OR
			n.username = $2) AND
			n.is_active = true
	`
)
