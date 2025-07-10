package internal

// displays a progress bar for the total time of the test
// displays the active number of VUs at any point of time
// finally displays the value of all the metrics in a table

import (
	"context"
	// "fmt"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kernelpanic77/hellfire/internal/metrics"
)

var (
	colTitles = [10]string{"Metric", "Type", "Sum", "Rates if any", "Min", "Max", "Avg", "Med", "P(90)", "P(95)"}
	// rowHeader         = table.Row{colTitleIndex, colTitleFirstName, colTitleLastName, colTitleSalary}
	// row1              = table.Row{1, "Arya", "Stark", "3000µs"}
	// row2              = table.Row{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"}
	// row3              = table.Row{300, "Tyrion", "Lannister", 5000}
	// rowFooter         = table.Row{"", "", "Total", 10000}
)

func ConvertFloatSliceToStringSlice(floatSlice []float64) []interface{} {
	// Create an empty slice of strings with the same length as the float slice
	strSlice := make([]interface{}, len(floatSlice))

	// Iterate through the float slice and convert each element to a string
	for i, v := range floatSlice {
		strSlice[i] = (strconv.FormatFloat(v, 'f', -1, 64))
	}

	return strSlice
}

// Depending upon Strategy show progress bar
// Progress bar for total test duration
// Metric avg min med max p(90) p(95)
type Report struct {
	Registry    *metrics.MetricsRegistry
	tableWriter table.Writer
	context     context.Context
}

func NewReport(ctx context.Context, registry *metrics.MetricsRegistry) *Report {
	return &Report{
		Registry:    registry,
		tableWriter: table.NewWriter(),
		context:     ctx,
	}
}

func (report *Report) createRowHeaders() table.Row {
	// iterate over all metrics in
	rows := table.Row{}
	for i := 0; i < len(colTitles); i++ {
		rows = append(rows, colTitles[i])
	}
	return rows
}

func (report *Report) GenerateReport() string {
	colConfigs := report.createRowHeaders()
	report.tableWriter.AppendHeader(colConfigs)
	// report.tableWriter.SetColumnConfigs(colConfigs)
	report.tableWriter.SetAutoIndex(true)
	for name, mlist := range report.Registry.Metrics {
		r := table.Row{name}
		r = append(r, "")

		// for _, m := range mlist {
		// 	fmt.Println(m)
		// }

		// append 7 empty columns
		for i := 0; i < 9; i++ {
			r = append(r, "_")
		}
		for gidx, m := range mlist {
			var metric metrics.MetricType = m.Type
			r[1] = r[1].(string) + metrics.MetricTypeString(metric)
			if gidx != len(mlist)-1 {
				r[1] = r[1].(string) + ","
			}
			if m.Type == metrics.Counter {
				val := ConvertFloatSliceToStringSlice([]float64{m.Sink.FetchSampleValue()})
				r[2] = val
			} else if m.Type == metrics.Rate {
				val := ConvertFloatSliceToStringSlice([]float64{m.Sink.FetchSampleValue()})
				r[3] = val
			} else if m.Type == metrics.Guage {
				val := ConvertFloatSliceToStringSlice((m.Sink).(*metrics.GuageSink).FetchGuage())
				start_idx := 4
				for idx, mval := range val {
					r[start_idx+idx] = mval
				}
			} else if m.Type == metrics.Trend {
				val := ConvertFloatSliceToStringSlice((m.Sink).(*metrics.TrendSink).FetchTrends())
				// r = append(r, val...)
				start_idx := 6
				for idx, mval := range val {
					r[start_idx+idx] = mval
				}
			}
		}

		report.tableWriter.AppendRow(r)
	}
	report.tableWriter.SetCaption("Hellfire Report!")
	return report.tableWriter.Render()
}
