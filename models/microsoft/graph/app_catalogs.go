package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AppCatalogs struct {
    additionalData map[string]interface{};
    teamsApps []TeamsApp;
}
func NewAppCatalogs()(*AppCatalogs) {
    m := &AppCatalogs{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AppCatalogs) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AppCatalogs) GetTeamsApps()([]TeamsApp) {
    if m == nil {
        return nil
    } else {
        return m.teamsApps
    }
}
func (m *AppCatalogs) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["teamsApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsApp() })
        if err != nil {
            return err
        }
        res := make([]TeamsApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*TeamsApp))
        }
        m.SetTeamsApps(res)
        return nil
    }
    return res
}
func (m *AppCatalogs) IsNil()(bool) {
    return m == nil
}
func (m *AppCatalogs) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTeamsApps()))
        for i, v := range m.GetTeamsApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("teamsApps", cast)
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
func (m *AppCatalogs) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AppCatalogs) SetTeamsApps(value []TeamsApp)() {
    m.teamsApps = value
}
