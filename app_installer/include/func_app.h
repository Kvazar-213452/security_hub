#ifndef FUNC_APP_H
#define FUNC_APP_H

#include <string>

void unzip(const std::string& zipFilePath, const std::string& destDir);
void CreateShortcut(const std::wstring& shortcutPath, const std::wstring& targetPath, const std::wstring& workingDir);
void runCommandInBackground(const char* command);

#endif