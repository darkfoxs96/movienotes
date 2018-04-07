{{template "bundle/main.tpl" .}}
      
{{define "body"}}
        <div class="container-fluid">
            <div class="container">
                <div class="row justify-content-center">
                    <div class="col-sm-10 col-md-8 col-lg-6 col-xl-6 pt-4"  style="padding-top: 5px; padding-bottom: 5px; text-align: justify;">
                        <div class="alert alert-success" role="alert">
                          {{ .Massege }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
{{end}}