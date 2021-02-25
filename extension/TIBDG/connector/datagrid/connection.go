package datagrid

import (
	"errors"

	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/connection"
	"github.com/project-flogo/core/support/log"
)

var ConnectorLogger = log.ChildLogger(log.RootLogger(), "datagrid.connection")
var factory = &Factory{}

type Factory struct {
}

var (
	ErrorInvalidConnection       = errors.New("invalid connection object set at runtime")
	ErrorInvalidEmptyDestination = errors.New("invalid empty destination set")
)

type Settings struct {
	Name     string `md:"name,required"`
	GridName     string `md:"gridName,required"`
	RealmURL      string `md:"realmUrl,required"`
}

func (*Factory) Type() string {
	return "Datagrid"
}

func init() {
	err := connection.RegisterManagerFactory(factory)
	if err != nil {
		panic(err)
	}
}

func (*Factory) NewManager(settings map[string]interface{}) (connection.Manager, error) {

	sharedConn := &ClientManager{
	}
	var err error
	s := &Settings{}

	err = metadata.MapToStruct(settings, s, true)

	if err != nil {
		return nil, err
	}

	/*opts := ems.NewClientOptions()
	opts.SetServerUrl(s.URL).SetUsername(s.UserName).SetPassword(s.Password)

	client := ems.NewClient(opts)*/

	sharedConn.RealmURL = s.RealmURL
	sharedConn.GridName = s.GridName

	ConnectorLogger.Debugf("New Manager : %v %v",sharedConn.GridName,sharedConn.RealmURL)

	return sharedConn, nil
}

type ClientManager struct {
	GridName   string
	RealmURL    string
}

func (k *ClientManager) Type() string {
	return "Datagrid"
}

func (k *ClientManager) GetConnection() interface{} {
	ConnectorLogger.Debugf("Get Conn : %v %v",k.GridName,k.RealmURL)
	return k
}

func (k *ClientManager) ReleaseConnection(connection interface{}) {
}

/*func (k *ClientManager) Start() error {
	ConnectorLogger.Debugf("Creating connection with TIBCO EMS for - [%s]", k.name)
	err := k.client.Connect()
	if err != nil {
		ConnectorLogger.Errorf("Failed to create connection for [%s] due to error - {%s}", k.name, err.Error())
	}
	return err
}

func (k *ClientManager) Stop() error {
	ConnectorLogger.Debugf("Stopping connection with TIBCO EMS for - [%s]", k.name)
	if k.client != nil {
		_ = k.client.Disconnect()
	}
	return nil
}*/
