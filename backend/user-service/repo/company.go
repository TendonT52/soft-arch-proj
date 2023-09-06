package repo

import (
	"context"

	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
)

func (r *userRepository) CreateCompany(ctx context.Context, company *pbv1.CreateCompanyRequest) error {
	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into Table users
	query := "INSERT INTO users (email, password, role) VALUES ($1, $2, $3) RETURNING id"
	var id int64
	err = tx.QueryRowContext(ctx, query, company.Email, company.Password, "company").Scan(&id)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into Table companies
	query = "INSERT INTO companies (cid, name, description, location, phone, category) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = tx.ExecContext(ctx, query, id, company.Name, company.Description, company.Location, company.Phone, company.Category)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Commit the transaction if all insertions were successful
	err = tx.Commit()
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}

func (r *userRepository) GetCompanyByID(ctx context.Context, id int64) (*pbv1.Company, error) {
	query := "SELECT users.id, companies.name, users.email, companies.description, companies.location, companies.phone, companies.category, companies.status FROM users INNER JOIN companies ON users.id = companies.cid WHERE users.id = $1"
	var company pbv1.Company
	err := r.db.QueryRowContext(ctx, query, id).Scan(&company.Id, &company.Name, &company.Email, &company.Description, &company.Location, &company.Phone, &company.Category, &company.Status)
	if err != nil {
		return &pbv1.Company{}, domain.ErrInternal.From(err.Error(), err)
	}
	return &company, nil
}

func (r *userRepository) GetAllCompany(ctx context.Context) ([]*pbv1.Company, error) {
	query := `
		SELECT users.id, companies.name, users.email, companies.description, companies.location, companies.phone, companies.category, companies.status 
		FROM users INNER JOIN companies ON users.id = companies.cid
		ORDER BY companies.name`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()

	companies := make([]*pbv1.Company, 0)
	for rows.Next() {
		var company pbv1.Company
		err := rows.Scan(&company.Id, &company.Name, &company.Email, &company.Description, &company.Location, &company.Phone, &company.Category, &company.Status)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
		companies = append(companies, &company)
	}
	return companies, nil
}

func (r *userRepository) GetApprovedCompany(ctx context.Context, search string) ([]*pbv1.Company, error) {
	query :=
		`	SELECT 
			users.id, companies.name, users.email, companies.description, companies.location, companies.phone, companies.category 
			FROM 
				users INNER JOIN companies ON users.id = companies.cid, 
				to_tsvector(companies.category || companies.name) document,
				to_tsquery($1) query,
				NULLIF(ts_rank(to_tsvector(companies.category), query), 0) rank_category,
				NULLIF(ts_rank(to_tsvector(companies.name), query), 0) rank_name,
				SIMILARITY($1, companies.category || companies.name) similarity
			WHERE 
				users.verified = true AND 
				companies.status = 'Approve' AND
				query @@ document OR similarity > 0
			ORDER BY rank_category, rank_name, similarity DESC NULLS LAST
		`

	rows, err := r.db.QueryContext(ctx, query, search)
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()

	companies := make([]*pbv1.Company, 0)
	for rows.Next() {
		var company pbv1.Company
		err := rows.Scan(&company.Id, &company.Name, &company.Email, &company.Description, &company.Location, &company.Phone, &company.Category)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
		companies = append(companies, &company)
	}
	return companies, nil
}

func (r *userRepository) UpdateCompanyByID(ctx context.Context, id int64, req *pbv1.Company) error {
	query := "UPDATE companies SET name = $1, description = $2, location = $3, phone = $4, category = $5 WHERE cid = $6"
	_, err := r.db.ExecContext(ctx, query, req.Name, req.Description, req.Location, req.Phone, req.Category, id)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}

func (r *userRepository) UpdateCompanyStatus(ctx context.Context, id int64, status string) error {
	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Update Table users
	var verified bool
	if status == "Approve" {
		verified = true
	} else {
		verified = false
	}
	query := "UPDATE users SET verified = $2, updated_at = current_timestamp WHERE id = $1"
	_, err = tx.ExecContext(ctx, query, id, verified)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Update Table companies
	query = "UPDATE companies SET status = $2, updated_at = current_timestamp WHERE cid = $1"
	_, err = tx.ExecContext(ctx, query, id, status)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Commit the transaction if all insertions were successful
	err = tx.Commit()
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}

func (r *userRepository) DeleteCompany(ctx context.Context, id int64) error {
	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete from Table companies
	query := "DELETE FROM companies WHERE cid = $1"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete from Table users
	query = "DELETE FROM users WHERE id = $1"
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Commit the transaction if all insertions were successful
	err = tx.Commit()
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	return nil
}
