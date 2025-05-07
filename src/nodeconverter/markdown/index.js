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

function convertMarkdown(path) {
    const md = createMarkdownIt();
    const markdown = fs.readFileSync(path, 'utf8');
    return md.render(markdown);
}

exports.convertMarkdown = convertMarkdown;
