/* global db */

var indexCount = 8;
var urlCount = 3;

function zeroFill(num, length) {
    length = length || 3;
    var pref = (new Array(length)).join(0);
    return (pref + num).slice(-length);
}
