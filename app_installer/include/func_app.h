#ifndef FUNC_APP_H
#define FUNC_APP_H

#include <string>

void copyFileToDirectory(const std::string& sourceFile, const std::string& targetDir);
void CreateShortcut(const std::wstring& shortcutPath, const std::wstring& targetPath, const std::wstring& workingDir);
void runCommandInBackground(const char* command);

#endif