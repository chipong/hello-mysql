// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Chocolate is an object representing the database table.
type Chocolate struct {
	Dog string `boil:"dog" json:"dog" toml:"dog" yaml:"dog"`

	R *chocolateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L chocolateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChocolateColumns = struct {
	Dog string
}{
	Dog: "dog",
}

var ChocolateTableColumns = struct {
	Dog string
}{
	Dog: "chocolate.dog",
}

// Generated where

var ChocolateWhere = struct {
	Dog whereHelperstring
}{
	Dog: whereHelperstring{field: "`chocolate`.`dog`"},
}

// ChocolateRels is where relationship names are stored.
var ChocolateRels = struct {
}{}

// chocolateR is where relationships are stored.
type chocolateR struct {
}

// NewStruct creates a new relationship struct
func (*chocolateR) NewStruct() *chocolateR {
	return &chocolateR{}
}

// chocolateL is where Load methods for each relationship are stored.
type chocolateL struct{}

var (
	chocolateAllColumns            = []string{"dog"}
	chocolateColumnsWithoutDefault = []string{"dog"}
	chocolateColumnsWithDefault    = []string{}
	chocolatePrimaryKeyColumns     = []string{"dog"}
	chocolateGeneratedColumns      = []string{}
)

type (
	// ChocolateSlice is an alias for a slice of pointers to Chocolate.
	// This should almost always be used instead of []Chocolate.
	ChocolateSlice []*Chocolate
	// ChocolateHook is the signature for custom Chocolate hook methods
	ChocolateHook func(context.Context, boil.ContextExecutor, *Chocolate) error

	chocolateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	chocolateType                 = reflect.TypeOf(&Chocolate{})
	chocolateMapping              = queries.MakeStructMapping(chocolateType)
	chocolatePrimaryKeyMapping, _ = queries.BindMapping(chocolateType, chocolateMapping, chocolatePrimaryKeyColumns)
	chocolateInsertCacheMut       sync.RWMutex
	chocolateInsertCache          = make(map[string]insertCache)
	chocolateUpdateCacheMut       sync.RWMutex
	chocolateUpdateCache          = make(map[string]updateCache)
	chocolateUpsertCacheMut       sync.RWMutex
	chocolateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var chocolateAfterSelectMu sync.Mutex
var chocolateAfterSelectHooks []ChocolateHook

var chocolateBeforeInsertMu sync.Mutex
var chocolateBeforeInsertHooks []ChocolateHook
var chocolateAfterInsertMu sync.Mutex
var chocolateAfterInsertHooks []ChocolateHook

var chocolateBeforeUpdateMu sync.Mutex
var chocolateBeforeUpdateHooks []ChocolateHook
var chocolateAfterUpdateMu sync.Mutex
var chocolateAfterUpdateHooks []ChocolateHook

var chocolateBeforeDeleteMu sync.Mutex
var chocolateBeforeDeleteHooks []ChocolateHook
var chocolateAfterDeleteMu sync.Mutex
var chocolateAfterDeleteHooks []ChocolateHook

var chocolateBeforeUpsertMu sync.Mutex
var chocolateBeforeUpsertHooks []ChocolateHook
var chocolateAfterUpsertMu sync.Mutex
var chocolateAfterUpsertHooks []ChocolateHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Chocolate) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chocolateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Chocolate) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chocolateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Chocolate) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chocolateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Chocolate) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chocolateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Chocolate) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chocolateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Chocolate) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chocolateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Chocolate) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chocolateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Chocolate) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chocolateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Chocolate) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range chocolateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChocolateHook registers your hook function for all future operations.
func AddChocolateHook(hookPoint boil.HookPoint, chocolateHook ChocolateHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		chocolateAfterSelectMu.Lock()
		chocolateAfterSelectHooks = append(chocolateAfterSelectHooks, chocolateHook)
		chocolateAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		chocolateBeforeInsertMu.Lock()
		chocolateBeforeInsertHooks = append(chocolateBeforeInsertHooks, chocolateHook)
		chocolateBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		chocolateAfterInsertMu.Lock()
		chocolateAfterInsertHooks = append(chocolateAfterInsertHooks, chocolateHook)
		chocolateAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		chocolateBeforeUpdateMu.Lock()
		chocolateBeforeUpdateHooks = append(chocolateBeforeUpdateHooks, chocolateHook)
		chocolateBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		chocolateAfterUpdateMu.Lock()
		chocolateAfterUpdateHooks = append(chocolateAfterUpdateHooks, chocolateHook)
		chocolateAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		chocolateBeforeDeleteMu.Lock()
		chocolateBeforeDeleteHooks = append(chocolateBeforeDeleteHooks, chocolateHook)
		chocolateBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		chocolateAfterDeleteMu.Lock()
		chocolateAfterDeleteHooks = append(chocolateAfterDeleteHooks, chocolateHook)
		chocolateAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		chocolateBeforeUpsertMu.Lock()
		chocolateBeforeUpsertHooks = append(chocolateBeforeUpsertHooks, chocolateHook)
		chocolateBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		chocolateAfterUpsertMu.Lock()
		chocolateAfterUpsertHooks = append(chocolateAfterUpsertHooks, chocolateHook)
		chocolateAfterUpsertMu.Unlock()
	}
}

