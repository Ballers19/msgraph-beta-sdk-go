package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchHit 
type SearchHit struct {
    // The _id property
    _id *string
    // The _score property
    _score *int32
    // The _source property
    _source Entityable
    // The _summary property
    _summary *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the content source which the externalItem is part of .
    contentSource *string
    // The internal identifier for the item.
    hitId *string
    // The rank or the order of the result.
    rank *int32
    // The resource property
    resource Entityable
    // ID of the result template for rendering the search result. This ID must map to a display layout in the resultTemplates dictionary, included in the searchresponse as well.
    resultTemplateId *string
    // A summary of the result, if a summary is available.
    summary *string
}
// NewSearchHit instantiates a new searchHit and sets the default values.
func NewSearchHit()(*SearchHit) {
    m := &SearchHit{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSearchHitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchHitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchHit(), nil
}
// Get_id gets the _id property value. The _id property
func (m *SearchHit) Get_id()(*string) {
    if m == nil {
        return nil
    } else {
        return m._id
    }
}
// Get_score gets the _score property value. The _score property
func (m *SearchHit) Get_score()(*int32) {
    if m == nil {
        return nil
    } else {
        return m._score
    }
}
// Get_source gets the _source property value. The _source property
func (m *SearchHit) Get_source()(Entityable) {
    if m == nil {
        return nil
    } else {
        return m._source
    }
}
// Get_summary gets the _summary property value. The _summary property
func (m *SearchHit) Get_summary()(*string) {
    if m == nil {
        return nil
    } else {
        return m._summary
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchHit) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentSource gets the contentSource property value. The name of the content source which the externalItem is part of .
func (m *SearchHit) GetContentSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentSource
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchHit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.Set_id(val)
        }
        return nil
    }
    res["_score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.Set_score(val)
        }
        return nil
    }
    res["_source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.Set_source(val.(Entityable))
        }
        return nil
    }
    res["_summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.Set_summary(val)
        }
        return nil
    }
    res["contentSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentSource(val)
        }
        return nil
    }
    res["hitId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHitId(val)
        }
        return nil
    }
    res["rank"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(Entityable))
        }
        return nil
    }
    res["resultTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultTemplateId(val)
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    return res
}
// GetHitId gets the hitId property value. The internal identifier for the item.
func (m *SearchHit) GetHitId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hitId
    }
}
// GetRank gets the rank property value. The rank or the order of the result.
func (m *SearchHit) GetRank()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rank
    }
}
// GetResource gets the resource property value. The resource property
func (m *SearchHit) GetResource()(Entityable) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResultTemplateId gets the resultTemplateId property value. ID of the result template for rendering the search result. This ID must map to a display layout in the resultTemplates dictionary, included in the searchresponse as well.
func (m *SearchHit) GetResultTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resultTemplateId
    }
}
// GetSummary gets the summary property value. A summary of the result, if a summary is available.
func (m *SearchHit) GetSummary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
// Serialize serializes information the current object
func (m *SearchHit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("_id", m.Get_id())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("_score", m.Get_score())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("_source", m.Get_source())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("_summary", m.Get_summary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentSource", m.GetContentSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hitId", m.GetHitId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("rank", m.GetRank())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resultTemplateId", m.GetResultTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
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
// Set_id sets the _id property value. The _id property
func (m *SearchHit) Set_id(value *string)() {
    if m != nil {
        m._id = value
    }
}
// Set_score sets the _score property value. The _score property
func (m *SearchHit) Set_score(value *int32)() {
    if m != nil {
        m._score = value
    }
}
// Set_source sets the _source property value. The _source property
func (m *SearchHit) Set_source(value Entityable)() {
    if m != nil {
        m._source = value
    }
}
// Set_summary sets the _summary property value. The _summary property
func (m *SearchHit) Set_summary(value *string)() {
    if m != nil {
        m._summary = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchHit) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentSource sets the contentSource property value. The name of the content source which the externalItem is part of .
func (m *SearchHit) SetContentSource(value *string)() {
    if m != nil {
        m.contentSource = value
    }
}
// SetHitId sets the hitId property value. The internal identifier for the item.
func (m *SearchHit) SetHitId(value *string)() {
    if m != nil {
        m.hitId = value
    }
}
// SetRank sets the rank property value. The rank or the order of the result.
func (m *SearchHit) SetRank(value *int32)() {
    if m != nil {
        m.rank = value
    }
}
// SetResource sets the resource property value. The resource property
func (m *SearchHit) SetResource(value Entityable)() {
    if m != nil {
        m.resource = value
    }
}
// SetResultTemplateId sets the resultTemplateId property value. ID of the result template for rendering the search result. This ID must map to a display layout in the resultTemplates dictionary, included in the searchresponse as well.
func (m *SearchHit) SetResultTemplateId(value *string)() {
    if m != nil {
        m.resultTemplateId = value
    }
}
// SetSummary sets the summary property value. A summary of the result, if a summary is available.
func (m *SearchHit) SetSummary(value *string)() {
    if m != nil {
        m.summary = value
    }
}
