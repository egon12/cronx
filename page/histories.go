package page

import (
	"html/template"
	"sync"
)

const historyTemplate = `
<!--
	This page is only being used for development to restructure the code,
	the real html page is on histories.go.
-->
<!DOCTYPE html>
<html lang="en">
<head>
	<!-- Standard Meta -->
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
	<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
	<!-- Site Properties -->
	<title>Cronx</title>
	<link
	   rel="stylesheet"
	   type="text/css"
	   href="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.2/dist/semantic.min.css">
	<script
	   src="https://code.jquery.com/jquery-3.1.1.min.js"
	   integrity="sha256-hVVnYaiADRTO2PzUGmuLJr8BLUSjGIZsDYGmIJLv2b8="
	   crossorigin="anonymous"></script>
	<script
	   src="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.2/dist/semantic.min.js"
	   crossorigin="anonymous"></script>
	<script
	   src="https://cdnjs.cloudflare.com/ajax/libs/html2canvas/0.5.0-beta4/html2canvas.min.js"
	   integrity="sha512-OqcrADJLG261FZjar4Z6c4CfLqd861A3yPNMb+vRQ2JwzFT49WT4lozrh3bcKxHxtDTgNiqgYbEUStzvZQRfgQ=="
	   crossorigin="anonymous"></script>
	<script src="https://cdn.jsdelivr.net/npm/canvas2image@1.0.5/canvas2image.min.js"></script>
	<script type='text/javascript'>
		function screenshot() {
			html2canvas(document.querySelector('#data_table')).then(function(canvas) {
				Canvas2Image.saveAsPNG(canvas, canvas.width, canvas.height);
			});
		}
	</script>
	<style>
        body > .ui.container {
            margin-top: 3em;
            padding-bottom: 3em;
        }
	</style>
	<title>Cronx</title>
</head>
<body>
<div class="ui container">
	<div class="ui left fixed vertical stackable inverted main menu">
		<div class="header item">
			<i class="stopwatch icon"></i>
			Cronx
		</div>
		<a class="item" href="/jobs">
			<i class="tasks icon"></i>
			Jobs
		</a>
		<a class="item active" href="javascript:window.location.reload()">
			<i class="history icon"></i>
			Histories
		</a>
		<div class="item" onclick="screenshot()">
			<button class="fluid ui labeled inverted green icon button">
				<i class="camera icon"></i>
				<div class="left aligned">Screenshot</div>
			</button>
		</div>
	</div>
	<div id="data_table">
		<table class="ui sortable selectable center aligned celled table">
			<thead>
            {{if or .Pagination.PreviousURI .Pagination.NextURI}}
				<tr>
					<th colspan="6">
						<div class="ui right floated pagination menu">
                            {{if .Pagination.PreviousURI}}
								<a class="icon item" href="{{.Pagination.PreviousURI}}">
									<i class="left chevron icon"></i>
								</a>
                            {{end}}
                            {{if .Pagination.NextURI}}
								<a class="icon item" href="{{.Pagination.NextURI}}">
									<i class="right chevron icon"></i>
								</a>
                            {{end}}
						</div>
					</th>
				</tr>
            {{end}}
			<tr>
				<th id="id"
                        {{if eq (index .Sort.Columns "id") "ASC"}} class="sorted ascending"
                        {{else if eq (index .Sort.Columns "id") "DESC"}} class="sorted descending"
                        {{end}}
				>ID
				</th>
				<th id="name"
                        {{if eq (index .Sort.Columns "name") "ASC"}} class="sorted ascending"
                        {{else if eq (index .Sort.Columns "name") "DESC"}} class="sorted descending"
                        {{end}}
				>Name
				</th>
				<th id="status"
                        {{if eq (index .Sort.Columns "status") "ASC"}} class="sorted ascending"
                        {{else if eq (index .Sort.Columns "status") "DESC"}} class="sorted descending"
                        {{end}}
				>Status
				</th>
				<th id="started_at"
                        {{if eq (index .Sort.Columns "started_at") "ASC"}} class="sorted ascending"
                        {{else if eq (index .Sort.Columns "started_at") "DESC"}} class="sorted descending"
                        {{end}}
				>Started at
				</th>
				<th id="finished_at"
                        {{if eq (index .Sort.Columns "finished_at") "ASC"}} class="sorted ascending"
                        {{else if eq (index .Sort.Columns "finished_at") "DESC"}} class="sorted descending"
                        {{end}}
				>Finished at
				</th>
				<th id="latency"
                        {{if eq (index .Sort.Columns "latency") "ASC"}} class="sorted ascending"
                        {{else if eq (index .Sort.Columns "latency") "DESC"}} class="sorted descending"
                        {{end}}
				>Latency
				</th>
			</tr>
			</thead>
			<tbody>
            {{if not .Data}}
				<tr>
					<td colspan="6" class="center aligned"><b><i>No records found.</i></b></td>
				</tr>
            {{end}}
            {{range .Data}}
				<tr
                        {{if eq .Status "SUCCESS"}} class="positive"
                        {{else if eq .Status "ERROR"}} class="error"
                        {{end}}
				>
					<td>{{.ID}}</td>
					<td class="left aligned">
                        {{if gt .Metadata.TotalWave 1 }}
                            {{.Name}} ({{.Metadata.Wave}}/{{.Metadata.TotalWave}})
                        {{else}}
                            {{.Name}}
                        {{end}}

                        {{if eq .Status "ERROR"}}
							<br/>
							<br/>
							err = {{.Error.Err}}<br/>
                            {{if .Error.Fields}}
								fields = {{.Error.Fields}}<br/>
                            {{end}}
                            {{if .Error.Code}}
								code = {{.Error.Code}}<br/>
                            {{end}}
                            {{if .Error.MetricStatus}}
								metric_status = {{.Error.MetricStatus}}<br/>
                            {{end}}
                            {{if .Error.Message}}
								message = {{.Error.Message}}<br/>
                            {{end}}
                            {{if .Error.Line}}
								line = {{.Error.Line}}<br/>
                            {{end}}
                            {{if .Error.OpTraces}}
								op_traces = {{.Error.OpTraces}}<br/>
                            {{end}}
                        {{end}}
					</td>
					<td>
                        {{if eq .Status "SUCCESS"}}
							<div class="ui green label">
								SUCCESS
							</div>
                        {{else if eq .Status "ERROR"}}
							<div class="ui red label">
								FAILED
							</div>
                        {{else}}
							<div class="ui label">
								<i class="arrow up icon"></i>
                                {{.Status}}
							</div>
                        {{end}}
					</td>
					<td>{{.StartedAt.Format "2006-01-02 15:04:05"}}</td>
					<td>{{.FinishedAt.Format "2006-01-02 15:04:05"}}</td>
					<td>{{.LatencyText}}</td>
				</tr>
            {{end}}
			</tbody>
			<tfoot>
            {{if or .Pagination.PreviousURI .Pagination.NextURI}}
				<tr>
					<th colspan="6">
						<div class="ui right floated pagination menu">
                            {{if .Pagination.PreviousURI}}
								<a class="icon item" href="{{.Pagination.PreviousURI}}">
									<i class="left chevron icon"></i>
								</a>
                            {{end}}
                            {{if .Pagination.NextURI}}
								<a class="icon item" href="{{.Pagination.NextURI}}">
									<i class="right chevron icon"></i>
								</a>
                            {{end}}
						</div>
					</th>
				</tr>
            {{end}}
			</tfoot>
		</table>
	</div>
</div>
</body>
</html>
`

var (
	historyPageOnce  sync.Once
	historyPage      *template.Template
	historyPageError error
)

func GetHistoryTemplate() (*template.Template, error) {
	historyPageOnce.Do(func() {
		t := template.New(historiesTemplateName)
		historyPage, historyPageError = t.Parse(historyTemplate)
	})

	return historyPage, historyPageError
}
