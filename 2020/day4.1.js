const fs = require('fs');

let count = 0;
const fieldsFoundInEachPassport = [];
       
fs.readFile('./files/day4.txt', (err, data) => {

    if (err) {
        throw err;
    }
    
    const passports = data.toString().split('\n\n');

    for (let i = 0; i < passports.length; i++) {
        passports[i] = passports[i].replace(/\n/g, ' ');

        const passportFields = passports[i].split(' ');
        const fieldsFound = [];


        for (const fields of passportFields) {
            const fieldName = fields.split(':')[0];

            switch (fieldName) {
                case 'byr':
                    fieldsFound.push('byr');        
                    break;
            
                case 'iyr':
                    fieldsFound.push('iyr');       
                    break;
            
                case 'eyr':
                    fieldsFound.push('eyr');        
                    break;
            
                case 'hgt':
                    fieldsFound.push('hgt');       
                    break;
            
                case 'hcl':
                    fieldsFound.push('hcl');
                    break;
            
                case 'ecl':
                    fieldsFound.push('ecl');        
                    break;
            
                case 'pid':
                    fieldsFound.push('pid');        
                    break;
            
                case 'cid':
                    fieldsFound.push('cid'); 
                    break;
            }
        }
        
        fieldsFoundInEachPassport.push(fieldsFound);
    }

    for (const result of fieldsFoundInEachPassport) {
        if (result.length == 7 && !result.includes('cid')) {
            count++;
        }

        if (result.length == 8) {
            count++;
        }
    }

    console.log(count);

    
});
