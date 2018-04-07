{{template "bundle/main.tpl" .}}
      
{{define "body"}}
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
                {{ range .Notes }}
                <div  id="note-{{ .Id }}" class="row">
                    <div class="col-sm-5 col-md-4 col-lg-3 col-xl-3" style="padding-top: 5px; padding-bottom: 5px">
                        <i class="fa fa-pencil fa-lg fafaEdit" aria-hidden="true" onclick="EditNoteOpen('{{ .Id }}')"></i>
                        <i class="fa fa-trash fa-lg fafaDelete" aria-hidden="true" onclick="DeleteNote('{{ .Id }}')"></i>
                        <img src="/static/content/{{ .Id }}.jpg" alt="" class="w-100 imgColorBorder-{{ .Rating }}">
                    </div>
                    <div class="col-sm-7 col-md-8 col-lg-9 col-xl-9"  style="padding-top: 5px; padding-bottom: 5px">
                        <h3 id="nameNote-{{ .Id }}" style="font-size: 18px;">{{ .Name }}</h3>
                        <a class="badge badge-light">{{ .Rating }}/5</a>
                        <a id="ratingNote-{{ .Id }}" style="display: none">{{ .Rating }}</a>
                        <p id="descriptionNote-{{ .Id }}" style="font-size: 14px">{{ .Note }}</p>
                    </div>
                </div>
                {{ end }}
            </div>
        </div>
{{end}}