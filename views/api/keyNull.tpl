        <div class="container-fluid">
            <div class="container">
                <div class="row justify-content-center text-center">
                    <div class="col-sm-4 col-md-4 col-lg-4 col-xl-4"  style="padding-top: 5px; padding-bottom: 5px">
                        <h3 style="font-size: 18px;" id="logo_key">{{ .KeyViews }}</h3>
                    </div>
                </div>
                <div class="row justify-content-center text-center">
                    <div class="col-sm-4 col-md-4 col-lg-4 col-xl-4"  style="padding-top: 5px; padding-bottom: 20px">
                        <button type="button" class="btn btn-primary" onclick="AddNoteOpen()">Добавить запись</button>
                    </div>
                </div>
            </div>
        </div>
        <noindex style="display: none" id="version">{{ .Version }}</noindex>