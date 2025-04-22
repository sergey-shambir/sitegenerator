const fs = require('fs');

if (process.argv.length < 1) {
    console.error(`Usage: ${process.argv[0]} <path-to-markdown>`);
    process.exit(1);
}

const inputFilePath = process.argv[0];

const data = fs.readFileSync(inputFilePath, 'utf8');
console.log(data);
