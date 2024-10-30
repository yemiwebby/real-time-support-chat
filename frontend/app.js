const chatBox = document.getElementById("chat-box");
const messageInput = document.getElementById("message-input");
const sendBtn = document.getElementById("send-btn");

sendBtn.addEventListener("click", () => {
  const message = messageInput.value.trim();
  if (message !== "") {
    addMessageToChat("You", message);
    sendMessage(message);
    messageInput.value = "";
  }
});

function addMessageToChat(sender, message) {
  const messageElement = document.createElement("div");
  messageElement.classList.add("message");
  if (sender === "You") {
    messageElement.classList.add("user-message");
  }
  messageElement.textContent = `${sender}: ${message}`;
  chatBox.appendChild(messageElement);
}

async function sendMessage(message) {
  const response = await fetch("http://localhost:8080/send-message", {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: `channelSID=<YOUR_CHANNEL_SID>&from=webUser&body=${encodeURIComponent(message)}`,
  });

  if (response.ok) {
    console.log("Message sent successfully");
  } else {
    console.error("Failed to send message");
  }
}
