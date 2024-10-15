#include <iostream>
#include <filesystem>
#include <fstream>
#include <string>
#include <unordered_set>
#include <openssl/evp.h> // Потрібно для EVP

namespace fs = std::filesystem;

std::string calculate_sha256(const std::string& path) {
    unsigned char hash[EVP_MAX_MD_SIZE];
    unsigned int hash_len;

    EVP_MD_CTX* mdctx = EVP_MD_CTX_new();
    EVP_DigestInit_ex(mdctx, EVP_sha256(), nullptr);

    std::ifstream file(path, std::ios::binary);
    if (!file.is_open()) {
        EVP_MD_CTX_free(mdctx);
        return "";
    }

    char buffer[4096];
    while (file.read(buffer, sizeof(buffer))) {
        EVP_DigestUpdate(mdctx, buffer, file.gcount());
    }
    EVP_DigestUpdate(mdctx, buffer, file.gcount());
    EVP_DigestFinal_ex(mdctx, hash, &hash_len);
    EVP_MD_CTX_free(mdctx);

    // Перетворюємо хеш на рядок
    std::string result;
    for (unsigned int i = 0; i < hash_len; i++) {
        char buf[3]; // два символи для хеша + нульовий символ
        snprintf(buf, sizeof(buf), "%02x", hash[i]);
        result += buf;
    }
    return result;
}

bool is_malicious(const std::string& hash, const std::unordered_set<std::string>& virus_signatures) {
    return virus_signatures.find(hash) != virus_signatures.end();
}

int main() {
    std::unordered_set<std::string> virus_signatures = {
        "abc123", // Приклад підпису вірусу
        // Додайте інші підписи
    };

    std::string directory_path = "C:/path/to/scan";
    for (const auto& entry : fs::directory_iterator(directory_path)) {
        if (entry.is_regular_file()) {
            std::string file_path = entry.path().string();
            std::string file_hash = calculate_sha256(file_path);

            if (is_malicious(file_hash, virus_signatures)) {
                std::cout << "Malicious file detected: " << file_path << std::endl;
                // Вжити заходів, наприклад, видалити файл
            }
        }
    }

    return 0;
}

