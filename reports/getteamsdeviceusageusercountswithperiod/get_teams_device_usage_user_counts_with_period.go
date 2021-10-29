package getteamsdeviceusageusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetTeamsDeviceUsageUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The number of users who were active on the Teams mobile client for Android.
    androidPhone *int64;
    // The number of users who were active in the Teams desktop client on a ChromeOS computer.
    chromeOS *int64;
    // The number of users who were active on the Teams mobile client for iOS.
    ios *int64;
    // The number of users who were active in the Teams desktop client on a Linux computer.
    linux *int64;
    // The number of users who were active in the Teams desktop client on a macOS computer.
    mac *int64;
    // The date on which the users performed the activities.
    reportDate *string;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The number of users who were active in the Teams web client on devices.
    web *int64;
    // The number of users who were active in the Teams desktop client on a Windows-based computer.
    windows *int64;
    // The number of users who were active on the Teams mobile client for Windows phone.
    windowsPhone *int64;
}
// Instantiates a new getTeamsDeviceUsageUserCountsWithPeriod and sets the default values.
func NewGetTeamsDeviceUsageUserCountsWithPeriod()(*GetTeamsDeviceUsageUserCountsWithPeriod) {
    m := &GetTeamsDeviceUsageUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the androidPhone property value. The number of users who were active on the Teams mobile client for Android.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetAndroidPhone()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.androidPhone
    }
}
// Gets the chromeOS property value. The number of users who were active in the Teams desktop client on a ChromeOS computer.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetChromeOS()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.chromeOS
    }
}
// Gets the ios property value. The number of users who were active on the Teams mobile client for iOS.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetIos()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.ios
    }
}
// Gets the linux property value. The number of users who were active in the Teams desktop client on a Linux computer.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetLinux()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.linux
    }
}
// Gets the mac property value. The number of users who were active in the Teams desktop client on a macOS computer.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mac
    }
}
// Gets the reportDate property value. The date on which the users performed the activities.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// Gets the reportPeriod property value. The number of days the report covers.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. The latest date of the content.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the web property value. The number of users who were active in the Teams web client on devices.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetWeb()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.web
    }
}
// Gets the windows property value. The number of users who were active in the Teams desktop client on a Windows-based computer.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetWindows()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows
    }
}
// Gets the windowsPhone property value. The number of users who were active on the Teams mobile client for Windows phone.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetWindowsPhone()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windowsPhone
    }
}
// The deserialization information for the current model
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAndroidPhone(val)
        return nil
    }
    res["chromeOS"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetChromeOS(val)
        return nil
    }
    res["ios"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetIos(val)
        return nil
    }
    res["linux"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetLinux(val)
        return nil
    }
    res["mac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMac(val)
        return nil
    }
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportDate(val)
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportPeriod(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["web"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetWeb(val)
        return nil
    }
    res["windows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetWindows(val)
        return nil
    }
    res["windowsPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetWindowsPhone(val)
        return nil
    }
    return res
}
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("androidPhone", m.GetAndroidPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("chromeOS", m.GetChromeOS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("ios", m.GetIos())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("linux", m.GetLinux())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("mac", m.GetMac())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportDate", m.GetReportDate())
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
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("web", m.GetWeb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("windows", m.GetWindows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("windowsPhone", m.GetWindowsPhone())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the androidPhone property value. The number of users who were active on the Teams mobile client for Android.
// Parameters:
//  - value : Value to set for the androidPhone property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetAndroidPhone(value *int64)() {
    m.androidPhone = value
}
// Sets the chromeOS property value. The number of users who were active in the Teams desktop client on a ChromeOS computer.
// Parameters:
//  - value : Value to set for the chromeOS property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetChromeOS(value *int64)() {
    m.chromeOS = value
}
// Sets the ios property value. The number of users who were active on the Teams mobile client for iOS.
// Parameters:
//  - value : Value to set for the ios property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetIos(value *int64)() {
    m.ios = value
}
// Sets the linux property value. The number of users who were active in the Teams desktop client on a Linux computer.
// Parameters:
//  - value : Value to set for the linux property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetLinux(value *int64)() {
    m.linux = value
}
// Sets the mac property value. The number of users who were active in the Teams desktop client on a macOS computer.
// Parameters:
//  - value : Value to set for the mac property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetMac(value *int64)() {
    m.mac = value
}
// Sets the reportDate property value. The date on which the users performed the activities.
// Parameters:
//  - value : Value to set for the reportDate property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// Sets the reportPeriod property value. The number of days the report covers.
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. The latest date of the content.
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the web property value. The number of users who were active in the Teams web client on devices.
// Parameters:
//  - value : Value to set for the web property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetWeb(value *int64)() {
    m.web = value
}
// Sets the windows property value. The number of users who were active in the Teams desktop client on a Windows-based computer.
// Parameters:
//  - value : Value to set for the windows property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetWindows(value *int64)() {
    m.windows = value
}
// Sets the windowsPhone property value. The number of users who were active on the Teams mobile client for Windows phone.
// Parameters:
//  - value : Value to set for the windowsPhone property.
func (m *GetTeamsDeviceUsageUserCountsWithPeriod) SetWindowsPhone(value *int64)() {
    m.windowsPhone = value
}
