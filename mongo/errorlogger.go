package mongo

import (
	"log"
)

func logFindError(c Collection, query, projection, data interface{}) {
	if c.flag {
		if projection != nil {
			log.Printf("\nCollection : %s\nQuery : %v\nProjection : %v\nData : %T", c.Name(), query, projection, data)
		} else {
			log.Printf("\nCollection : %s\nQuery : %v\nData : %T", c.Name(), query, data)
		}
	}
}

func logError(c Collection, query interface{}, data []interface{}) {
	if c.flag {
		if (query != nil) && (data != nil) {
			log.Printf("\nCollection : %s\nQuery : %v\nData : %T", c.Name(), query, data)
		}
		if data == nil {
			log.Printf("\nCollection : %s\nQuery : %v", c.Name(), query)
		} else {
			log.Printf("\nCollection : %s\nData : %s", c.Name(), data)

		}
	}
}

func logResult(c Collection, query, data, res interface{}) {
	if c.flag {
		if (data != nil) && (query != nil) {
			log.Printf("\nCollection : %s\nQuery : %v\nData : %T\nResult : %v", c.Name(), query, data, res)
		}
		if query != nil {
			log.Printf("\nCollection : %s\nQuery : %v\nResult : %v", c.Name(), query, res)
		} else {
			log.Printf("\nCollection : %s\nData : %s\nInsertResult : %v", c.Name(), data, res)
		}
	}
}
