<div class="container-fluid">
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-sm-6 col-md-4 col-lg-3 col-xl-3" style="padding-top: 5px; padding-bottom: 5px">
                <img src="/static/img/icon.svg" alt="" class="w-100">
            </div>
        </div>
        <div class="row text-center justify-content-center">
            <div class="col-sm-4 col-md-4 col-lg-4 col-xl-4" style="padding-top: 5px; padding-bottom: 5px">
                <h3 style="font-size: 18px;">Movie Notes</h3>
            </div>
        </div>
        <div class="row justify-content-center">
            <div class="col-sm-10 col-md-8 col-lg-6 col-xl-6 searchMobile" style="padding-top: 5px; padding-bottom: 5px; text-align: justify;">
                <p style="font-size: 14px">Сервис позволяющий создавать заметки понравившимся или нет Вам фильмам, что бы в будущем решать что пересмотреть, а что не стоит, в заметке достаточно указать название фильма и оценку (от 1 до 5) в зависимости от оценки, записи будут подсвечиваться определённым цветом. Изображение подберётся автоматически либо можно вставить своё. Описание не обязательно.</p>
            </div>
        </div>
        <div class="row justify-content-center">
            <div class="col-sm-10 col-md-8 col-lg-6 col-xl-6" style="padding-top: 5px; padding-bottom: 5px; text-align: justify;">
                <p style="font-size: 14px">Все заметки хранятся в ключе, просмотр записей не требует пароля (так как информации о владельце нет, не кто не знает чьи это заметки). Для редактирования либо добавления записи уже понадобиться знать пароль.</p>
            </div>
        </div>
        <div class="row justify-content-center">
            <div class="col-sm-10 col-md-8 col-lg-6 col-xl-6" style="padding-top: 5px; padding-bottom: 5px; text-align: justify;">
                <p style="font-size: 14px">При нажатии на кнопку дом: будет предложено ввести ключ. После ввода ключа при повторном нажатии на дом вас перенаправят на введённый ранее ключ. Если нажать дважды на дом: предложит изменить ключ.</p>
            </div>
        </div>
        <div class="row justify-content-center text-center">
            <div class="col-sm-10 col-md-8 col-lg-6 col-xl-6 p-4">
                <button type="button" class="btn btn-primary" onclick="CreateKeyOpen()">Создать ключ</button>
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
