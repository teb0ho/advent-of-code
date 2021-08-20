const fs = require('fs');
       
fs.readFile('./files/day7.txt', (err, data) => {

    if (err) {
        throw err;
    }

    const bagsContainingGold = new Set();
    const lines = data.toString().split('\n');


    for (const line of lines) {
        if (line.includes(' shiny gold')) {
            bagsContainingGold.add(line.split(' bags ')[0]);
        }
    }

    let found = true;
    while(found)  {
        found = false;
        for (const bag of bagsContainingGold) {
            for (const line of lines) {
                if (line.includes(` ${bag}`)) {
                    if (!bagsContainingGold.has(line.split(' bags ')[0])) {
                        found = true;
                        bagsContainingGold.add(line.split(' bags ')[0]);
                    }

                }
            }
        }
    }

    console.log(bagsContainingGold);
});
