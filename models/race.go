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

// Race is an object representing the database table.
type Race struct {
	ID       int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	RaceDate null.Time   `boil:"race_date" json:"race_date,omitempty" toml:"race_date" yaml:"race_date,omitempty"`
	Track    null.String `boil:"track" json:"track,omitempty" toml:"track" yaml:"track,omitempty"`

	R *raceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L raceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RaceColumns = struct {
	ID       string
	RaceDate string
	Track    string
}{
	ID:       "id",
	RaceDate: "race_date",
	Track:    "track",
}

var RaceTableColumns = struct {
	ID       string
	RaceDate string
	Track    string
}{
	ID:       "race.id",
	RaceDate: "race.race_date",
	Track:    "race.track",
}

// Generated where

var RaceWhere = struct {
	ID       whereHelperint
	RaceDate whereHelpernull_Time
	Track    whereHelpernull_String
}{
	ID:       whereHelperint{field: "`race`.`id`"},
	RaceDate: whereHelpernull_Time{field: "`race`.`race_date`"},
	Track:    whereHelpernull_String{field: "`race`.`track`"},
}

// RaceRels is where relationship names are stored.
var RaceRels = struct {
	RaceResults string
}{
	RaceResults: "RaceResults",
}

// raceR is where relationships are stored.
type raceR struct {
	RaceResults RaceResultSlice `boil:"RaceResults" json:"RaceResults" toml:"RaceResults" yaml:"RaceResults"`
}

// NewStruct creates a new relationship struct
func (*raceR) NewStruct() *raceR {
	return &raceR{}
}

func (r *raceR) GetRaceResults() RaceResultSlice {
	if r == nil {
		return nil
	}
	return r.RaceResults
}

// raceL is where Load methods for each relationship are stored.
type raceL struct{}

var (
	raceAllColumns            = []string{"id", "race_date", "track"}
	raceColumnsWithoutDefault = []string{"id", "race_date", "track"}
	raceColumnsWithDefault    = []string{}
	racePrimaryKeyColumns     = []string{"id"}
	raceGeneratedColumns      = []string{}
)

type (
	// RaceSlice is an alias for a slice of pointers to Race.
	// This should almost always be used instead of []Race.
	RaceSlice []*Race
	// RaceHook is the signature for custom Race hook methods
	RaceHook func(context.Context, boil.ContextExecutor, *Race) error

	raceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	raceType                 = reflect.TypeOf(&Race{})
	raceMapping              = queries.MakeStructMapping(raceType)
	racePrimaryKeyMapping, _ = queries.BindMapping(raceType, raceMapping, racePrimaryKeyColumns)
	raceInsertCacheMut       sync.RWMutex
	raceInsertCache          = make(map[string]insertCache)
	raceUpdateCacheMut       sync.RWMutex
	raceUpdateCache          = make(map[string]updateCache)
	raceUpsertCacheMut       sync.RWMutex
	raceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var raceAfterSelectMu sync.Mutex
var raceAfterSelectHooks []RaceHook

var raceBeforeInsertMu sync.Mutex
var raceBeforeInsertHooks []RaceHook
var raceAfterInsertMu sync.Mutex
var raceAfterInsertHooks []RaceHook

var raceBeforeUpdateMu sync.Mutex
var raceBeforeUpdateHooks []RaceHook
var raceAfterUpdateMu sync.Mutex
var raceAfterUpdateHooks []RaceHook

var raceBeforeDeleteMu sync.Mutex
var raceBeforeDeleteHooks []RaceHook
var raceAfterDeleteMu sync.Mutex
var raceAfterDeleteHooks []RaceHook

var raceBeforeUpsertMu sync.Mutex
var raceBeforeUpsertHooks []RaceHook
var raceAfterUpsertMu sync.Mutex
var raceAfterUpsertHooks []RaceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Race) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Race) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Race) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Race) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Race) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Race) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Race) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Race) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Race) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRaceHook registers your hook function for all future operations.
func AddRaceHook(hookPoint boil.HookPoint, raceHook RaceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		raceAfterSelectMu.Lock()
		raceAfterSelectHooks = append(raceAfterSelectHooks, raceHook)
		raceAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		raceBeforeInsertMu.Lock()
		raceBeforeInsertHooks = append(raceBeforeInsertHooks, raceHook)
		raceBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		raceAfterInsertMu.Lock()
		raceAfterInsertHooks = append(raceAfterInsertHooks, raceHook)
		raceAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		raceBeforeUpdateMu.Lock()
		raceBeforeUpdateHooks = append(raceBeforeUpdateHooks, raceHook)
		raceBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		raceAfterUpdateMu.Lock()
		raceAfterUpdateHooks = append(raceAfterUpdateHooks, raceHook)
		raceAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		raceBeforeDeleteMu.Lock()
		raceBeforeDeleteHooks = append(raceBeforeDeleteHooks, raceHook)
		raceBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		raceAfterDeleteMu.Lock()
		raceAfterDeleteHooks = append(raceAfterDeleteHooks, raceHook)
		raceAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		raceBeforeUpsertMu.Lock()
		raceBeforeUpsertHooks = append(raceBeforeUpsertHooks, raceHook)
		raceBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		raceAfterUpsertMu.Lock()
		raceAfterUpsertHooks = append(raceAfterUpsertHooks, raceHook)
		raceAfterUpsertMu.Unlock()
	}
}

