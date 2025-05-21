const { join } = require("path");
const { calculateFileHash } = require("./calculatefilehash");

const testDataDir = join(__dirname, 'test_data');

describe("calculateFileHash function", () => {
    it("can calculate PNG image hash", () => {
        const imagePath = join(testDataDir, 'img', 'image.png');
        const imageHash = calculateFileHash(imagePath);

        expect(imageHash).toBe('7633de907692432289b47b6b885f3ee7');
    });

    it("can calculate JPEG image hash", () => {
        const imagePath = join(testDataDir, 'img', 'image.jpg');
        const imageHash = calculateFileHash(imagePath);

        expect(imageHash).toBe('9eac7817096bcdc0a63dfa179d76df0f');
    });
})
