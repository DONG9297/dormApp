const express = require('express')
const router = express.Router()
const request = require('request')

router.get('/choose', (req, res) => {
    request({
        url: "http://dorm:10713/get_dormList",
        method: "POST",
        json: true,
        headers: {
            'content-type': 'application/json;charset=UTF-8',
            "Cookie": "jwt=" + req.cookies.jwt
        },
    }, function (error, response) {
        if (!error && response.statusCode === 200) {
            // 获取宿舍列表成功
            res.render('choose.html', {
                dorm_list: response.body.data.dorm_list,
                building_list: response.body.data.building_list
            })
        } else if (response.statusCode === 401) {
            res.render("login.html");
        } else {
            // 获取宿舍列表失败
            res.render('choose.html', {
                dorm_list: null,
                building_list: null
            })
        }
    })
})

router.post('/choose', (req, res) => {
    let data = {
        building: req.body.building,
        stucode0: req.body.stucode0,
        stucode1: req.body.stucode1,
        stucode2: req.body.stucode2
    };
    // 判断格式
    let flag2 = data.building != null;
    let flag3 = data.stucode0 === "" || /^[0-9]{8}$/.test(data.stucode0);
    let flag4 = data.stucode1 === "" || /^[0-9]{8}$/.test(data.stucode1);
    let flag5 = data.stucode2 === "" || /^[0-9]{8}$/.test(data.stucode2);

    if (flag2 && flag3 && flag4 && flag5) {
        // 发送请求
        request({
                url: "http://dorm:10713/choose",
                method: "POST",
                json: true,
                headers: {
                    'content-type': 'application/json;charset=UTF-8',
                    "Cookie": "jwt=" + req.cookies.jwt
                },
                body: data
            }, function (error, response) {
                if (!error && response.statusCode === 200) {
                    res.send("选宿舍提交成功，<a href='/'>去首页</a>，或<a href='/order/result'\>查看结果</a>");
                } else if (response.statusCode === 401) {
                    res.render("login.html");
                } else {
                    res.send("选宿舍提交失败，请重新提交");
                }
            }
        )
    } else {
        res.send("请填写正确的信息");
    }
})

router.get('/result', (req, res) => {
    request({
        url: "http://dorm:10713/get_result",
        method: "POST",
        json: true,
        headers: {
            'content-type': 'application/json;charset=UTF-8',
            "Cookie": "jwt=" + req.cookies.jwt
        },
    }, function (error, response, body) {
        if (!error && response.statusCode === 200) {
            res.render('result.html', {
                user_name: body.data.user.name,
                status: true,
                dorm_info: body.data.dorm_info,
            })
        } else if (response.statusCode === 401) {
            res.render("login.html");
        } else {
            res.render('result.html', {
                user_name: body.data.user.name,
                status: false,
                dorm_info: null,
            })
        }
    })
})

module.exports = router;