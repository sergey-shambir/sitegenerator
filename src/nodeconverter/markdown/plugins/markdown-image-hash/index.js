const {resolve} = require('path');
const {calculateFileHash} = require('./calculatefilehash');

/**
 * Adds file hash to markdown image URL for assets versioning.
 *   - finds image in directory passed through `env.assetsDir`
 * 
 * Options:
 *   - addHashFn - function that adds hash to URL, default function transforms "img.png" into "img.png?v=fa028bc3"
 *   - hashLength - truncated hash length, note that original hash has 32 characters
 * 
 * 
 * 
 * @param {*} md 
 * @param {{
 *  addHashFn: Function | undefined,
 *  hashLength: number | undefined
 * }} options 
 */
function imageAttrsPlugin(md, options) {
    options = options || {};
    const addHashFn = options.addHashFn || ((src, hash) => src + '?v=' + hash);
    const hashLength = options.hashLength || 8;

    function applyImageAttrs(state) {
        // do not process first and last token
        for (let i = 1, l = state.tokens.length; i < (l - 1); ++i) {
            const token = state.tokens[i];

            if (token.type !== 'inline') {
                continue;
            }
            // children: image alone, or link_open -> image -> link_close
            if (!token.children || (token.children.length !== 1 && token.children.length !== 3)) {
                continue;
            }
            // one child, should be img
            if (token.children.length === 1 && token.children[0].type !== 'image') {
                continue;
            }
            // three children, should be image enclosed in link
            if (token.children.length === 3) {
                const [childrenA, childrenB, childrenC] = token.children;
                const isEnclosed = childrenA.type !== 'link_open' ||
                    childrenB.type !== 'image' ||
                    childrenC.type !== 'link_close';

                if (isEnclosed) {
                    continue;
                }
            }
            // prev token is paragraph open
            if (i !== 0 && state.tokens[i - 1].type !== 'paragraph_open') {
                continue;
            }
            // next token is paragraph close
            if (i !== (l - 1) && state.tokens[i + 1].type !== 'paragraph_close') {
                continue;
            }

            // for linked images, image is one off
            let image = token.children.length === 1 ? token.children[0] : token.children[1];

            const assetsDir = state.env ? state.env.assetsDir : undefined;
            const srcAttr = image.attrs.find(([k]) => k === 'src');
            const src = srcAttr[1];
            const imagePath = resolveImagePath(src, assetsDir);
            const imageHash = calculateFileHash(imagePath).substring(0, hashLength);
            srcAttr[1] = addHashFn(src, imageHash);
        }
    }

    md.core.ruler.before('linkify', 'image_attrs', applyImageAttrs);
}

function resolveImagePath(imagePath, assetsDir) {
    return assetsDir ? resolve(assetsDir, imagePath) : resolve(imagePath);
}

module.exports = imageAttrsPlugin;