// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"test/dal/model"
)

func newProduct(db *gorm.DB, opts ...gen.DOOption) product {
	_product := product{}

	_product.productDo.UseDB(db, opts...)
	_product.productDo.UseModel(&model.Product{})

	tableName := _product.productDo.TableName()
	_product.ALL = field.NewAsterisk(tableName)
	_product.ID = field.NewInt64(tableName, "id")
	_product.CreatedAt = field.NewTime(tableName, "created_at")
	_product.UpdatedAt = field.NewTime(tableName, "updated_at")
	_product.DeletedAt = field.NewField(tableName, "deleted_at")
	_product.Code = field.NewString(tableName, "code")
	_product.Price = field.NewInt64(tableName, "price")

	_product.fillFieldMap()

	return _product
}

type product struct {
	productDo productDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Code      field.String
	Price     field.Int64

	fieldMap map[string]field.Expr
}

func (p product) Table(newTableName string) *product {
	p.productDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p product) As(alias string) *product {
	p.productDo.DO = *(p.productDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *product) updateTableName(table string) *product {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Code = field.NewString(table, "code")
	p.Price = field.NewInt64(table, "price")

	p.fillFieldMap()

	return p
}

func (p *product) WithContext(ctx context.Context) IProductDo { return p.productDo.WithContext(ctx) }

func (p product) TableName() string { return p.productDo.TableName() }

func (p product) Alias() string { return p.productDo.Alias() }

func (p product) Columns(cols ...field.Expr) gen.Columns { return p.productDo.Columns(cols...) }

func (p *product) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *product) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["code"] = p.Code
	p.fieldMap["price"] = p.Price
}

func (p product) clone(db *gorm.DB) product {
	p.productDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p product) replaceDB(db *gorm.DB) product {
	p.productDo.ReplaceDB(db)
	return p
}

type productDo struct{ gen.DO }

type IProductDo interface {
	gen.SubQuery
	Debug() IProductDo
	WithContext(ctx context.Context) IProductDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProductDo
	WriteDB() IProductDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProductDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProductDo
	Not(conds ...gen.Condition) IProductDo
	Or(conds ...gen.Condition) IProductDo
	Select(conds ...field.Expr) IProductDo
	Where(conds ...gen.Condition) IProductDo
	Order(conds ...field.Expr) IProductDo
	Distinct(cols ...field.Expr) IProductDo
	Omit(cols ...field.Expr) IProductDo
	Join(table schema.Tabler, on ...field.Expr) IProductDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProductDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProductDo
	Group(cols ...field.Expr) IProductDo
	Having(conds ...gen.Condition) IProductDo
	Limit(limit int) IProductDo
	Offset(offset int) IProductDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProductDo
	Unscoped() IProductDo
	Create(values ...*model.Product) error
	CreateInBatches(values []*model.Product, batchSize int) error
	Save(values ...*model.Product) error
	First() (*model.Product, error)
	Take() (*model.Product, error)
	Last() (*model.Product, error)
	Find() ([]*model.Product, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Product, err error)
	FindInBatches(result *[]*model.Product, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Product) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProductDo
	Assign(attrs ...field.AssignExpr) IProductDo
	Joins(fields ...field.RelationField) IProductDo
	Preload(fields ...field.RelationField) IProductDo
	FirstOrInit() (*model.Product, error)
	FirstOrCreate() (*model.Product, error)
	FindByPage(offset int, limit int) (result []*model.Product, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProductDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p productDo) Debug() IProductDo {
	return p.withDO(p.DO.Debug())
}

func (p productDo) WithContext(ctx context.Context) IProductDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p productDo) ReadDB() IProductDo {
	return p.Clauses(dbresolver.Read)
}

func (p productDo) WriteDB() IProductDo {
	return p.Clauses(dbresolver.Write)
}

func (p productDo) Session(config *gorm.Session) IProductDo {
	return p.withDO(p.DO.Session(config))
}

func (p productDo) Clauses(conds ...clause.Expression) IProductDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p productDo) Returning(value interface{}, columns ...string) IProductDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p productDo) Not(conds ...gen.Condition) IProductDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p productDo) Or(conds ...gen.Condition) IProductDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p productDo) Select(conds ...field.Expr) IProductDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p productDo) Where(conds ...gen.Condition) IProductDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p productDo) Order(conds ...field.Expr) IProductDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p productDo) Distinct(cols ...field.Expr) IProductDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p productDo) Omit(cols ...field.Expr) IProductDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p productDo) Join(table schema.Tabler, on ...field.Expr) IProductDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p productDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProductDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p productDo) RightJoin(table schema.Tabler, on ...field.Expr) IProductDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p productDo) Group(cols ...field.Expr) IProductDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p productDo) Having(conds ...gen.Condition) IProductDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p productDo) Limit(limit int) IProductDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p productDo) Offset(offset int) IProductDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p productDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProductDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p productDo) Unscoped() IProductDo {
	return p.withDO(p.DO.Unscoped())
}

func (p productDo) Create(values ...*model.Product) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p productDo) CreateInBatches(values []*model.Product, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p productDo) Save(values ...*model.Product) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p productDo) First() (*model.Product, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) Take() (*model.Product, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) Last() (*model.Product, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) Find() ([]*model.Product, error) {
	result, err := p.DO.Find()
	return result.([]*model.Product), err
}

func (p productDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Product, err error) {
	buf := make([]*model.Product, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p productDo) FindInBatches(result *[]*model.Product, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p productDo) Attrs(attrs ...field.AssignExpr) IProductDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p productDo) Assign(attrs ...field.AssignExpr) IProductDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p productDo) Joins(fields ...field.RelationField) IProductDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p productDo) Preload(fields ...field.RelationField) IProductDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p productDo) FirstOrInit() (*model.Product, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) FirstOrCreate() (*model.Product, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Product), nil
	}
}

func (p productDo) FindByPage(offset int, limit int) (result []*model.Product, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p productDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p productDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p productDo) Delete(models ...*model.Product) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *productDo) withDO(do gen.Dao) *productDo {
	p.DO = *do.(*gen.DO)
	return p
}
