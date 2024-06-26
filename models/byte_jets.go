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

// ByteJet is an object representing the database table.
type ByteJet struct {
	ID            []byte      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name          null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	BytePilotID   null.Bytes  `boil:"byte_pilot_id" json:"byte_pilot_id,omitempty" toml:"byte_pilot_id" yaml:"byte_pilot_id,omitempty"`
	ByteAirportID null.Bytes  `boil:"byte_airport_id" json:"byte_airport_id,omitempty" toml:"byte_airport_id" yaml:"byte_airport_id,omitempty"`

	R *byteJetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L byteJetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ByteJetColumns = struct {
	ID            string
	Name          string
	BytePilotID   string
	ByteAirportID string
}{
	ID:            "id",
	Name:          "name",
	BytePilotID:   "byte_pilot_id",
	ByteAirportID: "byte_airport_id",
}

var ByteJetTableColumns = struct {
	ID            string
	Name          string
	BytePilotID   string
	ByteAirportID string
}{
	ID:            "byte_jets.id",
	Name:          "byte_jets.name",
	BytePilotID:   "byte_jets.byte_pilot_id",
	ByteAirportID: "byte_jets.byte_airport_id",
}

// Generated where

type whereHelpernull_Bytes struct{ field string }

func (w whereHelpernull_Bytes) EQ(x null.Bytes) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bytes) NEQ(x null.Bytes) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bytes) LT(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bytes) LTE(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bytes) GT(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bytes) GTE(x null.Bytes) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Bytes) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bytes) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var ByteJetWhere = struct {
	ID            whereHelper__byte
	Name          whereHelpernull_String
	BytePilotID   whereHelpernull_Bytes
	ByteAirportID whereHelpernull_Bytes
}{
	ID:            whereHelper__byte{field: "`byte_jets`.`id`"},
	Name:          whereHelpernull_String{field: "`byte_jets`.`name`"},
	BytePilotID:   whereHelpernull_Bytes{field: "`byte_jets`.`byte_pilot_id`"},
	ByteAirportID: whereHelpernull_Bytes{field: "`byte_jets`.`byte_airport_id`"},
}

// ByteJetRels is where relationship names are stored.
var ByteJetRels = struct {
	BytePilot   string
	ByteAirport string
}{
	BytePilot:   "BytePilot",
	ByteAirport: "ByteAirport",
}

// byteJetR is where relationships are stored.
type byteJetR struct {
	BytePilot   *BytePilot   `boil:"BytePilot" json:"BytePilot" toml:"BytePilot" yaml:"BytePilot"`
	ByteAirport *ByteAirport `boil:"ByteAirport" json:"ByteAirport" toml:"ByteAirport" yaml:"ByteAirport"`
}

// NewStruct creates a new relationship struct
func (*byteJetR) NewStruct() *byteJetR {
	return &byteJetR{}
}

func (r *byteJetR) GetBytePilot() *BytePilot {
	if r == nil {
		return nil
	}
	return r.BytePilot
}

func (r *byteJetR) GetByteAirport() *ByteAirport {
	if r == nil {
		return nil
	}
	return r.ByteAirport
}

// byteJetL is where Load methods for each relationship are stored.
type byteJetL struct{}

var (
	byteJetAllColumns            = []string{"id", "name", "byte_pilot_id", "byte_airport_id"}
	byteJetColumnsWithoutDefault = []string{"id", "name", "byte_pilot_id", "byte_airport_id"}
	byteJetColumnsWithDefault    = []string{}
	byteJetPrimaryKeyColumns     = []string{"id"}
	byteJetGeneratedColumns      = []string{}
)

