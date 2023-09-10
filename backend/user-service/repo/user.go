package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	"github.com/TikhampornSky/go-auth-verifiedMail/initializers"
	"github.com/TikhampornSky/go-auth-verifiedMail/port"
	"github.com/redis/go-redis/v9"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
}

type TX interface {
	Commit() error
	Rollback() error
}

type Redis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type userRepository struct {
	db    DBTX
	redis Redis
}

func NewUserRepository(db DBTX, redis Redis) port.UserRepoPort {
	return &userRepository{db: db, redis: redis}
}

func (r *userRepository) CreateAdmin(ctx context.Context, admin *pbv1.CreateAdminRequest, createTime int64) (int64, error) {
	query := "INSERT INTO users (email, password, role, verified, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	var id int64
	err := r.db.QueryRowContext(ctx, query, admin.Email, admin.Password, "admin", true, createTime, createTime).Scan(&id)
	if err != nil {
		return 0, domain.ErrInternal.From(err.Error(), err)
	}

	return id, nil
}

func (r *userRepository) CheckEmailExist(ctx context.Context, email string) error {
	query := "SELECT id FROM users WHERE email = $1"
	var id int64
	err := r.db.QueryRowContext(ctx, query, email).Scan(&id)
	if err != sql.ErrNoRows && err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}
	if id != 0 {
		return domain.ErrDuplicateEmail.With("email already exists")
	}
	return nil
}

func (r *userRepository) GetPassword(ctx context.Context, req *pbv1.LoginRequest) (int64, string, int64, error) {
	queryVerify := "SELECT id, verified, password, created_at FROM users WHERE email = $1;"
	var id int64
	var verified bool
	var password string
	var created_at int64
	err := r.db.QueryRowContext(ctx, queryVerify, req.Email).Scan(&id, &verified, &password, &created_at)
	if err != nil {
		return 0, "", 0, domain.ErrInternal.From(err.Error(), err)
	}
	if !verified {
		return 0, "", 0, domain.ErrNotVerified.With("user not verified")
	}

	return id, password, created_at, nil
}

func (r *userRepository) CheckUserIDExist(ctx context.Context, id int64) (string, error) {
	query := "SELECT id, role FROM users WHERE id = $1"
	var idUser int64
	var role string
	err := r.db.QueryRowContext(ctx, query, id).Scan(&idUser, &role)
	if err != nil {
		return role, domain.ErrInternal.From(err.Error(), err)
	}
	return role, nil
}

func (r *userRepository) CheckIfAdmin(ctx context.Context, id int64) error {
	query := "SELECT id FROM users WHERE id = $1 AND role = 'admin'"
	var idUser int64
	err := r.db.QueryRowContext(ctx, query, id).Scan(&idUser)
	if err == sql.ErrNoRows {
		return domain.ErrNotAuthorized.With("user not admin")
	}
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}
	return nil
}

// Redis Zone //
func (r *userRepository) SetValueRedis(ctx context.Context, key string, value string) error {
	config, _ := initializers.LoadConfig("..")
	err := r.redis.Set(ctx, key, value, time.Duration(config.REDISTimeout)*time.Minute).Err()
	if err != nil {
		return domain.ErrInternal.From(err.Error(), err)
	}
	return nil
}

func (r *userRepository) GetValueRedis(ctx context.Context, key string) (string, error) {
	value, err := r.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", domain.ErrRedisNotFound.With("key %s not found", key)
	}
	if err != nil {
		return "", domain.ErrInternal.From(err.Error(), err)
	}
	return value, nil
}
