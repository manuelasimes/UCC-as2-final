package config

import (
	"fmt"
)

var (
	SOLRHOST       = "localhost"
	SOLRPORT       = 8983
	SOLRCOLLECTION = "hotelSearch"

	HOTELSHOST = "localhost"
	HOTELSPORT = 8090

	QUEUENAME = "worker_solr"
	EXCHANGE  = "hotels"

	LBHOST = "lbbusqueda"
	LBPORT = 80

	RABBITUSER     = "user"
	RABBITPASSWORD = "password"
	RABBITHOST     = "localhost"
	RABBITPORT     = 5672

	AMPQConnectionURL = fmt.Sprintf("amqp://%s:%s@%s:%d/", RABBITUSER, RABBITPASSWORD, RABBITHOST, RABBITPORT)

	USERAPIHOST = "localhost"
	USERAPIPORT = 8070
)