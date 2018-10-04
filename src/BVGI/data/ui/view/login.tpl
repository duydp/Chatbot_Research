{{define "content"}}
<link rel="stylesheet" href="resources/login/style.css">
<script src="/resources/login/index.js"></script>
<div class="form">
    <div class="row">
        <div class="col-sm-3">
            <a class="navbar-brand px-50" href="/faq">
                <img src="/resources/images/ic-logo.png" width="84px"></img>
            </a>
        </div>
        <div>
            <h1>{{.botname}}</h1>
        </div>
    </div>
    <div class="tab-content">
        <div id="login">
            <form action="/login" method="post">
                <div class="field-wrap">
                    <input name="username" id="name" type="text" required autocomplete="off" />
                </div>
                <div class="field-wrap">
                    <input id="password" name="password" type="password" required autocomplete="off" />
                </div>
                <!-- <p class="forgot"><a href="#">Forgot Password?</a></p> -->
                <button class="btn btn-lg btn-primary col-sm-12" />Log In</button>
            </form>
        </div>
        <div id="signup">
            <!-- Implement later -->
        </div>
    </div>
</div>
{{end}}