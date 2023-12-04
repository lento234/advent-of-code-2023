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
    var result = try part1(data);
    print("Part 1: {}\n", .{result});
}

test "part 1" {
    const input =
        \\ 1abc2
        \\ pqr3stu8vwx
        \\ a1b2c3d4e5f
        \\ treb7uchet
    ;
    try expect(@as(u32, 142) == part1(input) catch 0);
}

test "part 2" {
    const input =
        \\ two1nine
        \\ eightwothree
        \\ abcone2threexyz
        \\ xtwone3four
        \\ 4nineeightseven2
        \\ zoneight234
        \\ 7pqrstsixteen
    ;
    try expect(@as(u32, 281) == part2(input) catch 0);
}

fn part2(data: []const u8) !u32 {

    // Split data to lines
    var lines = std.mem.tokenize(u8, data, "\n");

    var i: u8 = 0;
    while (lines.next()) |line| {
        if (std.mem.indexOf(u8, "one", 0) != null) {
            print("{}: {s}\n", .{ i, line });
        }
        i += 1;
    }

    return 0;
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
