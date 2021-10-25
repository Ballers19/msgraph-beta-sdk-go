package reminderviewwithstartdatetimewithenddatetime

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type ReminderViewWithStartDateTimeWithEndDateTime struct {
    additionalData map[string]interface{};
    changeKey *string;
    eventEndTime *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone;
    eventId *string;
    eventLocation *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location;
    eventStartTime *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone;
    eventSubject *string;
    eventWebLink *string;
    reminderFireTime *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone;
}
func NewReminderViewWithStartDateTimeWithEndDateTime()(*ReminderViewWithStartDateTimeWithEndDateTime) {
    m := &ReminderViewWithStartDateTimeWithEndDateTime{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventEndTime()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.eventEndTime
    }
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventId
    }
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventLocation()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location) {
    if m == nil {
        return nil
    } else {
        return m.eventLocation
    }
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventStartTime()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.eventStartTime
    }
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventSubject
    }
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventWebLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventWebLink
    }
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetReminderFireTime()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.reminderFireTime
    }
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["changeKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChangeKey(val)
        return nil
    }
    res["eventEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetEventEndTime(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone))
        return nil
    }
    res["eventId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEventId(val)
        return nil
    }
    res["eventLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewLocation() })
        if err != nil {
            return err
        }
        m.SetEventLocation(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location))
        return nil
    }
    res["eventStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetEventStartTime(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone))
        return nil
    }
    res["eventSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEventSubject(val)
        return nil
    }
    res["eventWebLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEventWebLink(val)
        return nil
    }
    res["reminderFireTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetReminderFireTime(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone))
        return nil
    }
    return res
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) IsNil()(bool) {
    return m == nil
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventEndTime", m.GetEventEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventId", m.GetEventId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventLocation", m.GetEventLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventStartTime", m.GetEventStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventSubject", m.GetEventSubject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventWebLink", m.GetEventWebLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reminderFireTime", m.GetReminderFireTime())
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
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetChangeKey(value *string)() {
    m.changeKey = value
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventEndTime(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone)() {
    m.eventEndTime = value
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventId(value *string)() {
    m.eventId = value
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventLocation(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Location)() {
    m.eventLocation = value
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventStartTime(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone)() {
    m.eventStartTime = value
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventSubject(value *string)() {
    m.eventSubject = value
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventWebLink(value *string)() {
    m.eventWebLink = value
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetReminderFireTime(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DateTimeTimeZone)() {
    m.reminderFireTime = value
}
