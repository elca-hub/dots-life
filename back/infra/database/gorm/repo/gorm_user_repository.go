package repo

import (
	"back/domain/model"
	"back/infra/database/gorm/orm"
	"context"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

type transactionKey struct{}

var transactionContextKey = transactionKey{}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (g *GormUserRepository) Create(ctx context.Context, u *model.User) error {
	tx := g.db.WithContext(ctx)
	userOrm := convertToORM(u)

	return gorm.G[orm.User](tx).Create(ctx, userOrm)
}

func (g *GormUserRepository) ExistsEmail(ctx context.Context, e model.Email) (bool, error) {
	res, err := gorm.G[orm.User](g.db).Where("email = ?", e.String()).Count(ctx, "email")

	if err != nil {
		return false, err
	}

	return res > 0, nil
}

func (g *GormUserRepository) Find(ctx context.Context, id string) (*model.User, error) {
	u, err := gorm.G[orm.User](g.db).Where("id = ?", id).First(ctx)

	if err != nil {
		return nil, err
	}

	return convertToDomain(&u), nil
}

func (g *GormUserRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
	tx := g.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := fn(context.WithValue(ctx, transactionContextKey, tx.Statement.Context)); err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit().Error
}

func convertToORM(model *model.User) *orm.User {
	return &orm.User{
		ID:    model.ID(),
		Name:  model.Name().String(),
		Email: model.Email().String(),
	}
}

func convertToDomain(o *orm.User) *model.User {
	email, _ := model.NewEmail(o.Email)
	name, _ := model.NewName(o.Name)

	u, _ := model.NewUser(
		o.ID,
		name,
		email,
		o.CreatedAt,
		o.UpdatedAt,
	)

	return u
}
