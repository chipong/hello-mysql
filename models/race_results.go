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

// RaceResult is an object representing the database table.
type RaceResult struct {
	ID     int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	RaceID null.Int    `boil:"race_id" json:"race_id,omitempty" toml:"race_id" yaml:"race_id,omitempty"`
	Name   null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`

	R *raceResultR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L raceResultL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RaceResultColumns = struct {
	ID     string
	RaceID string
	Name   string
}{
	ID:     "id",
	RaceID: "race_id",
	Name:   "name",
}

var RaceResultTableColumns = struct {
	ID     string
	RaceID string
	Name   string
}{
	ID:     "race_results.id",
	RaceID: "race_results.race_id",
	Name:   "race_results.name",
}

// Generated where

var RaceResultWhere = struct {
	ID     whereHelperint
	RaceID whereHelpernull_Int
	Name   whereHelpernull_String
}{
	ID:     whereHelperint{field: "`race_results`.`id`"},
	RaceID: whereHelpernull_Int{field: "`race_results`.`race_id`"},
	Name:   whereHelpernull_String{field: "`race_results`.`name`"},
}

// RaceResultRels is where relationship names are stored.
var RaceResultRels = struct {
	Race                        string
	ResultRaceResultScratchings string
}{
	Race:                        "Race",
	ResultRaceResultScratchings: "ResultRaceResultScratchings",
}

// raceResultR is where relationships are stored.
type raceResultR struct {
	Race                        *Race                     `boil:"Race" json:"Race" toml:"Race" yaml:"Race"`
	ResultRaceResultScratchings RaceResultScratchingSlice `boil:"ResultRaceResultScratchings" json:"ResultRaceResultScratchings" toml:"ResultRaceResultScratchings" yaml:"ResultRaceResultScratchings"`
}

// NewStruct creates a new relationship struct
func (*raceResultR) NewStruct() *raceResultR {
	return &raceResultR{}
}

func (r *raceResultR) GetRace() *Race {
	if r == nil {
		return nil
	}
	return r.Race
}

func (r *raceResultR) GetResultRaceResultScratchings() RaceResultScratchingSlice {
	if r == nil {
		return nil
	}
	return r.ResultRaceResultScratchings
}

// raceResultL is where Load methods for each relationship are stored.
type raceResultL struct{}

var (
	raceResultAllColumns            = []string{"id", "race_id", "name"}
	raceResultColumnsWithoutDefault = []string{"id", "race_id", "name"}
	raceResultColumnsWithDefault    = []string{}
	raceResultPrimaryKeyColumns     = []string{"id"}
	raceResultGeneratedColumns      = []string{}
)

type (
	// RaceResultSlice is an alias for a slice of pointers to RaceResult.
	// This should almost always be used instead of []RaceResult.
	RaceResultSlice []*RaceResult
	// RaceResultHook is the signature for custom RaceResult hook methods
	RaceResultHook func(context.Context, boil.ContextExecutor, *RaceResult) error

	raceResultQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	raceResultType                 = reflect.TypeOf(&RaceResult{})
	raceResultMapping              = queries.MakeStructMapping(raceResultType)
	raceResultPrimaryKeyMapping, _ = queries.BindMapping(raceResultType, raceResultMapping, raceResultPrimaryKeyColumns)
	raceResultInsertCacheMut       sync.RWMutex
	raceResultInsertCache          = make(map[string]insertCache)
	raceResultUpdateCacheMut       sync.RWMutex
	raceResultUpdateCache          = make(map[string]updateCache)
	raceResultUpsertCacheMut       sync.RWMutex
	raceResultUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var raceResultAfterSelectMu sync.Mutex
var raceResultAfterSelectHooks []RaceResultHook

var raceResultBeforeInsertMu sync.Mutex
var raceResultBeforeInsertHooks []RaceResultHook
var raceResultAfterInsertMu sync.Mutex
var raceResultAfterInsertHooks []RaceResultHook

var raceResultBeforeUpdateMu sync.Mutex
var raceResultBeforeUpdateHooks []RaceResultHook
var raceResultAfterUpdateMu sync.Mutex
var raceResultAfterUpdateHooks []RaceResultHook

var raceResultBeforeDeleteMu sync.Mutex
var raceResultBeforeDeleteHooks []RaceResultHook
var raceResultAfterDeleteMu sync.Mutex
var raceResultAfterDeleteHooks []RaceResultHook

var raceResultBeforeUpsertMu sync.Mutex
var raceResultBeforeUpsertHooks []RaceResultHook
var raceResultAfterUpsertMu sync.Mutex
var raceResultAfterUpsertHooks []RaceResultHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RaceResult) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceResultAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RaceResult) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceResultBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RaceResult) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceResultAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RaceResult) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceResultBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RaceResult) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceResultAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RaceResult) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceResultBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RaceResult) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceResultAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RaceResult) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceResultBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RaceResult) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range raceResultAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRaceResultHook registers your hook function for all future operations.
func AddRaceResultHook(hookPoint boil.HookPoint, raceResultHook RaceResultHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		raceResultAfterSelectMu.Lock()
		raceResultAfterSelectHooks = append(raceResultAfterSelectHooks, raceResultHook)
		raceResultAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		raceResultBeforeInsertMu.Lock()
		raceResultBeforeInsertHooks = append(raceResultBeforeInsertHooks, raceResultHook)
		raceResultBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		raceResultAfterInsertMu.Lock()
		raceResultAfterInsertHooks = append(raceResultAfterInsertHooks, raceResultHook)
		raceResultAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		raceResultBeforeUpdateMu.Lock()
		raceResultBeforeUpdateHooks = append(raceResultBeforeUpdateHooks, raceResultHook)
		raceResultBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		raceResultAfterUpdateMu.Lock()
		raceResultAfterUpdateHooks = append(raceResultAfterUpdateHooks, raceResultHook)
		raceResultAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		raceResultBeforeDeleteMu.Lock()
		raceResultBeforeDeleteHooks = append(raceResultBeforeDeleteHooks, raceResultHook)
		raceResultBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		raceResultAfterDeleteMu.Lock()
		raceResultAfterDeleteHooks = append(raceResultAfterDeleteHooks, raceResultHook)
		raceResultAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		raceResultBeforeUpsertMu.Lock()
		raceResultBeforeUpsertHooks = append(raceResultBeforeUpsertHooks, raceResultHook)
		raceResultBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		raceResultAfterUpsertMu.Lock()
		raceResultAfterUpsertHooks = append(raceResultAfterUpsertHooks, raceResultHook)
		raceResultAfterUpsertMu.Unlock()
	}
}

// One returns a single raceResult record from the query.
func (q raceResultQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RaceResult, error) {
	o := &RaceResult{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for race_results")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RaceResult records from the query.
func (q raceResultQuery) All(ctx context.Context, exec boil.ContextExecutor) (RaceResultSlice, error) {
	var o []*RaceResult

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RaceResult slice")
	}

	if len(raceResultAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RaceResult records in the query.
func (q raceResultQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count race_results rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q raceResultQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if race_results exists")
	}

	return count > 0, nil
}

// Race pointed to by the foreign key.
func (o *RaceResult) Race(mods ...qm.QueryMod) raceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.RaceID),
	}

	queryMods = append(queryMods, mods...)

	return Races(queryMods...)
}

// ResultRaceResultScratchings retrieves all the race_result_scratching's RaceResultScratchings with an executor via results_id column.
func (o *RaceResult) ResultRaceResultScratchings(mods ...qm.QueryMod) raceResultScratchingQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`race_result_scratchings`.`results_id`=?", o.ID),
	)

	return RaceResultScratchings(queryMods...)
}

// LoadRace allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (raceResultL) LoadRace(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRaceResult interface{}, mods queries.Applicator) error {
	var slice []*RaceResult
	var object *RaceResult

	if singular {
		var ok bool
		object, ok = maybeRaceResult.(*RaceResult)
		if !ok {
			object = new(RaceResult)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeRaceResult)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeRaceResult))
			}
		}
	} else {
		s, ok := maybeRaceResult.(*[]*RaceResult)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeRaceResult)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeRaceResult))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &raceResultR{}
		}
		if !queries.IsNil(object.RaceID) {
			args[object.RaceID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &raceResultR{}
			}

			if !queries.IsNil(obj.RaceID) {
				args[obj.RaceID] = struct{}{}
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
		qm.From(`race`),
		qm.WhereIn(`race.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Race")
	}

	var resultSlice []*Race
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Race")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for race")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for race")
	}

	if len(raceAfterSelectHooks) != 0 {
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
		object.R.Race = foreign
		if foreign.R == nil {
			foreign.R = &raceR{}
		}
		foreign.R.RaceResults = append(foreign.R.RaceResults, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.RaceID, foreign.ID) {
				local.R.Race = foreign
				if foreign.R == nil {
					foreign.R = &raceR{}
				}
				foreign.R.RaceResults = append(foreign.R.RaceResults, local)
				break
			}
		}
	}

	return nil
}

