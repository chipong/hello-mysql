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

// EventTwo is an object representing the database table.
type EventTwo struct {
	ID   uint64           `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name null.String      `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Face EventTwoNullFace `boil:"face" json:"face,omitempty" toml:"face" yaml:"face,omitempty"`

	R *eventTwoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L eventTwoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EventTwoColumns = struct {
	ID   string
	Name string
	Face string
}{
	ID:   "id",
	Name: "name",
	Face: "face",
}

var EventTwoTableColumns = struct {
	ID   string
	Name string
	Face string
}{
	ID:   "event_two.id",
	Name: "event_two.name",
	Face: "event_two.face",
}

// Generated where

type whereHelperEventTwoNullFace struct{ field string }

func (w whereHelperEventTwoNullFace) EQ(x EventTwoNullFace) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelperEventTwoNullFace) NEQ(x EventTwoNullFace) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelperEventTwoNullFace) LT(x EventTwoNullFace) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperEventTwoNullFace) LTE(x EventTwoNullFace) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperEventTwoNullFace) GT(x EventTwoNullFace) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperEventTwoNullFace) GTE(x EventTwoNullFace) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperEventTwoNullFace) IN(slice []EventTwoNullFace) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperEventTwoNullFace) NIN(slice []EventTwoNullFace) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelperEventTwoNullFace) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelperEventTwoNullFace) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var EventTwoWhere = struct {
	ID   whereHelperuint64
	Name whereHelpernull_String
	Face whereHelperEventTwoNullFace
}{
	ID:   whereHelperuint64{field: "`event_two`.`id`"},
	Name: whereHelpernull_String{field: "`event_two`.`name`"},
	Face: whereHelperEventTwoNullFace{field: "`event_two`.`face`"},
}

// EventTwoRels is where relationship names are stored.
var EventTwoRels = struct {
}{}

// eventTwoR is where relationships are stored.
type eventTwoR struct {
}

// NewStruct creates a new relationship struct
func (*eventTwoR) NewStruct() *eventTwoR {
	return &eventTwoR{}
}

// eventTwoL is where Load methods for each relationship are stored.
type eventTwoL struct{}

var (
	eventTwoAllColumns            = []string{"id", "name", "face"}
	eventTwoColumnsWithoutDefault = []string{"name", "face"}
	eventTwoColumnsWithDefault    = []string{"id"}
	eventTwoPrimaryKeyColumns     = []string{"id"}
	eventTwoGeneratedColumns      = []string{}
)

type (
	// EventTwoSlice is an alias for a slice of pointers to EventTwo.
	// This should almost always be used instead of []EventTwo.
	EventTwoSlice []*EventTwo
	// EventTwoHook is the signature for custom EventTwo hook methods
	EventTwoHook func(context.Context, boil.ContextExecutor, *EventTwo) error

	eventTwoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	eventTwoType                 = reflect.TypeOf(&EventTwo{})
	eventTwoMapping              = queries.MakeStructMapping(eventTwoType)
	eventTwoPrimaryKeyMapping, _ = queries.BindMapping(eventTwoType, eventTwoMapping, eventTwoPrimaryKeyColumns)
	eventTwoInsertCacheMut       sync.RWMutex
	eventTwoInsertCache          = make(map[string]insertCache)
	eventTwoUpdateCacheMut       sync.RWMutex
	eventTwoUpdateCache          = make(map[string]updateCache)
	eventTwoUpsertCacheMut       sync.RWMutex
	eventTwoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var eventTwoAfterSelectMu sync.Mutex
var eventTwoAfterSelectHooks []EventTwoHook

var eventTwoBeforeInsertMu sync.Mutex
var eventTwoBeforeInsertHooks []EventTwoHook
var eventTwoAfterInsertMu sync.Mutex
var eventTwoAfterInsertHooks []EventTwoHook

var eventTwoBeforeUpdateMu sync.Mutex
var eventTwoBeforeUpdateHooks []EventTwoHook
var eventTwoAfterUpdateMu sync.Mutex
var eventTwoAfterUpdateHooks []EventTwoHook

var eventTwoBeforeDeleteMu sync.Mutex
var eventTwoBeforeDeleteHooks []EventTwoHook
var eventTwoAfterDeleteMu sync.Mutex
var eventTwoAfterDeleteHooks []EventTwoHook

var eventTwoBeforeUpsertMu sync.Mutex
var eventTwoBeforeUpsertHooks []EventTwoHook
var eventTwoAfterUpsertMu sync.Mutex
var eventTwoAfterUpsertHooks []EventTwoHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EventTwo) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTwoAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EventTwo) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTwoBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EventTwo) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTwoAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EventTwo) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTwoBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EventTwo) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTwoAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EventTwo) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTwoBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EventTwo) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTwoAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EventTwo) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTwoBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EventTwo) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTwoAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEventTwoHook registers your hook function for all future operations.
func AddEventTwoHook(hookPoint boil.HookPoint, eventTwoHook EventTwoHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		eventTwoAfterSelectMu.Lock()
		eventTwoAfterSelectHooks = append(eventTwoAfterSelectHooks, eventTwoHook)
		eventTwoAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		eventTwoBeforeInsertMu.Lock()
		eventTwoBeforeInsertHooks = append(eventTwoBeforeInsertHooks, eventTwoHook)
		eventTwoBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		eventTwoAfterInsertMu.Lock()
		eventTwoAfterInsertHooks = append(eventTwoAfterInsertHooks, eventTwoHook)
		eventTwoAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		eventTwoBeforeUpdateMu.Lock()
		eventTwoBeforeUpdateHooks = append(eventTwoBeforeUpdateHooks, eventTwoHook)
		eventTwoBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		eventTwoAfterUpdateMu.Lock()
		eventTwoAfterUpdateHooks = append(eventTwoAfterUpdateHooks, eventTwoHook)
		eventTwoAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		eventTwoBeforeDeleteMu.Lock()
		eventTwoBeforeDeleteHooks = append(eventTwoBeforeDeleteHooks, eventTwoHook)
		eventTwoBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		eventTwoAfterDeleteMu.Lock()
		eventTwoAfterDeleteHooks = append(eventTwoAfterDeleteHooks, eventTwoHook)
		eventTwoAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		eventTwoBeforeUpsertMu.Lock()
		eventTwoBeforeUpsertHooks = append(eventTwoBeforeUpsertHooks, eventTwoHook)
		eventTwoBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		eventTwoAfterUpsertMu.Lock()
		eventTwoAfterUpsertHooks = append(eventTwoAfterUpsertHooks, eventTwoHook)
		eventTwoAfterUpsertMu.Unlock()
	}
}

