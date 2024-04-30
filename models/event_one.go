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

// EventOne is an object representing the database table.
type EventOne struct {
	ID   uint64          `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name null.String     `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Day  EventOneNullDay `boil:"day" json:"day,omitempty" toml:"day" yaml:"day,omitempty"`

	R *eventOneR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L eventOneL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EventOneColumns = struct {
	ID   string
	Name string
	Day  string
}{
	ID:   "id",
	Name: "name",
	Day:  "day",
}

var EventOneTableColumns = struct {
	ID   string
	Name string
	Day  string
}{
	ID:   "event_one.id",
	Name: "event_one.name",
	Day:  "event_one.day",
}

// Generated where

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint64) IN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint64) NIN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperEventOneNullDay struct{ field string }

func (w whereHelperEventOneNullDay) EQ(x EventOneNullDay) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelperEventOneNullDay) NEQ(x EventOneNullDay) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelperEventOneNullDay) LT(x EventOneNullDay) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperEventOneNullDay) LTE(x EventOneNullDay) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperEventOneNullDay) GT(x EventOneNullDay) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperEventOneNullDay) GTE(x EventOneNullDay) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperEventOneNullDay) IN(slice []EventOneNullDay) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperEventOneNullDay) NIN(slice []EventOneNullDay) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelperEventOneNullDay) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelperEventOneNullDay) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var EventOneWhere = struct {
	ID   whereHelperuint64
	Name whereHelpernull_String
	Day  whereHelperEventOneNullDay
}{
	ID:   whereHelperuint64{field: "`event_one`.`id`"},
	Name: whereHelpernull_String{field: "`event_one`.`name`"},
	Day:  whereHelperEventOneNullDay{field: "`event_one`.`day`"},
}

// EventOneRels is where relationship names are stored.
var EventOneRels = struct {
}{}

// eventOneR is where relationships are stored.
type eventOneR struct {
}

// NewStruct creates a new relationship struct
func (*eventOneR) NewStruct() *eventOneR {
	return &eventOneR{}
}

// eventOneL is where Load methods for each relationship are stored.
type eventOneL struct{}

var (
	eventOneAllColumns            = []string{"id", "name", "day"}
	eventOneColumnsWithoutDefault = []string{"name", "day"}
	eventOneColumnsWithDefault    = []string{"id"}
	eventOnePrimaryKeyColumns     = []string{"id"}
	eventOneGeneratedColumns      = []string{}
)

type (
	// EventOneSlice is an alias for a slice of pointers to EventOne.
	// This should almost always be used instead of []EventOne.
	EventOneSlice []*EventOne
	// EventOneHook is the signature for custom EventOne hook methods
	EventOneHook func(context.Context, boil.ContextExecutor, *EventOne) error

	eventOneQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	eventOneType                 = reflect.TypeOf(&EventOne{})
	eventOneMapping              = queries.MakeStructMapping(eventOneType)
	eventOnePrimaryKeyMapping, _ = queries.BindMapping(eventOneType, eventOneMapping, eventOnePrimaryKeyColumns)
	eventOneInsertCacheMut       sync.RWMutex
	eventOneInsertCache          = make(map[string]insertCache)
	eventOneUpdateCacheMut       sync.RWMutex
	eventOneUpdateCache          = make(map[string]updateCache)
	eventOneUpsertCacheMut       sync.RWMutex
	eventOneUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var eventOneAfterSelectMu sync.Mutex
var eventOneAfterSelectHooks []EventOneHook

var eventOneBeforeInsertMu sync.Mutex
var eventOneBeforeInsertHooks []EventOneHook
var eventOneAfterInsertMu sync.Mutex
var eventOneAfterInsertHooks []EventOneHook

var eventOneBeforeUpdateMu sync.Mutex
var eventOneBeforeUpdateHooks []EventOneHook
var eventOneAfterUpdateMu sync.Mutex
var eventOneAfterUpdateHooks []EventOneHook

var eventOneBeforeDeleteMu sync.Mutex
var eventOneBeforeDeleteHooks []EventOneHook
var eventOneAfterDeleteMu sync.Mutex
var eventOneAfterDeleteHooks []EventOneHook

var eventOneBeforeUpsertMu sync.Mutex
var eventOneBeforeUpsertHooks []EventOneHook
var eventOneAfterUpsertMu sync.Mutex
var eventOneAfterUpsertHooks []EventOneHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EventOne) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventOneAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EventOne) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventOneBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EventOne) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventOneAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EventOne) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventOneBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EventOne) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventOneAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EventOne) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventOneBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EventOne) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventOneAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EventOne) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventOneBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EventOne) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventOneAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEventOneHook registers your hook function for all future operations.
func AddEventOneHook(hookPoint boil.HookPoint, eventOneHook EventOneHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		eventOneAfterSelectMu.Lock()
		eventOneAfterSelectHooks = append(eventOneAfterSelectHooks, eventOneHook)
		eventOneAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		eventOneBeforeInsertMu.Lock()
		eventOneBeforeInsertHooks = append(eventOneBeforeInsertHooks, eventOneHook)
		eventOneBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		eventOneAfterInsertMu.Lock()
		eventOneAfterInsertHooks = append(eventOneAfterInsertHooks, eventOneHook)
		eventOneAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		eventOneBeforeUpdateMu.Lock()
		eventOneBeforeUpdateHooks = append(eventOneBeforeUpdateHooks, eventOneHook)
		eventOneBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		eventOneAfterUpdateMu.Lock()
		eventOneAfterUpdateHooks = append(eventOneAfterUpdateHooks, eventOneHook)
		eventOneAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		eventOneBeforeDeleteMu.Lock()
		eventOneBeforeDeleteHooks = append(eventOneBeforeDeleteHooks, eventOneHook)
		eventOneBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		eventOneAfterDeleteMu.Lock()
		eventOneAfterDeleteHooks = append(eventOneAfterDeleteHooks, eventOneHook)
		eventOneAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		eventOneBeforeUpsertMu.Lock()
		eventOneBeforeUpsertHooks = append(eventOneBeforeUpsertHooks, eventOneHook)
		eventOneBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		eventOneAfterUpsertMu.Lock()
		eventOneAfterUpsertHooks = append(eventOneAfterUpsertHooks, eventOneHook)
		eventOneAfterUpsertMu.Unlock()
	}
}

