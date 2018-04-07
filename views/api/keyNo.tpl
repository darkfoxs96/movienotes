        <div class="container-fluid">
            <div class="container">
                <div class="row justify-content-center">
                    <div class="col-sm-10 col-md-8 col-lg-6 col-xl-6 pt-4"  style="padding-top: 5px; padding-bottom: 5px; text-align: justify;">
                        <div class="alert alert-danger" role="alert">
                          Такой ключ не существует!
                        </div>
                    </div>
                </div>
                <div class="row justify-content-center text-center">
                    <div class="col-sm-10 col-md-8 col-lg-6 col-xl-6 p-4">
                        <button type="button" class="btn btn-primary" onclick="CreateKeyOpen()">Создать его</button>
                    </div>
                </div>
            </div>
        </div>
    
      
        <div class="modal fade" id="create_key" tabindex="-1" role="dialog" aria-labelledby="create_keyModalLabel" aria-hidden="true">
          <div class="modal-dialog" role="document">
            <div class="modal-content">
              <div class="modal-header">
                <h5 class="modal-title" id="exampleModalLabel">Создать ключ</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                  <span aria-hidden="true">&times;</span>
                </button>
              </div>
              <div class="modal-body">
                  <form name="form_createkey" action="javascript:void(null);" onsubmit="CreateKey()" method="POST">
                       <div class="form-group"> 
                        <label for="id">Ключ</label>
                        <input type="text" class="form-control" placeholder="Ключ" id="createId" name="id" maxlength="30" required>
                       </div> 
                       <div class="form-group"> 
                        <label for="password">Пароль</label>
                        <input type="password" class="form-control" placeholder="Пароль не иенее 6 символов" name="password" id="password" maxlength="30" minlength="6" required>
                       </div>
                       <div id="alert_CreateKey" class="alert alert-danger alert-dismissable" role="alert">
                           <strong>Error!</strong> Некорректные данные или ключ занят.
                       </div>  
                      <button type="submit" class="btn btn-primary">Создать</button>
                    </form>       
              </div>
            </div>
          </div>
        </div>