$(document).ready(function () {

    navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia;
    window.URL = window.URL || window.webkitURL;
    if (!navigator.getUserMedia) {
        alert("カメラ未対応のブラウザです。");
    }

    var canvas = document.getElementById("canvasA"),
        context = canvas.getContext("2d"),
        video = document.getElementById("video"),
        btnStart = document.getElementById("start"),
        btnStop = document.getElementById("stop"),
        btnPhoto = document.getElementById("photo"),
        videoObj = {
            video: true,
            audio: false
        };

    //TODO
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

            //停止ボタン押下時の処理
            btnStop.addEventListener("click", function () {
                localMediaStream.stop();
                //location.reload();
            });

            //撮影、解析ボタン押下時の処理
            btnPhoto.addEventListener("click", function () {

                $('#msg-disp').show();
                $('#picasoText').text("撮影したファイルは http://localhost:3001/img/hey.jpg ");

                //カメラの映像をキャンバスに描画
                context.drawImage(video, 0, 0, canvas.width, canvas.height);

                var png = canvas.toDataURL('image/jpeg');
                var formData = new FormData();
                formData.append('file', png);

                $.ajax({
                    type: "POST",
                    url: "/api_hey",
                    data: {
                        file: png
                    }
                }).done(function (data) {
                    console.log(data);
                    var alchemyResult = data.hack;
                    if(data.hack.length > 0){
                        var result = data.hack[1];
                        console.log(result);
                    }
                });
            });
        }
    });

});