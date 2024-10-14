#include <iostream>
#include <fstream>
#include <unordered_set>
#include <string>

std::unordered_set<std::string> loadVirusSignatures(const std::string& filename) {
    std::unordered_set<std::string> virus_signatures;
    std::ifstream file(filename);
    
    if (!file.is_open()) {
        std::cerr << "Не вдалося відкрити файл: " << filename << std::endl;
        return virus_signatures;
    }

    std::string line;
    while (std::getline(file, line)) {
        virus_signatures.insert(line);
    }

    file.close();
    return virus_signatures;
}

int main() {
    std::unordered_set<std::string> virus_signatures = loadVirusSignatures("hashes.txt");
    
    // Перевірка наявності хешу
    std::string test_hash = "5eb63bbbe01eeed093cb22bb8f5acdd2"; // Введіть тестовий хеш
    if (virus_signatures.find(test_hash) != virus_signatures.end()) {
        std::cout << "Вірус виявлено!" << std::endl;
    } else {
        std::cout << "Вірус не виявлено." << std::endl;
    }

    return 0;
}
