package sql

import (
	"errors"
)

func (this *engine) executeSelect() ([]map[string]interface{}, error) { //SQL
	rows, err := this.db.Query(this.query.String(), this.args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	if len(columns) <= 0 {
		return nil, errors.New("column is empty")
	}
	// Make a slice for the values
	values := make([]interface{}, len(columns))
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := 0; i < len(values); i++ {
		scanArgs[i] = &values[i]

	}
	var results []map[string]interface{}
	for rows.Next() {

		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		t := make(map[string]interface{})
		//log.Println(len(selectKeys))
		for i, v := range values {
			t[columns[i]] = v

		}

		results = append(results, t)
	}
	//rows.Close()
	return results, nil

}
