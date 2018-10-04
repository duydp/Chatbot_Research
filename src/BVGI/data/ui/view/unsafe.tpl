{{define "content"}} {{$intents := index $ "intents"}}

<div class="jumbotron mb-0 pb-5">
    <div class="container">
        <div class="col-sm-12 col-center-block mb-5">
            <h1 align="center" class="text-primary">UNSAFE</h1>
        </div>

        <div class="row justify-content-md-center" style="min-height: 400px;">
            <div class="row col-sm-6">
                <div class="col-sm-6">
                     <button type="button" id="clean" class="col-sm-11 btn btn-danger btn-lg" style="height: 50px;">CLEAN
                        <i id="loadingClean" class="fa fa-spinner fa-spin ml-2" style="font-size:24px; color:white; display: none;"></i>
                    </button>
                </div>
                
                <div class="col-sm-6">
                     <button type="button" id="trainall" disabled="true" class="col-sm-11 btn btn-primary btn-lg" style="height: 50px;">TRAIN
                        <i id="loadingTrain" class="fa fa-spinner fa-spin ml-2" style="font-size:24px; color:white; display: none;"></i>
                    </button>
                </div>
            </div>
        </div>

    </div>
</div>

</body>
<script type="text/javascript">
$(document).ready(function() {

    var isActive = true;
    pollServer();
    function pollServer()
    {
        if (isActive){
            $('#loadingTrain').show();
            window.setTimeout(function () {
                $.ajax({
                    url: "/trainstatus",
                    type: "POST",
                    success: function (result) {
                        console.log(result+"\n");
                        if(result.trim() == 'Done'){
                            isActive = false;
                            $('#loadingTrain').hide();
                            $('#trainall').prop('disabled', true);
                            $('#clean').prop('disabled', false);

                            return;
                        }else{
                            $('#loadingTrain').show();
                            $('#trainall').prop('disabled', true);
                            $('#clean').prop('disabled', true);
                            
                            pollServer();    
                        }                            
                    },
                    error: function () {
                    }});
            }, 1000);
        }else{
            $('#loadingTrain').hide();
            $('#trainall').prop('disabled', true);
            $('#clean').prop('disabled', false);
        }
    }



    $("#trainall").on("click", function(){
        var ok = confirm("Are you sure?");
        if (ok) {
            $('#loadingTrain').show();
            $('#trainall').prop('disabled', true);
            $('#clean').prop('disabled', true);
            isActive = true;
            pollServer();   
            $.ajax({
                type: "POST",
                url: "/unsafe",
                data: "",
                cache: false,
                success: function(result) {
                    // $('#loadingTrain').hide();
                    // $('#clean').prop('disabled', false);
                }
            });
            
        }
        return false;
     });

    $("#clean").on("click", function(){
        var ok = confirm("Are you sure?");
        if (ok) {
            $('#loadingClean').show();
            $('#clean').prop('disabled', true);
            $.ajax({
                type: "DELETE",
                url: "/unsafe",
                data: "",
                cache: false,
                success: function(result) {
                    $('#loadingClean').hide();
                    $('#trainall').prop('disabled', false);
                    $('#clean').prop('disabled', false);
                }
            });   
        }
        return false;
     });
    
});
</script>

{{end}}