type (
	// ByteJetSlice is an alias for a slice of pointers to ByteJet.
	// This should almost always be used instead of []ByteJet.
	ByteJetSlice []*ByteJet
	// ByteJetHook is the signature for custom ByteJet hook methods
	ByteJetHook func(context.Context, boil.ContextExecutor, *ByteJet) error

	byteJetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	byteJetType                 = reflect.TypeOf(&ByteJet{})
	byteJetMapping              = queries.MakeStructMapping(byteJetType)
	byteJetPrimaryKeyMapping, _ = queries.BindMapping(byteJetType, byteJetMapping, byteJetPrimaryKeyColumns)
	byteJetInsertCacheMut       sync.RWMutex
	byteJetInsertCache          = make(map[string]insertCache)
	byteJetUpdateCacheMut       sync.RWMutex
	byteJetUpdateCache          = make(map[string]updateCache)
	byteJetUpsertCacheMut       sync.RWMutex
	byteJetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var byteJetAfterSelectMu sync.Mutex
var byteJetAfterSelectHooks []ByteJetHook

var byteJetBeforeInsertMu sync.Mutex
var byteJetBeforeInsertHooks []ByteJetHook
var byteJetAfterInsertMu sync.Mutex
var byteJetAfterInsertHooks []ByteJetHook

var byteJetBeforeUpdateMu sync.Mutex
var byteJetBeforeUpdateHooks []ByteJetHook
var byteJetAfterUpdateMu sync.Mutex
var byteJetAfterUpdateHooks []ByteJetHook

var byteJetBeforeDeleteMu sync.Mutex
var byteJetBeforeDeleteHooks []ByteJetHook
var byteJetAfterDeleteMu sync.Mutex
var byteJetAfterDeleteHooks []ByteJetHook

var byteJetBeforeUpsertMu sync.Mutex
var byteJetBeforeUpsertHooks []ByteJetHook
var byteJetAfterUpsertMu sync.Mutex
var byteJetAfterUpsertHooks []ByteJetHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ByteJet) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range byteJetAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ByteJet) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range byteJetBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ByteJet) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range byteJetAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ByteJet) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range byteJetBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ByteJet) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range byteJetAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ByteJet) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range byteJetBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ByteJet) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range byteJetAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ByteJet) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range byteJetBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ByteJet) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range byteJetAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddByteJetHook registers your hook function for all future operations.
func AddByteJetHook(hookPoint boil.HookPoint, byteJetHook ByteJetHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		byteJetAfterSelectMu.Lock()
		byteJetAfterSelectHooks = append(byteJetAfterSelectHooks, byteJetHook)
		byteJetAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		byteJetBeforeInsertMu.Lock()
		byteJetBeforeInsertHooks = append(byteJetBeforeInsertHooks, byteJetHook)
		byteJetBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		byteJetAfterInsertMu.Lock()
		byteJetAfterInsertHooks = append(byteJetAfterInsertHooks, byteJetHook)
		byteJetAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		byteJetBeforeUpdateMu.Lock()
		byteJetBeforeUpdateHooks = append(byteJetBeforeUpdateHooks, byteJetHook)
		byteJetBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		byteJetAfterUpdateMu.Lock()
		byteJetAfterUpdateHooks = append(byteJetAfterUpdateHooks, byteJetHook)
		byteJetAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		byteJetBeforeDeleteMu.Lock()
		byteJetBeforeDeleteHooks = append(byteJetBeforeDeleteHooks, byteJetHook)
		byteJetBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		byteJetAfterDeleteMu.Lock()
		byteJetAfterDeleteHooks = append(byteJetAfterDeleteHooks, byteJetHook)
		byteJetAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		byteJetBeforeUpsertMu.Lock()
		byteJetBeforeUpsertHooks = append(byteJetBeforeUpsertHooks, byteJetHook)
		byteJetBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		byteJetAfterUpsertMu.Lock()
		byteJetAfterUpsertHooks = append(byteJetAfterUpsertHooks, byteJetHook)
		byteJetAfterUpsertMu.Unlock()
	}
}

// One returns a single byteJet record from the query.
func (q byteJetQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ByteJet, error) {
	o := &ByteJet{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for byte_jets")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ByteJet records from the query.
func (q byteJetQuery) All(ctx context.Context, exec boil.ContextExecutor) (ByteJetSlice, error) {
	var o []*ByteJet

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ByteJet slice")
	}

	if len(byteJetAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ByteJet records in the query.
func (q byteJetQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count byte_jets rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q byteJetQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if byte_jets exists")
	}

	return count > 0, nil
}

// BytePilot pointed to by the foreign key.
func (o *ByteJet) BytePilot(mods ...qm.QueryMod) bytePilotQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.BytePilotID),
	}

	queryMods = append(queryMods, mods...)

	return BytePilots(queryMods...)
}

// ByteAirport pointed to by the foreign key.
func (o *ByteJet) ByteAirport(mods ...qm.QueryMod) byteAirportQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ByteAirportID),
	}

	queryMods = append(queryMods, mods...)

	return ByteAirports(queryMods...)
}

