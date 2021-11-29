// 参考 https://segmentfault.com/a/1190000013308591

const express = require('express');
const app = express();
let cookieParser = require('cookie-parser');
app.use(cookieParser());
app.use(express.urlencoded({extended: true}))

app.set('views', './views');
app.set('view engine', 'ejs');
app.engine('.html', require('ejs').__express);

const index = require('./routes/index')
const user = require('./routes/user')
const order = require('./routes/order')
app.use("/", index);
app.use("/user", user);
app.use("/order", order);


app.listen(8080, (res, req) => {
    console.log('Node app start at port 8080');
});
