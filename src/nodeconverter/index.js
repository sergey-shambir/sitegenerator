const fs = require('fs');
const MarkdownIt = require('markdown-it');
const markdownItMeta = require('markdown-it-meta');
const markdownItTitle = require('markdown-it-title');
const markdownItHighlightJs = require('markdown-it-highlightjs');
const markdownItExternalLinks = require('markdown-it-external-links');
const markdownItAnchor = require('markdown-it-anchor');
const markdownItTocDoneRight = require('markdown-it-toc-done-right');
const markdownItImageFigures = require('markdown-it-image-figures');
const { default: slugify } = require('@sindresorhus/slugify');

function createMarkdownIt() {
    return MarkdownIt({ html: true })
        .use(markdownItMeta)
        .use(markdownItTitle)
        .use(markdownItHighlightJs)
        .use(markdownItExternalLinks, {
            externalClassName: 'external-link',
            internalClassName: 'internal-link',
            externalTarget: '_blank',
        })
        .use(markdownItAnchor, { slugify: s => slugify(s) })
        .use(markdownItTocDoneRight, { slugify: s => slugify(s) })
        .use(markdownItImageFigures, {
            dataType: true,
            figcaption: "title",
            lazy: true,
            async: true
        });
}

const args = process.argv.slice(2);
if (args.length != 1) {
    console.error(`Usage: ${process.argv[0]} ${process.argv[1]} <path-to-markdown>`);
    process.exit(1);
}

const inputFilePath = args[0];
const md = createMarkdownIt();

const markdown = fs.readFileSync(inputFilePath, 'utf8');
const html = md.render(markdown)
console.log(html);