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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Elephant is an object representing the database table.
type Elephant struct {
	ID      []byte     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name    []byte     `boil:"name" json:"name" toml:"name" yaml:"name"`
	TigerID null.Bytes `boil:"tiger_id" json:"tiger_id,omitempty" toml:"tiger_id" yaml:"tiger_id,omitempty"`

	R *elephantR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L elephantL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ElephantColumns = struct {
	ID      string
	Name    string
	TigerID string
}{
	ID:      "id",
	Name:    "name",
	TigerID: "tiger_id",
}

var ElephantTableColumns = struct {
	ID      string
	Name    string
	TigerID string
}{
	ID:      "elephants.id",
	Name:    "elephants.name",
	TigerID: "elephants.tiger_id",
}

// Generated where

var ElephantWhere = struct {
	ID      whereHelper__byte
	Name    whereHelper__byte
	TigerID whereHelpernull_Bytes
}{
	ID:      whereHelper__byte{field: "`elephants`.`id`"},
	Name:    whereHelper__byte{field: "`elephants`.`name`"},
	TigerID: whereHelpernull_Bytes{field: "`elephants`.`tiger_id`"},
}

// ElephantRels is where relationship names are stored.
var ElephantRels = struct {
	Tiger string
}{
	Tiger: "Tiger",
}

// elephantR is where relationships are stored.
type elephantR struct {
	Tiger *Tiger `boil:"Tiger" json:"Tiger" toml:"Tiger" yaml:"Tiger"`
}

// NewStruct creates a new relationship struct
func (*elephantR) NewStruct() *elephantR {
	return &elephantR{}
}

func (r *elephantR) GetTiger() *Tiger {
	if r == nil {
		return nil
	}
	return r.Tiger
}

// elephantL is where Load methods for each relationship are stored.
type elephantL struct{}

var (
	elephantAllColumns            = []string{"id", "name", "tiger_id"}
	elephantColumnsWithoutDefault = []string{"id", "name", "tiger_id"}
	elephantColumnsWithDefault    = []string{}
	elephantPrimaryKeyColumns     = []string{"id"}
	elephantGeneratedColumns      = []string{}
)

type (
	// ElephantSlice is an alias for a slice of pointers to Elephant.
	// This should almost always be used instead of []Elephant.
	ElephantSlice []*Elephant
	// ElephantHook is the signature for custom Elephant hook methods
	ElephantHook func(context.Context, boil.ContextExecutor, *Elephant) error

	elephantQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	elephantType                 = reflect.TypeOf(&Elephant{})
	elephantMapping              = queries.MakeStructMapping(elephantType)
	elephantPrimaryKeyMapping, _ = queries.BindMapping(elephantType, elephantMapping, elephantPrimaryKeyColumns)
	elephantInsertCacheMut       sync.RWMutex
	elephantInsertCache          = make(map[string]insertCache)
	elephantUpdateCacheMut       sync.RWMutex
	elephantUpdateCache          = make(map[string]updateCache)
	elephantUpsertCacheMut       sync.RWMutex
	elephantUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var elephantAfterSelectMu sync.Mutex
var elephantAfterSelectHooks []ElephantHook

var elephantBeforeInsertMu sync.Mutex
var elephantBeforeInsertHooks []ElephantHook
var elephantAfterInsertMu sync.Mutex
var elephantAfterInsertHooks []ElephantHook

var elephantBeforeUpdateMu sync.Mutex
var elephantBeforeUpdateHooks []ElephantHook
var elephantAfterUpdateMu sync.Mutex
var elephantAfterUpdateHooks []ElephantHook

var elephantBeforeDeleteMu sync.Mutex
var elephantBeforeDeleteHooks []ElephantHook
var elephantAfterDeleteMu sync.Mutex
var elephantAfterDeleteHooks []ElephantHook

var elephantBeforeUpsertMu sync.Mutex
var elephantBeforeUpsertHooks []ElephantHook
var elephantAfterUpsertMu sync.Mutex
var elephantAfterUpsertHooks []ElephantHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Elephant) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range elephantAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Elephant) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range elephantBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Elephant) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range elephantAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Elephant) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range elephantBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Elephant) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range elephantAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Elephant) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range elephantBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Elephant) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range elephantAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Elephant) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range elephantBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Elephant) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range elephantAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddElephantHook registers your hook function for all future operations.
func AddElephantHook(hookPoint boil.HookPoint, elephantHook ElephantHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		elephantAfterSelectMu.Lock()
		elephantAfterSelectHooks = append(elephantAfterSelectHooks, elephantHook)
		elephantAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		elephantBeforeInsertMu.Lock()
		elephantBeforeInsertHooks = append(elephantBeforeInsertHooks, elephantHook)
		elephantBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		elephantAfterInsertMu.Lock()
		elephantAfterInsertHooks = append(elephantAfterInsertHooks, elephantHook)
		elephantAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		elephantBeforeUpdateMu.Lock()
		elephantBeforeUpdateHooks = append(elephantBeforeUpdateHooks, elephantHook)
		elephantBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		elephantAfterUpdateMu.Lock()
		elephantAfterUpdateHooks = append(elephantAfterUpdateHooks, elephantHook)
		elephantAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		elephantBeforeDeleteMu.Lock()
		elephantBeforeDeleteHooks = append(elephantBeforeDeleteHooks, elephantHook)
		elephantBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		elephantAfterDeleteMu.Lock()
		elephantAfterDeleteHooks = append(elephantAfterDeleteHooks, elephantHook)
		elephantAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		elephantBeforeUpsertMu.Lock()
		elephantBeforeUpsertHooks = append(elephantBeforeUpsertHooks, elephantHook)
		elephantBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		elephantAfterUpsertMu.Lock()
		elephantAfterUpsertHooks = append(elephantAfterUpsertHooks, elephantHook)
		elephantAfterUpsertMu.Unlock()
	}
}

