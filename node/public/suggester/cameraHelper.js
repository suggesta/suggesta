$(document).ready(function () {

    //APIを格納
    navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia;

    //window.URLのAPIをすべてwindow.URLに統一
    window.URL = window.URL || window.webkitURL;

    if (!navigator.getUserMedia) {
        alert("カメラ未対応のブラウザです。");
    }

    // 変数
    var canvas = document.getElementById("canvasA"),
        context = canvas.getContext("2d"),
        video = document.getElementById("video"),
        btnStart = document.getElementById("start"),
        btnStop = document.getElementById("btnStop"),
        btnPhoto = document.getElementById("photo"),
        videoObj = {
            video: true,
            audio: false
        };

    //キャンバスのサイズ指定は、CSSだけだと、録画した際の幅がずれ、不十分なので、明示的にサイズを指定
    canvas.width = 320;
    canvas.height = 240;

    //再生ボタン押下時
    btnStart.addEventListener("click", function () {
        var localMediaStream;

        if (navigator.getUserMedia) {

            $('#mode-img-disp').hide();

            //カメラの使用の許可
            navigator.getUserMedia(videoObj, function (stream) {
                localMediaStream = stream;
                video.src = window.URL.createObjectURL(localMediaStream);
            }, function (error) {
                alert("カメラの設定を確認してください。");
                console.error("getUserMedia error: ", error.code);
            });




        }


    });

    //停止ボタン押下時の処理
    btnStop.addEventListener("click", function () {
        //localMediaStream.stop();
        location.reload();
    });

    function rect(argHey, argPx, argSec) {

        setInterval(function(){
            argHey.animate({
                marginTop: '-=' + argPx + 'px'
            }, argSec).animate({
                marginTop: '+=' + argPx + 'px'
            }, argSec);
        },1000);


        //setTimeout(rect(argHey), 5000); //アニメーションを繰り返す間隔
    }

    $('#btnDemo').click(function(){
        setTimeout(function () {
            setTimeout(function () {
                var hey = $(kuroArea).find('img')[0];
                console.log($(hey));
                $(hey).show();
                setTimeout(rect($(hey), 30, 1800));
            }, 1000);

            setTimeout(function () {
                var hey = $(kuroArea).find('img')[1];
                console.log($(hey));
                $(hey).show();
                setTimeout(rect($(hey), 15, 2800));
            }, 5500);

            setTimeout(function () {
                var hey = $(kuroArea).find('img')[2];
                console.log($(hey));
                $(hey).show();
                setTimeout(rect($(hey), 20, 3800));
            }, 8000);

            setTimeout(function () {
                var hey = $(kuroArea).find('img')[3];
                console.log($(hey));
                $(hey).show();
                setTimeout(rect($(hey), 12, 2800));
            }, 10000);

            setTimeout(function () {
                var hey = $(kuroArea).find('img')[4];
                console.log($(hey));
                $(hey).show();
                //setTimeout(rect($(hey), 112, 800));
            }, 15000);


            console.log($(kuroArea).find('img')[0]);
        }, 3000);



    });


});