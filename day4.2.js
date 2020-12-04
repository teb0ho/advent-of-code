const fs = require('fs');

let count = 0;
let passports;
       
fs.readFile('./files/day4.txt', (err, data) => {
    

    passports = data.toString().split('\n\n');

    for (let i = 0; i < passports.length; i++) {
        passports[i] = passports[i].replace(/\n/g, ' ');

        const passportFields = passports[i].split(' ');
        const fieldsFound = [];


        for (const fields of passportFields) {
            const fieldName = fields.split(':');

            switch (fieldName[0]) {
                case 'byr':
                    if (fieldName[1] >= 1920 && fieldName[1] <= 2002) {
                        fieldsFound.push('byr');
                    }
                    break;
            
                case 'iyr':
                    if (fieldName[1] >= 2010 && fieldName[1] <= 2020) {
                        fieldsFound.push('iyr');      
                    }
                    break;
            
                case 'eyr':
                    if (fieldName[1] >= 2020 && fieldName[1] <= 2030) {
                        fieldsFound.push('eyr'); 
                    }        
                    break;
            
                case 'hgt':
                    if (fieldName[1].includes('cm') || fieldName[1].includes('in')) {
                        
                        if (fieldName[1].includes('cm') && (fieldName[1].slice(0, -2) >= 150 && 
                        fieldName[1].slice(0, -2) <= 193)) { 
                            fieldsFound.push('hgt'); 
                        }
                        else if (fieldName[1].includes('in') && (fieldName[1].slice(0, 2) >= 59 && 
                        fieldName[1].slice(0, -2) <= 76)) { 
                            fieldsFound.push('hgt');
                        }
          
                    }       
                    break;
            
                case 'hcl':
                    if (fieldName[1].includes('#')) {
                        let re = /[0-9,a-f]/g;
                        const matches = fieldName[1].match(re);
                        if (matches.length == 6) {
                            fieldsFound.push('hcl');
                        }
                    }
                    break;
            
                case 'ecl':

                    if (fieldName[1] == 'amb' || fieldName[1] == 'blu' ||
                    fieldName[1] == 'brn' || fieldName[1] == 'gry' ||
                    fieldName[1] == 'grn' || fieldName[1] == 'hzl' ||
                    fieldName[1] == 'oth') {
                        fieldsFound.push('ecl');  
                    }        
                    break;
            
                case 'pid':
                    if (/^[0-9]{9}$/g.test(fieldName[1])) {
                        fieldsFound.push('pid');
                    }        
                    break;
            }
        }
        
        
        if (fieldsFound.length == 7) {
            count++;
        }
    }
    
    console.log(count);
});
