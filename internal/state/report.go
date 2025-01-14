package internal

// displays a progress bar for the total time of the test
// displays the active number of VUs at any point of time
// finally displays the value of all the metrics in a table

import (
	"context"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kernelpanic77/hellfire/internal/metrics"
)

var (
	colTitles = [9]string{"Metric", "Type", "Sum", "Min", "Max", "Avg", "Med", "P(90)", "P(95)"}
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
	registry *metrics.MetricsRegistry
	tableWriter  table.Writer
	context context.Context
}


func NewReport(ctx context.Context, registry *metrics.MetricsRegistry) *Report {
	return &Report{
		registry: registry,  
		tableWriter: table.NewWriter(),
		context: ctx,
	}
}

func (report *Report) createColumnConfigs() []table.ColumnConfig {
	// iterate over all metrics in 
	cols := []table.ColumnConfig{}
	for i := 0; i < len(colTitles); i++ {
		cols = append(cols, table.ColumnConfig{Name: colTitles[i]})
	}
	return cols
}

func (report *Report) GenerateReport() string {
	colConfigs := report.createColumnConfigs()
	report.tableWriter.SetColumnConfigs(colConfigs)
	report.tableWriter.SetAutoIndex(true) 
	for name, mlist := range report.registry.Metrics {
		r := table.Row{name}
		for _, m := range mlist {
			if m.Type == metrics.Counter || m.Type == metrics.Rate {
				val := ConvertFloatSliceToStringSlice([]float64{m.Sink.FetchSampleValue()})
				r = append(r, val...)
			} else if m.Type == metrics.Guage {
				val := ConvertFloatSliceToStringSlice((m.Sink).(*metrics.GuageSink).FetchGuage())
				r = append(r, val...)
			} else if m.Type == metrics.Trend {
				val := ConvertFloatSliceToStringSlice((m.Sink).(*metrics.TrendSink).FetchTrends())
				r = append(r, val...)
			}
		}
		report.tableWriter.AppendRow(r)
	}
	report.tableWriter.SetCaption("Hellfire Report!")
	return report.tableWriter.Render()
}