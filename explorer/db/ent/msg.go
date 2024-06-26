// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/omni-network/omni/explorer/db/ent/msg"
)

// Msg is the model entity for the Msg schema.
type Msg struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Sender holds the value of the "sender" field.
	Sender []byte `json:"sender,omitempty"`
	// To holds the value of the "to" field.
	To []byte `json:"to,omitempty"`
	// Data holds the value of the "data" field.
	Data []byte `json:"data,omitempty"`
	// GasLimit holds the value of the "gas_limit" field.
	GasLimit uint64 `json:"gas_limit,omitempty"`
	// SourceChainID holds the value of the "source_chain_id" field.
	SourceChainID uint64 `json:"source_chain_id,omitempty"`
	// DestChainID holds the value of the "dest_chain_id" field.
	DestChainID uint64 `json:"dest_chain_id,omitempty"`
	// Offset holds the value of the "offset" field.
	Offset uint64 `json:"offset,omitempty"`
	// TxHash holds the value of the "tx_hash" field.
	TxHash []byte `json:"tx_hash,omitempty"`
	// ReceiptHash holds the value of the "receipt_hash" field.
	ReceiptHash []byte `json:"receipt_hash,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MsgQuery when eager-loading is set.
	Edges        MsgEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MsgEdges holds the relations/edges for other nodes in the graph.
type MsgEdges struct {
	// Block holds the value of the block edge.
	Block []*Block `json:"block,omitempty"`
	// Receipts holds the value of the receipts edge.
	Receipts []*Receipt `json:"receipts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BlockOrErr returns the Block value or an error if the edge
// was not loaded in eager-loading.
func (e MsgEdges) BlockOrErr() ([]*Block, error) {
	if e.loadedTypes[0] {
		return e.Block, nil
	}
	return nil, &NotLoadedError{edge: "block"}
}

// ReceiptsOrErr returns the Receipts value or an error if the edge
// was not loaded in eager-loading.
func (e MsgEdges) ReceiptsOrErr() ([]*Receipt, error) {
	if e.loadedTypes[1] {
		return e.Receipts, nil
	}
	return nil, &NotLoadedError{edge: "receipts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Msg) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case msg.FieldSender, msg.FieldTo, msg.FieldData, msg.FieldTxHash, msg.FieldReceiptHash:
			values[i] = new([]byte)
		case msg.FieldID, msg.FieldGasLimit, msg.FieldSourceChainID, msg.FieldDestChainID, msg.FieldOffset:
			values[i] = new(sql.NullInt64)
		case msg.FieldStatus:
			values[i] = new(sql.NullString)
		case msg.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Msg fields.
func (m *Msg) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case msg.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case msg.FieldSender:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field sender", values[i])
			} else if value != nil {
				m.Sender = *value
			}
		case msg.FieldTo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value != nil {
				m.To = *value
			}
		case msg.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil {
				m.Data = *value
			}
		case msg.FieldGasLimit:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gas_limit", values[i])
			} else if value.Valid {
				m.GasLimit = uint64(value.Int64)
			}
		case msg.FieldSourceChainID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source_chain_id", values[i])
			} else if value.Valid {
				m.SourceChainID = uint64(value.Int64)
			}
		case msg.FieldDestChainID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dest_chain_id", values[i])
			} else if value.Valid {
				m.DestChainID = uint64(value.Int64)
			}
		case msg.FieldOffset:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field offset", values[i])
			} else if value.Valid {
				m.Offset = uint64(value.Int64)
			}
		case msg.FieldTxHash:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tx_hash", values[i])
			} else if value != nil {
				m.TxHash = *value
			}
		case msg.FieldReceiptHash:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field receipt_hash", values[i])
			} else if value != nil {
				m.ReceiptHash = *value
			}
		case msg.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = value.String
			}
		case msg.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Msg.
// This includes values selected through modifiers, order, etc.
func (m *Msg) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryBlock queries the "block" edge of the Msg entity.
func (m *Msg) QueryBlock() *BlockQuery {
	return NewMsgClient(m.config).QueryBlock(m)
}

// QueryReceipts queries the "receipts" edge of the Msg entity.
func (m *Msg) QueryReceipts() *ReceiptQuery {
	return NewMsgClient(m.config).QueryReceipts(m)
}

// Update returns a builder for updating this Msg.
// Note that you need to call Msg.Unwrap() before calling this method if this Msg
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Msg) Update() *MsgUpdateOne {
	return NewMsgClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Msg entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Msg) Unwrap() *Msg {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Msg is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Msg) String() string {
	var builder strings.Builder
	builder.WriteString("Msg(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("sender=")
	builder.WriteString(fmt.Sprintf("%v", m.Sender))
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(fmt.Sprintf("%v", m.To))
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", m.Data))
	builder.WriteString(", ")
	builder.WriteString("gas_limit=")
	builder.WriteString(fmt.Sprintf("%v", m.GasLimit))
	builder.WriteString(", ")
	builder.WriteString("source_chain_id=")
	builder.WriteString(fmt.Sprintf("%v", m.SourceChainID))
	builder.WriteString(", ")
	builder.WriteString("dest_chain_id=")
	builder.WriteString(fmt.Sprintf("%v", m.DestChainID))
	builder.WriteString(", ")
	builder.WriteString("offset=")
	builder.WriteString(fmt.Sprintf("%v", m.Offset))
	builder.WriteString(", ")
	builder.WriteString("tx_hash=")
	builder.WriteString(fmt.Sprintf("%v", m.TxHash))
	builder.WriteString(", ")
	builder.WriteString("receipt_hash=")
	builder.WriteString(fmt.Sprintf("%v", m.ReceiptHash))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(m.Status)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Msgs is a parsable slice of Msg.
type Msgs []*Msg