// One returns a single race record from the query.
func (q raceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Race, error) {
	o := &Race{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for race")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Race records from the query.
func (q raceQuery) All(ctx context.Context, exec boil.ContextExecutor) (RaceSlice, error) {
	var o []*Race

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Race slice")
	}

	if len(raceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Race records in the query.
func (q raceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count race rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q raceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if race exists")
	}

	return count > 0, nil
}

// RaceResults retrieves all the race_result's RaceResults with an executor.
func (o *Race) RaceResults(mods ...qm.QueryMod) raceResultQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`race_results`.`race_id`=?", o.ID),
	)

	return RaceResults(queryMods...)
}

// LoadRaceResults allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (raceL) LoadRaceResults(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRace interface{}, mods queries.Applicator) error {
	var slice []*Race
	var object *Race

	if singular {
		var ok bool
		object, ok = maybeRace.(*Race)
		if !ok {
			object = new(Race)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeRace)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeRace))
			}
		}
	} else {
		s, ok := maybeRace.(*[]*Race)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeRace)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeRace))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &raceR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &raceR{}
			}
			args[obj.ID] = struct{}{}
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
		qm.From(`race_results`),
		qm.WhereIn(`race_results.race_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load race_results")
	}

	var resultSlice []*RaceResult
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice race_results")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on race_results")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for race_results")
	}

	if len(raceResultAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.RaceResults = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &raceResultR{}
			}
			foreign.R.Race = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.RaceID) {
				local.R.RaceResults = append(local.R.RaceResults, foreign)
				if foreign.R == nil {
					foreign.R = &raceResultR{}
				}
				foreign.R.Race = local
				break
			}
		}
	}

	return nil
}

// AddRaceResults adds the given related objects to the existing relationships
// of the race, optionally inserting them as new records.
// Appends related to o.R.RaceResults.
// Sets related.R.Race appropriately.
func (o *Race) AddRaceResults(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RaceResult) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.RaceID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `race_results` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"race_id"}),
				strmangle.WhereClause("`", "`", 0, raceResultPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.RaceID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &raceR{
			RaceResults: related,
		}
	} else {
		o.R.RaceResults = append(o.R.RaceResults, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &raceResultR{
				Race: o,
			}
		} else {
			rel.R.Race = o
		}
	}
	return nil
}

// SetRaceResults removes all previously related items of the
// race replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Race's RaceResults accordingly.
// Replaces o.R.RaceResults with related.
// Sets related.R.Race's RaceResults accordingly.
func (o *Race) SetRaceResults(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RaceResult) error {
	query := "update `race_results` set `race_id` = null where `race_id` = ?"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.RaceResults {
			queries.SetScanner(&rel.RaceID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Race = nil
		}
		o.R.RaceResults = nil
	}

	return o.AddRaceResults(ctx, exec, insert, related...)
}

// RemoveRaceResults relationships from objects passed in.
// Removes related items from R.RaceResults (uses pointer comparison, removal does not keep order)
// Sets related.R.Race.
func (o *Race) RemoveRaceResults(ctx context.Context, exec boil.ContextExecutor, related ...*RaceResult) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.RaceID, nil)
		if rel.R != nil {
			rel.R.Race = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("race_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.RaceResults {
			if rel != ri {
				continue
			}

			ln := len(o.R.RaceResults)
			if ln > 1 && i < ln-1 {
				o.R.RaceResults[i] = o.R.RaceResults[ln-1]
			}
			o.R.RaceResults = o.R.RaceResults[:ln-1]
			break
		}
	}

	return nil
}

// Races retrieves all the records using an executor.
func Races(mods ...qm.QueryMod) raceQuery {
	mods = append(mods, qm.From("`race`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`race`.*"})
	}

	return raceQuery{q}
}

// FindRace retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRace(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Race, error) {
	raceObj := &Race{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `race` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, raceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from race")
	}

	if err = raceObj.doAfterSelectHooks(ctx, exec); err != nil {
		return raceObj, err
	}

	return raceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Race) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no race provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(raceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	raceInsertCacheMut.RLock()
	cache, cached := raceInsertCache[key]
	raceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			raceAllColumns,
			raceColumnsWithDefault,
			raceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(raceType, raceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(raceType, raceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `race` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `race` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `race` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, racePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into race")
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
		return errors.Wrap(err, "models: unable to populate default values for race")
	}

CacheNoHooks:
	if !cached {
		raceInsertCacheMut.Lock()
		raceInsertCache[key] = cache
		raceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Race.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Race) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	raceUpdateCacheMut.RLock()
	cache, cached := raceUpdateCache[key]
	raceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			raceAllColumns,
			racePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update race, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `race` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, racePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(raceType, raceMapping, append(wl, racePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update race row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for race")
	}

	if !cached {
		raceUpdateCacheMut.Lock()
		raceUpdateCache[key] = cache
		raceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q raceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for race")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for race")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RaceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), racePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `race` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, racePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in race slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all race")
	}
	return rowsAff, nil
}

var mySQLRaceUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Race) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no race provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(raceColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRaceUniqueColumns, o)

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

	raceUpsertCacheMut.RLock()
	cache, cached := raceUpsertCache[key]
	raceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			raceAllColumns,
			raceColumnsWithDefault,
			raceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			raceAllColumns,
			racePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert race, could not build update column list")
		}

		ret := strmangle.SetComplement(raceAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`race`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `race` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(raceType, raceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(raceType, raceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for race")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(raceType, raceMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for race")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for race")
	}

CacheNoHooks:
	if !cached {
		raceUpsertCacheMut.Lock()
		raceUpsertCache[key] = cache
		raceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Race record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Race) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Race provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), racePrimaryKeyMapping)
	sql := "DELETE FROM `race` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from race")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for race")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q raceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no raceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from race")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for race")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RaceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(raceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), racePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `race` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, racePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from race slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for race")
	}

	if len(raceAfterDeleteHooks) != 0 {
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
func (o *Race) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRace(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RaceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RaceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), racePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `race`.* FROM `race` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, racePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RaceSlice")
	}

	*o = slice

	return nil
}

// RaceExists checks if the Race row exists.
func RaceExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `race` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if race exists")
	}

	return exists, nil
}

// Exists checks if the Race row exists.
func (o *Race) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RaceExists(ctx, exec, o.ID)
}
