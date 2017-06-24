var express = require('express');
var router = express.Router();
var fs = require("fs");

/* GET home page. */
router.get('/test', function(req, res, next) {
    console.log("this is just a test.....");
    console.log(req);
    console.log(res);
});

router.get('/api', function(req, res, next) {
    console.log("this is just a test.....");
    console.log(req);
    console.log(res);
});

router.post('/api_hey', function (req, res) {

    var isDebug = true;

    /** POST されたデータ*/
    var postedData = req.body;

    /** POSTされてきた画像データ */
    var imgData = postedData.file;

    /** POST　送信元に返却するデータ */
    var dataList = {
        file: postedData.file,
        memo: ""
    };

    if (imgData) {

        //画像ファイルを生成
        var base64Data = imgData.replace(/^data:image\/jpeg;base64,/, '');
        var filePath = "public/img/hey.jpg"; //画像保存場所
        console.log('file path is ' + filePath);
        fs.writeFile(filePath, base64Data, 'base64', function (err) {
            if (err) {
                console.log("File write error..." + err);
            } else {
                console.log("file write success...aaa");

                if (!isDebug) {
                    //生成した画像または。保存したimg url をへGOサーバーへ送信
                    //public/img/hey.jpg
                } else {
                    //デバッグの場合 決め打ち
                    dataList['hack'] = [
                        {score: '0.9', text: 'test1'},
                        {score: '0.2', text: 'test2'}
                    ];
                    res.send(dataList);
                }
            }
        });
    } else {
        console.log("no image data...");
    }
});

module.exports = router;
