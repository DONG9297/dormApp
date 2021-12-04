const express = require('express');
const router = express.Router();
const request = require("request");

router.get('/', (req, res) => {
    request({
        url: "http://login:10712/is_logged",
        method: "POST",
        json: true,
        headers: {
            'content-type': 'application/json;charset=UTF-8',
            "Cookie": "jwt="+req.cookies.jwt
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

module.exports = router;