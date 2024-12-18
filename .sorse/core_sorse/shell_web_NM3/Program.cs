using System;
using System.Windows.Forms;
using Microsoft.Web.WebView2.WinForms;

namespace WebViewExample {
    static class Program {
        [STAThread]
        static void Main() {
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);

            var form = new Form {
                Text = "Browser - Security Hub",
                Width = 1024,
                Height = 768
            };

            var webView = new WebView2 {
                Dock = DockStyle.Fill
            };
            form.Controls.Add(webView);

            webView.CoreWebView2InitializationCompleted += (sender, e) => {
                if (e.IsSuccess) {
                    webView.CoreWebView2.Navigate("https://github.com/Kvazar-213452/security_hub");
                }
                else {
                    MessageBox.Show($"Помилка ініціалізації WebView2: {e.InitializationException.Message}", "Помилка", MessageBoxButtons.OK, MessageBoxIcon.Error);
                }
            };

            webView.EnsureCoreWebView2Async();

            Application.Run(form);
        }
    }
}
