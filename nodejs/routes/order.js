const express = require('express')
const router = express.Router()
const request = require('request')
const {response} = require("express");

router.get('/choose', (req, res) => {
    isLogged(req.cookies, function (flag, response) {
        let gender = response.body.data.user.gender;
        if (flag) {
            // 获取宿舍列表
            request({
                url: "http://dorm:10713/get_dormList",
                method: "POST",
                json: true,
                headers: {
                    'content-type': 'application/json;charset=UTF-8',
                },
                body: {
                    gender: gender,
                }
            }, function (error, response, body) {
                if (!error && response.statusCode == 200) {
                    // 获取宿舍列表成功
                    res.render('choose.html', {
                        gender: gender,
                        dorm_list: response.body.data.dorm_list
                    })
                } else {
                    res.render('choose.html', {
                        gender: null,
                        dorm_list: null
                    })
                }
            })
        } else {
            res.send("未登录，去<a href='/user/login'\>登录</a>或<a href='/user/register'\>注册</a>");
        }
    })
})

router.post('/choose', (req, res) => {
    isLogged(req.cookies, function (flag, response) {
        let data = {
            user_id: response.body.data.user.id,
            gender: response.body.data.user.gender,
            building : req.body.building,
            stucode0 : req.body.stucode0,
            stucode1 : req.body.stucode1,
            stucode2 : req.body.stucode2
        };
        if (flag) {
            // 判断格式
            let flag2 = data.building != null;
            let flag3 = data.stucode0 == "" || /^[0-9]{8}$/.test(data.stucode0);
            let flag4 = data.stucode1 == "" || /^[0-9]{8}$/.test(data.stucode1);
            let flag5 = data.stucode2 == "" || /^[0-9]{8}$/.test(data.stucode2);

            if (flag2 && flag3 && flag4 && flag5) {
                // 发送请求
                request({
                    url: "http://dorm:10713/choose",
                    method: "POST",
                    json: true,
                    headers: {
                        'content-type': 'application/json;charset=UTF-8',
                    },
                    body: data
                }, function (error, resp, body) {
                    if (!error && resp.statusCode == 200) {
                        res.send("选宿舍提交成功，<a href='/'>去首页</a>，或<a href='/order/result'\>查看结果</a>");
                    } else {
                        res.send(resp.body.message);
                    }
                })
            } else {
                res.send("请填写正确的信息");
            }
        }
    })
})

router.get('/result', (req, res) => {
    isLogged(req.cookies, function (flag, response) {
        if (flag) {
            request({
                url: "http://dorm:10713/get_result",
                method: "POST",
                json: true,
                headers: {
                    'content-type': 'application/json;charset=UTF-8',
                },
                body: {
                    user_id: response.body.data.user.id,
                }
            }, function (error, resp, body) {
                if (!error && resp.statusCode == 200) {
                    res.render('result.html', {
                        user_name: response.body.data.user.name,
                        status: true,
                        dorm_info: resp.body.data.dorm_info,
                    })
                } else {
                    res.render('result.html', {
                        user_name: response.body.data.user.name,
                        status: false,
                        dorm_info: null,
                    })
                }
            })
        } else {
            res.send("未登录，去<a href='/user/login'\>登录</a>或<a href='/user/register'\>注册</a>");
        }
    })
})

// 判断是否登录
function isLogged(cookies, fn) {
    if (cookies == null) {
        fn(false, response);
    }
    var cookieValue = cookies.name;
    request({
        url: "http://login:10712/is_logged",
        method: "POST",
        json: true,
        headers: {
            'content-type': 'application/json;charset=UTF-8',
        },
        body: {
            session_id: cookieValue
        }
    }, function (error, response, body) {
        if (!error && response.statusCode == 200) {
            fn(true, response);
        } else {
            fn(false, response);
        }
    })
}

module.exports = router;