// LoadBytePilot allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (byteJetL) LoadBytePilot(ctx context.Context, e boil.ContextExecutor, singular bool, maybeByteJet interface{}, mods queries.Applicator) error {
	var slice []*ByteJet
	var object *ByteJet

	if singular {
		var ok bool
		object, ok = maybeByteJet.(*ByteJet)
		if !ok {
			object = new(ByteJet)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeByteJet)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeByteJet))
			}
		}
	} else {
		s, ok := maybeByteJet.(*[]*ByteJet)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeByteJet)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeByteJet))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &byteJetR{}
		}
		if !queries.IsNil(object.BytePilotID) {
			args[object.BytePilotID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &byteJetR{}
			}

			if !queries.IsNil(obj.BytePilotID) {
				args[obj.BytePilotID] = struct{}{}
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
		qm.From(`byte_pilots`),
		qm.WhereIn(`byte_pilots.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load BytePilot")
	}

	var resultSlice []*BytePilot
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice BytePilot")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for byte_pilots")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for byte_pilots")
	}

	if len(bytePilotAfterSelectHooks) != 0 {
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
		object.R.BytePilot = foreign
		if foreign.R == nil {
			foreign.R = &bytePilotR{}
		}
		foreign.R.ByteJet = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.BytePilotID, foreign.ID) {
				local.R.BytePilot = foreign
				if foreign.R == nil {
					foreign.R = &bytePilotR{}
				}
				foreign.R.ByteJet = local
				break
			}
		}
	}

	return nil
}

// LoadByteAirport allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (byteJetL) LoadByteAirport(ctx context.Context, e boil.ContextExecutor, singular bool, maybeByteJet interface{}, mods queries.Applicator) error {
	var slice []*ByteJet
	var object *ByteJet

	if singular {
		var ok bool
		object, ok = maybeByteJet.(*ByteJet)
		if !ok {
			object = new(ByteJet)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeByteJet)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeByteJet))
			}
		}
	} else {
		s, ok := maybeByteJet.(*[]*ByteJet)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeByteJet)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeByteJet))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &byteJetR{}
		}
		if !queries.IsNil(object.ByteAirportID) {
			args[object.ByteAirportID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &byteJetR{}
			}

			if !queries.IsNil(obj.ByteAirportID) {
				args[obj.ByteAirportID] = struct{}{}
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
		qm.From(`byte_airports`),
		qm.WhereIn(`byte_airports.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ByteAirport")
	}

	var resultSlice []*ByteAirport
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ByteAirport")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for byte_airports")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for byte_airports")
	}

	if len(byteAirportAfterSelectHooks) != 0 {
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
		object.R.ByteAirport = foreign
		if foreign.R == nil {
			foreign.R = &byteAirportR{}
		}
		foreign.R.ByteJets = append(foreign.R.ByteJets, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ByteAirportID, foreign.ID) {
				local.R.ByteAirport = foreign
				if foreign.R == nil {
					foreign.R = &byteAirportR{}
				}
				foreign.R.ByteJets = append(foreign.R.ByteJets, local)
				break
			}
		}
	}

	return nil
}

// SetBytePilot of the byteJet to the related item.
// Sets o.R.BytePilot to related.
// Adds o to related.R.ByteJet.
func (o *ByteJet) SetBytePilot(ctx context.Context, exec boil.ContextExecutor, insert bool, related *BytePilot) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `byte_jets` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"byte_pilot_id"}),
		strmangle.WhereClause("`", "`", 0, byteJetPrimaryKeyColumns),
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

	queries.Assign(&o.BytePilotID, related.ID)
	if o.R == nil {
		o.R = &byteJetR{
			BytePilot: related,
		}
	} else {
		o.R.BytePilot = related
	}

	if related.R == nil {
		related.R = &bytePilotR{
			ByteJet: o,
		}
	} else {
		related.R.ByteJet = o
	}

	return nil
}

// RemoveBytePilot relationship.
// Sets o.R.BytePilot to nil.
// Removes o from all passed in related items' relationships struct.
func (o *ByteJet) RemoveBytePilot(ctx context.Context, exec boil.ContextExecutor, related *BytePilot) error {
	var err error

	queries.SetScanner(&o.BytePilotID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("byte_pilot_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.BytePilot = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	related.R.ByteJet = nil
	return nil
}

// SetByteAirport of the byteJet to the related item.
// Sets o.R.ByteAirport to related.
// Adds o to related.R.ByteJets.
func (o *ByteJet) SetByteAirport(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ByteAirport) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `byte_jets` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"byte_airport_id"}),
		strmangle.WhereClause("`", "`", 0, byteJetPrimaryKeyColumns),
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

	queries.Assign(&o.ByteAirportID, related.ID)
	if o.R == nil {
		o.R = &byteJetR{
			ByteAirport: related,
		}
	} else {
		o.R.ByteAirport = related
	}

	if related.R == nil {
		related.R = &byteAirportR{
			ByteJets: ByteJetSlice{o},
		}
	} else {
		related.R.ByteJets = append(related.R.ByteJets, o)
	}

	return nil
}