// One returns a single elephant record from the query.
func (q elephantQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Elephant, error) {
	o := &Elephant{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for elephants")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Elephant records from the query.
func (q elephantQuery) All(ctx context.Context, exec boil.ContextExecutor) (ElephantSlice, error) {
	var o []*Elephant

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Elephant slice")
	}

	if len(elephantAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Elephant records in the query.
func (q elephantQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count elephants rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q elephantQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if elephants exists")
	}

	return count > 0, nil
}

// Tiger pointed to by the foreign key.
func (o *Elephant) Tiger(mods ...qm.QueryMod) tigerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.TigerID),
	}

	queryMods = append(queryMods, mods...)

	return Tigers(queryMods...)
}

// LoadTiger allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (elephantL) LoadTiger(ctx context.Context, e boil.ContextExecutor, singular bool, maybeElephant interface{}, mods queries.Applicator) error {
	var slice []*Elephant
	var object *Elephant

	if singular {
		var ok bool
		object, ok = maybeElephant.(*Elephant)
		if !ok {
			object = new(Elephant)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeElephant)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeElephant))
			}
		}
	} else {
		s, ok := maybeElephant.(*[]*Elephant)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeElephant)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeElephant))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &elephantR{}
		}
		if !queries.IsNil(object.TigerID) {
			args[object.TigerID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &elephantR{}
			}

			if !queries.IsNil(obj.TigerID) {
				args[obj.TigerID] = struct{}{}
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`tigers`),
		qm.WhereIn(`tigers.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Tiger")
	}

	var resultSlice []*Tiger
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Tiger")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tigers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tigers")
	}

	if len(tigerAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Tiger = foreign
		if foreign.R == nil {
			foreign.R = &tigerR{}
		}
		foreign.R.Elephant = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.TigerID, foreign.ID) {
				local.R.Tiger = foreign
				if foreign.R == nil {
					foreign.R = &tigerR{}
				}
				foreign.R.Elephant = local
				break
			}
		}
	}

	return nil
}

// SetTiger of the elephant to the related item.
// Sets o.R.Tiger to related.
// Adds o to related.R.Elephant.
func (o *Elephant) SetTiger(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Tiger) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `elephants` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"tiger_id"}),
		strmangle.WhereClause("`", "`", 0, elephantPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.TigerID, related.ID)
	if o.R == nil {
		o.R = &elephantR{
			Tiger: related,
		}
	} else {
		o.R.Tiger = related
	}

	if related.R == nil {
		related.R = &tigerR{
			Elephant: o,
		}
	} else {
		related.R.Elephant = o
	}

	return nil
}

// RemoveTiger relationship.
// Sets o.R.Tiger to nil.
// Removes o from all passed in related items' relationships struct.
func (o *Elephant) RemoveTiger(ctx context.Context, exec boil.ContextExecutor, related *Tiger) error {
	var err error

	queries.SetScanner(&o.TigerID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("tiger_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Tiger = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	related.R.Elephant = nil
	return nil
}

// Elephants retrieves all the records using an executor.
func Elephants(mods ...qm.QueryMod) elephantQuery {
	mods = append(mods, qm.From("`elephants`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`elephants`.*"})
	}

	return elephantQuery{q}
}

// FindElephant retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindElephant(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*Elephant, error) {
	elephantObj := &Elephant{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `elephants` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, elephantObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from elephants")
	}

	if err = elephantObj.doAfterSelectHooks(ctx, exec); err != nil {
		return elephantObj, err
	}

	return elephantObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Elephant) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no elephants provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(elephantColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	elephantInsertCacheMut.RLock()
	cache, cached := elephantInsertCache[key]
	elephantInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			elephantAllColumns,
			elephantColumnsWithDefault,
			elephantColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(elephantType, elephantMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(elephantType, elephantMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `elephants` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `elephants` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `elephants` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, elephantPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into elephants")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for elephants")
	}

CacheNoHooks:
	if !cached {
		elephantInsertCacheMut.Lock()
		elephantInsertCache[key] = cache
		elephantInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Elephant.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Elephant) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	elephantUpdateCacheMut.RLock()
	cache, cached := elephantUpdateCache[key]
	elephantUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			elephantAllColumns,
			elephantPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update elephants, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `elephants` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, elephantPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(elephantType, elephantMapping, append(wl, elephantPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update elephants row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for elephants")
	}

	if !cached {
		elephantUpdateCacheMut.Lock()
		elephantUpdateCache[key] = cache
		elephantUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q elephantQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for elephants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for elephants")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ElephantSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), elephantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `elephants` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, elephantPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in elephant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all elephant")
	}
	return rowsAff, nil
}

var mySQLElephantUniqueColumns = []string{
	"id",
	"tiger_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Elephant) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no elephants provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(elephantColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLElephantUniqueColumns, o)

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

	elephantUpsertCacheMut.RLock()
	cache, cached := elephantUpsertCache[key]
	elephantUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			elephantAllColumns,
			elephantColumnsWithDefault,
			elephantColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			elephantAllColumns,
			elephantPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert elephants, could not build update column list")
		}

		ret := strmangle.SetComplement(elephantAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`elephants`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `elephants` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(elephantType, elephantMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(elephantType, elephantMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for elephants")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(elephantType, elephantMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for elephants")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for elephants")
	}

CacheNoHooks:
	if !cached {
		elephantUpsertCacheMut.Lock()
		elephantUpsertCache[key] = cache
		elephantUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Elephant record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Elephant) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Elephant provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), elephantPrimaryKeyMapping)
	sql := "DELETE FROM `elephants` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from elephants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for elephants")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q elephantQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no elephantQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from elephants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for elephants")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ElephantSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(elephantBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), elephantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `elephants` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, elephantPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from elephant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for elephants")
	}

	if len(elephantAfterDeleteHooks) != 0 {
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
func (o *Elephant) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindElephant(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ElephantSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ElephantSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), elephantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `elephants`.* FROM `elephants` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, elephantPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ElephantSlice")
	}

	*o = slice

	return nil
}

// ElephantExists checks if the Elephant row exists.
func ElephantExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `elephants` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if elephants exists")
	}

	return exists, nil
}

// Exists checks if the Elephant row exists.
func (o *Elephant) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ElephantExists(ctx, exec, o.ID)
}
