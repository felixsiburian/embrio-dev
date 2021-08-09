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

	QueryUpdateSaldo = `
		UPDATE
			rekenings
		SET
			saldo = $1,
			modified_by = $2,
			modified_date = now()
		WHERE
			nasabah_id = $3 AND
			no_rekening = $4
`

	QueryGetNasabahByID = "SELECT " +
		"n.username, n.fullname, n.email, n.phone_number, n.alamat, n.created_date " +
		"FROM " +
		"nasabahs n " +
		"WHERE " +
		"n.nasabah_id = $1"

	QueryGetSaldoNasabah = "SELECT " +
		"n.fullname, r.no_rekening, r.cabang_bank, r.saldo, r.is_active, r.created_date " +
		"FROM " +
		"rekenings r " +
		"LEFT JOIN nasabahs n " +
		"on r.nasabah_id = n.nasabah_id " +
		"WHERE " +
		"r.nasabah_id = $1 AND " +
		"r.no_rekening = $2"
)
