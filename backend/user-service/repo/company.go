package repo

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/lib/pq"
)

func (r *userRepository) CreateCompany(ctx context.Context, company *pbv1.CreateCompanyRequest, createTime int64) (int64, error) {
	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into Table users
	query := "INSERT INTO users (email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int64
	err = tx.QueryRowContext(ctx, query, company.Email, company.Password, domain.CompanyRole, createTime, createTime).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Insert into Table companies
	query = "INSERT INTO companies (cid, name, description, location, phone, category, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err = tx.ExecContext(ctx, query, id, company.Name, company.Description, company.Location, company.Phone, company.Category, createTime, createTime)
	if err != nil {
		tx.Rollback()
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	// Commit the transaction if all insertions were successful
	err = tx.Commit()
	if err != nil {
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	return id, nil
}

func (r *userRepository) GetCompanyByID(ctx context.Context, id int64) (*pbv1.Company, error) {
	query := "SELECT users.id, companies.name, users.email, companies.description, companies.location, companies.phone, companies.category, companies.status FROM users INNER JOIN companies ON users.id = companies.cid WHERE users.id = $1"
	var company pbv1.Company
	err := r.db.QueryRowContext(ctx, query, id).Scan(&company.Id, &company.Name, &company.Email, &company.Description, &company.Location, &company.Phone, &company.Category, &company.Status)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, domain.ErrUserIDNotFound.From(err.Error(), err)
	}
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
	parts := strings.Fields(search)
	tquery := strings.Join(parts, " | ")
	var rows *sql.Rows
	var err error
	if search == "" {
		query := `
			SELECT 
			users.id, companies.name, users.email, companies.description, companies.location, companies.phone, companies.category 
			FROM 
				users INNER JOIN companies ON users.id = companies.cid
			WHERE 
				users.verified = true 
				AND companies.status = 'Approve'
		`
		rows, err = r.db.QueryContext(ctx, query)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
	} else {
		query :=
			`	SELECT 
				users.id, companies.name, users.email, companies.description, companies.location, companies.phone, companies.category 
				FROM 
					users INNER JOIN companies ON users.id = companies.cid, 
					to_tsvector(companies.category || companies.name) document,
					to_tsquery($1) query,
					NULLIF(ts_rank(to_tsvector(companies.category), query), 0) rank_category,
					NULLIF(ts_rank(to_tsvector(companies.name), query), 0) rank_name,
					SIMILARITY($2, companies.category || companies.name) similarity
				WHERE 
					users.verified = true 
					AND query @@ document 
					OR similarity > 0
					AND companies.status = 'Approve'
				ORDER BY rank_category DESC, rank_name DESC, similarity DESC NULLS LAST
			`
		rows, err = r.db.QueryContext(ctx, query, tquery, search)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
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

func (r *userRepository) UpdateCompanyByID(ctx context.Context, id int64, req *pbv1.UpdatedCompany) error {
	current_timestamp := time.Now().Unix()
	query := "UPDATE companies SET name = $1, description = $2, location = $3, phone = $4, category = $5, updated_at = $6 WHERE cid = $7"
	_, err := r.db.ExecContext(ctx, query, req.Name, req.Description, req.Location, req.Phone, req.Category, current_timestamp, id)
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
	if status == domain.ComapanyStatusApprove {
		verified = true
	} else {
		verified = false
	}
	current_timestamp := time.Now().Unix()
	query := "UPDATE users SET verified = $2, updated_at = $3 WHERE id = $1"
	_, err = tx.ExecContext(ctx, query, id, verified, current_timestamp)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Update Table companies
	query = "UPDATE companies SET status = $2, updated_at = $3 WHERE cid = $1"
	_, err = tx.ExecContext(ctx, query, id, status, current_timestamp)
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

func (r *userRepository) DeleteCompanies(ctx context.Context) error {
	// Start a transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete from Table companies
	query := "DELETE FROM companies WHERE cid > 0"
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return domain.ErrInternal.From(err.Error(), err)
	}

	// Delete from Table users
	query = "DELETE FROM users WHERE role = 'company'"
	_, err = tx.ExecContext(ctx, query)
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

func (r *userRepository) GetCompanies(ctx context.Context, ids []int64) ([]*pbv1.CompanyInfo, error) {
	query := "SELECT users.id, companies.name FROM users INNER JOIN companies ON users.id = companies.cid WHERE users.id = ANY($1)"
	rows, err := r.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, domain.ErrInternal.From(err.Error(), err)
	}
	defer rows.Close()

	var companies []*pbv1.CompanyInfo
	for rows.Next() {
		var company pbv1.CompanyInfo
		err := rows.Scan(&company.Id, &company.Name)
		if err != nil {
			return nil, domain.ErrInternal.From(err.Error(), err)
		}
		companies = append(companies, &company)
	}

	return companies, nil
}
