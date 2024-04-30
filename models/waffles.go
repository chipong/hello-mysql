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

// Waffle is an object representing the database table.
type Waffle struct {
	Cat string `boil:"cat" json:"cat" toml:"cat" yaml:"cat"`

	R *waffleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L waffleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WaffleColumns = struct {
	Cat string
}{
	Cat: "cat",
}

var WaffleTableColumns = struct {
	Cat string
}{
	Cat: "waffles.cat",
}

// Generated where

var WaffleWhere = struct {
	Cat whereHelperstring
}{
	Cat: whereHelperstring{field: "`waffles`.`cat`"},
}

// WaffleRels is where relationship names are stored.
var WaffleRels = struct {
}{}

// waffleR is where relationships are stored.
type waffleR struct {
}

// NewStruct creates a new relationship struct
func (*waffleR) NewStruct() *waffleR {
	return &waffleR{}
}

// waffleL is where Load methods for each relationship are stored.
type waffleL struct{}

var (
	waffleAllColumns            = []string{"cat"}
	waffleColumnsWithoutDefault = []string{"cat"}
	waffleColumnsWithDefault    = []string{}
	wafflePrimaryKeyColumns     = []string{"cat"}
	waffleGeneratedColumns      = []string{}
)

type (
	// WaffleSlice is an alias for a slice of pointers to Waffle.
	// This should almost always be used instead of []Waffle.
	WaffleSlice []*Waffle
	// WaffleHook is the signature for custom Waffle hook methods
	WaffleHook func(context.Context, boil.ContextExecutor, *Waffle) error

	waffleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	waffleType                 = reflect.TypeOf(&Waffle{})
	waffleMapping              = queries.MakeStructMapping(waffleType)
	wafflePrimaryKeyMapping, _ = queries.BindMapping(waffleType, waffleMapping, wafflePrimaryKeyColumns)
	waffleInsertCacheMut       sync.RWMutex
	waffleInsertCache          = make(map[string]insertCache)
	waffleUpdateCacheMut       sync.RWMutex
	waffleUpdateCache          = make(map[string]updateCache)
	waffleUpsertCacheMut       sync.RWMutex
	waffleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var waffleAfterSelectMu sync.Mutex
var waffleAfterSelectHooks []WaffleHook

var waffleBeforeInsertMu sync.Mutex
var waffleBeforeInsertHooks []WaffleHook
var waffleAfterInsertMu sync.Mutex
var waffleAfterInsertHooks []WaffleHook

var waffleBeforeUpdateMu sync.Mutex
var waffleBeforeUpdateHooks []WaffleHook
var waffleAfterUpdateMu sync.Mutex
var waffleAfterUpdateHooks []WaffleHook

var waffleBeforeDeleteMu sync.Mutex
var waffleBeforeDeleteHooks []WaffleHook
var waffleAfterDeleteMu sync.Mutex
var waffleAfterDeleteHooks []WaffleHook

var waffleBeforeUpsertMu sync.Mutex
var waffleBeforeUpsertHooks []WaffleHook
var waffleAfterUpsertMu sync.Mutex
var waffleAfterUpsertHooks []WaffleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Waffle) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waffleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Waffle) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waffleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Waffle) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waffleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Waffle) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waffleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Waffle) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waffleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Waffle) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waffleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Waffle) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waffleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Waffle) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waffleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Waffle) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range waffleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWaffleHook registers your hook function for all future operations.
func AddWaffleHook(hookPoint boil.HookPoint, waffleHook WaffleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		waffleAfterSelectMu.Lock()
		waffleAfterSelectHooks = append(waffleAfterSelectHooks, waffleHook)
		waffleAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		waffleBeforeInsertMu.Lock()
		waffleBeforeInsertHooks = append(waffleBeforeInsertHooks, waffleHook)
		waffleBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		waffleAfterInsertMu.Lock()
		waffleAfterInsertHooks = append(waffleAfterInsertHooks, waffleHook)
		waffleAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		waffleBeforeUpdateMu.Lock()
		waffleBeforeUpdateHooks = append(waffleBeforeUpdateHooks, waffleHook)
		waffleBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		waffleAfterUpdateMu.Lock()
		waffleAfterUpdateHooks = append(waffleAfterUpdateHooks, waffleHook)
		waffleAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		waffleBeforeDeleteMu.Lock()
		waffleBeforeDeleteHooks = append(waffleBeforeDeleteHooks, waffleHook)
		waffleBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		waffleAfterDeleteMu.Lock()
		waffleAfterDeleteHooks = append(waffleAfterDeleteHooks, waffleHook)
		waffleAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		waffleBeforeUpsertMu.Lock()
		waffleBeforeUpsertHooks = append(waffleBeforeUpsertHooks, waffleHook)
		waffleBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		waffleAfterUpsertMu.Lock()
		waffleAfterUpsertHooks = append(waffleAfterUpsertHooks, waffleHook)
		waffleAfterUpsertMu.Unlock()
	}
}

