const fs = require('fs');

fs.readFile('./files/day9.txt', (err, data) => {

    if (err) {
        throw err;
    }
    const number = new Set();
    const lines = data.toString().split('\n');

    console.log(+lines[0] + +lines[1]);
    for (let i = 25; i < lines.length; i++) {
        for (let j = i - 25; j < i; j++) {
            for (let k = i - 25; k < i; k++) {
                if (j !== k && lines[j] !== lines[k]) {
                    if (+lines[j] + +lines[k] === +lines[i]) {
                        number.add(lines[i]);
                        break;
                    }
                }

            }

        }

    }

    let finalAnswer = 0;
    for (let i = 25; i < lines.length; i++) {
        if (!number.has(lines[i])) {
            finalAnswer = lines[i];
            break;
        }

    }

    console.log(finalAnswer);

});
