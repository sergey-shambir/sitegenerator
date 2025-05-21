const crypto = require('crypto');
const fs = require('fs');

/**
 * 
 * @param {string} filePath
 * @returns 
 */
function calculateFileHash(filePath) {
    const bytes = fs.readFileSync(filePath);
    const hash = crypto.createHash('md5');
    hash.update(bytes);
    return hash.digest('hex');
}

module.exports.calculateFileHash = calculateFileHash
