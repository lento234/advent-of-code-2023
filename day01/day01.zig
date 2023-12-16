const std = @import("std");

// aliases
const expect = std.testing.expect;
const print = std.debug.print;

// allocator
var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
// defer arena.deinit();
const alloc = arena.allocator();

fn isNumber(char: u8) bool {
    return (char >= '0') and (char <= '9');
}

pub fn main() !void {
    print("Advent of code: day {}\n", .{1});

    // Parse input
    var buffer: [32768]u8 = undefined;
    const data = try std.fs.cwd().readFile("input.txt", &buffer);

    // Part 1
    const result1 = try part1(data);
    print("Part 1: {}\n", .{result1});

    // Part 1
    const result2 = try part2(data);
    print("Part 2: {}\n", .{result2});
}

test "part 1" {
    const input =
        \\1abc2
        \\pqr3stu8vwx
        \\a1b2c3d4e5f
        \\treb7uchet
    ;
    try expect(@as(u32, 142) == part1(input) catch 0);
}

test "part 2" {
    const input =
        \\two1nine
        \\eightwothree
        \\abcone2threexyz
        \\xtwone3four
        \\4nineeightseven2
        \\zoneight234
        \\7pqrstsixteen
    ;
    try expect(@as(u32, 281) == part2(input) catch 0);
}

fn findDigit(string: []const u8, substr: []const u8) ?usize {
    return std.mem.indexOf(u8, string, substr);
}

fn findDigitPos(string: []const u8, start_index: usize, substr: []const u8) ?usize {
    return std.mem.indexOfPos(u8, string, start_index, substr);
}

fn part2(data: []const u8) !u32 {

    // Split data to lines
    var lines = std.mem.tokenize(u8, data, "\n");

    var wDigits = std.mem.split(u8, "one two three four five six seven eight nine", " ");
    var nDigits = std.mem.split(u8, "1 2 3 4 5 6 7 8 9", " ");

    var result: u32 = 0;

    var i: usize = 0;
    while (lines.next()) |line| {
        var list = std.ArrayList(u8).init(alloc);
        var loc = std.ArrayList(usize).init(alloc);
        defer loc.deinit();
        defer list.deinit();

        var k: u8 = 1;

        while (wDigits.next()) |digit| {
            var idx: usize = 0;
            while (findDigitPos(line, idx, digit)) |newidx| {
                try list.append(k);
                try loc.append(newidx);
                idx = newidx + 1;
            }
            k += 1;
        }
        wDigits.reset();

        k = 1;
        while (nDigits.next()) |digit| {
            var idx: usize = 0;
            while (findDigitPos(line, idx, digit)) |newidx| {
                try list.append(k);
                try loc.append(newidx);
                idx = newidx + 1;
            }
            k += 1;
        }
        nDigits.reset();

        const minIdx: usize = std.mem.indexOfMin(usize, loc.items);
        const maxIdx: usize = std.mem.indexOfMax(usize, loc.items);

        const value: u8 = list.items[minIdx] * 10 + list.items[maxIdx];
        // print("{d}: {d}\n", .{ i, value });
        result += @as(u32, value);
        i += 1;
    }

    return result;
}

fn part1(data: []const u8) !u32 {

    // Split data to lines
    var lines = std.mem.tokenize(u8, data, "\n");

    // Inspect text
    var result: u32 = 0;
    while (lines.next()) |line| {
        var list = std.ArrayList(u8).init(alloc);
        defer list.deinit();
        // Check character
        for (line) |char| {
            if (isNumber(char)) {
                try list.append(char - '0');
            }
        }

        const value: u8 = list.items[0] * 10 + list.items[list.items.len - 1];

        result += @as(u32, value);
    }

    return result;
}
