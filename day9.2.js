const fs = require('fs');

fs.readFile('./files/day9.txt', (err, data) => {

    if (err) {
        throw err;
    }
    const lines = data.toString().split('\n');
    let sumFound = false;
    const indices = [];
    let i = 0;
    let j = i + 1;

    while (i < lines.length) {
        console.log(i);
        while(j < lines.length) {
            let sum = 0;
            for (let k = i; k <= j; k++) {
                sum += +lines[k];
            }
            if (sum === 14360655) {
                console.log('found');
                sumFound = true;
                indices.push(i);
                indices.push(j)
                break;
            } else {
                j++;
            }
        }
        if (sumFound) {
            break;
        }
        i++;
        j = i + 1;
    }

    const elements = lines.slice(indices[0], indices[1] + 1);
    const answer = elements.reduce((a, b) => Math.min(a, b)) + elements.reduce((a, b) => Math.max(a, b));
    
    console.log(answer);
});
