import {useEffect, useState, useRef} from "react";

interface Message {
	username: string;
	content: string;
}

interface ChatProps {
	username: string;
}

const Chat = ({username}: ChatProps) => {
	const [messages, setMessages] = useState<Message[]>([]);
	const [input, setInput] = useState<string>("");
	const socketRef = useRef<WebSocket | null>(null);

	useEffect(() => {
		socketRef.current = new WebSocket(`ws://localhost:8080/ws?username=${encodeURIComponent(username)}`);

		socketRef.current.onopen = () => {
			console.log("✅ WebSocket connection opened");
		};

		socketRef.current.onmessage = (event) => {
			const msg: Message = JSON.parse(event.data);
			setMessages((prev) => [...prev, msg]);
		};

		socketRef.current.onerror = (error) => {
			console.error("❌ WebSocket error:", error);
		};

		socketRef.current.onclose = () => {
			console.log("❎ WebSocket connection closed");
		};

		return () => {
			socketRef.current?.close();
		};
	}, [username]);

	const sendMessage = () => {
		if (input.trim() && socketRef.current?.readyState === WebSocket.OPEN) {
			const message = {username, content: input};
			socketRef.current.send(JSON.stringify(message));
			setInput("");
		}
	};

	return (
		<div className="chat-container">
			<div className="messages">
				{messages.map((msg, idx) => (
					<div key={idx} className="message">
						<strong>{msg.username}: </strong>{msg.content}
					</div>
				))}
			</div>
			<div className="input-container">
				<input
					type="text"
					placeholder="Введите сообщение..."
					value={input}
					onChange={(e) => setInput(e.target.value)}
					onKeyDown={(e) => e.key === "Enter" && sendMessage()}
				/>
				<button onClick={sendMessage}>Отправить</button>
			</div>
		</div>
	);
};

export default Chat;