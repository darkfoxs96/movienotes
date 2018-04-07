<nav class="navbar navbar-expand-lg navbar-light bg-light sticky-top" id="navbar_top">
    <a href="#" class="navbar-brand" id="brend">
            <img src="/static/img/icon.svg" width="30" height="30" alt="logo">
            MovieNotes</a>
    <form class="form-inline my-2 my-lg-0 mr-auto" action="javascript:void(null);" onsubmit="SearchKey()" method="GET">
        <input type="text" id="key_input" name="keyid" class="form-control mr-2 ml-2 form-success" style="min-width: 300px" placeholder="Введите ключ" aria-label="Search" required autocomplete="on">
    </form>
    <a id="home_icon" class="pr-2" href="javascript:void(null);" onclick="KeyGo()"><i class="fa fa-home fa-2x navbarIcon" aria-hidden="true"></i></a>
    <a id="key_icon" class="pr-2 navbar_icon" href="javascript:void(null);"><i class="fa fa-{{ .IconKey }}  fa-2x navbarIcon key_icon_a" aria-hidden="true" onclick="KeyInOut('{{ .FunKey }}')"></i></a>
    <a id="key_info" class="badge badge-light key_info">Ключ: {{ .Key }}</a>
</nav>

<nav class="navbar justify-content-between navbar-light bg-light fixed-bottom" id="navbar_footer" style="display: none">
    <a href="#" class="navbar-brand">
        <img src="/static/img/icon.svg" width="30" height="30" alt="logo">
        MovieNotes</a>
    <a id="home_icon_footer" class="pr-2 navbar_icon" href="javascript:void(null);" onclick="KeyGo()"><i class="fa fa-home fa-2x navbarIcon" aria-hidden="true"></i></a>
    <a id="key_icon_footer" class="pr-2 navbar_icon" href="javascript:void(null);"><i class="fa fa-{{ .IconKey }} fa-2x navbarIcon key_icon_a" aria-hidden="true" onclick="KeyInOut('{{ .FunKey }}')"></i></a>
    <a id="key_info_footer" class="badge badge-light key_info">Ключ: {{ .Key }}</a>
</nav>


<div class="modal fade" id="home_key_modal" tabindex="-1" role="dialog" aria-labelledby="home_keyModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="exampleModalLabel">Запомнить ключ</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
            </div>
            <div class="modal-body">
                <form action="javascript:void(null);" onsubmit="SetHomeKey()">
                    <div class="form-group">
                        <label for="memoryid">Ключ</label>
                        <input type="text" class="form-control" placeholder="Ключ" id="memoryhomekey" name="memoryid" minlength="6" maxlength="30" required>
                    </div>
                    <button type="submit" class="btn btn-primary">Запомнить</button>
                </form>
            </div>
        </div>
    </div>
</div>

<div class="modal fade" id="in_key_modal" tabindex="-1" role="dialog" aria-labelledby="in_keyModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="exampleModalLabel">Авторизовать ключ</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
            </div>
            <div class="modal-body">
                <form name="form_login" action="javascript:void(null);" onsubmit="UserIn()" method="POST">
                    <div class="form-group">
                        <label for="id">Ключ</label>
                        <input type="text" class="form-control" placeholder="Ключ" id="in_id" name="id" maxlength="30" minlength="6" required>
                    </div>
                    <div class="form-group">
                        <label for="password">Пароль</label>
                        <input type="password" class="form-control" placeholder="Пароль не менее 6 символов" name="password" id="in_password" maxlength="30" minlength="6" required>
                    </div>
                    <div id="alert_LogIn" class="alert alert-danger alert-dismissable" role="alert">
                        <strong>Error!</strong> Неправильный логин или пароль.
                    </div>
                    <button type="submit" class="btn btn-primary">Авторизация</button>
                </form>
            </div>
        </div>
    </div>
</div>
