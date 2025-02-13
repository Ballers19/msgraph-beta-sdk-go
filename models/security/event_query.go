package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EventQuery 
type EventQuery struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The query property
    query *string
    // The queryType property
    queryType *QueryType
}
// NewEventQuery instantiates a new eventQuery and sets the default values.
func NewEventQuery()(*EventQuery) {
    m := &EventQuery{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEventQueryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEventQueryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEventQuery(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EventQuery) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EventQuery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["query"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val)
        }
        return nil
    }
    res["queryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseQueryType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryType(val.(*QueryType))
        }
        return nil
    }
    return res
}
// GetQuery gets the query property value. The query property
func (m *EventQuery) GetQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.query
    }
}
// GetQueryType gets the queryType property value. The queryType property
func (m *EventQuery) GetQueryType()(*QueryType) {
    if m == nil {
        return nil
    } else {
        return m.queryType
    }
}
// Serialize serializes information the current object
func (m *EventQuery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    if m.GetQueryType() != nil {
        cast := (*m.GetQueryType()).String()
        err := writer.WriteStringValue("queryType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EventQuery) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetQuery sets the query property value. The query property
func (m *EventQuery) SetQuery(value *string)() {
    if m != nil {
        m.query = value
    }
}
// SetQueryType sets the queryType property value. The queryType property
func (m *EventQuery) SetQueryType(value *QueryType)() {
    if m != nil {
        m.queryType = value
    }
}
