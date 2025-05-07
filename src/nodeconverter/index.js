const markdown = require('./markdown');
const sass = require('./sass');

function printUsageAndExit() {
    console.error(`Usage: ${process.argv[0]} ${process.argv[1]} <format> <input-path>`);
    console.error('Supported formats: markdown, sass');
    process.exit(1);
}

const args = process.argv.slice(2);
if (args.length != 2) {
    printUsageAndExit();
}

const format = args[0];
const path = args[1];

if (format === 'markdown') {
    console.log(markdown.convertMarkdown(path));
} else if (format === 'sass') {
    console.log(sass.convertSass(path));
} else {
    printUsageAndExit();
}