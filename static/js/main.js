//////////Size
WindousSize();

$(window).resize(function() {
    WindousSize();
});

function WindousSize() {
    var width = window.innerWidth;
    var height = window.innerHeight;
    $("#body_size").css("min-height", (height - 100));
    if (width < 610) {
        $("#key_input").css("min-width", (width - 50));
        $("#home_icon").css("display", "none");
        $("#key_icon").css("display", "none");
        $("#key_info").css("display", "none");
        $("#brend").css("display", "none");
        $("#navbar_footer").css("display", "flex");
    } else {
        if (width > 800) {
            $("#key_input").css("min-width", (width - 550));
        } else {
            $("#key_input").css("min-width", 200);
        }
        $("#home_icon").css("display", "block");
        $("#key_icon").css("display", "block");
        $("#key_info").css("display", "block");
        $("#brend").css("display", "block");
        $("#navbar_footer").css("display", "none");
    }
    if (height < 510) {
        $("#navbar_top").removeClass("sticky-top");
        $("#navbar_footer").removeClass("fixed-bottom");
    } else {
        $("#navbar_top").addClass("sticky-top");
        $("#navbar_footer").addClass("fixed-bottom");
    }
}

////////Memory
var memoryGo = false;
var memoryHttp = [];
GetStaticPage();

function AddSource(http, source, version) {
    var body = new Map();
    body.set("source", source);
    if (version == "") {
        body.set("version", "" + $("#version").html());
    } else {
        body.set("version", version);
    }
    var unix_time = window.performance && window.performance.now && window.performance.timing && window.performance.timing.navigationStart ? window.performance.now() + window.performance.timing.navigationStart : Date.now();
    body.set("time", unix_time);
    var saveMap = JSON.stringify(Array.from(body.entries()));
    try {
        localStorage.setItem(http, saveMap);
    } catch (e) {
        localStorage.clear();
        GetStaticPage()
        localStorage.setItem(http, saveMap);
    }
}

function GetSource(http) {
    return new Map(JSON.parse(localStorage.getItem(http)));
}

function GetStaticPage() {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
    if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
        previous = $("#search_input").val();
        try {
            var items = JSON.parse(xmlhttp.responseText);
            var body = new Map();
            body.set("source", items[0]);
            body.set("version", "");
            var unix_time = window.performance && window.performance.now && window.performance.timing && window.performance.timing.navigationStart ? window.performance.now() + window.performance.timing.navigationStart : Date.now();
            body.set("time", unix_time);
            var saveMap = JSON.stringify(Array.from(body.entries()));
            try {
               localStorage.setItem("/", saveMap);
            } catch (e) {
               localStorage.clear();
               localStorage.setItem("/", saveMap);
            }
            body.set("source", items[1]);
            body.set("version", "");
            saveMap = JSON.stringify(Array.from(body.entries()));
            localStorage.setItem("/key/addnote", saveMap);
            body.set("source", items[2]);
            body.set("version", "");
            saveMap = JSON.stringify(Array.from(body.entries()));
            localStorage.setItem("/key/editnote", saveMap);
            body.set("source", items[3]);
            body.set("version", "-1");
            saveMap = JSON.stringify(Array.from(body.entries()));
            localStorage.setItem("/key/keyno", saveMap);
            body.set("source", items[4]);
            body.set("version", "0");
            saveMap = JSON.stringify(Array.from(body.entries()));
            localStorage.setItem("/key/keynan", saveMap);
        } catch(err) {
            console.log(err.message + " in " + xmlhttp.responseText);
            return;
        }
    }
    };
    xmlhttp.open("GET", "/api/page", false);
    xmlhttp.send(null);
}

