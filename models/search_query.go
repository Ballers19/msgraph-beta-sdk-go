package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchQuery 
type SearchQuery struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The query_string property
    query_string SearchQueryStringable
    // The search query containing the search terms. Required.
    queryString *string
    // Provides a way to decorate the query string. Supports both KQL and query variables. Optional.
    queryTemplate *string
}
// NewSearchQuery instantiates a new searchQuery and sets the default values.
func NewSearchQuery()(*SearchQuery) {
    m := &SearchQuery{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSearchQueryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchQueryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchQuery(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchQuery) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchQuery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["query_string"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchQueryStringFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery_string(val.(SearchQueryStringable))
        }
        return nil
    }
    res["queryString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryString(val)
        }
        return nil
    }
    res["queryTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryTemplate(val)
        }
        return nil
    }
    return res
}
// GetQuery_string gets the query_string property value. The query_string property
func (m *SearchQuery) GetQuery_string()(SearchQueryStringable) {
    if m == nil {
        return nil
    } else {
        return m.query_string
    }
}
// GetQueryString gets the queryString property value. The search query containing the search terms. Required.
func (m *SearchQuery) GetQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryString
    }
}
// GetQueryTemplate gets the queryTemplate property value. Provides a way to decorate the query string. Supports both KQL and query variables. Optional.
func (m *SearchQuery) GetQueryTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryTemplate
    }
}
// Serialize serializes information the current object
func (m *SearchQuery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("query_string", m.GetQuery_string())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryString", m.GetQueryString())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryTemplate", m.GetQueryTemplate())
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
func (m *SearchQuery) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetQuery_string sets the query_string property value. The query_string property
func (m *SearchQuery) SetQuery_string(value SearchQueryStringable)() {
    if m != nil {
        m.query_string = value
    }
}
// SetQueryString sets the queryString property value. The search query containing the search terms. Required.
func (m *SearchQuery) SetQueryString(value *string)() {
    if m != nil {
        m.queryString = value
    }
}
// SetQueryTemplate sets the queryTemplate property value. Provides a way to decorate the query string. Supports both KQL and query variables. Optional.
func (m *SearchQuery) SetQueryTemplate(value *string)() {
    if m != nil {
        m.queryTemplate = value
    }
}
