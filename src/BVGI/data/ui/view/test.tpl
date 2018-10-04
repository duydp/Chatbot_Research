{{define "content"}} {{$intents := index $ "intents"}}
<style>
    .scrollbar {
        height: auto;
        max-height: 180px;
        overflow-x: hidden;
    }
</style>

<div class="jumbotron mb-0 pb-5">
    <div class="container">
        <div class="col-sm-12 col-center-block mb-5">
            <h1 align="center" class="text-primary">Test</h1>
        </div>

        <div class="row justify-content-md-center" style="min-height: 400px;">
            <form class="form col-sm-10">
                <div class="form-inline form-group" onKeyPress="return submit(event)">
                    <label class="control-label col-sm-2">Question</label>
                    <input type="text" id="question" class="form-control col-sm-8" style="height:50px;" placeholder="Enter question">

                    <button type="button" class="ml-3 btn btn-primary btn-md" style="height: 50px;" onclick="check()">Check
                        <i id="loadingIntent" class="fa fa-spinner fa-spin ml-2" style="font-size:24px; color:white; display: none;"></i>
                    </button>
                    
                </div>

                <div class="form-inline form-group my-5">
                    <label class="control-label col-sm-2 pr-3">Intent</label>
                    <select class="selection dropdown col-sm-8" id="search-select" style="height:50px;">
                        <option value="_">unknown</option>
                        {{range $i, $v := $intents}}
                        <option value="{{$v}}">{{$v}}</option>
                        {{end}}
                    </select>
                    <button type="button" id="correctBtn" class="ml-3 btn btn-primary btn-md" disabled="true" style="height: 50px;" onclick="correctIntent()">Correct
                    <i id="loadingCorrect" class="fa fa-spinner fa-spin" style="font-size:24px; color:white; display: none;"></i>
                    </button>
                </div>
            </form>
        </div>
    </div>

</div>

</body>
<script type="text/javascript">
    $('#select')
        .dropdown();

    function searchIntent(){
        var question = $("#question").val();
        var dataString = 'question=' + question;
        $('#loadingIntent').show();

        var status = "";
        $.ajax({
            type: "POST",
            url: "/detectintent",
            data: dataString,
            cache: false,
            success: function(result) {
                $('#loadingIntent').hide();
                $('#correctBtn').attr("disabled", false);
                status = $.trim(result);
                if (status.length == 0) {
                    status = "_";
                }
                var arr = document.getElementById("search-select").options;
                for (var i = 0; i < arr.length; i++) {
                    if (arr[i].value == status) {
                        document.getElementById("search-select").options.selectedIndex = i;
                        break;
                    }
                }
            }
        });
    }   
    function check(){
        searchIntent();
    }
    function submit(e) {
        if (e.keyCode == 13) {
            searchIntent();
        } else {
            return true;
        }
        return false;
    };

    function correctIntent() {
        var question = $("#question").val();
        var intent = $("#search-select").val();
        $('#loadingCorrect').show();
        if (intent == "_") {
            alert("cannot using unknow as intent");
        } else {
            var dataString = 'question=' + question + "&intent=" + intent;
            var status = "";
            $.ajax({
                type: "POST",
                url: "/correcttest",
                data: dataString,
                cache: false,
                success: function(result) {
                    $('#loadingCorrect').hide();
                }
            });
        }
        return false;
    }
</script>

{{end}}