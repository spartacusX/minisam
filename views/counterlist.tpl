<div class="row-fluid">
<table class="table table-hover table-bordered table-striped">
	<thead>
		<tr class="caption">
			<th>
				<input type="checkbox"></input>
			</th>
			<th>
				Id
			</th>
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
		<tr>
			<td>
				<img src="static/ico/{{.Status}}.gif"/>
			</td>
			<td>
				{{.Id}}
			</td>
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
</table> 



<div class="pagination offset5">
		<li><a href="#">Prev</a></li>		
		{{.Test}}
		<li><a href="">Next</a></li>
</div>

<div class="row-fluid">
	<div class="panel panel-default col-lg-5">
		<div class="panel-heading">Information</div>
		
			<div class="panel-body">
				<form class="bs-example">
					<div class="row-fluid span6">
					    <div class="row-fluid">
					      <label class="control-label span6" for="disabledInput">Name:</label>
					      <input class="form-control span6" id="disabledInput" type="text" placeholder="0" disabled="">
					    </div>
					    <div class="row-fluid">
					      <label class="control-label span6" for="disabledInput">Code:</label>
					      <input class="form-control span6" id="disabledInput" type="text" placeholder="0" disabled="">
					    </div>
					    <div class="row-fluid">
					      <label class="control-label span6" for="disabledInput">Nature:</label>
					      <input class="form-control span6" id="disabledInput" type="text" placeholder="0" disabled="">
					    </div>
					    <div class="row-fluid">
					      <label class="control-label span6" for="disabledInput">Group by:</label>
					      <input class="form-control span6" id="disabledInput" type="text" placeholder="0" disabled="">
					    </div>
					    <div class="row-fluid">
					      <label class="control-label span6" for="disabledInput">Supervisor:</label>
					      <input class="form-control span6" id="disabledInput" type="text" placeholder="0" disabled="">
					    </div>
					    <div class="row-fluid">
					      <label class="control-label span6" for="disabledInput">Scope of application:</label>
					      <input class="form-control span6" id="disabledInput" type="text" placeholder="0" disabled="">
					    </div>
					    <div class="row-fluid">
					      <label class="control-label span6" for="disabledInput">License contract model:</label>
					      <input class="form-control span6" id="disabledInput" type="text" placeholder="0" disabled="">
					    </div>
					</div>
					<div class="row-fluid span6">
						<div class="checkbox">
	                        <label>
	                          <input type="checkbox"> Use as template
	                        </label>
                      	</div>
                      	<div class="checkbox">
	                        <label>
	                          <input type="checkbox"> Is part of corporate software management
	                        </label>
                      	</div>
                      	<div class="checkbox">
	                        <label>
	                          <input type="checkbox"> Do not include in the compliance reports
	                        </label>
                      	</div>
                      	<div class="checkbox">
	                        <label>
	                          <input type="checkbox"> Edit the counter using the wizards
	                        </label>
                      	</div>
                      	<div class="checkbox">
	                        <label>
	                          <input type="checkbox"> Rights count
	                        </label>
                      	</div>
                      	<div class="checkbox">
	                        <label>
	                          <input type="checkbox"> Entitlements count
	                        </label>
                      	</div>
                      	<div class="checkbox">
	                        <label>
	                          <input type="checkbox"> Installations/Utilizations count
	                        </label>
                      	</div>
                      	<div class="checkbox">
	                        <label>
	                          <input type="checkbox"> Software upgrade counter
	                        </label>
                      	</div>
					</div>
				</form>
			</div>
	</div>
	<div class="panel panel-default col-lg-7">
            <div class="bs-example">
              <ul class="nav nav-tabs" style="margin-bottom: 15px;">
                <li class="active"><a href="#general" data-toggle="tab">General</a></li>
                <li><a href="#rights" data-toggle="tab">Rights</a></li>
                <li><a href="#installs" data-toggle="tab">Installations/Utilizations</a></li>                               
                <li class="dropdown">
                  <a class="dropdown-toggle" data-toggle="dropdown" href="#">
                    Dropdown <span class="caret"></span>
                  </a>
                  <ul class="dropdown-menu">
                    <li><a href="#dropdown1" data-toggle="tab">Action</a></li>
                    <li class="divider"></li>
                    <li><a href="#dropdown2" data-toggle="tab">Another action</a></li>
                  </ul>
                </li>
              </ul>
              <div id="myTabContent" class="tab-content">
                <div class="tab-pane fade active in" id="general">
                	<div class="panel panel-default">
                	  	<div class="panel-heading">
                	  		Result of calculations
                	  	</div>
		              	<div class="panel-body">
			              	<form class="bs-example">
				                <div class="row-fluid">
				                  <label class="control-label span3" for="disabledInput">Rights count:</label>
				                  <input class="form-control span3" id="disabledInput" type="text" placeholder="0" disabled="">
				                </div>
				                <div class="row-fluid">
				                  <label class="control-label span3" for="disabledInput">Installations/Utilizations count:</label>
				                  <input class="form-control span3" id="disabledInput" type="text" placeholder="0" disabled="">
				                </div>
				                <div class="row-fluid">
				                  <label class="control-label span3" for="disabledInput">Unused Installations:</label>
				                  <input class="form-control span3" id="disabledInput" type="text" placeholder="0" disabled="">
				                </div>
				                <div class="row-fluid">
				                  <label class="control-label span3" for="disabledInput">Compliance:</label>
				                  <input class="form-control span3" id="disabledInput" type="text" placeholder="0" disabled="">
				                </div>
				            </form>
			            </div>
			            <div class="panel-heading">
			            	Comment
			            </div>
		              	<div class="panel-body">
		              		<input class="form-control" id="disabledInput" type="text" placeholder="comment" disabled="">
		              	</div>
		            </div>
		        </div>
		        <div class="tab-pane fade" id="rights">
                  	<div class="panel panel-default">
		                	 <div class="panel-heading">
		                	  	Criteria
		                	 </div>
		              	<div class="panel-body">
		              		<form class="bs-example">
				                <div class="row-fluid">
				                  	<label class="control-label span3" for="disabledInput">Rights counter context:</label>
				                  	<input class="form-control span3" id="disabledInput" type="text" placeholder="0" disabled="">
				                </div>
				                <div class="row-fluid">
				                  	<label class="control-label span3" for="disabledInput">Scope of the rights to be counted:</label>
				                  	<input class="form-control span3" id="disabledInput" type="text" placeholder="0" disabled="">
				                </div>
				                <div class="row-fluid">
				                  	<label class="control-label span3" for="disabledInput">Rights -> Group By link:</label>
				                  	<input class="form-control span3" id="disabledInput" type="text" placeholder="0" disabled="">
				                </div>
				                <div class="row-fluid">
				                  	<label class="control-label span3" for="disabledInput">License type:</label>
				                  	<input class="form-control span3" id="disabledInput" type="text" placeholder="0" disabled="">
				                </div>
			                </form>
			            </div>
			        
			            <div class="panel-heading">
			            	Calculation
			            </div>
		              	<div class="panel-body">
		              		<div class="row-fluid">
				                <label class="control-label span3" for="disabledInput">Rights-calculation mode:</label>
				                <input class="form-control span3" id="disabledInput" type="text" placeholder="0" disabled="">
				            </div>
				            <div class="row-fluid offset3">
				                <input class="form-control span3" id="disabledInput" type="text" placeholder="0" disabled="">
				            </div>
		              	</div>
		             </div>
                </div>
                <div class="tab-pane fade" id="installs">
                  	<div class="panel panel-default">
                	  <div class="panel-heading">
                	  	Criteria
                	  </div>
		              <div class="panel-body">
		              	<form class="bs-example">
			                <div class="row-fluid">
			                  <label class="control-label span5" for="disabledInput">Installation/Utilization counter context:</label>
			                  <input class="form-control span5" id="disabledInput" type="text" placeholder="0" disabled="">
			                </div>
			                <div class="row-fluid">
			                  <label class="control-label span5" for="disabledInput">Scope of the Installations/Utilizations to be counted:</label>
			                  <input class="form-control span5" id="disabledInput" type="text" placeholder="0" disabled="">
			                </div>
			                <div class="row-fluid">
			                  <label class="control-label span5" for="disabledInput">Installations/Utilizations-> Group By link:</label>
			                  <input class="form-control span5" id="disabledInput" type="text" placeholder="0" disabled="">
			                </div>
			            </form>
			           </div>
			            <div class="panel-heading">
			            	Calculation
			            </div>
		              	<div class="panel-body">
		              		<div class="row-fluid">
				                <label class="control-label span5" for="disabledInput">Installations/Utilizations-calculation mode:</label>
				                <input class="form-control span5" id="disabledInput" type="text" placeholder="0" disabled="">
				            </div>
				            <div class="row-fluid offset3">
				                <input class="form-control span5" id="disabledInput" type="text" placeholder="0" disabled="">
				            </div>
		              	</div>
		            </div>
                </div>
                <div class="tab-pane fade" id="dropdown2">
                  <p>Trust fund seitan letterpress, keytar raw denim keffiyeh etsy art party before they sold out master cleanse gluten-free squid scenester freegan cosby sweater. Fanny pack portland seitan DIY, art party locavore wolf cliche high life echo park Austin. Cred vinyl keffiyeh DIY salvia PBR, banh mi before they sold out farm-to-table VHS viral locavore cosby sweater.</p>
                </div>
                  
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