var keyCreateFunc = "";
function GoBodyPage(http) {
    var body = GetSource(http);
    var keyid = ""; 
    if(body.get("time") == undefined) {
        keyid = http.substr(http.indexOf('keyid=') + 6)
        GetBodyApi(http, "")
        var source = "";
        source = $('#converter').text();
        if(source == "-1") {
            var bodySource = GetSource("/key/keyno");
            $("#body_content").empty();
            $("#body_content").append(bodySource.get("source"));
            keyCreateFunc = keyid;
            AddSource(http, "", "-1");
        } else {
            $("#body_content").empty();
            $("#body_content").append(source);
            AddSource(http, source, "");
        }
        return
    }
    var time = window.performance && window.performance.now && window.performance.timing && window.performance.timing.navigationStart ? window.performance.now() + window.performance.timing.navigationStart : Date.now();
    switch (body.get("version")) {
        case "0":
            keyid = http.substr(http.indexOf('keyid=') + 6)
            if((body.get("time") + 30000) < time){
                GetBodyApi(http, "0")
                var source = "";
                source = $('#converter').text();
                if(source == "") {
                    var bodySource = GetSource("/key/keynan");
                    $("#body_content").empty();
                    $("#body_content").append(bodySource.get("source"));
                    $("#logo_key").html(keyid);
                    AddSource(http, "", "0");
                    return
                }
                if(source == "-1") {
                    var bodySource = GetSource("/key/keyno");
                    $("#body_content").empty();
                    $("#body_content").append(bodySource.get("source"));
                    keyCreateFunc = keyid;
                    AddSource(http, "", "-1");
                    return
                }
                $("#body_content").empty();
                $("#body_content").append(source);
                AddSource(http, source, "")
            } else {
                var bodySource = GetSource("/key/keynan");
                $("#body_content").empty();
                $("#body_content").append(bodySource.get("source"));
                $("#logo_key").html(keyid);
            }
            break;
        case "-1":
            keyid = http.substr(http.indexOf('keyid=') + 6)
            if((body.get("time") + 30000) < time){
                GetBodyApi(http, "-1")
                var source = "";
                source = $('#converter').text();
                if(source == "") {
                    var bodySource = GetSource("/key/keyno");
                    $("#body_content").empty();
                    $("#body_content").append(bodySource.get("source"));
                    keyCreateFunc = keyid;
                    AddSource(http, "", "-1");
                    return
                }
                if(source == "0") {
                    var bodySource = GetSource("/key/keynan");
                    $("#body_content").empty();
                    $("#body_content").append(bodySource.get("source"));
                    $("#logo_key").html(keyid);
                    AddSource(http, "", "0");
                    return
                }
                $("#body_content").empty();
                $("#body_content").append(source);
                AddSource(http, source, "")
            } else {
                var bodySource = GetSource("/key/keyno");
                $("#body_content").empty();
                $("#body_content").append(bodySource.get("source"));
                keyCreateFunc = keyid;
            }
            break;
        case "-2":
            keyid = http.substr(http.indexOf('keyid=') + 6)
            GetBodyApi(http, "")
            var source = "";
            source = $('#converter').text();
            if(source == "-1") {
                var bodySource = GetSource("/key/keynan");
                $("#body_content").empty();
                $("#body_content").append(bodySource.get("source"));
                keyCreateFunc = keyid;
                AddSource(http, "", "-1");
            } else {
                $("#body_content").empty();
                $("#body_content").append(source);
                AddSource(http, source, "");
            }
            break;
        case "":
            if((body.get("time") + 86400000) < time){
                GetStaticPage();
            }
            $("#body_content").empty();
            $("#body_content").append(body.get("source"));
            return
            break;
        default:
            keyid = http.substr(http.indexOf('keyid=') + 6)
            if((body.get("time") + 15000) < time){
                GetBodyApi(http, body.get("version"))
                var source = "";
                source = $('#converter').text();
                if(source == "") {
                    $("#body_content").empty();
                    $("#body_content").append(body.get("source"));
                    AddSource(http, body.get("source"), "");
                    return
                }
                if(source == "0") {
                    var bodySource = GetSource("/key/keynan");
                    $("#body_content").empty();
                    $("#body_content").append(bodySource.get("source"));
                    $("#logo_key").html(keyid);
                    AddSource(http, "", "0");
                    return
                }
                if(source == "-1") {
                    var bodySource = GetSource("/key/keynan");
                    $("#body_content").empty();
                    $("#body_content").append(bodySource.get("source"));
                    keyCreateFunc = keyid;
                    AddSource(http, "", "-1");
                    return
                }
                $("#body_content").empty();
                $("#body_content").append(source);
                AddSource(http, source, "");
            } else {
                $("#body_content").empty();
                $("#body_content").append(body.get("source"));
            }
            break;
    }
}

