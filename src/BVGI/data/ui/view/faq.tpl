{{define "content"}} {{$faqs := index $ "faqs"}}
<div class="jumbotron mb-0 pb-5">
    <div class="container">
        <div class="col-xs-12 col-center-block mb-5">
            <h1 align="center" class="text-primary">FAQs</h1>
        </div>
        <div class="row">
            <button class="btn btn-success btn-lg pull-left mt-2 mb-2" onclick="location.href='newfaq';">New question</button>
            <button class="btn btn-primary btn-lg pull-left mt-2 mb-2 ml-3" id="btnExport">Export Excel</button>
        </div>
        <div class="row">
            <table class="table col-sm-12 table-bordered table-hover .table-striped px-0 mx-0" style="word-wrap: break-word;">
                <thead class="thead-inverse col-sm-12 px-0 mx-0">
                    <tr class="row col-sm-12 px-0 mx-0">
                        <th class="col-sm-1 text-center">ID</th>
                        <th class="col-sm-3 text-center">Intent</th>
                        <th class="col-sm-5 text-center">Answer</th>
                        <th class="col-sm-1 text-center">Samples</th>
                        <th class="col-sm-2 text-center">Action</th>
                    </tr>
                </thead>
                <tbody ic-confirm="Are you sure?" ic-target="closest tr" ic-replace-target="true">
                    {{ range $idx, $faq := $faqs}}
                    <tr class="row col-sm-12 table-active px-0 mx-0">
                        <td class="col-sm-1">{{ $idx }}</td>
                        <td class="col-sm-3">{{ $faq.IntentName }}</td>
                        <td class="col-sm-5">{{ $faq.Answer }}</td>
                        <td class="col-sm-1">{{ len $faq.Questions}}</td>
                        <td class="col-sm-2">
                            <button class="btn btn-primary btn-md" style="margin-right: 0.5em;" onclick="location.href='/editview?id={{$faq.ID}}';">Edit</button>
                            <button class="btn btn-danger btn-md delete" id="delete" ic-delete-from="/faq/{{$faq.ID}}">Delete</button>
                        </td>
                    </tr>

                    {{end}}

                </tbody>
            </table>
        </div>
    </div>
</div>

<script type="text/javascript">
    $("#btnExport").click(function() {
       var isOK = confirm("Are you sure?");
       if(isOK){
            window.location = "/export";
       }
        
    });
</script>

{{end}}