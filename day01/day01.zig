const std = @import("std");

fn isNumber(char: u8) bool {
    return switch (char) {
        '0'...'9' => true,
        else => false,
    };
}

pub fn main() !void {
    // std.debug.print("Advent of code: day {d}\n", .{1});
    std.log.info("Advent of code: day {}", .{1});

    // Read file
    var buffer: [32768]u8 = undefined;
    // const data = try std.fs.cwd().readFile("test_input.txt", &buffer);
    const data = try std.fs.cwd().readFile("input.txt", &buffer);
    const len = data.len;
    std.log.info("{} bytes read!", .{len});

    // Split data to lines
    var lines = std.mem.tokenize(u8, data, "\n");

    // allocator
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const alloc = arena.allocator();

    // Inspect text
    // var i: u32 = 0;
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

    std.log.info("Result: {}", .{result});
}