// One returns a single eventTwo record from the query.
func (q eventTwoQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EventTwo, error) {
	o := &EventTwo{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for event_two")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EventTwo records from the query.
func (q eventTwoQuery) All(ctx context.Context, exec boil.ContextExecutor) (EventTwoSlice, error) {
	var o []*EventTwo

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EventTwo slice")
	}

	if len(eventTwoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EventTwo records in the query.
func (q eventTwoQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count event_two rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q eventTwoQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if event_two exists")
	}

	return count > 0, nil
}

// EventTwos retrieves all the records using an executor.
func EventTwos(mods ...qm.QueryMod) eventTwoQuery {
	mods = append(mods, qm.From("`event_two`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`event_two`.*"})
	}

	return eventTwoQuery{q}
}

// FindEventTwo retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEventTwo(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*EventTwo, error) {
	eventTwoObj := &EventTwo{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `event_two` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, eventTwoObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from event_two")
	}

	if err = eventTwoObj.doAfterSelectHooks(ctx, exec); err != nil {
		return eventTwoObj, err
	}

	return eventTwoObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EventTwo) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no event_two provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventTwoColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	eventTwoInsertCacheMut.RLock()
	cache, cached := eventTwoInsertCache[key]
	eventTwoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			eventTwoAllColumns,
			eventTwoColumnsWithDefault,
			eventTwoColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(eventTwoType, eventTwoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(eventTwoType, eventTwoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `event_two` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `event_two` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `event_two` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, eventTwoPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into event_two")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == eventTwoMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for event_two")
	}

CacheNoHooks:
	if !cached {
		eventTwoInsertCacheMut.Lock()
		eventTwoInsertCache[key] = cache
		eventTwoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EventTwo.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EventTwo) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	eventTwoUpdateCacheMut.RLock()
	cache, cached := eventTwoUpdateCache[key]
	eventTwoUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			eventTwoAllColumns,
			eventTwoPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update event_two, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `event_two` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, eventTwoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(eventTwoType, eventTwoMapping, append(wl, eventTwoPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update event_two row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for event_two")
	}

	if !cached {
		eventTwoUpdateCacheMut.Lock()
		eventTwoUpdateCache[key] = cache
		eventTwoUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q eventTwoQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for event_two")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for event_two")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EventTwoSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventTwoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `event_two` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventTwoPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in eventTwo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all eventTwo")
	}
	return rowsAff, nil
}

var mySQLEventTwoUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EventTwo) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no event_two provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventTwoColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEventTwoUniqueColumns, o)

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

	eventTwoUpsertCacheMut.RLock()
	cache, cached := eventTwoUpsertCache[key]
	eventTwoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			eventTwoAllColumns,
			eventTwoColumnsWithDefault,
			eventTwoColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			eventTwoAllColumns,
			eventTwoPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert event_two, could not build update column list")
		}

		ret := strmangle.SetComplement(eventTwoAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`event_two`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `event_two` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(eventTwoType, eventTwoMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(eventTwoType, eventTwoMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for event_two")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == eventTwoMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(eventTwoType, eventTwoMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for event_two")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for event_two")
	}

CacheNoHooks:
	if !cached {
		eventTwoUpsertCacheMut.Lock()
		eventTwoUpsertCache[key] = cache
		eventTwoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EventTwo record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EventTwo) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EventTwo provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), eventTwoPrimaryKeyMapping)
	sql := "DELETE FROM `event_two` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from event_two")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for event_two")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q eventTwoQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no eventTwoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from event_two")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for event_two")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EventTwoSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(eventTwoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventTwoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `event_two` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventTwoPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from eventTwo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for event_two")
	}

	if len(eventTwoAfterDeleteHooks) != 0 {
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
func (o *EventTwo) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEventTwo(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventTwoSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EventTwoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventTwoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `event_two`.* FROM `event_two` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventTwoPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EventTwoSlice")
	}

	*o = slice

	return nil
}

// EventTwoExists checks if the EventTwo row exists.
func EventTwoExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `event_two` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if event_two exists")
	}

	return exists, nil
}

// Exists checks if the EventTwo row exists.
func (o *EventTwo) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EventTwoExists(ctx, exec, o.ID)
}
