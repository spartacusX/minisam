<div class="container-fluid">
	<div class="row-fluid">
		<div class="span7">
			<table class="table table-hover table-bordered table-striped">
				<thead>
					<tr class="caption">
						<th>
							Name
						</th>
						<th>
							Rights
						</th>
						<th>
							Entitlements
						</th>
						<th>
							Installs
						</th>
						<th>
							Unused
						</th>
					</tr>
				</thead>
				<tbody>
					{{range .CountResult}}
					<tr class="{{.Status}}">
						<td id="qqq">
							<a href="javascript:" onclick="javascritp:onShowDashBoard();">{{.Name}}</a>
						</td>
						<td>
							{{.LicUseRights}}
						</td>
						<td>
							{{.EntCount}}
						</td>
						<td>
							{{.SoftInstallCount}}
						</td>
						<td>
							{{.UnusedInstall}}
						</td>
					</tr>
					{{end}}
				</tbody>
			</table> <!-- <button class="btn btn-info" type="button" onclick="location.href='http://localhost:8080/new'">New</button>
			<a href="#myModal" role="button" class="btn" data-toggle="modal">NewCounter</a> -->

			<div class="pagination offset4">
			  <ul>
			    <li><a href="#">Prev</a></li>
			    	
			    	{{.Test}}
			    <li><a href="">Next</a></li>
			  </ul>
			</div>
		</div>

		

	</div>
</div>

<script type="text/javascript">
	function onShowDashBoard(){
		var str=document.getElementById("qqq").value;
		location.href = "http://localhost:8080/dashboard?Name=xxx&LicUseRights=30&EntCount=10&SoftInstallCount=20&UnusedInstall=10";
	}
</script>



<!-- Button to trigger modal -->

<!-- <button type="button" id="xxx" onclick="test()">xxx</button> -->
 
<!-- Modal -->
<div id="myModal" class="modal hide fade" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
  <div class="modal-header">
    <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
    <h3 id="myModalLabel">Adding Counter...</h3>
  </div>
  <div class="modal-body">
    <form method="post">
      <div class="control-group">
        <div class="container"></div>
<div class="control-group">
	<label class="control-label" id="" for="">
		<div class="control-group">
			<label class="control-label">Text</label>
			<div class="controls">
				<input type="number" id="Id" class="" placeholder="Id" name="Id">
			</div>
		</div>
		Choose wisely:
	</label>
	<div class="controls">
		<select name="Model" id="Model">
			<option value="MS Office">MS Office</option>
			<option value="Oracle">Oracle</option>
			<option value="Adobe">Adobe</option>
			<option value="AM">AM</option>
			<option value="SM">SM</option>
			<option value="IBM">IBM</option>
		</select>
		<input type="number" id="LicenseNum" name="LicenseNum" placeholder="LicenseNum">
		<input type="number" id="InstallationNum" name="InstallationNum" placeholder="InstallationNum">
		<div class="control-group">
			<div class="controls"></div>
		</div>
	</div>
</div>
      </div>

<!--       <div class="btn-toolbar">
            <button class="btn btn-primary" type="submit"><i class="icon-save"></i> Save </button>
      </div> -->
    </form>
  </div>
  <div class="modal-footer">
    <button class="btn" data-dismiss="modal" aria-hidden="true">关闭</button>
    <button class="btn btn-primary">Save changes</button>
  </div>
</div>