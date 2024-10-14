#ifndef FUNC_WRITE_H
#define FUNC_WRITE_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* read_name_from_file(const char* file_path) {
    FILE* file = fopen(file_path, "r");
    if (!file) {
        perror("Failed to open file");
        return NULL;
    }

    char line[256];
    while (fgets(line, sizeof(line), file)) {
        if (strncmp(line, "name =", 6) == 0) {
            char* name = strtok(line + 6, "\r\n");
            fclose(file);
            return strdup(name);
        }
    }

    fclose(file);
    return NULL;
}

char* read_file_html(const char* file_path) {
    FILE* file = fopen(file_path, "r");
    if (!file) {
        perror("Failed to open file");
        return NULL;
    }

    char line[256];
    while (fgets(line, sizeof(line), file)) {
        if (strncmp(line, "html =", 6) == 0) {
            char* html = strtok(line + 6, "\r\n");
            fclose(file);
            return strdup(html);
        }
    }

    fclose(file);
    return NULL;
}

int read_window_height(const char* file_path) {
    FILE* file = fopen(file_path, "r");
    if (!file) {
        perror("Failed to open file");
        return -1;
    }

    char line[256];
    while (fgets(line, sizeof(line), file)) {
        if (strncmp(line, "window_h =", 10) == 0) {
            char* value_str = strtok(line + 10, "\r\n");
            if (value_str) {
                int height = atoi(value_str);
                fclose(file);
                return height;
            }
        }
    }

    fclose(file);
    return -1;
}

int read_window_width(const char* file_path) {
    FILE* file = fopen(file_path, "r");
    if (!file) {
        perror("Failed to open file");
        return -1;
    }

    char line[256];
    while (fgets(line, sizeof(line), file)) {
        if (strncmp(line, "window_w =", 10) == 0) {
            char* value_str = strtok(line + 10, "\r\n");
            if (value_str) {
                int width = atoi(value_str);
                fclose(file);
                return width;
            }
        }
    }

    fclose(file);
    return -1;
}

#endif 