// RemoveByteAirport relationship.
// Sets o.R.ByteAirport to nil.
// Removes o from all passed in related items' relationships struct.
func (o *ByteJet) RemoveByteAirport(ctx context.Context, exec boil.ContextExecutor, related *ByteAirport) error {
	var err error

	queries.SetScanner(&o.ByteAirportID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("byte_airport_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.ByteAirport = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ByteJets {
		if queries.Equal(o.ByteAirportID, ri.ByteAirportID) {
			continue
		}

		ln := len(related.R.ByteJets)
		if ln > 1 && i < ln-1 {
			related.R.ByteJets[i] = related.R.ByteJets[ln-1]
		}
		related.R.ByteJets = related.R.ByteJets[:ln-1]
		break
	}
	return nil
}

// ByteJets retrieves all the records using an executor.
func ByteJets(mods ...qm.QueryMod) byteJetQuery {
	mods = append(mods, qm.From("`byte_jets`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`byte_jets`.*"})
	}

	return byteJetQuery{q}
}

// FindByteJet retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindByteJet(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*ByteJet, error) {
	byteJetObj := &ByteJet{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `byte_jets` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, byteJetObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from byte_jets")
	}

	if err = byteJetObj.doAfterSelectHooks(ctx, exec); err != nil {
		return byteJetObj, err
	}

	return byteJetObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ByteJet) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no byte_jets provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(byteJetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	byteJetInsertCacheMut.RLock()
	cache, cached := byteJetInsertCache[key]
	byteJetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			byteJetAllColumns,
			byteJetColumnsWithDefault,
			byteJetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(byteJetType, byteJetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(byteJetType, byteJetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `byte_jets` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `byte_jets` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `byte_jets` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, byteJetPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into byte_jets")
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
		return errors.Wrap(err, "models: unable to populate default values for byte_jets")
	}

CacheNoHooks:
	if !cached {
		byteJetInsertCacheMut.Lock()
		byteJetInsertCache[key] = cache
		byteJetInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ByteJet.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ByteJet) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	byteJetUpdateCacheMut.RLock()
	cache, cached := byteJetUpdateCache[key]
	byteJetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			byteJetAllColumns,
			byteJetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update byte_jets, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `byte_jets` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, byteJetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(byteJetType, byteJetMapping, append(wl, byteJetPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update byte_jets row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for byte_jets")
	}

	if !cached {
		byteJetUpdateCacheMut.Lock()
		byteJetUpdateCache[key] = cache
		byteJetUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q byteJetQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for byte_jets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for byte_jets")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ByteJetSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), byteJetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `byte_jets` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, byteJetPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in byteJet slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all byteJet")
	}
	return rowsAff, nil
}

var mySQLByteJetUniqueColumns = []string{
	"id",
	"byte_pilot_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ByteJet) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no byte_jets provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(byteJetColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLByteJetUniqueColumns, o)

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

	byteJetUpsertCacheMut.RLock()
	cache, cached := byteJetUpsertCache[key]
	byteJetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			byteJetAllColumns,
			byteJetColumnsWithDefault,
			byteJetColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			byteJetAllColumns,
			byteJetPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert byte_jets, could not build update column list")
		}

		ret := strmangle.SetComplement(byteJetAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`byte_jets`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `byte_jets` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(byteJetType, byteJetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(byteJetType, byteJetMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for byte_jets")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(byteJetType, byteJetMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for byte_jets")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for byte_jets")
	}

CacheNoHooks:
	if !cached {
		byteJetUpsertCacheMut.Lock()
		byteJetUpsertCache[key] = cache
		byteJetUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ByteJet record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ByteJet) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ByteJet provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), byteJetPrimaryKeyMapping)
	sql := "DELETE FROM `byte_jets` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from byte_jets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for byte_jets")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q byteJetQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no byteJetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from byte_jets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for byte_jets")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ByteJetSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(byteJetBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), byteJetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `byte_jets` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, byteJetPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from byteJet slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for byte_jets")
	}

	if len(byteJetAfterDeleteHooks) != 0 {
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
func (o *ByteJet) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindByteJet(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ByteJetSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ByteJetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), byteJetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `byte_jets`.* FROM `byte_jets` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, byteJetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ByteJetSlice")
	}

	*o = slice

	return nil
}

// ByteJetExists checks if the ByteJet row exists.
func ByteJetExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `byte_jets` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if byte_jets exists")
	}

	return exists, nil
}

// Exists checks if the ByteJet row exists.
func (o *ByteJet) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ByteJetExists(ctx, exec, o.ID)
}