// One returns a single chocolate record from the query.
func (q chocolateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Chocolate, error) {
	o := &Chocolate{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for chocolate")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Chocolate records from the query.
func (q chocolateQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChocolateSlice, error) {
	var o []*Chocolate

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Chocolate slice")
	}

	if len(chocolateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Chocolate records in the query.
func (q chocolateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count chocolate rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q chocolateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if chocolate exists")
	}

	return count > 0, nil
}

// Chocolates retrieves all the records using an executor.
func Chocolates(mods ...qm.QueryMod) chocolateQuery {
	mods = append(mods, qm.From("`chocolate`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`chocolate`.*"})
	}

	return chocolateQuery{q}
}

// FindChocolate retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChocolate(ctx context.Context, exec boil.ContextExecutor, dog string, selectCols ...string) (*Chocolate, error) {
	chocolateObj := &Chocolate{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `chocolate` where `dog`=?", sel,
	)

	q := queries.Raw(query, dog)

	err := q.Bind(ctx, exec, chocolateObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from chocolate")
	}

	if err = chocolateObj.doAfterSelectHooks(ctx, exec); err != nil {
		return chocolateObj, err
	}

	return chocolateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Chocolate) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no chocolate provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chocolateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	chocolateInsertCacheMut.RLock()
	cache, cached := chocolateInsertCache[key]
	chocolateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			chocolateAllColumns,
			chocolateColumnsWithDefault,
			chocolateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(chocolateType, chocolateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(chocolateType, chocolateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `chocolate` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `chocolate` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `chocolate` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, chocolatePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into chocolate")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Dog,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for chocolate")
	}

CacheNoHooks:
	if !cached {
		chocolateInsertCacheMut.Lock()
		chocolateInsertCache[key] = cache
		chocolateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Chocolate.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Chocolate) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	chocolateUpdateCacheMut.RLock()
	cache, cached := chocolateUpdateCache[key]
	chocolateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			chocolateAllColumns,
			chocolatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update chocolate, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `chocolate` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, chocolatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(chocolateType, chocolateMapping, append(wl, chocolatePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update chocolate row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for chocolate")
	}

	if !cached {
		chocolateUpdateCacheMut.Lock()
		chocolateUpdateCache[key] = cache
		chocolateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q chocolateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for chocolate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for chocolate")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChocolateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chocolatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `chocolate` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, chocolatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in chocolate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all chocolate")
	}
	return rowsAff, nil
}

var mySQLChocolateUniqueColumns = []string{
	"dog",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Chocolate) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no chocolate provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(chocolateColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLChocolateUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	chocolateUpsertCacheMut.RLock()
	cache, cached := chocolateUpsertCache[key]
	chocolateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			chocolateAllColumns,
			chocolateColumnsWithDefault,
			chocolateColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			chocolateAllColumns,
			chocolatePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert chocolate, could not build update column list")
		}

		ret := strmangle.SetComplement(chocolateAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`chocolate`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `chocolate` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(chocolateType, chocolateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(chocolateType, chocolateMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for chocolate")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(chocolateType, chocolateMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for chocolate")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for chocolate")
	}

CacheNoHooks:
	if !cached {
		chocolateUpsertCacheMut.Lock()
		chocolateUpsertCache[key] = cache
		chocolateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Chocolate record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Chocolate) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Chocolate provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), chocolatePrimaryKeyMapping)
	sql := "DELETE FROM `chocolate` WHERE `dog`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from chocolate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for chocolate")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q chocolateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no chocolateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from chocolate")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for chocolate")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChocolateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(chocolateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chocolatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `chocolate` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, chocolatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from chocolate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for chocolate")
	}

	if len(chocolateAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Chocolate) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChocolate(ctx, exec, o.Dog)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChocolateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChocolateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), chocolatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `chocolate`.* FROM `chocolate` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, chocolatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChocolateSlice")
	}

	*o = slice

	return nil
}

// ChocolateExists checks if the Chocolate row exists.
func ChocolateExists(ctx context.Context, exec boil.ContextExecutor, dog string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `chocolate` where `dog`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, dog)
	}
	row := exec.QueryRowContext(ctx, sql, dog)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if chocolate exists")
	}

	return exists, nil
}

// Exists checks if the Chocolate row exists.
func (o *Chocolate) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ChocolateExists(ctx, exec, o.Dog)
}
