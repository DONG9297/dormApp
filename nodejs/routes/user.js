const express = require('express')
const router = express.Router()
const request = require('request')

const md5 = require('md5-node')
const {response} = require("express");

// 注册
router.get('/register', (req, res) => {
    res.render('register.html');
})
router.post('/register', (req, res) => {
    var phone = req.body.phone;
    var stu_id = req.body.stu_id;
    var name = req.body.name;
    var gender = req.body.gender;
    var password = req.body.password;

    var flag1 = /^(13|14|15|17|18)[0-9]{9}$/.test(phone);
    var flag2 = /^[0-9]{10}$/.test(stu_id);
    var flag3 = /^[\u4e00-\u9fa5]{0,10}$/.test(name);
    var flag4 = gender == '男' || gender == '女';
    var flag5 = /^([\w_]{6,20})$/.test(password);

    if (flag1 && flag2 && flag3 && flag4 && flag5) {
        password = md5(password)
        request({
            url: "http://register:10711/register",
            method: "POST",
            json: true,
            headers: {
                'content-type': 'application/json;charset=UTF-8',
            },
            body: {
                phone: phone,
                stu_id: stu_id,
                name: name,
                gender: gender,
                password: password
            }
        }, function (error, response, body) {
            if (!error && response.statusCode == 200) {
                res.send("注册成功，去<a href='/user/login'\>登录</a>");
            } else if (response.body.message == "手机号已注册") {
                res.send("手机号已存在，去<a href='/user/login'\>登录</a>");
            } else if (response.body.message == "学号已注册") {
                res.send("学号已存在，去<a href='/user/login'\>登录</a>");
            }
        })
    } else {
        res.send("注册失败，请重新注册")
    }
})

// 登录
router.get('/login', (req, res) => {
    isLogged(req.cookies, function (flag) {
        if (flag) {
            res.send("已登录，<a href='/'>去首页</a>，或<a href='/user/logout'\>退出</a>");
        } else {
            res.render('login.html');
        }
    })
})
router.post('/login', (req, res) => {
    var phone = req.body.phone;
    var password = req.body.password;
    var reg1 = /^(13|14|15|17|18)[0-9]{9}$/.test(phone)
    var reg2 = /^([\w_]{6,20})$/.test(password)

    if (reg1 && reg2) {
        password = md5(password)
        request({
            url: "http://login:10712/login",
            method: "POST",
            json: true,
            headers: {
                'content-type': 'application/json;charset=UTF-8',
            },
            body: {
                phone: phone,
                password: password
            }
        }, function (error, response, body) {
            if (!error && response.statusCode == 200) {
                // 设置 Cookie
                var cookieValue = response.body.data.session.session_id;
                res.cookie("name", cookieValue);//,{maxAge: 6000}
                res.send("登录成功，<a href='/'>去首页</a>，或<a href='/user/logout'\>退出</a>");
            }
        })
    } else {
        res.send("登录失败，请重新登录")
    }
})

// 退出
router.get('/logout', (req, res) => {
    isLogged(req.cookies, function (flag, response) {
        if (flag) {
            request({
                url: "http://login:10712/logout",
                method: "POST",
                json: true,
                headers: {
                    'content-type': 'application/json;charset=UTF-8',
                },
                body: {
                    session_id: req.cookies.name
                }
            }, function (error, response, body) {
                if (!error && response.statusCode == 200) {
                    // 设置 Cookie
                    res.cookie("name", req.cookies.name, {maxAge: 0});
                    res.send("退出成功，请重新<a href='/user/login'\>登录</a>");
                } else {
                    res.send("未登录，去<a href='/user/login'\>登录</a>");
                }
            })
        } else {
            res.send("未登录，去<a href='/user/login'\>登录</a>");
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
