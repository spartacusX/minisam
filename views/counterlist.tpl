<!--List screen-->
<div class="row-fluid">
<table class="table table-hover table-bordered table-striped">
	<thead>
		<tr class="caption">
			{{range .ColumnList}}
			<th class="{{.ColumnClassAtt}}">
				{{.ColumnName}}
			</th>
			{{end}}
		</tr>
	</thead>
	<tbody>
		{{range .CountResult}}
		<tr>
			<td>
				<img src="static/ico/{{.Status}}.png"/>
			</td>
			<td id="counterId" class="hide">
				{{.LCounterId}}
			</td>
			<td id="qqq">
				<a href="javascript:" onclick="javascript:loadXMLDoc(event);">{{.Name}}</a>
			</td>
			<td>
				{{.LLicUseRights}}
			</td>
			<td>
				{{.LEntCount}}
			</td>
			<td>
				{{.LSoftInstallCount}}
			</td>
			<td>
				{{.LUnusedInstall}}
			</td>
		</tr>
		{{end}}
	</tbody>
</table> 

<div class="pagination offset5">
		<li><a href="/index?PageNum=-1">Prev</a></li>		
		{{.Pagination}}
		<li><a href="/index?PageNum=-2">Next</a></li>
</div>

<!--Detail screen-->
<div class="row-fluid">
	<div class="panel panel-default col-lg-5">
		<div class="panel-heading">Information</div>
		
			<div class="panel-body">
				<form class="bs-example">
					<div class="row-fluid span6">
						{{range .DetailInformationInput}}
						<div class="row-fluid">
					      <label class="control-label span6" for="{{.Id}}">{{.Desc}}</label>
					      <input class="form-control span6" id="{{.Id}}" type="text" placeholder="">
					    </div>
						{{end}}
					</div>
					<div class="row-fluid span6">
						{{range .DetailInformationCheckBox}}
						<div class="checkbox">
	                        <label>
	                          <input type="checkbox" id="{{.Id}}"> {{.Desc}}
	                        </label>
                      	</div>
						{{end}}
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
			              		{{range .DetailGeneralTab}}
			              		<div class="row-fluid">
				                  <label class="control-label span3" for="{{.Id}}">{{.Desc}}:</label>
				                  <input class="form-control span3" id="{{.Id}}" type="text" placeholder="0">
				                </div>
			              		{{end}}
				            </form>
			            </div>
			            <div class="panel-heading">
			            	Comment
			            </div>
		              	<div class="panel-body">
		              		<input class="form-control" id="disabledInput" type="text" placeholder="comment">
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
				                  	<input class="form-control span3" id="disabledInput" type="text" placeholder="0">
				                </div>
				                <div class="row-fluid">
				                  	<label class="control-label span3" for="disabledInput">Scope of the rights to be counted:</label>
				                  	<input class="form-control span3" id="disabledInput" type="text" placeholder="0">
				                </div>
				                <div class="row-fluid">
				                  	<label class="control-label span3" for="disabledInput">Rights -> Group By link:</label>
				                  	<input class="form-control span3" id="disabledInput" type="text" placeholder="0">
				                </div>
				                <div class="row-fluid">
				                  	<label class="control-label span3" for="disabledInput">License type:</label>
				                  	<input class="form-control span3" id="disabledInput" type="text" placeholder="0">
				                </div>
			                </form>
			            </div>
			        
			            <div class="panel-heading">
			            	Calculation
			            </div>
		              	<div class="panel-body">
		              		<div class="row-fluid">
				                <label class="control-label span3" for="disabledInput">Rights-calculation mode:</label>
				                <input class="form-control span3" id="disabledInput" type="text" placeholder="0">
				            </div>
				            <div class="row-fluid">
				            	<label class="control-label span3" for="disabledInput"></label>
				                <input class="form-control span3" id="disabledInput" type="text" placeholder="0">
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
			                  <input class="form-control span5" id="disabledInput" type="text" placeholder="0">
			                </div>
			                <div class="row-fluid">
			                  <label class="control-label span5" for="disabledInput">Scope of the Installations/Utilizations to be counted:</label>
			                  <input class="form-control span5" id="disabledInput" type="text" placeholder="0">
			                </div>
			                <div class="row-fluid">
			                  <label class="control-label span5" for="disabledInput">Installations/Utilizations-> Group By link:</label>
			                  <input class="form-control span5" id="disabledInput" type="text" placeholder="0">
			                </div>
			            </form>
			           </div>
			            <div class="panel-heading">
			            	Calculation
			            </div>
		              	<div class="panel-body">
		              		<div class="row-fluid">
				                <label class="control-label span5" for="disabledInput">Installations/Utilizations-calculation mode:</label>
				                <input class="form-control span5" id="disabledInput" type="text" placeholder="0">
				            </div>
				            <div class="row-fluid">
				            	<label class="control-label span5" for="disabledInput"></label>
				                <input class="form-control span5" id="disabledInput" type="text" placeholder="0">
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

<div id="xxContainer1" class="span2" style="min-width: 310px; height: 300px; margin: 0 auto"></div>
<script type="text/javascript">
	{{.Dashboard}}
</script>

<script type="text/javascript">
	function loadXMLDoc(e){
		//alert(e);
		//alert(e.target.parentNode.parentNode.cells[1].innerHTML);
		var counterId = e.target.parentNode.parentNode.cells[1].innerHTML;
		var rights = e.target.parentNode.parentNode.cells[3].innerHTML;
		//var entitlements = e.target.parentNode.parentNode.cells[4].innerHTML;
		var installs = e.target.parentNode.parentNode.cells[5].innerHTML;
		var unuseds = e.target.parentNode.parentNode.cells[6].innerHTML;

  		document.getElementById('DetailGeneralTab_dLicUseRights').placeholder = rights;
  		//document.getElementById('DetailGeneralTab_dEntCount').placeholder = entitlements;
  		document.getElementById('DetailGeneralTab_dSoftInstallCount').placeholder = installs;
  		document.getElementById('DetailGeneralTab_dUnusedInstall').placeholder = unuseds;

		var url = "/index?CounterId=" + counterId;
		//alert(url);
		var xmlhttp;
		var x;
		if (window.XMLHttpRequest) {// code for IE7+, Firefox, Chrome, Opera, Safari
		  	xmlhttp = new XMLHttpRequest();
		} else {// code for IE6, IE5
		  	xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
		}
		xmlhttp.onreadystatechange = function() 
		{
		  if (xmlhttp.readyState==4 && xmlhttp.status==200) 
		  {
		  	{{range .DetailInformationInput}}
		    x = xmlhttp.responseXML.documentElement.getElementsByTagName("{{.SqlName}}");
		    try{
		    	document.getElementById('{{.Id}}').placeholder = x.item(0).textContent;		    	
		    }catch(er){
		    	document.getElementById('{{.Id}}').placeholder = "0";
		    }
		    {{end}}
		    
		    {{range .DetailInformationCheckBox}}
		    x = xmlhttp.responseXML.documentElement.getElementsByTagName("{{.SqlName}}");
		    var bType = x.item(0).textContent;
	    	if (bType == "true"){
	    		document.getElementById('{{.Id}}').setAttribute("checked", "checked")
	    	}
	    	{{end}}
		  }
		}
		xmlhttp.open("GET", url, true);
		xmlhttp.send();
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