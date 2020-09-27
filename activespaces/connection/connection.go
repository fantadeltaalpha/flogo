package connection

import (
	"github.com/pkg/errors"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/connection"
	"github.com/project-flogo/core/support/log"
	"tibco.com/tibdg"
)

func init() {
	_ = connection.RegisterManager("activespacesConnection", &ActivespacesConnection{})
	_ = connection.RegisterManagerFactory(&Factory{})
}

//Settings struct
type Settings struct {
	RealmURL 	string `md:"realmUrl,required"`
	GridName string `md:"gridName,required"`
	ClientLabel string `md:"clientLabel,required"`
	ConnectWaitTime float64 `md:"gridName,required"`
}

//ActivespacesConnection struct
type ActivespacesConnection struct {
	connection *tibdg.Connection
	properties *tibdg.Props
}

//Factory struct
type Factory struct {
}

//Type function
func (*Factory) Type() string {

	return "pulsar"
}

//NewManager function
func (*Factory) NewManager(settings map[string]interface{}) (connection.Manager, error) {
	logger := log.ChildLogger(log.RootLogger(), "kafka-shared-conn")
	s := &Settings{}
	logger.Debugf("Setting: %+q\n", s)
	err := metadata.MapToStruct(settings, s, true)
	if err != nil {
		return nil, err
	}

	props := tibdg.Props{}
	props[tibdg.ConnectionPropertyStringClientLabel] = s.ClientLabel

	if s.ConnectWaitTime > 0.0 {
		props[tibdg.ConnectionPropertyDoubleConnectWaitTime] = s.ConnectWaitTime
	}

	conn, err := tibdg.NewConnection(s.RealmURL, s.GridName, props)
	if err != nil {
		return nil, errors.Wrapf(err,"Failed to connect to datagrid %s at %s",s.GridName,s.RealmURL)
	}

	return &ActivespacesConnection{connection:conn,properties:&props},nil

	/*auth := getAuthentication(s)

	clientOps := pulsar.ClientOptions{
		URL:            s.URL,
		Authentication: auth,
	}
	client, err := pulsar.NewClient(clientOps)

	if err != nil {
		return nil, err
	}*/

	//return &PulsarConnection{client: client}, nil
}

//Type function
func (p *ActivespacesConnection) Type() string {

	return "activespaces"
}

//GetConnection function
func (p *ActivespacesConnection) GetConnection() interface{} {

	return p.connection
}

//Stop function
func (p *ActivespacesConnection) Stop() error {
	p.connection.Close()
	return nil
}

//Start function
func (p *ActivespacesConnection) Start() error {
	return nil
}

//ReleaseConnection function
func (p *ActivespacesConnection) ReleaseConnection(connection interface{}) {

}

/*func getAuthentication(s *Settings) pulsar.Authentication {
	if len(s.AthenzAuthentication) != 0 {
		return pulsar.NewAuthenticationAthenz(s.AthenzAuthentication)
	}

	if s.CertFile != "" && s.KeyFile != "" {
		return pulsar.NewAuthenticationTLS(s.CertFile, s.KeyFile)
	}

	return nil
}*/
