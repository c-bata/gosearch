/* global db */
load('fixtures/common.js');

var data = [];

function createData() {
    for (var i = 0; i < indexCount; i++) {
        var num = zeroFill(i);

        var keyword = 'keyword' + num;
        var url = [];

        for (var j=0; j < urlCount; j++) {
            url.push('http://example.com/' + num + zeroFill(j))
        }

        var obj = {
            _id: 'index' + num,
            Keyword: keyword,
            Url: url
        };
        data.push(obj);
    }
}

createData();

db.index.drop();
db.index.insert(data);
