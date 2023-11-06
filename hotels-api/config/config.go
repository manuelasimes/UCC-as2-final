package config

import (
	"fmt"
)

var (
	SOLRHOST       = "solr"
	SOLRPORT       = 8983
	SOLRCOLLECTION = "hotelSearch"

	HOTELSHOST = "hotels-api"
	HOTELSPORT = 8060

	QUEUENAME = "worker_solr"
	EXCHANGE  = "hotels"

	LBHOST = "lbbusqueda"
	LBPORT = 80

	RABBITUSER     = "user"
	RABBITPASSWORD = "password"
	RABBITHOST     = "rabbit"
	RABBITPORT     = 5672

	AMPQConnectionURL = fmt.Sprintf("amqp://%s:%s@%s:%d/", RABBITUSER, RABBITPASSWORD, RABBITHOST, RABBITPORT)

	USERAPIHOST = "user-res-api"
	USERAPIPORT = 8070
)