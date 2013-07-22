<div class="container-fluid">
	<div class="row-fluid">
		<div class="span12">
			<table class="table table-hover table-striped">
				<thead>
					<tr>
						<th>
							CounterID
						</th>
						<th>
							Model
						</th>
						<th>
							License
						</th>
						<th>
							Installation
						</th>
					</tr>
				</thead>
				<tbody>
					{{range .Counters}}
					<tr>
						<td>
							{{.Id}}
						</td>
						<td>
							{{.Model}}
						</td>
						<td>
							{{.LicenseNum}}
						</td>
						<td>
							{{.InstallationNum}}
						</td>
					</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</div>
</div>