// One returns a single waffle record from the query.
func (q waffleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Waffle, error) {
	o := &Waffle{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for waffles")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Waffle records from the query.
func (q waffleQuery) All(ctx context.Context, exec boil.ContextExecutor) (WaffleSlice, error) {
	var o []*Waffle

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Waffle slice")
	}

	if len(waffleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Waffle records in the query.
func (q waffleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count waffles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q waffleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if waffles exists")
	}

	return count > 0, nil
}

// Waffles retrieves all the records using an executor.
func Waffles(mods ...qm.QueryMod) waffleQuery {
	mods = append(mods, qm.From("`waffles`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`waffles`.*"})
	}

	return waffleQuery{q}
}

// FindWaffle retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWaffle(ctx context.Context, exec boil.ContextExecutor, cat string, selectCols ...string) (*Waffle, error) {
	waffleObj := &Waffle{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `waffles` where `cat`=?", sel,
	)

	q := queries.Raw(query, cat)

	err := q.Bind(ctx, exec, waffleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from waffles")
	}

	if err = waffleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return waffleObj, err
	}

	return waffleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Waffle) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no waffles provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(waffleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	waffleInsertCacheMut.RLock()
	cache, cached := waffleInsertCache[key]
	waffleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			waffleAllColumns,
			waffleColumnsWithDefault,
			waffleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(waffleType, waffleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(waffleType, waffleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `waffles` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `waffles` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `waffles` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, wafflePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into waffles")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Cat,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for waffles")
	}

CacheNoHooks:
	if !cached {
		waffleInsertCacheMut.Lock()
		waffleInsertCache[key] = cache
		waffleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Waffle.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Waffle) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	waffleUpdateCacheMut.RLock()
	cache, cached := waffleUpdateCache[key]
	waffleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			waffleAllColumns,
			wafflePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update waffles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `waffles` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, wafflePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(waffleType, waffleMapping, append(wl, wafflePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update waffles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for waffles")
	}

	if !cached {
		waffleUpdateCacheMut.Lock()
		waffleUpdateCache[key] = cache
		waffleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q waffleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for waffles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for waffles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WaffleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wafflePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `waffles` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, wafflePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in waffle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all waffle")
	}
	return rowsAff, nil
}

var mySQLWaffleUniqueColumns = []string{
	"cat",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Waffle) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no waffles provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(waffleColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLWaffleUniqueColumns, o)

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

	waffleUpsertCacheMut.RLock()
	cache, cached := waffleUpsertCache[key]
	waffleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			waffleAllColumns,
			waffleColumnsWithDefault,
			waffleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			waffleAllColumns,
			wafflePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert waffles, could not build update column list")
		}

		ret := strmangle.SetComplement(waffleAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`waffles`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `waffles` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(waffleType, waffleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(waffleType, waffleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for waffles")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(waffleType, waffleMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for waffles")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for waffles")
	}

CacheNoHooks:
	if !cached {
		waffleUpsertCacheMut.Lock()
		waffleUpsertCache[key] = cache
		waffleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Waffle record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Waffle) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Waffle provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), wafflePrimaryKeyMapping)
	sql := "DELETE FROM `waffles` WHERE `cat`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from waffles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for waffles")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q waffleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no waffleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from waffles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for waffles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WaffleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(waffleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wafflePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `waffles` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, wafflePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from waffle slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for waffles")
	}

	if len(waffleAfterDeleteHooks) != 0 {
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
func (o *Waffle) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWaffle(ctx, exec, o.Cat)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WaffleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WaffleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wafflePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `waffles`.* FROM `waffles` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, wafflePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in WaffleSlice")
	}

	*o = slice

	return nil
}

// WaffleExists checks if the Waffle row exists.
func WaffleExists(ctx context.Context, exec boil.ContextExecutor, cat string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `waffles` where `cat`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, cat)
	}
	row := exec.QueryRowContext(ctx, sql, cat)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if waffles exists")
	}

	return exists, nil
}

// Exists checks if the Waffle row exists.
func (o *Waffle) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return WaffleExists(ctx, exec, o.Cat)
}