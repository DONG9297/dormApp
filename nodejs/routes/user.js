const express = require('express')
const router = express.Router()
const request = require('request')

const md5 = require('md5-node')

// 注册
router.get('/register', (req, res) => {
    res.render('register.html');
})
router.post('/register', (req, res) => {
    let data = {
        "phone": req.body.phone,
        "stu_id": req.body.stu_id,
        "name": req.body.name,
        "gender": req.body.gender,
        "password": req.body.password
    }
    let flag1 = /^(13|14|15|17|18)[0-9]{9}$/.test(data.phone);
    let flag2 = /^[0-9]{10}$/.test(data.stu_id);
    let flag3 = /^[\u4e00-\u9fa5]{0,10}$/.test(data.name);
    let flag4 = data.gender === '男' || data.gender === '女';
    let flag5 = /^([\w_]{6,20})$/.test(data.password);

    if (flag1 && flag2 && flag3 && flag4 && flag5) {
        data.password = md5(data.password)
        request({
            "url": "http://register:10711/register",
            "method": "POST",
            "json": true,
            "headers": {
                'content-type': 'application/json;charset=UTF-8',
            },
            "body": data
        }, function (error, response, body) {
            if (!error && response.statusCode === 200) {
                res.send("注册成功，去<a href='/user/login'\>登录</a>");
            } else if (body.message === "手机号已注册") {
                res.send("手机号已存在，去<a href='/user/login'\>登录</a>");
            } else if (body.message === "学号已注册") {
                res.send("学号已存在，去<a href='/user/login'\>登录</a>");
            }
        })
    } else {
        res.send("请填写正确的信息")
    }
})

// 登录
router.get('/login', (req, res) => {
    request({
        url: "http://login:10712/is_logged",
        method: "POST",
        json: true,
        headers: {
            'content-type': 'application/json;charset=UTF-8',
            "Cookie": "jwt=" + req.cookies.jwt
        },
    }, function (error, response, body) {
        if (!error && response.statusCode === 200) {
            res.render('index.html', {
                user: body.data.user
            });
        } else {
            res.render('login.html');
        }
    })
})
router.post('/login', (req, res) => {
    let data = {
        "phone": req.body.phone,
        "password": req.body.password
    }
    let reg1 = /^(13|14|15|17|18)[0-9]{9}$/.test(data.phone)
    let reg2 = /^([\w_]{6,20})$/.test(data.password)

    if (reg1 && reg2) {
        data.password = md5(data.password)
        request({
            url: "http://login:10712/login",
            method: "POST",
            json: true,
            headers: {
                'content-type': 'application/json;charset=UTF-8',
            },
            body: data
        }, function (error, response, body) {
            if (!error && response.statusCode === 200) {
                // 设置 Cookie
                let cookieValue = response.body.data.token;
                res.cookie("jwt", cookieValue);//,{maxAge: 6000}
                res.send("登录成功，<a href='/'>去首页</a>，或<a href='/user/logout'\>退出</a>");
            }
        })
    } else {
        res.send("用户名或密码不正确")
    }
})

// 退出
router.get('/logout', (req, res) => {
    request({
        url: "http://login:10712/logout",
        method: "POST",
        json: true,
        headers: {
            'content-type': 'application/json;charset=UTF-8',
            "Cookie": "jwt=" + req.cookies.jwt
        },
    }, function (error, response, body) {
        if (!error && response.statusCode === 200) {
            // 设置 Cookie
            res.cookie("jwt", req.cookies.jwt, {maxAge: 0});
            res.send("退出成功，请重新<a href='/user/login'\>登录</a>");
        } else {
            res.render("login.html");
        }
    })
})

module.exports = router;
