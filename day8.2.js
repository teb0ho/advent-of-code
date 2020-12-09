const fs = require('fs');
       
fs.readFile('./files/day8.txt', (err, data) => {

    if (err) {
        throw err;
    }

    let acc = 0;
    const lines = data.toString().split('\n');
    found = false;
    index = 0;
    while (!found) {
        let linesModified = [];
        for (let i = index; i < lines.length; i++) {
            linesModified = lines.slice();
            const instruction = linesModified[i].split(' ');
            if (instruction[0] === 'jmp') {
                linesModified[i] = linesModified[i].replace('jmp', 'nop');
                index = i+1;
                break;               
            } 
        }
        let visits = new Array(lines.length);
        let j = 0;
        acc = 0;
        while (j < lines.length) {
            if (visits[j]) {
                break;
            } 
    
            if (j  === linesModified.length - 1) {
                console.log('break');
                found = true;
                break;
            }
    
            const instructionValue = linesModified[j].split(' ');
    
            if (instructionValue[0] === 'jmp') {
                visits[j] = 1;
                j = inBounds(j, +instructionValue[1], visits.length);
            }
    
            if (instructionValue[0] === 'nop') {
                visits[j] = 1;
                j++;
            }
            else if (instructionValue[0] === 'acc') {
                acc += +instructionValue[1];
                visits[j] = 1;
                j++;
            }
        }
    }

    console.log(acc);
    console.log(index);

});

function inBounds(currentIndex, acc, arraySize) {
    if (acc > 0) {
        let i = 0;
        while (i < acc) {
            i++;
            currentIndex++;
            if (currentIndex == arraySize - 1) {
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


