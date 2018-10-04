<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>FPT Telecom Chatbot</title>

    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css" integrity="sha384-rwoIResjU2yc3z8GV/NPeZWAv56rSmLldC3R/AZzGRnGxQQKnKkoFVhFQhNUwEyJ" crossorigin="anonymous">

    <script src="https://code.jquery.com/jquery-3.1.1.slim.min.js" integrity="sha384-A7FZj7v+d/sdmMqp/nOQwliLvUsJfDHW+k9Omg/a/EheAdgtzNs3hpfag6Ed950n" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js" integrity="sha384-DztdAPBWPRXSA/3eYEEUWrWCy7G5KFbe8fFjk5JAIxUYHKkDx6Qin1DkWx51bBrb" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js" integrity="sha384-vBWWzlZJ8ea9aCX4pEW3rVHjgjt7zpkNpZk+02D9phzyeVkE+jo0ieGizqPLForn" crossorigin="anonymous"></script>

    <link rel="shortcut icon" href="/resources/images/favicon.ico" type="image/x-icon">

    <!-- Custom styles for this template -->
    <link href="/resources/jumbotron.css" rel="stylesheet">
    
    <script src="/resources/escape.js"></script>
    <script src="/resources/dist/dropdown.js"></script>
    <link rel="stylesheet" href="/resources/footer-distributed.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <link rel="stylesheet" type="text/css" href="/resources/dist/dropdown.min.css">

    <script src="/resources/jquery.min.js"></script>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.2.0/css/font-awesome.min.css">

    <meta name="intercoolerjs:use-actual-http-method" content="true" />
    
    <script src="https://intercoolerreleases-leaddynocom.netdna-ssl.com/intercooler-1.1.1.min.js"></script>

     <script src="https://cdnjs.cloudflare.com/ajax/libs/tether/1.2.0/js/tether.min.js" integrity="sha384-Plbmg8JY28KFelvJVai01l8WyZzrYWG825m+cZ0eDDS1f7d/js6ikvy1+X+guPIB" crossorigin="anonymous"></script>

     <script src="/resources/jquery.twbsPagination.js" type="text/javascript"></script>

</head>

<body>
    {{with not .botname}}
    <nav class="navbar navbar-toggleable-md navbar-inverse fixed-top bg-inverse pl-5 pr-5">
        <a id="top" class="top"></a>
        <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="px-50">
            <a class="navbar-brand px-50" href="/faq">
                <img src="/resources/images/ic-logo.png" width="64px"></img>
            </a>
        </div>

        <div class="collapse navbar-collapse" id="navbarsExampleDefault">
            <ul class="navbar-nav mr-auto" id="menu">
                <li class="nav-item" data-id="1">
                    <a class="nav-link" href="/faq">FAQ</a>
                </li>
                <li class="nav-item" data-id="2">
                    <a class="nav-link" href="/history/1">History</a>
                </li>
                <li class="nav-item" data-id="3">
                    <a class="nav-link" href="/staff">Staff</a>
                </li>
                <li class="nav-item" data-id="4">
                    <a class="nav-link" href="/test">Test</a>
                </li>
                <li class="nav-item" data-id="5">
                    <a class="nav-link" href="/help">Help</a>
                </li>
            </ul>
            <div class="pull-xs-right rightBtns">
                <button class="btn btn-outline-success my-2 my-sm-0" type="submit" id="btnTrain">Train</button>
                <button class="btn btn-outline-danger my-2 my-sm-0" type="submit" onclick="javascript:location.href='/logout'">Logout</button>
            </div>
        </div>
    </nav>

    <script type="text/javascript">
        var isActive = true;
        pollServer();
        function pollServer()
        {
            if (isActive)
            {
                window.setTimeout(function () {
                    $.ajax({
                        url: "/trainstatus",
                        type: "POST",
                        success: function (result) {
                            // $('#status').append(result.trim() + "_");
                            if(result.trim() == 'Done'){
                                isActive = false;
                                $('#btnTrain').attr("disabled", false);
                                $('#btnTrain').text("Train");
                                return;
                            }else{
                                $('#btnTrain').attr("disabled", true);
                                $('#btnTrain').text("Training");
                                pollServer();    
                            }                            
                        },
                        error: function () {
                        }});
                }, 1000);
            }
        }
        $('#btnTrain').on("click", function(e){
            isActive = true;
            $('#btnTrain').attr("disabled", true);
            $('#btnTrain').text("Training");
             $.ajax({
                    type: "POST",
                    url: "/train",
                    data: "",
                    async: true,
                    cache: false,
                    success: function(result) {
                    }
            });
            pollServer();

        });

        $('#menu').on("click","li", function(e) {
            var self = $(this);
            self.addClass('active');
            e.preventdefault();
        });
    </script>
    {{end}} 

    {{template "content" .}} 

    {{with not .botname}}
    <footer class="footer-distributed my-0">
        <div class="footer-left">
            <p class="footer-links">
                <a href="/faq">FAQ</a> ·
                <a href="/history">History</a> ·
                <a href="/staff">Staff</a> ·
                <a href="/test">Test</a> ·
                <a href="/help">Help</a>
            </p>

            <p>FPT.AI &copy; 2017</p>
        </div>

    </footer>
    </div>
    <!-- /container -->
    {{end}}

</body>

</html>