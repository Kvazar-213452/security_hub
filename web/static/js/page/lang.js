function change_lang_now() {
    $.ajax({
        url: "/config_global",
        type: "POST",
        contentType: "application/json",
        data: JSON.stringify(null),
        success: function (response) {
            lang_global = response['lang'];
            lang_change_page(lang_global);
        },
        error: function (xhr, status, error) {
            console.error("Помилка при відправці:", status, error);
        }
    });
}

function lang_change_page(lang) {
    if (lang === "en") {
        // wifi
        $('#lang_wefsdeeeeee').html("WIFI management");
        $('#lang_nfefdfdvghyt').html("Get more information about yours WIFI");
        $('#lang_dqwedddcwww').html("Connection:");

        // system
        $('#lang_system_wdeewds').html("Data os");
        $('#lang_system_fesrdfyyyyy').html("library data");
        $('#lang_system_wfr839wefsff').html("NetworkAdapters data");
        $('#lang_system_v00qwdweee').html("Open programs");
        $('#lang_system_vfd8723ed').html("System data");
        $('#lang_system_verdfvcww').html("There is something interesting here");

        // antivirus
        $('#antivirus_errfee2').html("Antivirus");
        $('#antivirus_vef093f').html("Be in information security");
        $('#dwdc21e12d').html("Check the site");
        $('#dwdc21e12d1').html("Check the file");
        $('#dwdc21e12d2').html("On the background");
        $('#antivirus_894534ffvvv').html("Check");
        $('#h3ruiwefer24f').html("Description");
        $('#fewrvw243rgefvcc').html("Background features are enabled automatically at startup");
        $('#vbs612dwes655').html("Monitoring flash drives");
        $('#vb92354gu04ttg').html("This setting activates the USB media monitoring function, which notifies users when new flash drives are connected. After enabling this function, the program constantly scans available USB media and reports each new device that is connected.");
        $('#bbv612ee3dwe').html("Set the cmd command that will be launched after connecting the USB drive");
        $('#bg_dqwderfd').html("Turn on");
        $('#upload-button').html("Quick check");
        $('#upload-button1').html("Detailed check");

        // cleaning
        $('#vc728i32000').html("Cleaning pc");
        $('#bvh9iuweddd').html("Unfortunately, the script does not clean the System32 folder. Sorry");

        $('#bx66210doddddw').html("Cleaning information");
        $('#bo9ho2_0dddd').html("Clearing DNS and ARP tables");
        $('#d_odw2983fff').html("File system logs");
        $('#b8cq82d3333').html("Deleting the history of open paths");
        $('#vc972odwe9').html("Clearing the internet cache");
        $('#b89bdq9822222').html("Deleting specific files");
        $('#bc92idw092222d').html("Deleting temporary files, cache and logs");
        $('#vwer43erfdfcff').html("Additional settings");
        $('#bvpuwe033230_233').html("Remove Windows backup copies");
        $('#v21qp9wd000222s').html('The "Remove Windows backup copies" setting deletes saved system backups that were created for restoring Windows in case of failure or data loss. Backups may include system images (system backups), restore points, as well as automatic file archives.');
        $('.cbq98222fz233').html("Install");
        $('#cbwq982fff45t5').html("Delete all Wi-Fi profiles");
        $('#cbqw278d436666').html('The "Delete all Wi-Fi profiles" setting will remove saved wireless network profiles that your device has connected to.');
        $('#bv89qwdasdwededddd').html("Delete remote desktop connection settings.");
        $('#vb0928923ee3ddd').html('The "Delete remote desktop connection settings" setting deletes saved configurations and connection history for remote desktops.');
        $('#bviubiwdd3333r').html("Restart the doskey command processor");
        $('#bciu292ed45675663').html(`
            The "Restart doskey command processor" setting refreshes the doskey command processor.
            <br><br>
            Restarting doskey clears the command history and resets all defined macros. This can be useful for resetting the command line environment to its default state.
        `);
        $('#bf0qwp32r4t5651222').html("Clean the computer");
    } else if (lang === "uk") {
        // wifi
        $('#lang_wefsdeeeeee').html("Вайфай менеджер");
        $('#lang_nfefdfdvghyt').html("Отримайте більше інформації про ваш вайфай");
        $('#lang_dqwedddcwww').html("З'єднання:");

        // system
        $('#lang_system_wdeewds').html("Дані операційної системи");
        $('#lang_system_fesrdfyyyyy').html("Дані бібліотек");
        $('#lang_system_wfr839wefsff').html("Дані мережевих адаптерів");
        $('#lang_system_v00qwdweee').html("Відкриті програми");
        $('#lang_system_vfd8723ed').html("Системні дані");
        $('#lang_system_verdfvcww').html("Тут є дещо цікаве можливо");
        
        // antivirus
        $('#antivirus_errfee2').html("Антивірус");
        $('#antivirus_vef093f').html("Будьте в інформаційній безпеці");
        $('#dwdc21e12d').html("Перевірити сайт");
        $('#dwdc21e12d1').html("Перевірити файл");
        $('#dwdc21e12d2').html("На фоні");
        $('#antivirus_894534ffvvv').html("Перевірити");
        $('#h3ruiwefer24f').html("Опис");
        $('#fewrvw243rgefvcc').html("Функції на фоні вмикаються автоматично при запуску");
        $('#vbs612dwes655').html("Моніторинг флешок");
        $('#vb92354gu04ttg').html("Це налаштування активує функцію моніторингу USB-носіїв, яка попереджає користувача при підключенні нових флешок. Після ввімкнення цієї функції програма постійно сканує доступні USB-носії і повідомляє про кожен новий пристрій, що підключається.");
        $('#bbv612ee3dwe').html("Встановити cmd команду яка буде запускатись після підключення USB-носія");
        $('#bg_dqwderfd').html("Увімкнути");
        $('#upload-button').html("Швидка перевірка");
        $('#upload-button1').html("Детальна перевірка");

        // cleaning
        $('#vc728i32000').html("Очищення ПК");
        $('#bvh9iuweddd').html("На жаль скрипт не очищує папку System32. Вибачте");
        $('#bx66210doddddw').html("Інформація очистки");
        $('#bo9ho2_0dddd').html("Очищення DNS та ARP-таблиць");
        $('#d_odw2983fff').html("Журнали файлової системи");
        $('#b8cq82d3333').html("Видалення історії відкритих шляхів");
        $('#vc972odwe9').html("Очищення інтернет-кешу");
        $('#b89bdq9822222').html("Видалення специфічних файлів");
        $('#bc92idw092222d').html("Видалення тимчасових файлів, кешу та логів");
        $('#vwer43erfdfcff').html("Додаткові налаштування");
        $('#bvpuwe033230_233').html("Видалити резервні копії Windows");
        $('#v21qp9wd000222s').html('Налаштування "Видалити резервні копії Windows" видаляє збережені резервні копії системи, які створювалися для відновлення Windows у випадку збою або втрати даних. Резервні копії можуть включати зображення системи (системні копії), точки відновлення, а також автоматичні архіви файлів.');
        $('.cbq98222fz233').html("Встановити");
        $('#cbwq982fff45t5').html("Видалити всі профілі Wi-Fi");
        $('#cbqw278d436666').html('Налаштування "Видалити всі профілі Wi-Fi" призведе до видалення збережених профілів бездротових мереж, до яких підключався ваш пристрій.');
        $('#bv89qwdasdwededddd').html("Видалити налаштування підключення до віддаленого робочого столу.");
        $('#vb0928923ee3ddd').html('Налаштування "Видалити налаштування підключення до віддаленого робочого столу" видаляє збережені конфігурації та історію підключень до віддалених робочих столів.');
        $('#bviubiwdd3333r').html("Перезапустити командний процесор doskey");
        $('#bciu292ed45675663').html(`
        Налаштування "Перезапускає командний процесор doskey" оновлює роботу командного процесора doskey
                            <br><br>
                            Перезапуск doskey очищає історію команд та скидає всі визначені макроси. Це може бути корисним для скидання середовища командного рядка до початкового стану.
        `);
        $('#bf0qwp32r4t5651222').html("Очистити комп'ютер");
    }
}