function GetBodyApi(http, version) {
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
    if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
        previous = $("#search_input").val();
        try {
            $('#converter').text(this.responseText);
            return;  
        } catch(err) {
            console.log(err.message + " in " + xmlhttp.responseText);
            return;
        }
    }
    };
    xmlhttp.open("GET", "/api" + http + "&version=" + version, false);
    xmlhttp.send(null);
}

function StepGo(http) {
    memoryHttp.push(http)
    history.pushState(null, null, http);
    GoBodyPage(http);
    memoryGo = false;
    setTimeout(function() {memoryGo = true;}, 20);
}

function StepBack() {
    if(memoryHttp.length > 1) {
        memoryHttp.pop();
        history.pushState(null, null, memoryHttp[(memoryHttp.length - 1)]);
        GoBodyPage(memoryHttp[(memoryHttp.length - 1)]);
    }
}

$(window).bind('popstate', function() {
    if(memoryGo) {
    StepBack();
    }
});

//NAVBAR
var HomeKey = "";
HomeKey = localStorage.getItem("homekey");

var colClickHome = 0;
function KeyGo() {
    colClickHome++;
    if(colClickHome == 2) {
        $('#home_key_modal').modal('show');
        colClickHome = 0;
        return;
    }
    setTimeout(function() {
        if(colClickHome == 1) {
            if(HomeKey == undefined) {
            $('#home_key_modal').modal('show');
            } else {
                if(HomeKey == "") {
                    $('#home_key_modal').modal('show');
                } else {
                    StepGo("/key/?keyid=" + HomeKey);
                }    
            }
            colClickHome = 0;
        }
    }, 200);
}

function SetHomeKey() {
    HomeKey = $("#memoryhomekey").val();
    localStorage.setItem("homekey", HomeKey)
    $('#home_key_modal').modal('hide');
}

function KeyInOut(inOut) {
    if(inOut == "in"){
        $("#alert_LogIn").hide();
        $('#in_key_modal').modal('show');
    }
    if(inOut == "out"){
        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if (this.readyState == 4) {
                    if (this.status == 200) {
                        $(".key_icon_a").attr('onclick', 'KeyInOut("in")');
                        $(".key_icon_a").removeClass("fa-sign-out");
                        $(".key_icon_a").addClass("fa-key");
                        MyKey = "";
                        $(".key_info").html("Ключ: Нет")
                        HomeKey = "";
                        localStorage.setItem("homekey", "")
                    }
                }
        };
        xhr.open("GET", '/user/outkey', false);
        xhr.send();
    }
}

var MyKey = "";
MyKey = $(".key_info").html();
if(MyKey != "Ключ: Нет") {
   MyKey =  MyKey.substr(6)
   if (MyKey.length > 7) {
        $(".key_info").html("Ключ: " + MyKey.substr(0, 7) + "...");
    } else {
        $(".key_info").html("Ключ: " + MyKey);
    } 
} else {
    MyKey = "";
}

function UserIn() {
    $("#alert_LogIn").hide();
    var formData = new FormData(document.forms.form_login);
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4) {
                if (this.status == 200) {
                    $(".key_icon_a").attr('onclick', 'KeyInOut("out")');
                    $(".key_icon_a").removeClass("fa-key");
                    $(".key_icon_a").addClass("fa-sign-out");
                    MyKey = $("#in_id").val();
                    if (MyKey.length > 7) {
                        $(".key_info").html("Ключ: " + MyKey.substr(0, 7) + "...");
                    } else {
                        $(".key_info").html("Ключ: " + MyKey);
                    }
                    $('#in_key_modal').modal('hide');
                } else {
                    $("#alert_LogIn").hide().show('medium');
                }
            }
    };
    xhr.open("POST", '/user/inkey', true);
    xhr.send(formData);
}

$('.navbar-brand').bind('click', function(){
  StepGo("/");
});

function SearchKey() {
    var val = "";
    val = $("#key_input").val();
    if(val != "") {
        $("#key_input").blur();
        StepGo("/key/?keyid=" + val);
    }
}

