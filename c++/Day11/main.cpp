#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>
#include <set>
#include <fstream>
#include <cstdint>
#include <chrono>

struct Point {
    int64_t x, y;
    static int64_t Distance(const Point& p1, const Point& p2) {
        auto dx = std::abs(p2.x - p1.x);
        auto dy = std::abs(p2.y - p1.y);
        return dx + dy;
    }
};

constexpr int64_t kExpansion = 999999;

std::vector<Point> Expand(const std::vector<Point> &initial_galaxies, const std::set<int64_t > &empty_columns,
                          const std::set<int64_t> &empty_rows) {
    std::vector<Point> expansion(initial_galaxies.size());
    for (auto i: empty_columns) {
        for (int j = 0; j < initial_galaxies.size(); j++) {
            if (initial_galaxies[j].x > i) {
                expansion[j].x += 1;
            }
        }
    }
    for (auto i: empty_rows) {
        for (int j = 0; j < initial_galaxies.size(); j++) {
            if (initial_galaxies[j].y > i) {
                expansion[j].y += 1;
            }
        }
    }
    std::vector<Point> expanded(initial_galaxies.size());
    for (int i = 0; i < initial_galaxies.size(); i++) {
        expanded[i] = Point{initial_galaxies[i].x + (expansion[i].x * kExpansion), initial_galaxies[i].y + (expansion[i].y * kExpansion)};
    }
    return expanded;
}

int main() {
    auto time_start = std::chrono::high_resolution_clock::now();

    std::fstream input("../Day11/input.txt");
    if (!input.is_open()) {
        std::cout << "ERROR: file not open!" << std::endl;
        return 1;
    }

    std::set<int64_t> empty_rows, empty_columns;
    std::vector<Point> galaxies;

    std::string line;
    for (int row = 0; std::getline(input, line); row++) {
        if (empty_columns.empty()) {
            // at the beginning assume all columns are empty
            for (int i = 0; i < line.length(); i++) {
                empty_columns.insert(i);
            }
        }

        bool no_galaxies = true;
        for (int column = 0; column < line.length(); column++) {
            if (line[column] == '#') {
                no_galaxies = false;
                empty_columns.erase(column);
                galaxies.push_back(Point(column, row));
            }
        }

        if (no_galaxies) {
            empty_rows.insert(row);
        }

    }


    auto expanded_galaxies = Expand(galaxies, empty_columns, empty_rows);
    int64_t sum = 0;
    for (int i = 0; i < expanded_galaxies.size(); i++) {
        for (int j = i; j < expanded_galaxies.size(); j++) {
            if (i == j) {
                continue;
            }
            sum += Point::Distance(expanded_galaxies[i], expanded_galaxies[j]);
        }
    }

    std::cout << "Sum: " << sum << std::endl;
    std::cout << "Took " << std::chrono::duration_cast<std::chrono::microseconds>(std::chrono::high_resolution_clock::now() - time_start).count() << "us" << std::endl;
    return 0;
}
