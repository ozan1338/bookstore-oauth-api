package cassandra

import (
	"os"
	"time"

	"github.com/gocql/gocql"
)

const (
	cassandra_host="cassandra_host"
)

var (
	session *gocql.Session
	host = os.Getenv(cassandra_host)
)

func init() {
	//Connect to Cassandra cluster
	cluster := gocql.NewCluster(host)
	cluster.ProtoVersion = 4
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout  = time.Second * 10
	cluster.DisableInitialHostLookup = true

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}


}

func GetSession() *gocql.Session {
	return session
}