//BODY
function CreateKeyOpen() {
    var keyinput = "";
    keyinput = $("#key_input").val();
    if(keyinput != ""){
        $('#createId').val(keyinput);   
    }
    $("#alert_CreateKey").hide();
    $('#create_key').modal('show');
}

function CreateKey() {
    $("#alert_CreateKey").hide();
    var formData = new FormData(document.forms.form_createkey);
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4) {
            if (this.status == 200) {
                $('#create_key').modal('hide');
                setTimeout(function() {
                    $(".key_icon_a").attr('onclick', 'KeyInOut("out")');
                    $(".key_icon_a").removeClass("fa-key");
                    $(".key_icon_a").addClass("fa-sign-out");
                    MyKey = $("#createId").val();
                    if (MyKey.length > 7) {
                        $(".key_info").html("Ключ: " + MyKey.substr(0, 7) + "...")
                    } else {
                        $(".key_info").html("Ключ: " + MyKey)
                    }
                    AddSource("/key/?keyid=" + MyKey, "", "-2")
                    StepGo("/key/?keyid=" + MyKey);
                }, 400);
            } else {
                $("#alert_CreateKey").hide().show('medium');
            }
        }
    };
    xhr.open("POST", '/key/createkey', true);
    xhr.send(formData);
}

function AddNoteOpen() {
    keyaddid = $("#logo_key").html();
    if(keyaddid == MyKey) {
        StepGo("/key/addnote");
        $("#idkey").val(keyaddid);
        addBlurNameAdd();
        $("#alert_addNote").hide();
    } else {
        $("#alert_LogIn").hide();
        $('#in_key_modal').modal('show');
    }
}

function AddNote() {
    $("#alert_addNote").hide();
    var formData = new FormData(document.forms.form_addNote);
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4) {
            if (this.status == 200) {
                AddSource("/key/?keyid=" + MyKey, "", "-2");
                StepGo("/key/?keyid=" + MyKey);
            } else {
                $("#alert_addNote").hide().show('medium');
            }
        }
    };
    xhr.open("POST", '/key/createnote', true);
    xhr.send(formData);
}

function EditNoteOpen(idNote) {
    keyaddid = $("#logo_key").html();
    if(keyaddid == MyKey) {
        var name = $("#nameNote-" + idNote).html();
        var rating = $("#ratingNote-" + idNote).html();
        var description = $("#descriptionNote-" + idNote).html();
        StepGo("/key/editnote");
        $("#alert_editNote").hide();
        $("#name").val(name);
        $("#rating").val(rating);
        $("#description").val(description);
        $("#idnote").val(idNote);
    } else {
        $("#alert_LogIn").hide();
        $('#in_key_modal').modal('show');
    }
}

function EditNote() {
    $("#alert_editNote").hide();
    var formData = new FormData(document.forms.form_editNote);
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4) {
            if (this.status == 200) {
                AddSource("/key/?keyid=" + MyKey, "", "-2");
                StepGo("/key/?keyid=" + MyKey);
            } else {
                $("#alert_editNote").hide().show('medium');
            }
        }
    };
    xhr.open("PUT", '/key/editnote', true);
    xhr.send(formData);    
}

function DeleteNote(idNote) {
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
        if (this.readyState == 4) {
                if (this.status == 200) {
                    $("#note-" + idNote).remove();
                } else {
                    $('#in_id').val($("#logo_key").html());
                    $("#alert_LogIn").hide();
                    $('#in_key_modal').modal('show');
                }
            }
    };
    xhr.open("DELETE", '/key/deletenote?idnote=' + idNote, true);
    xhr.send(null);
}

function addBlurNameAdd() {
    $("#name_add_in").blur(function() {
        var name = "";
        name = $("#name_add_in").val();
        if(name != "") {
            var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if (this.readyState == 4) {
                    if (this.status == 200) {
                        $('#dudu').text(this.responseText);
                        $("#img_add").prop("src", "/static/content/" + $('#dudu').text());
                    }
                }
        };
        xhr.open("GET", '/key/createnote/loadimgapi?name=' + name, true);
        xhr.send(null);
        }
    });
}

function imgMy() {
    $("#img_add").prop("src", "/static/img/noimg.jpg");
}