// LoadResultRaceResultScratchings allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (raceResultL) LoadResultRaceResultScratchings(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRaceResult interface{}, mods queries.Applicator) error {
	var slice []*RaceResult
	var object *RaceResult

	if singular {
		var ok bool
		object, ok = maybeRaceResult.(*RaceResult)
		if !ok {
			object = new(RaceResult)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeRaceResult)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeRaceResult))
			}
		}
	} else {
		s, ok := maybeRaceResult.(*[]*RaceResult)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeRaceResult)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeRaceResult))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &raceResultR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &raceResultR{}
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
		qm.From(`race_result_scratchings`),
		qm.WhereIn(`race_result_scratchings.results_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load race_result_scratchings")
	}

	var resultSlice []*RaceResultScratching
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice race_result_scratchings")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on race_result_scratchings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for race_result_scratchings")
	}

	if len(raceResultScratchingAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ResultRaceResultScratchings = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &raceResultScratchingR{}
			}
			foreign.R.Result = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ResultsID {
				local.R.ResultRaceResultScratchings = append(local.R.ResultRaceResultScratchings, foreign)
				if foreign.R == nil {
					foreign.R = &raceResultScratchingR{}
				}
				foreign.R.Result = local
				break
			}
		}
	}

	return nil
}

// SetRace of the raceResult to the related item.
// Sets o.R.Race to related.
// Adds o to related.R.RaceResults.
func (o *RaceResult) SetRace(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Race) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `race_results` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"race_id"}),
		strmangle.WhereClause("`", "`", 0, raceResultPrimaryKeyColumns),
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

	queries.Assign(&o.RaceID, related.ID)
	if o.R == nil {
		o.R = &raceResultR{
			Race: related,
		}
	} else {
		o.R.Race = related
	}

	if related.R == nil {
		related.R = &raceR{
			RaceResults: RaceResultSlice{o},
		}
	} else {
		related.R.RaceResults = append(related.R.RaceResults, o)
	}

	return nil
}

// RemoveRace relationship.
// Sets o.R.Race to nil.
// Removes o from all passed in related items' relationships struct.
func (o *RaceResult) RemoveRace(ctx context.Context, exec boil.ContextExecutor, related *Race) error {
	var err error

	queries.SetScanner(&o.RaceID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("race_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Race = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.RaceResults {
		if queries.Equal(o.RaceID, ri.RaceID) {
			continue
		}

		ln := len(related.R.RaceResults)
		if ln > 1 && i < ln-1 {
			related.R.RaceResults[i] = related.R.RaceResults[ln-1]
		}
		related.R.RaceResults = related.R.RaceResults[:ln-1]
		break
	}
	return nil
}

// AddResultRaceResultScratchings adds the given related objects to the existing relationships
// of the race_result, optionally inserting them as new records.
// Appends related to o.R.ResultRaceResultScratchings.
// Sets related.R.Result appropriately.
func (o *RaceResult) AddResultRaceResultScratchings(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RaceResultScratching) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ResultsID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `race_result_scratchings` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"results_id"}),
				strmangle.WhereClause("`", "`", 0, raceResultScratchingPrimaryKeyColumns),
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

			rel.ResultsID = o.ID
		}
	}

	if o.R == nil {
		o.R = &raceResultR{
			ResultRaceResultScratchings: related,
		}
	} else {
		o.R.ResultRaceResultScratchings = append(o.R.ResultRaceResultScratchings, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &raceResultScratchingR{
				Result: o,
			}
		} else {
			rel.R.Result = o
		}
	}
	return nil
}

// RaceResults retrieves all the records using an executor.
func RaceResults(mods ...qm.QueryMod) raceResultQuery {
	mods = append(mods, qm.From("`race_results`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`race_results`.*"})
	}

	return raceResultQuery{q}
}

// FindRaceResult retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRaceResult(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RaceResult, error) {
	raceResultObj := &RaceResult{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `race_results` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, raceResultObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from race_results")
	}

	if err = raceResultObj.doAfterSelectHooks(ctx, exec); err != nil {
		return raceResultObj, err
	}

	return raceResultObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RaceResult) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no race_results provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(raceResultColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	raceResultInsertCacheMut.RLock()
	cache, cached := raceResultInsertCache[key]
	raceResultInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			raceResultAllColumns,
			raceResultColumnsWithDefault,
			raceResultColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(raceResultType, raceResultMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(raceResultType, raceResultMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `race_results` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `race_results` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `race_results` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, raceResultPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into race_results")
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
		return errors.Wrap(err, "models: unable to populate default values for race_results")
	}

CacheNoHooks:
	if !cached {
		raceResultInsertCacheMut.Lock()
		raceResultInsertCache[key] = cache
		raceResultInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RaceResult.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RaceResult) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	raceResultUpdateCacheMut.RLock()
	cache, cached := raceResultUpdateCache[key]
	raceResultUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			raceResultAllColumns,
			raceResultPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update race_results, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `race_results` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, raceResultPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(raceResultType, raceResultMapping, append(wl, raceResultPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update race_results row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for race_results")
	}

	if !cached {
		raceResultUpdateCacheMut.Lock()
		raceResultUpdateCache[key] = cache
		raceResultUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q raceResultQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for race_results")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for race_results")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RaceResultSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), raceResultPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `race_results` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, raceResultPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in raceResult slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all raceResult")
	}
	return rowsAff, nil
}

var mySQLRaceResultUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RaceResult) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no race_results provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(raceResultColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRaceResultUniqueColumns, o)

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

	raceResultUpsertCacheMut.RLock()
	cache, cached := raceResultUpsertCache[key]
	raceResultUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			raceResultAllColumns,
			raceResultColumnsWithDefault,
			raceResultColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			raceResultAllColumns,
			raceResultPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert race_results, could not build update column list")
		}

		ret := strmangle.SetComplement(raceResultAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`race_results`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `race_results` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(raceResultType, raceResultMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(raceResultType, raceResultMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for race_results")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(raceResultType, raceResultMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for race_results")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for race_results")
	}

CacheNoHooks:
	if !cached {
		raceResultUpsertCacheMut.Lock()
		raceResultUpsertCache[key] = cache
		raceResultUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RaceResult record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RaceResult) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RaceResult provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), raceResultPrimaryKeyMapping)
	sql := "DELETE FROM `race_results` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from race_results")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for race_results")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q raceResultQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no raceResultQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from race_results")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for race_results")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RaceResultSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(raceResultBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), raceResultPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `race_results` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, raceResultPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from raceResult slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for race_results")
	}

	if len(raceResultAfterDeleteHooks) != 0 {
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
func (o *RaceResult) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRaceResult(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RaceResultSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RaceResultSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), raceResultPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `race_results`.* FROM `race_results` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, raceResultPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RaceResultSlice")
	}

	*o = slice

	return nil
}

// RaceResultExists checks if the RaceResult row exists.
func RaceResultExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `race_results` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if race_results exists")
	}

	return exists, nil
}

// Exists checks if the RaceResult row exists.
func (o *RaceResult) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RaceResultExists(ctx, exec, o.ID)
}