package com.example;

import javafx.application.Application;
import javafx.scene.Scene;
import javafx.scene.web.WebView;
import javafx.scene.web.WebEngine;
import javafx.stage.Stage;

// .sorse/core_sorse/shell_web_NM3/src/main/java/com/example/Main.java

public class Main extends Application {
    @Override
    public void start(Stage primaryStage) {
        String[] args = getParameters().getRaw().toArray(new String[0]);

        String url = "null";

        if (args.length > 0) {
            url = args[0];
        }

        WebView webView = new WebView();
        WebEngine webEngine = webView.getEngine();
        webEngine.load(url);

        Scene scene = new Scene(webView, 800, 600);
        primaryStage.setScene(scene);
        primaryStage.setTitle("SPX");

        primaryStage.setOnCloseRequest(event -> {
            System.exit(0);
        });

        primaryStage.show();
    }

    public static void main(String[] args) {
        launch(args);
    }
}
