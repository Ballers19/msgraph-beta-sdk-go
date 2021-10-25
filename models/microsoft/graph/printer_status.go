package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrinterStatus struct {
    additionalData map[string]interface{};
    description *string;
    details []PrinterProcessingStateDetail;
    processingState *PrinterProcessingState;
    processingStateDescription *string;
    processingStateReasons []PrinterProcessingStateReason;
    state *PrinterProcessingState;
}
func NewPrinterStatus()(*PrinterStatus) {
    m := &PrinterStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PrinterStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PrinterStatus) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *PrinterStatus) GetDetails()([]PrinterProcessingStateDetail) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
func (m *PrinterStatus) GetProcessingState()(*PrinterProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.processingState
    }
}
func (m *PrinterStatus) GetProcessingStateDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.processingStateDescription
    }
}
func (m *PrinterStatus) GetProcessingStateReasons()([]PrinterProcessingStateReason) {
    if m == nil {
        return nil
    } else {
        return m.processingStateReasons
    }
}
func (m *PrinterStatus) GetState()(*PrinterProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *PrinterStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrinterProcessingStateDetail)
        if err != nil {
            return err
        }
        res := make([]PrinterProcessingStateDetail, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrinterProcessingStateDetail))
        }
        m.SetDetails(res)
        return nil
    }
    res["processingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrinterProcessingState)
        if err != nil {
            return err
        }
        cast := val.(PrinterProcessingState)
        m.SetProcessingState(&cast)
        return nil
    }
    res["processingStateDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProcessingStateDescription(val)
        return nil
    }
    res["processingStateReasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrinterProcessingStateReason)
        if err != nil {
            return err
        }
        res := make([]PrinterProcessingStateReason, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrinterProcessingStateReason))
        }
        m.SetProcessingStateReasons(res)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrinterProcessingState)
        if err != nil {
            return err
        }
        cast := val.(PrinterProcessingState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *PrinterStatus) IsNil()(bool) {
    return m == nil
}
func (m *PrinterStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("details", SerializePrinterProcessingStateDetail(m.GetDetails()))
        if err != nil {
            return err
        }
    }
    if m.GetProcessingState() != nil {
        cast := m.GetProcessingState().String()
        err := writer.WriteStringValue("processingState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("processingStateDescription", m.GetProcessingStateDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("processingStateReasons", SerializePrinterProcessingStateReason(m.GetProcessingStateReasons()))
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *PrinterStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PrinterStatus) SetDescription(value *string)() {
    m.description = value
}
func (m *PrinterStatus) SetDetails(value []PrinterProcessingStateDetail)() {
    m.details = value
}
func (m *PrinterStatus) SetProcessingState(value *PrinterProcessingState)() {
    m.processingState = value
}
func (m *PrinterStatus) SetProcessingStateDescription(value *string)() {
    m.processingStateDescription = value
}
func (m *PrinterStatus) SetProcessingStateReasons(value []PrinterProcessingStateReason)() {
    m.processingStateReasons = value
}
func (m *PrinterStatus) SetState(value *PrinterProcessingState)() {
    m.state = value
}
