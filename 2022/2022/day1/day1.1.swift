//
//  day1.1.swift
//  2022
//
//  Created by Teboho Mphure on 2022/12/01.
//

import Foundation

class Day1_1 {
    var total = 0
    var totalCalories: [Int] = []

    func PrintMaxCalories() {

        let lines = FileUtils.ReadFile(directory: "repos/advent-of-code/2022/2022/files/day1.1.txt")

        for line in lines {
            let num = Int(line) ?? 0
            
            if num != 0 {
                total += num
            } else {
                totalCalories.append(total)
                total = 0
            }
        }

        print(totalCalories.max()!)
    }
}

