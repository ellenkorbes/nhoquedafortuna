var express = require('express')
var mongo = require('mongodb').MongoClient
var path = require('path')
var math = require('mathjs')
var bodyParser = require('body-parser')
var app = express()
var moment = require('moment-timezone')
var mongouri = 'mongodb://' + process.env.USER + ':' + process.env.PASS + '@' + process.env.HOST + ':' + process.env.PORT + '/' + process.env.DB;

app.use(bodyParser.json())

app.use(bodyParser.urlencoded({
  extended: true
}))

app.post('/form', function (request, response) {
  mongo.connect(mongouri, function (err, db) {
    if (err) throw err
    console.log(request.body)
    var now = moment().tz("America/Sao_Paulo").format()
    var newEntry = [{
      "name": request.body.name,
      "email": request.body.email,
      "ip": request.headers['x-forwarded-for'].split(',')[0],
      "time": now
    }]
    db.collection('leads').insert(newEntry, function (err, result) {
      if (err) throw err
      console.log(result)
      db.close()
    })
    response.redirect("http://www.nhoquedafortuna.com/thanks.html")
  }
  )
})


app.get("/", function (request, response) {
  response.redirect("http://www.nhoquedafortuna.com")
  })

var listener = app.listen(3000, () => {
  console.log('Listening on port ' + listener.address().port);
})

app.get("/teste", function (request, response) {
  response.redirect("http://www.nhoquedafortuna.com") // nhoque.glitch.me
})