// One returns a single eventOne record from the query.
func (q eventOneQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EventOne, error) {
	o := &EventOne{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for event_one")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EventOne records from the query.
func (q eventOneQuery) All(ctx context.Context, exec boil.ContextExecutor) (EventOneSlice, error) {
	var o []*EventOne

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EventOne slice")
	}

	if len(eventOneAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EventOne records in the query.
func (q eventOneQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count event_one rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q eventOneQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if event_one exists")
	}

	return count > 0, nil
}

// EventOnes retrieves all the records using an executor.
func EventOnes(mods ...qm.QueryMod) eventOneQuery {
	mods = append(mods, qm.From("`event_one`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`event_one`.*"})
	}

	return eventOneQuery{q}
}

// FindEventOne retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEventOne(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*EventOne, error) {
	eventOneObj := &EventOne{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `event_one` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, eventOneObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from event_one")
	}

	if err = eventOneObj.doAfterSelectHooks(ctx, exec); err != nil {
		return eventOneObj, err
	}

	return eventOneObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EventOne) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no event_one provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventOneColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	eventOneInsertCacheMut.RLock()
	cache, cached := eventOneInsertCache[key]
	eventOneInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			eventOneAllColumns,
			eventOneColumnsWithDefault,
			eventOneColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(eventOneType, eventOneMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(eventOneType, eventOneMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `event_one` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `event_one` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `event_one` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, eventOnePrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into event_one")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == eventOneMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for event_one")
	}

CacheNoHooks:
	if !cached {
		eventOneInsertCacheMut.Lock()
		eventOneInsertCache[key] = cache
		eventOneInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EventOne.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EventOne) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	eventOneUpdateCacheMut.RLock()
	cache, cached := eventOneUpdateCache[key]
	eventOneUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			eventOneAllColumns,
			eventOnePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update event_one, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `event_one` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, eventOnePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(eventOneType, eventOneMapping, append(wl, eventOnePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update event_one row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for event_one")
	}

	if !cached {
		eventOneUpdateCacheMut.Lock()
		eventOneUpdateCache[key] = cache
		eventOneUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q eventOneQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for event_one")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for event_one")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EventOneSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventOnePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `event_one` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventOnePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in eventOne slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all eventOne")
	}
	return rowsAff, nil
}

var mySQLEventOneUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EventOne) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no event_one provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventOneColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEventOneUniqueColumns, o)

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

	eventOneUpsertCacheMut.RLock()
	cache, cached := eventOneUpsertCache[key]
	eventOneUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			eventOneAllColumns,
			eventOneColumnsWithDefault,
			eventOneColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			eventOneAllColumns,
			eventOnePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert event_one, could not build update column list")
		}

		ret := strmangle.SetComplement(eventOneAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`event_one`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `event_one` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(eventOneType, eventOneMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(eventOneType, eventOneMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for event_one")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == eventOneMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(eventOneType, eventOneMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for event_one")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for event_one")
	}

CacheNoHooks:
	if !cached {
		eventOneUpsertCacheMut.Lock()
		eventOneUpsertCache[key] = cache
		eventOneUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EventOne record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EventOne) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EventOne provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), eventOnePrimaryKeyMapping)
	sql := "DELETE FROM `event_one` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from event_one")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for event_one")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q eventOneQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no eventOneQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from event_one")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for event_one")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EventOneSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(eventOneBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventOnePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `event_one` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventOnePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from eventOne slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for event_one")
	}

	if len(eventOneAfterDeleteHooks) != 0 {
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
func (o *EventOne) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEventOne(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventOneSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EventOneSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventOnePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `event_one`.* FROM `event_one` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventOnePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EventOneSlice")
	}

	*o = slice

	return nil
}

// EventOneExists checks if the EventOne row exists.
func EventOneExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `event_one` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if event_one exists")
	}

	return exists, nil
}

// Exists checks if the EventOne row exists.
func (o *EventOne) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EventOneExists(ctx, exec, o.ID)
}
