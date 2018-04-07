<div class="container-fluid">
    <div class="container">
        <div class="row justify-content-center text-center">
            <div class="col-sm-8 col-md-6 col-lg-5 col-xl-5">
                <h3 style="font-size: 18px; padding: 10px">Редактирование записи</h3>
            </div>
        </div>
        <div class="row justify-content-center text-center">
            <div class="col-sm-8 col-md-6 col-lg-5 col-xl-5">
                <form name="form_editNote" action="javascript:void(null);" onsubmit="EditNote()" method="PUT" style="padding-bottom: 10px;">
                    <div class="form-group">
                        <label for="name">Название</label>
                        <input type="text" class="form-control" name="name" value="" placeholder="Название" id="name" maxlength="20" required>
                    </div>
                    <div class="form-group">
                        <label for="rating">Оценка</label>
                        <input type="number" class="form-control" name="rating" value="" min="1" max="5" id="rating" required>
                    </div>
                    <div class="form-group">
                        <label for="note">Заметка</label>
                        <textarea class="form-control" name="note" id="description" rows="7" maxlength="255"></textarea>
                    </div>
                    <div class="form-group" style="display: none">
                        <label for="idnote">idNote</label>
                        <input type="number" class="form-control" name="idnote" value="" id="idnote" maxlength="30" required>
                    </div>
                    <div id="alert_editNote" class="alert alert-danger alert-dismissable" role="alert">
                        <strong>Error!</strong> Некорректные данные.
                    </div>
                    <button type="submit" class="btn btn-primary">Изменить</button>
                </form>
            </div>
        </div>
    </div>
</div>
