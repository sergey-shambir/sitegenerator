const sass = require('sass');

function convertSass(path) {
    return sass.compile(path, {
        sourceMap: false
    }).css;
}

exports.convertSass = convertSass;