#include <iostream>
#include <string>
#include <map>
#include <algorithm>
#include <fstream>

std::map<std::string, bool> dict;
char sep = ' ';

bool hasNonAlpha(std::string word) {
    for (std::string::size_type i = 0; i < word.length(); ++i) {
        char c = word[i];
        if (c < 'a' || c > 'z') return true;
    }
    return false;
}

void printWord(std::string word, std::string lower) {
    if (hasNonAlpha(lower) || dict.find(lower) == dict.end()) {
        std::cout << '<' << word << '>';
    } else {
        std::cout << word;
    }
}

int main(int argc, char ** argv) {
    std::string dictpath;
    if (argc == 1) {
        dictpath = "/usr/share/dict/words";
    } else {
        dictpath = argv[1];
    }
    std::ifstream infile(dictpath.c_str());
    std::string word;
    while (infile >> word) {
        dict[word] = true;
    }

    std::string line;
    while (std::getline(std::cin, line)) {
        if (line.empty()) {
            std::cout << std::endl;
        } else {
            size_t pos = 0;
            bool first = true;
            while ((pos = line.find(sep)) != std::string::npos) {
                if (! first) {
                    std::cout << ' ';
                }
                first = false;
                std::string token = line.substr(0, pos);
                std::string lower = token;
                std::transform(lower.begin(), lower.end(), lower.begin(), ::tolower);
                printWord(token, lower);
                line.erase(0, pos + 1);
            }
            if (!first) std::cout << ' ';
            printWord(line, line);
            std::cout << std::endl;
        }
    }
}
