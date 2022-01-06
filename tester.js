const WS_SERVER = "ws://127.0.0.1:8080/ws";
let socket_connected = false;
let id_client = 1036;
let socket = null;

document.addEventListener(
	"DOMContentLoaded",
	() => {
		console.log("creo conexion web socket");
		socket = new WebSocket(WS_SERVER);
		/// logica de websockets

		console.log("Attempting Connection...");

		socket.onopen = () => {
			console.log("Successfully Connected");
			socket_connected = true;
			/*
			const login_message = {
				client_id: 123,
			};

			socket.send(JSON.stringify(login_message));
			*/
		};

		socket.onclose = (event) => {
			console.log("Socket Closed Connection: ", event);
			socket.send("Client Closed!");
		};

		socket.onerror = (error) => {
			console.log("Socket Error: ", error);
		};

		socket.onmessage = (message) => {
			console.log(message);
		};
	},
	false
);

const login = (client_id, socket) => {
	const message = {
		type: "login",
		client_id: parseInt(client_id, 10),
	};
	console.log(message);

	if (socket_connected) {
		socket.send(JSON.stringify(message));
	}
};

document.getElementById("login").addEventListener("click", () => {
	console.log("Voy a registrar el cliente");
	let client_id = document.getElementById("client_id").value;

	if (socket_connected) {
		login(client_id, socket);
	}
});

document.getElementById("send_message").addEventListener("click", () => {
	console.log("click 2");

	let text_message = document.getElementById("message").value;
	let peer = document.getElementById("peer").value;
	console.log("Message : ", text_message);

	const message = {
		type: "message",
		client_id: parseInt(peer, 10),
		message: text_message,
	};

	if (socket_connected) {
		socket.send(JSON.stringify(message));
	}
});
