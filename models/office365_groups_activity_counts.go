package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365GroupsActivityCounts 
type Office365GroupsActivityCounts struct {
    Entity
    // The number of emails received by Group mailboxes.
    exchangeEmailsReceived *int64
    // The date on which a number of emails were sent to a group mailbox or a number of messages were posted, read, or liked in a Yammer group
    reportDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The number of days the report covers.
    reportPeriod *string
    // The latest date of the content.
    reportRefreshDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The number of messages liked in Yammer groups.
    yammerMessagesLiked *int64
    // The number of messages posted to Yammer groups.
    yammerMessagesPosted *int64
    // The number of messages read in Yammer groups.
    yammerMessagesRead *int64
}
// NewOffice365GroupsActivityCounts instantiates a new Office365GroupsActivityCounts and sets the default values.
func NewOffice365GroupsActivityCounts()(*Office365GroupsActivityCounts) {
    m := &Office365GroupsActivityCounts{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOffice365GroupsActivityCountsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOffice365GroupsActivityCountsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOffice365GroupsActivityCounts(), nil
}
// GetExchangeEmailsReceived gets the exchangeEmailsReceived property value. The number of emails received by Group mailboxes.
func (m *Office365GroupsActivityCounts) GetExchangeEmailsReceived()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeEmailsReceived
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365GroupsActivityCounts) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchangeEmailsReceived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeEmailsReceived(val)
        }
        return nil
    }
    res["reportDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportDate(val)
        }
        return nil
    }
    res["reportPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["yammerMessagesLiked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerMessagesLiked(val)
        }
        return nil
    }
    res["yammerMessagesPosted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerMessagesPosted(val)
        }
        return nil
    }
    res["yammerMessagesRead"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerMessagesRead(val)
        }
        return nil
    }
    return res
}
// GetReportDate gets the reportDate property value. The date on which a number of emails were sent to a group mailbox or a number of messages were posted, read, or liked in a Yammer group
func (m *Office365GroupsActivityCounts) GetReportDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityCounts) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityCounts) GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetYammerMessagesLiked gets the yammerMessagesLiked property value. The number of messages liked in Yammer groups.
func (m *Office365GroupsActivityCounts) GetYammerMessagesLiked()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerMessagesLiked
    }
}
// GetYammerMessagesPosted gets the yammerMessagesPosted property value. The number of messages posted to Yammer groups.
func (m *Office365GroupsActivityCounts) GetYammerMessagesPosted()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerMessagesPosted
    }
}
// GetYammerMessagesRead gets the yammerMessagesRead property value. The number of messages read in Yammer groups.
func (m *Office365GroupsActivityCounts) GetYammerMessagesRead()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerMessagesRead
    }
}
// Serialize serializes information the current object
func (m *Office365GroupsActivityCounts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("exchangeEmailsReceived", m.GetExchangeEmailsReceived())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("reportDate", m.GetReportDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerMessagesLiked", m.GetYammerMessagesLiked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerMessagesPosted", m.GetYammerMessagesPosted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerMessagesRead", m.GetYammerMessagesRead())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExchangeEmailsReceived sets the exchangeEmailsReceived property value. The number of emails received by Group mailboxes.
func (m *Office365GroupsActivityCounts) SetExchangeEmailsReceived(value *int64)() {
    if m != nil {
        m.exchangeEmailsReceived = value
    }
}
// SetReportDate sets the reportDate property value. The date on which a number of emails were sent to a group mailbox or a number of messages were posted, read, or liked in a Yammer group
func (m *Office365GroupsActivityCounts) SetReportDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.reportDate = value
    }
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityCounts) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityCounts) SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetYammerMessagesLiked sets the yammerMessagesLiked property value. The number of messages liked in Yammer groups.
func (m *Office365GroupsActivityCounts) SetYammerMessagesLiked(value *int64)() {
    if m != nil {
        m.yammerMessagesLiked = value
    }
}
// SetYammerMessagesPosted sets the yammerMessagesPosted property value. The number of messages posted to Yammer groups.
func (m *Office365GroupsActivityCounts) SetYammerMessagesPosted(value *int64)() {
    if m != nil {
        m.yammerMessagesPosted = value
    }
}
// SetYammerMessagesRead sets the yammerMessagesRead property value. The number of messages read in Yammer groups.
func (m *Office365GroupsActivityCounts) SetYammerMessagesRead(value *int64)() {
    if m != nil {
        m.yammerMessagesRead = value
    }
}
