//
//  FileUtils.swift
//  2022
//
//  Created by Teboho Mphure on 2022/12/02.
//

import Foundation

class FileUtils {
    // repos/advent-of-code/2022/2022/files/day1.1.txt
    static func ReadFile(directory: String) -> [String] {
        let filePath = FileManager.default.urls(for: .documentDirectory, in: .userDomainMask).first!.appendingPathComponent(directory)
        let text = try! String(contentsOf: filePath, encoding: .utf8)

        let lines = text.components(separatedBy: .newlines)
        
        return lines
    }
}
