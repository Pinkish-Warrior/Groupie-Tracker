"use strict";

document.getElementById("startButton").addEventListener("click", function () {
  fetch("/start-server", { method: "POST" })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Failed to start server");
      }
      console.log("Server started successfully");
    })
    .catch((error) => {
      console.error("Error starting server:", error);
    });
});

document.getElementById("stopButton").addEventListener("click", function () {
  fetch("/stop-server", { method: "POST" })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Failed to stop server");
      }
      console.log("Server stopped successfully");
    })
    .catch((error) => {
      console.error("Error stopping server:", error);
    });
});
