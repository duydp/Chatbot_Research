{{define "content"}} {{$histories := index $ "history"}} {{$intents := index $ "intents"}}

<div class="jumbotron mb-0 pb-5">
    <div class="container">
        <div class="col-sm-12 col-center-block mb-5">
            
            <h1 align="center" class="text-primary">History</h1>
            
        </div>
        <div class="row justify-content-md-center" style="min-height: 400px;">
            <table class="table col-sm-12 table-bordered table-hover .table-striped px-0 mx-0">
                <thead class="thead-inverse col-sm-12 px-0 mx-0">
                    <tr class="row col-sm-12 px-0 mx-0">
                        <th class="col-sm-1 text-center">ID</th>
                        <th class="col-sm-4 text-center">Question</th>
                        <th class="col-sm-4 text-center">Intent</th>
                        <th class="col-sm-2 text-center">Created</th>
                        <th class="col-sm-1 text-center">Action</th>
                    </tr>
                </thead>
                <tbody id="tbody" ic-confirm="Are you sure?" ic-target="closest tr" ic-replace-target="true">
                    {{ range $idx, $h := $histories }}
                    <tr class="row col-sm-12 table-active px-0 mx-0" data-id="row">
                        <td class="col-sm-1 text-center" id="qid" data-id="{{$h.ID}}">{{ $idx }}</td>
                        <td class="col-sm-4" id="question" data-id="{{$h.Question}}">{{ $h.Question }}&nbsp</td>
                        <td class="row col-sm-4 pl-3 mx-0" data-id="tt1">
                            <div class="form-inline form-group my-0 col-sm-12" data-id="tt2">
                                <select class="selection dropdown col-sm-9" id="search-select" style="height:50px;" data-id="tt3">
                                    <option value="_">unknown</option>
                                    {{range $i, $v := $intents}} {{if eq $v $h.Answer}}
                                    <option value="{{$v}}" selected="selected">{{$v}}</option>
                                    {{else}}
                                    <option value="{{$v}}">{{$v}}</option>
                                    {{end}} {{end}}
                                </select>

                                <button type="button" class="col-sm-2 btn btn-success correct ml-3" style="height: 50px;">
                                    <i data-id="loading" class="fa fa-floppy-o" style="font-size:24px; color:white;"></i>
                                </button>
                            </div>
                        </td>
                        <td class="col-sm-2 mx-0 px-0 text-center">{{ $h.AskTime }}</td>
                        <td class="col-sm-1 mx-0 px-0">
                            <button class="col-sm-10 mx-2 px-0 btn btn-danger btn-sm delete" id="delete" ic-delete-from="/history/{{$h.ID}}" style="height: 50px;">Delete</button>
                        </td>
                    </tr>

                    {{end}}
                </tbody>
            </table>
        </div>

        <div class="row justify-content-md-center py-5">
          <nav aria-label="Page navigation">
              <ul class="pagination" id="pagination"></ul>
          </nav>
        </div>
    </div>
</div>

<script type="text/javascript">
    $('#pagination').twbsPagination({
        totalPages: {{.pages}},
        visiblePages: 10,
        startPage:{{.page}},
        initiateStartPageClick: false,
        onPageClick: function (event, page) {
            window.location="/history/" + page;
        }
    });


    $(document).ready(function() {
        $('body').on('click', 'button.correct', function() {
            var btn = $(this);

            var qID = $(this).data("id");
            var dataString = 'id=' + qID;
            var element = $(this).parent().parent().parent();

            var question = $(element).find("#question").data("id");
            var intent = $(this).prev().find('option:selected').text();
            var qid = $(element).find("#qid").data("id");

            if (intent == "unknown") {
                alert("cannot using unknow as intent");
            } else {
                $(this).removeClass("btn-secondary");
                $(this).addClass("btn-success");
                $(this).empty();
                $(this).append('<i data-id="loading" class="fa fa-spinner fa-spin" style="font-size:24px; color:white;"></i>');

                var dataString = 'question=' + question + "&intent=" + intent + "&qid=" + qid;
                var status = "";
                $.ajax({
                    type: "POST",
                    url: "/correcthistory",
                    data: dataString,
                    cache: false,
                    success: function(result) {
                        $(btn).empty();
                        $(btn).append('<i data-id="loading" class="fa fa-check" style="font-size:24px; color:white;"></i>');
                    }
                });
            }
            return false;
        });
    });

    
</script>

{{end}}