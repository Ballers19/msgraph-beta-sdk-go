package getteamsdeviceusagedistributiontotalusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod struct {
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
// Instantiates a new getTeamsDeviceUsageDistributionTotalUserCountsWithPeriod and sets the default values.
func NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod()(*GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) {
    m := &GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the androidPhone property value. The number of users who were active on the Teams mobile client for Android.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetAndroidPhone()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.androidPhone
    }
}
// Gets the chromeOS property value. The number of users who were active in the Teams desktop client on a ChromeOS computer.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetChromeOS()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.chromeOS
    }
}
// Gets the ios property value. The number of users who were active on the Teams mobile client for iOS.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetIos()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.ios
    }
}
// Gets the linux property value. The number of users who were active in the Teams desktop client on a Linux computer.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetLinux()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.linux
    }
}
// Gets the mac property value. The number of users who were active in the Teams desktop client on a macOS computer.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mac
    }
}
// Gets the reportPeriod property value. The number of days the report covers.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. The latest date of the content.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the web property value. The number of users who were active in the Teams web client on devices.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetWeb()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.web
    }
}
// Gets the windows property value. The number of users who were active in the Teams desktop client on a Windows-based computer.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetWindows()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows
    }
}
// Gets the windowsPhone property value. The number of users who were active on the Teams mobile client for Windows phone.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetWindowsPhone()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windowsPhone
    }
}
// The deserialization information for the current model
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) SetAndroidPhone(value *int64)() {
    m.androidPhone = value
}
// Sets the chromeOS property value. The number of users who were active in the Teams desktop client on a ChromeOS computer.
// Parameters:
//  - value : Value to set for the chromeOS property.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) SetChromeOS(value *int64)() {
    m.chromeOS = value
}
// Sets the ios property value. The number of users who were active on the Teams mobile client for iOS.
// Parameters:
//  - value : Value to set for the ios property.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) SetIos(value *int64)() {
    m.ios = value
}
// Sets the linux property value. The number of users who were active in the Teams desktop client on a Linux computer.
// Parameters:
//  - value : Value to set for the linux property.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) SetLinux(value *int64)() {
    m.linux = value
}
// Sets the mac property value. The number of users who were active in the Teams desktop client on a macOS computer.
// Parameters:
//  - value : Value to set for the mac property.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) SetMac(value *int64)() {
    m.mac = value
}
// Sets the reportPeriod property value. The number of days the report covers.
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. The latest date of the content.
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the web property value. The number of users who were active in the Teams web client on devices.
// Parameters:
//  - value : Value to set for the web property.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) SetWeb(value *int64)() {
    m.web = value
}
// Sets the windows property value. The number of users who were active in the Teams desktop client on a Windows-based computer.
// Parameters:
//  - value : Value to set for the windows property.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) SetWindows(value *int64)() {
    m.windows = value
}
// Sets the windowsPhone property value. The number of users who were active on the Teams mobile client for Windows phone.
// Parameters:
//  - value : Value to set for the windowsPhone property.
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod) SetWindowsPhone(value *int64)() {
    m.windowsPhone = value
}
