package main

import (
	"github.com/influxdb/influxdb-go"
)

func main() {
	completedChan := make(chan bool)
	for i := 0; i < 4; i++ {
		go func(instanceIndex int) {

			config := &influxdb.ClientConfig{
				Host:     "localhost:8086",
				Username: "root",
				Password: "root",
				Database: "testy",
				IsSecure: false,
			}

			client, err := influxdb.NewClient(config)
			if err != nil {
				println(err.Error())
			}
			for j := 100000 * i; j < (100000*i + 100000); j++ {
				testSeries1 := influxdb.Series{
					Name:    "shortTermThing",
					Columns: []string{"attribute1", "attribute2"},
					Points:  [][]interface{}{[]interface{}{j, j + 1}, []interface{}{j + 2, j + 3}},
				}

				testSeries := []*influxdb.Series{&testSeries1}
				err = client.WriteSeries(testSeries)
				if err != nil {
					println(instanceIndex)
					println(err.Error())
				}
			}
			completedChan <- true
		}(i)
	}

	count := 0

Dance:
	for _ = range completedChan {
		count++
		if count == 4 {
			break Dance
		}
	}
}
