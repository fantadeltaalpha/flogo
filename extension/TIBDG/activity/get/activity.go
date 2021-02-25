package get

import (
	"github.com/fantadeltaalpha/flogo/extension/TIBDG/connector/datagrid"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityLogger = log.ChildLogger(log.RootLogger(), "datagrid.activity.get")

var activityMd = activity.ToMetadata(&Settings{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	activityLogger.Debugf("Connection Manager: %v", s.ConnectionRef)
	cm, err := coerce.ToConnection(s.ConnectionRef)
	if err != nil {
		return nil, err
	}

	AssertType(cm.GetConnection())
	

	activityLogger.Debugf("Connection Manager: %T %v", cm.GetConnection(),cm.GetConnection())
	c, ok := (cm.GetConnection()).(*datagrid.ClientManager)
	if !ok {
		activityLogger.Debugf("Type Assert Error: %T, Get Conn : %v %v",cm.GetConnection(),c.GridName,c.RealmURL)
		activityLogger.Error(datagrid.ErrorInvalidConnection.Error())
		return nil, datagrid.ErrorInvalidConnection
	}
	act := &Activity{settings: s, gridMame: c.GridName, realmURL: c.RealmURL}
	return act, nil
}

func AssertType(i interface{}) *datagrid.ClientManager {

    // Recover
    defer func() {
        if err := recover(); err != nil {
            activityLogger.Error(err)
        }
    }()

    return i.(*datagrid.ClientManager)
}

type Activity struct {
	settings *Settings
	gridMame   string
	realmURL	string
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func (a *Activity) Eval(ctx activity.Context) (bool, error) {

	activityLogger.Debugf("%v %v", a.settings, a.gridMame, a.realmURL)

	/*input := &Input{}
	err := ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	if input.Message.Destination == "" {
		activityLogger.Error(conn.ErrorInvalidEmptyDestination.Error())
		return false, conn.ErrorInvalidEmptyDestination
	}

	err = a.client.Send(input.Message.Destination, a.settings.DestinationType, input.Message.Content, input.Message.DeliveryDelay, a.settings.DeliveryMode, input.Message.Expiration)
	if err != nil {
		activityLogger.Errorf("Failed to send message due to error - %s", err.Error())
		return false, err
	}
	activityLogger.Debugf("Message successfully sent on %s destination [%s] by activity [%s] in flow [%s]", a.settings.DestinationType,input.Message.Destination, ctx.Name(), ctx.ActivityHost().Name())*/

	return true, nil
}
