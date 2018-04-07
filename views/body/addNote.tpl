{{template "bundle/main.tpl" .}}
      
{{define "body"}}     
           <div class="container-fluid">
            <div class="container">
                <div class="row justify-content-center text-center">  
                   <div class="col-sm-8 col-md-6 col-lg-5 col-xl-4">
                       <h3 style="font-size: 18px; padding: 10px">Добавление записи</h3>
                       <img src="/static/img/noimg.jpg" id="img_add" alt="" class="w-100">
                   </div>
                </div>   
                <div class="row justify-content-center text-center">
                    <div class="col-sm-8 col-md-6 col-lg-5 col-xl-5">
                     <form method="POST" name="form_addNote" action="javascript:void(null);" onsubmit="AddNote()" enctype="multipart/form-data" style="padding-bottom: 10px;">
                       <div class="form-group">
                        <label for="img">Выберите изображение(.jpeg/.jpg) до 2мб. или авто. найдёт по названию</label>
                        <input type="file"  name="img" class="form-control-file" accept=".jpeg,.jpg" id="exampleFormControlImg" onclick="imgMy()">
                       </div>
                       <div class="form-group"> 
                        <label for="name">Название</label>
                        <input type="text" class="form-control" name="name" placeholder="Название не более 20 символов" id="name_add_in" maxlength="20" required>
                       </div>
                       <div class="form-group"> 
                        <label for="rating">Оценка</label>
                        <input type="number" class="form-control" name="rating" value="1" min="1" max="5" id="rating" required>
                       </div>  
                       <div class="form-group"> 
                        <label for="note">Заметка</label>
                        <textarea class="form-control" name="note" id="description" rows="7" maxlength="255"></textarea>
                       </div>
                       <div class="form-group" style="display: none"> 
                        <label for="idkey">idKey</label>
                        <input type="text" class="form-control" name="idkey" value="{{ .IdKey }}" id="idkey" maxlength="30" required>
                       </div>
                        <div id="alert_addNote" class="alert alert-danger alert-dismissable" role="alert">
                           <strong>Error!</strong> Некорректные данные.
                        </div>   
                      <button type="submit" class="btn btn-primary">Добавить</button>
                    </form>
                    </div>
                </div>
            </div>
        </div>
        <h3 style="display: none" id="dudu"></h3>
{{end}}