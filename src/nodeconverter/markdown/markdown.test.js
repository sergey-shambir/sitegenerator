const { convertMarkdown } = require('./markdown');
const { join } = require('path');

const testDataDir = join(__dirname, 'test_data');

describe("convertMarkdown function", () => {
    it("can process images", () => {
        const markdownPath = join(testDataDir, 'images.md')
        const html = convertMarkdown(markdownPath);
        expect(html).toMatchSnapshot();
    })
})