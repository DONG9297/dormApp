const express = require('express');
const router = express.Router();
const request = require('request');
const {response} = require("express");

router.get('/', (req, res) => {
    isLogged(req.cookies, function (flag, response) {
        if (flag) {
            res.render('index.html', {user: response.body.data.user});
        } else {
            res.send("请<a href='/user/login'\>登录</a>或<a href='/user/register'\>注册</a>");
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
            session_id: cookieValue,
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