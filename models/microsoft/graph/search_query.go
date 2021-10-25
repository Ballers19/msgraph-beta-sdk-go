package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchQuery struct {
    additionalData map[string]interface{};
    query_string *SearchQueryString;
    queryString *string;
}
func NewSearchQuery()(*SearchQuery) {
    m := &SearchQuery{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchQuery) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchQuery) GetQuery_string()(*SearchQueryString) {
    if m == nil {
        return nil
    } else {
        return m.query_string
    }
}
func (m *SearchQuery) GetQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryString
    }
}
func (m *SearchQuery) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["query_string"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchQueryString() })
        if err != nil {
            return err
        }
        m.SetQuery_string(val.(*SearchQueryString))
        return nil
    }
    res["queryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQueryString(val)
        return nil
    }
    return res
}
func (m *SearchQuery) IsNil()(bool) {
    return m == nil
}
func (m *SearchQuery) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SearchQuery) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchQuery) SetQuery_string(value *SearchQueryString)() {
    m.query_string = value
}
func (m *SearchQuery) SetQueryString(value *string)() {
    m.queryString = value
}
