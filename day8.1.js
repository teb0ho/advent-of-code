const fs = require('fs');
       
fs.readFile('./files/day8.txt', (err, data) => {

    if (err) {
        throw err;
    }

    let acc = 0;
    const lines = data.toString().split('\n');
    const visits = new Array(lines.length);

    let i = 0;
    while (i < lines.length) {
        if (visits[i]) {
            break;
        }

        const instructionValue = lines[i].split(' ');

        if (instructionValue[0] === 'jmp') {
            visits[i] = 1;
            i = inBounds(i, +instructionValue[1], visits.length);
        }

        if (instructionValue[0] === 'nop') {
            visits[i] = 1;
            i++;
        }
        else if (instructionValue[0] === 'acc') {
            acc += +instructionValue[1];
            visits[i] = 1;
            i++;
        }
    }

    console.log(acc);

});

function inBounds(currentIndex, acc, arraySize) {
    if (acc > 0) {
        let i = 0;
        while (i < acc) {
            i++;
            currentIndex++;
            if (currentIndex == arraySize) {
                currentIndex = 0;
            }
        }
    } else {
        let i = acc;
        while (i < 0) {
            i++;
            if (currentIndex < 0) {
                currentIndex = arraySize - 1;
            } else {
                currentIndex--;
            }
        }
    }

    return currentIndex;

}


